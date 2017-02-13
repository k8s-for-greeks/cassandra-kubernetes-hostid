package hostid

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
	apps_v1beta1 "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	typed_v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/pkg/api/v1"
)

type K8sClientInterface interface {
	GetPod(podName string, podNamespace string) (*v1.Pod, error)
	Pods(namespace string) typed_v1.PodInterface
	StatefulSets(namespace string) apps_v1beta1.StatefulSetInterface
}

type K8sClient struct {
	k8sClient *kubernetes.Clientset
}

func CreateK8sClientInCluster() (*K8sClient, error) {

	k8sClient, err := GetClientInCluster()

	if err != nil {
		return nil, fmt.Errorf("Unable to create k8s client %v", err)
	}

	return &K8sClient{
		k8sClient: k8sClient,
	}, nil
}

func CreateK8sClientOutsideCluster(config string) (*K8sClient, error) {

	k8sClient, err := GetClientClusterOutOfCluster(config)

	if err != nil {
		return nil, fmt.Errorf("Unable to create k8s client %v", err)
	}

	return &K8sClient{
		k8sClient: k8sClient,
	}, nil
}

func (bw K8sClient) StatefulSets(namespace string) apps_v1beta1.StatefulSetInterface {
	return bw.k8sClient.StatefulSets(namespace)
}

func (bw K8sClient) Pods(namespace string) typed_v1.PodInterface {
	return bw.k8sClient.Pods(namespace)
}

func (bw K8sClient) GetPod(podName string, podNamespace string) (*v1.Pod, error) {
	pod, err := bw.k8sClient.CoreV1().Pods(podNamespace).Get(podName)

	if err != nil {
		return nil, fmt.Errorf("Unable to get my own pod %v", err)
	}

	return pod, nil
}
