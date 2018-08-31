package v1alpha1

import (
	//"github.com/zhanglianx111/pkg/apis/nfscontroller"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

//var SchemeGroupVersion = schema.GroupVersion{Group: nfscontroller.GroupName, Version: "v1alpha1"}
var SchemeGroupVersion = schema.GroupVersion{Group: "nfscontroller.yonghui.cn", Version: "v1alpha1"}

func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&NFS{},
		&NFSList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
