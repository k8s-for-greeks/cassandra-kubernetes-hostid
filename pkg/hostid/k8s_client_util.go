package hostid

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

)

func GetClientInCluster() (c *kubernetes.Clientset, err error) {

	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("")
	}

	// creates the clientset
	c, err = kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func GetClientClusterOutOfCluster(kubeconfigFile string) (c *kubernetes.Clientset, err error) {

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigFile)
	if err != nil {
		panic(err.Error())
	}

	return newForConfig(config)
}

func newForConfig(config *rest.Config) (c *kubernetes.Clientset, err error) {
	// creates the clientset
	c, err = kubernetes.NewForConfig(config)

	if err != nil {
		return nil, err
	}
	return c, nil
}
