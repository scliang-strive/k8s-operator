package main

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	confPath := `F:\GoProjects\src\k8s-operator\config\kube.config`
	config, err := clientcmd.BuildConfigFromFlags("", confPath)
	if err != nil {
		fmt.Println("error :", err.Error())
		panic(err)
	}
	config.GroupVersion = &v1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs
	clientRest, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}
	pod := v1.Pod{}
	err = clientRest.Get().Namespace("default").Resource("pods").Name("nginx").Do(context.TODO()).Into(&pod)
	if err != nil {
		fmt.Println("error: ", err.Error())
	} else {
		fmt.Println("pod name: ", pod.Name)
	}
}
