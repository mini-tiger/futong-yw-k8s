package ksc

import (
	appsv1beta1 "k8s.io/api/apps/v1beta1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	rbacv1 "k8s.io/api/rbac/v1"
	storagev1 "k8s.io/api/storage/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"
)

// getClientByGroupVersion
func (r *resourceHandler) getClientByGroupVersion(groupVersion schema.GroupVersionResource) rest.Interface {
	switch groupVersion.Group {
	case corev1.GroupName:
		return r.client.CoreV1().RESTClient()
	case appsv1beta1.GroupName:
		if groupVersion.Version == "v1" {
			return r.client.AppsV1().RESTClient()
		}
		return r.client.AppsV1beta1().RESTClient()
	case autoscalingv1.GroupName:
		return r.client.AutoscalingV1().RESTClient()
	case batchv1.GroupName:
		if groupVersion.Version == "v1beta1" {
			return r.client.BatchV1beta1().RESTClient()
		}
		return r.client.BatchV1().RESTClient()
	case extensionsv1beta1.GroupName:
		return r.client.ExtensionsV1beta1().RESTClient()
	case storagev1.GroupName:
		return r.client.StorageV1().RESTClient()
	case rbacv1.GroupName:
		return r.client.RbacV1().RESTClient()
	default:
		return r.client.CoreV1().RESTClient()
	}
}
