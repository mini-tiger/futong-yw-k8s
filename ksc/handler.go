package ksc

import (
	"context"
	"fmt"

	"ftk8s/base/cfg"
	"ftk8s/base/enum"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
)

type IResource interface {
	Create(kind string, namespace string, object *runtime.Unknown) (*runtime.Unknown, error)
	Update(kind string, namespace string, name string, object *runtime.Unknown) (*runtime.Unknown, error)
	Get(kind string, namespace string, name string) (runtime.Object, error)
	List(kind string, namespace string, labelSelector string) ([]runtime.Object, error)
	Delete(kind string, namespace string, name string, options *metav1.DeleteOptions) error
}

type resourceHandler struct {
	client  *kubernetes.Clientset
	ksCache *KsCache
}

func NewResourceHandler(kubeClient *kubernetes.Clientset, ksCache *KsCache) IResource {
	return &resourceHandler{
		client:  kubeClient,
		ksCache: ksCache,
	}
}

func (r *resourceHandler) Create(kind string, namespace string, object *runtime.Unknown) (*runtime.Unknown, error) {
	resource, err := r.getResource(kind)
	if err != nil {
		return nil, err
	}

	k8sCli := r.getClientByGroupVersion(resource.GroupVersionResourceKind.GroupVersionResource)
	req := k8sCli.Post().
		Resource(kind).
		SetHeader(enum.HeaderKeyJson, enum.HeaderValueJson).
		Body(object.Raw)
	if resource.Namespaced {
		req.Namespace(namespace)
	}
	var result runtime.Unknown
	err = req.Do(context.Background()).Into(&result)

	return &result, err
}

func (r *resourceHandler) Update(kind string, namespace string, name string, object *runtime.Unknown) (*runtime.Unknown, error) {
	resource, err := r.getResource(kind)
	if err != nil {
		return nil, err
	}

	k8sCli := r.getClientByGroupVersion(resource.GroupVersionResourceKind.GroupVersionResource)
	req := k8sCli.Put().
		Resource(kind).
		Name(name).
		SetHeader(enum.HeaderKeyJson, enum.HeaderValueJson).
		Body(object.Raw)
	if resource.Namespaced {
		req.Namespace(namespace)
	}

	var result runtime.Unknown
	err = req.Do(context.Background()).Into(&result)

	return &result, err
}

func (r *resourceHandler) Delete(kind string, namespace string, name string, options *metav1.DeleteOptions) error {
	resource, err := r.getResource(kind)
	if err != nil {
		return err
	}

	k8sCli := r.getClientByGroupVersion(resource.GroupVersionResourceKind.GroupVersionResource)
	req := k8sCli.Delete().
		Resource(kind).
		Name(name).
		Body(options)
	if resource.Namespaced {
		req.Namespace(namespace)
	}

	return req.Do(context.Background()).Error()
}

func (r *resourceHandler) Get(kind string, namespace string, name string) (runtime.Object, error) {
	resource, err := r.getResource(kind)
	if err != nil {
		return nil, err
	}

	genericInformer, err := r.ksCache.sharedInformerFactory.ForResource(resource.GroupVersionResourceKind.GroupVersionResource)
	if err != nil {
		return nil, err
	}
	lister := genericInformer.Lister()
	var result runtime.Object
	if resource.Namespaced {
		result, err = lister.ByNamespace(namespace).Get(name)
		if err != nil {
			return nil, err
		}
	} else {
		result, err = lister.Get(name)
		if err != nil {
			return nil, err
		}
	}
	result.GetObjectKind().SetGroupVersionKind(schema.GroupVersionKind{
		Group:   resource.GroupVersionResourceKind.Group,
		Version: resource.GroupVersionResourceKind.Version,
		Kind:    resource.GroupVersionResourceKind.Kind,
	})

	return result, nil
}

func (r *resourceHandler) List(kind string, namespace string, labelSelector string) ([]runtime.Object, error) {
	resource, err := r.getResource(kind)
	if err != nil {
		return nil, err
	}

	genericInformer, err := r.ksCache.sharedInformerFactory.ForResource(resource.GroupVersionResourceKind.GroupVersionResource)
	if err != nil {
		return nil, err
	}
	selectors, err := labels.Parse(labelSelector)
	if err != nil {
		cfg.Mlog.Error("failed to build label selector, error message: ", err.Error())
		return nil, err
	}

	lister := genericInformer.Lister()
	var objs []runtime.Object
	if resource.Namespaced {
		objs, err = lister.ByNamespace(namespace).List(selectors)
		if err != nil {
			return nil, err
		}
	} else {
		objs, err = lister.List(selectors)
		if err != nil {
			return nil, err
		}
	}

	for i := range objs {
		objs[i].GetObjectKind().SetGroupVersionKind(schema.GroupVersionKind{
			Group:   resource.GroupVersionResourceKind.Group,
			Version: resource.GroupVersionResourceKind.Version,
			Kind:    resource.GroupVersionResourceKind.Kind,
		})
	}

	return objs, nil
}

func (r *resourceHandler) getResource(kind string) (Resource, error) {
	resourceMap := GetResourceMap()
	resource, ok := resourceMap[kind]
	var err error
	if !ok {
		err = fmt.Errorf("resource kind (%s) not support yet", kind)
		return resource, err
	}
	return resource, err
}
