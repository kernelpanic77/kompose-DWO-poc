package main

import (
	"fmt"
	"log"

	"github.com/kubernetes/kompose/pkg/kobject"
	"github.com/kubernetes/kompose/pkg/loader"
	"github.com/kubernetes/kompose/pkg/transformer/kubernetes"
)

// func convertRuntimeObjectToKindMeta(in runtime.Object) v1.APIGroup {

// }

func main(){

	// var kubeconfig *string
	// if home := homedir.HomeDir(); home != "" {
	// 	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	// } else {
	// 	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	// }

	// fmt.Println(*kubeconfig)
	// flag.Parse()

	// fmt.Print("hello")
	// config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	// if err != nil {
	// 	panic(err)
	// }
	// clientset, err := kube.NewForConfig(config)
	// if err != nil {
	// 	panic(err)
	// }

	//fmt.Print(*clientset)

	l, err := loader.GetLoader("compose")
	if err != nil{
		log.Fatal(err)
	}

	//fmt.Println(*clientset)

	komposeObj := kobject.KomposeObject{
		ServiceConfigs: make(map[string]kobject.ServiceConfig),
	}


	fmt.Printf("%+v\n", komposeObj)

	opt := kobject.ConvertOptions{
		CreateD: true,
		CreateDeploymentConfig: true,
		Volumes: "persistentVolumeClaim",
		Replicas: 1,
		Provider: "kubernetes",
		InputFiles: []string{"docker-compose.yml"},
	}

	fmt.Printf("%+v\n", opt)

	// opt := kobject.ConvertOptions{ToStdout:false, CreateD:true, CreateRC:false, CreateDS:false, CreateDeploymentConfig:true, BuildRepo:, BuildBranch:, Build:none, PushImage:false,PushImageRegistry:, CreateChart:false, GenerateYaml:false, GenerateJSON:false, StoreManifest:false, EmptyVols:false, Volumes:persistentVolumeClaim, PVCRequestSize: InsecureRepository:false Replicas:1 InputFiles:[docker-compose.yaml] OutFile: Provider:kubernetes Namespace: Controller: IsDeploymentFlag:false IsDaemonSetFlag:false IsReplicationControllerFlag:false IsReplicaSetFlag:false IsDeploymentConfigFlag:false IsNamespaceFlag:false Server: YAMLIndent:2 WithKomposeAnnotation:true MultipleContainerMode:false ServiceGroupMode: ServiceGroupName:}

	// inputFiles := []string{"docker-compose.yml"}
	komposeObj, err = l.LoadFile(opt.InputFiles)

	if(err != nil){
		log.Fatal(err)
	}

	//fmt.Printf("komposeObj: %v\n", komposeObj)
	// keys := make([]string, 0, len(komposeObj.ServiceConfigs))
	// for k := range komposeObj.ServiceConfigs {
	// 	keys = append(keys, k)
	// }
	
	t := &kubernetes.Kubernetes{Opt: opt}
	
	objects, err := t.Transform(komposeObj, opt)
	
	if(err != nil){
		log.Fatal(err)
	}

	fmt.Println("runtime objects", objects)

	
	for i := 0; i < len(objects); i++ {
		fmt.Println(objects[i].GetObjectKind().GroupVersionKind())
	}

	// fmt.Printf("t1: %T\n", myobj)
	
}