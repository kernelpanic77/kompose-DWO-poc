package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"reflect"

	"github.com/kubernetes/kompose/pkg/kobject"
	"github.com/kubernetes/kompose/pkg/loader"
	"github.com/kubernetes/kompose/pkg/transformer/kubernetes"

	//"k8s.io/apiserver/pkg/admission/plugin/webhook/namespace"
	kube "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main(){

	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kube.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	l, err := loader.GetLoader("compose")
	if err != nil{
		log.Fatal(err)
	}

	komposeObj := kobject.KomposeObject{
		ServiceConfigs: make(map[string]kobject.ServiceConfig),
	}

	opt := kobject.ConvertOptions{
		CreateD: true,
		CreateDeploymentConfig: true,
		Volumes: "persistentVolumeClaim",
		Replicas: 1,
		Provider: "kubernetes",
		InputFiles: []string{"docker-compose.yml"},
	}

	komposeObj, err = l.LoadFile(opt.InputFiles)

	if(err != nil){
		log.Fatal(err)
	}

	t := &kubernetes.Kubernetes{Opt: opt}
	
	objects, err := t.Transform(komposeObj, opt)
	
	if(err != nil){
		log.Fatal(err)
	}

	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
		
	// dep :=  objects[2].(*appsv1.Deployment)
	// fmt.Printf("%v\n", dep)

	deploymentObjs := make([]*appsv1.Deployment, 0)
	for i := 0; i < len(objects); i++ {
		if(objects[i].GetObjectKind().GroupVersionKind().Kind == "Deployment"){
			depl := objects[i].(*appsv1.Deployment)
			deploymentObjs = append(deploymentObjs, depl)
		}
	}

	for i := 0; i < len(deploymentObjs); i++ {
		result, err := deploymentsClient.Create(context.TODO(), deploymentObjs[i], metav1.CreateOptions{})
		if err != nil {
			panic(err)
		}
		fmt.Printf("Create Deployment %q.\n", result.GetObjectMeta().GetName())
	}

	// creating service objects
	keys := reflect.ValueOf(komposeObj.ServiceConfigs).MapKeys()
	serviceObjs := make([]*apiv1.Service, 0)
	for _, key := range keys {
		svc := *t.CreateService(key.Interface().(string), komposeObj.ServiceConfigs[key.Interface().(string)])
		serviceObjs = append(serviceObjs, &svc)
	}
	
	servicesClient := clientset.CoreV1().Services(apiv1.NamespaceDefault)

	for i := 0; i < len(serviceObjs); i++ {
		result, err := servicesClient.Create(context.TODO(), serviceObjs[i], metav1.CreateOptions{})
		if err != nil {
			panic(err)
		}
		fmt.Printf("Create Service %q.\n", result.GetObjectMeta().GetName())
	}
}