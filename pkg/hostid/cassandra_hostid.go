// Copyright [2016] [Matthew Stump, Vorstella Corp.]

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hostid

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"regexp"

	"github.com/golang/glog"
	"k8s.io/client-go/pkg/api"
	"k8s.io/client-go/pkg/api/v1"
)

const (
	AnnotationsKey  string = "annotations"
	MetaDataKey     string = "metadata"
	NodetoolCommand string = "info"
	UuidRegex       string = "[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}"
)

type CasssandraHostId struct {
	Client            K8sClientInterface
	NodetoolPath      string
	Namespace         string
	PodName           string
	AnnotationsPrefix string
}

func CreateCasssandraHostId(nodeToolPath string, pod string, namespace string, annotationPrefix string) (*CasssandraHostId, error) {

	if annotationPrefix == "" {
		annotationPrefix = "cassandra"
	}

	c, err := CreateK8sClientInCluster()
	if err != nil {
		return nil, err
	}
	return &CasssandraHostId{
		NodetoolPath:      nodeToolPath, // FIXME how do we get this??
		PodName:           pod,
		Namespace:         namespace,
		AnnotationsPrefix: annotationPrefix,
		Client:            c,
	}, err
}

//nodetoolPath := flag.String("nodetool", "/usr/bin/nodetool", "path to cassandra nodetool")
//namespace := flag.String("namespace", "default", "the kubernetes namespace")
//podName := flag.String("pod", "", "the kubernetes pod name")
//populateHostId := flag.Bool("populate", false, "populate the k8s annotations with our host ID")
//fetchHostId := flag.Bool("fetch", false, "fetch our host ID from k8s annotations")
//annotationsPrefix := flag.String("prefix", "cassandra", "the prefix for the annotations tracking host IDs")

func (h CasssandraHostId) getSerializedReference() (ref *v1.ObjectReference, err error) {

	pod, err := h.Client.Pods(h.Namespace).Get(h.PodName)
	if err != nil {
		glog.Errorf("foo", h.PodName)
		return nil, err
	}

	annotations := pod.ObjectMeta.Annotations
	createdByJson, ok := annotations[api.CreatedByAnnotation]
	if !ok {
		return nil, err
	}

	var createdBySerialized v1.SerializedReference
	err = json.Unmarshal([]byte(createdByJson), &createdBySerialized)
	if err != nil {
		return nil, err
	}

	return &createdBySerialized.Reference, nil
}

func (h CasssandraHostId) GetHostId() (*string, error) {
	r, err := h.getSerializedReference()

	if err != nil {
		return nil, fmt.Errorf("Could not find StatefulSet: %s", err)
	}

	statefulSet, err := h.Client.StatefulSets(r.Namespace).Get(r.Name)
	if err != nil {
		return nil, fmt.Errorf("Could not find StatefulSet %s: %s", r.Name, err)
	}

	id, ok := statefulSet.ObjectMeta.Annotations[getAnnotationName(h.AnnotationsPrefix, h.PodName)]
	if !ok {
		return nil, fmt.Errorf("Host ID for %s not present in annotations for StatefulSet %s: %s", h.PodName, r.Name, err)
	}

	return &id, nil
}

func (h CasssandraHostId) SaveHostId() error {

	r, err := h.getSerializedReference()

	if err != nil {
		fmt.Errorf("Could not find StatefulSet: %s", err)
	}

	id, err := getCassandraHostId(h.NodetoolPath)
	if err != nil {
		return fmt.Errorf("Could not obtain Cassandra host ID: %s", err)
	}

	statefulSets := h.Client.StatefulSets(r.Namespace)
	statefulSet, err := statefulSets.Get(r.Name)
	if err != nil {
		return fmt.Errorf("Could not find StatefulSet %s, %s", r.Name, err)
	}

	patch := make(map[string]map[string]map[string]string)
	annotationsPatch := make(map[string]map[string]string)
	annotationsPatch[AnnotationsKey] = map[string]string{getAnnotationName(h.AnnotationsPrefix, h.PodName): *id}
	patch[MetaDataKey] = annotationsPatch

	patchBytes, err := json.Marshal(patch)

	if err != nil {
		return fmt.Errorf("Could not generate patch JSON: %s", err)
	}

	glog.Info("Patching StatefulSet: %s", string(patchBytes))

	result, err := statefulSets.Patch(statefulSet.Name, api.MergePatchType, patchBytes)
	if err != nil {
		return fmt.Errorf("Error patching StatefulSet: %s", err)
	}

	glog.Infof("Resulting annotations: %s", result.ObjectMeta.Annotations)
	return nil
}

func getAnnotationName(prefix string, name string) string {
	return fmt.Sprintf("%v/%v", prefix, name)
}

func getCassandraHostId(nodetool string) (*string, error) {
	nodetoolOutput, err := runCommand(nodetool, NodetoolCommand)
	if err != nil {
		return nil, err
	}
	uuidRegex, _ := regexp.Compile(UuidRegex)
	match := uuidRegex.FindString(nodetoolOutput)
	if len(match) == 0 {
		return nil, errors.New("couldn't fetch Cassandra host ID")
	}
	return &match, nil
}

func runCommand(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	var stdOut bytes.Buffer
	cmd.Stdout = &stdOut
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("unable to run command %s", err)
	}
	return stdOut.String(), nil
}
