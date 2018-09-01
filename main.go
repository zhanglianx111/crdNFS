package main

import (
	"flag"
	"time"

	"github.com/golang/glog"

	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/tools/clientcmd"

	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	clientset "github.com/zhanglianx111/nfscontroller/pkg/client/clientset/versioned"
	"github.com/zhanglianx111/nfscontroller/pkg/signals"
	//"k8s.io/apimachinery/pkg/apis/meta/v1"
	informers "github.com/zhanglianx111/nfscontroller/pkg/client/informers/externalversions"
	"k8s.io/client-go/kubernetes"
)

var (
	masterURL  string
	kubeconfig string
)

func main() {
	flag.Parse()

	// set up signals so we handle the first shutdown signal gracefully
	stopCh := signals.SetupSignalHandler()

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		glog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	nfsClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building example clientset: %s", err.Error())
	}


	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(kubeClient, time.Second*30)
	nfsInformerFactory := informers.NewSharedInformerFactory(nfsClient, time.Second*30)

	controller := NewController(kubeClient, nfsClient, kubeInformerFactory, nfsInformerFactory)

	go kubeInformerFactory.Start(stopCh)
	go nfsInformerFactory.Start(stopCh)

	if err = controller.Start(2, stopCh); err != nil {
		glog.Fatalf("Error running controller: %s", err.Error())
	}
	/*
		list, err := nfsClient.NfscontrollerV1alpha1().Foos("default").Get("foo", v1.GetOptions{})
		if err != nil {
			glog.Errorf("xxx")
		}
		glog.Info(list)
	*/
}

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}
