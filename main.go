package main

import (
	"fmt"
	"log"

	"github.com/kubernetes/kompose/pkg/kobject"
	"github.com/kubernetes/kompose/pkg/loader"
	"github.com/kubernetes/kompose/pkg/transformer/kubernetes"
)

func main(){
	l, err := loader.GetLoader("compose")
	if err != nil{
		log.Fatal(err)
	}

	komposeObj := kobject.KomposeObject{
		ServiceConfigs: make(map[string]kobject.ServiceConfig),
	}

	opt := kobject.ConvertOptions{
		Provider: "kubernetes",
		InputFiles: []string{"docker-compose.yml"},
	}

	// inputFiles := []string{"docker-compose.yml"}
	komposeObj, err = l.LoadFile(opt.InputFiles)

	if(err != nil){
		log.Fatal(err)
	}

	fmt.Printf("komposeObj: %v\n", komposeObj)
	// keys := make([]string, 0, len(komposeObj.ServiceConfigs))
	// for k := range komposeObj.ServiceConfigs {
	// 	keys = append(keys, k)
	// }
	
	t := &kubernetes.Kubernetes{Opt: opt}
	
	objects, err := t.Transform(komposeObj, opt)
	
	if(err != nil){
		log.Fatal(err)
	}


	// Print output
	err = kubernetes.PrintList(objects, opt)
	if err != nil {
		log.Fatalf(err.Error())
	}
}