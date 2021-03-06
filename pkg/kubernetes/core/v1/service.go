// Code generated by helmit-generate. DO NOT EDIT.

package v1

import (
	"github.com/onosproject/helmit/pkg/kubernetes/resource"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

var ServiceKind = resource.Kind{
	Group:   "core",
	Version: "v1",
	Kind:    "Service",
	Scoped:  true,
}

var ServiceResource = resource.Type{
	Kind: ServiceKind,
	Name: "services",
}

func NewService(service *corev1.Service, client resource.Client) *Service {
	return &Service{
		Resource:           resource.NewResource(service.ObjectMeta, ServiceKind, client),
		Object:             service,
		EndpointsReference: NewEndpointsReference(client, resource.NewUIDFilter(service.UID)),
	}
}

type Service struct {
	*resource.Resource
	Object *corev1.Service
	EndpointsReference
}

func (r *Service) Delete() error {
	return r.Clientset().
		CoreV1().
		RESTClient().
		Delete().
		NamespaceIfScoped(r.Namespace, ServiceKind.Scoped).
		Resource(ServiceResource.Name).
		Name(r.Name).
		VersionedParams(&metav1.DeleteOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do().
		Error()
}
