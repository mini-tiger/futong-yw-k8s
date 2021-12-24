package ksc

import (
	"fmt"

	"ftk8s/util"

	jsoniter "github.com/json-iterator/go"
	appsv1 "k8s.io/api/apps/v1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	rbacv1 "k8s.io/api/rbac/v1"
	storagev1 "k8s.io/api/storage/v1"
)

// ValidateTemplate verify that the resource template content format is correct
func ValidateTemplate(templateKind string, content string) error {
	switch templateKind {
	case KindNameConfigMap:
		configMap := corev1.ConfigMap{}
		return jsoniter.Unmarshal(util.StringToByteSlice(content), &configMap)
	case KindNameDaemonSet:
		daemonSet := appsv1.DaemonSet{}
		return jsoniter.Unmarshal(util.StringToByteSlice(content), &daemonSet)
	case KindNameDeployment:
		deployment := appsv1.Deployment{}
		return jsoniter.Unmarshal(util.StringToByteSlice(content), &deployment)
	case KindNameEvent:
		event := corev1.Event{}
		return jsoniter.Unmarshal(util.StringToByteSlice(content), &event)
	case KindNameHorizontalPodAutoscaler:
		horizontalPodAutoscaler := autoscalingv1.HorizontalPodAutoscaler{}
		return jsoniter.Unmarshal(util.StringToByteSlice(content), &horizontalPodAutoscaler)
	case KindNameIngress:
		ingress := extensionsv1beta1.Ingress{}
		return jsoniter.Unmarshal(util.StringToByteSlice(content), &ingress)
	case KindNameJob:
		job := batchv1.Job{}
		return jsoniter.Unmarshal(util.StringToByteSlice(content), &job)
	case KindNameCronJob:
		cronJob := batchv1beta1.CronJob{}
		return jsoniter.Unmarshal(util.StringToByteSlice(content), &cronJob)
	case KindNameNamespace:
		namespace := corev1.Namespace{}
		return jsoniter.Unmarshal(util.StringToByteSlice(content), &namespace)
	case KindNameNode:
		node := corev1.Node{}
		return jsoniter.Unmarshal(util.StringToByteSlice(content), &node)
	case KindNamePersistentVolumeClaim:
		persistentVolumeClaim := corev1.PersistentVolumeClaim{}
		return jsoniter.Unmarshal(util.StringToByteSlice(content), &persistentVolumeClaim)
	case KindNamePersistentVolume:
		persistentVolume := corev1.PersistentVolume{}
		return jsoniter.Unmarshal(util.StringToByteSlice(content), &persistentVolume)
	case KindNamePod:
		pod := corev1.Pod{}
		return jsoniter.Unmarshal(util.StringToByteSlice(content), &pod)
	case KindNameReplicaSet:
		replicaSet := appsv1.ReplicaSet{}
		return jsoniter.Unmarshal(util.StringToByteSlice(content), &replicaSet)
	case KindNameSecret:
		secret := corev1.Secret{}
		return jsoniter.Unmarshal(util.StringToByteSlice(content), &secret)
	case KindNameService:
		service := corev1.Service{}
		return jsoniter.Unmarshal(util.StringToByteSlice(content), &service)
	case KindNameStatefulSet:
		statefulSet := appsv1.StatefulSet{}
		return jsoniter.Unmarshal(util.StringToByteSlice(content), &statefulSet)
	case KindNameEndpoints:
		endpoints := corev1.Endpoints{}
		return jsoniter.Unmarshal(util.StringToByteSlice(content), &endpoints)
	case KindNameStorageClass:
		storageClass := storagev1.StorageClass{}
		return jsoniter.Unmarshal(util.StringToByteSlice(content), &storageClass)
	case KindNameRole:
		role := rbacv1.Role{}
		return jsoniter.Unmarshal(util.StringToByteSlice(content), &role)
	case KindNameRoleBinding:
		roleBinding := rbacv1.RoleBinding{}
		return jsoniter.Unmarshal(util.StringToByteSlice(content), &roleBinding)
	case KindNameClusterRole:
		clusterRole := rbacv1.ClusterRole{}
		return jsoniter.Unmarshal(util.StringToByteSlice(content), &clusterRole)
	case KindNameClusterRoleBinding:
		clusterRoleBinding := rbacv1.ClusterRoleBinding{}
		return jsoniter.Unmarshal(util.StringToByteSlice(content), &clusterRoleBinding)
	case KindNameServiceAccount:
		serviceAccount := corev1.ServiceAccount{}
		return jsoniter.Unmarshal(util.StringToByteSlice(content), &serviceAccount)
	default:
		return fmt.Errorf("unsupported resource template kind: %s", templateKind)
	}
}
