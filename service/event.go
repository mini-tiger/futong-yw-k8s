package service

import (
	"context"

	"ftk8s/base/cfg"
	"ftk8s/ksc"
	"ftk8s/model"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListEvent(reqObj *model.ReqListEvent) (data *corev1.EventList, err error) {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return data, err
	}

	data, err = k8sCli.CoreV1().Events(reqObj.NamespaceName).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return data, err
	}

	return data, err
}

func ListEventByResource(reqObj *model.ReqListEventByResource) (data *corev1.EventList, err error) {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return data, err
	}

	var involvedObjectName string
	var involvedObjectNamespace string
	var involvedObjectKind string
	var involvedObjectUID string
	switch reqObj.ResourceKind {
	case ksc.KindNameDeployment:
		deployment, err := k8sCli.AppsV1().Deployments(reqObj.NamespaceName).Get(context.Background(), reqObj.ResourceName, metav1.GetOptions{})
		if err != nil {
			return data, err
		}
		involvedObjectName = deployment.Name
		involvedObjectNamespace = deployment.Namespace
		involvedObjectKind = ksc.KindNameDeployment
		involvedObjectUID = string(deployment.UID)

	case ksc.KindNameStatefulSet:
		statefulset, err := k8sCli.AppsV1().StatefulSets(reqObj.NamespaceName).Get(context.Background(), reqObj.ResourceName, metav1.GetOptions{})
		if err != nil {
			cfg.Mlog.Error("failed to k8sCli Get, error message: ", err.Error())
			return data, err
		}
		involvedObjectName = statefulset.Name
		involvedObjectNamespace = statefulset.Namespace
		involvedObjectKind = ksc.KindNameStatefulSet
		involvedObjectUID = string(statefulset.UID)

	case ksc.KindNameDaemonSet:
		daemonset, err := k8sCli.AppsV1().DaemonSets(reqObj.NamespaceName).Get(context.Background(), reqObj.ResourceName, metav1.GetOptions{})
		if err != nil {
			cfg.Mlog.Error("failed to k8sCli Get, error message: ", err.Error())
			return data, err
		}
		involvedObjectName = daemonset.Name
		involvedObjectNamespace = daemonset.Namespace
		involvedObjectKind = ksc.KindNameDaemonSet
		involvedObjectUID = string(daemonset.UID)

	case ksc.KindNameJob:
		job, err := k8sCli.BatchV1().Jobs(reqObj.NamespaceName).Get(context.Background(), reqObj.ResourceName, metav1.GetOptions{})
		if err != nil {
			cfg.Mlog.Error("failed to k8sCli Get, error message: ", err.Error())
			return data, err
		}
		involvedObjectName = job.Name
		involvedObjectNamespace = job.Namespace
		involvedObjectKind = ksc.KindNameJob
		involvedObjectUID = string(job.UID)

	case ksc.KindNameCronJob:
		cronjob, err := k8sCli.BatchV1beta1().CronJobs(reqObj.NamespaceName).Get(context.Background(), reqObj.ResourceName, metav1.GetOptions{})
		if err != nil {
			cfg.Mlog.Error("failed to k8sCli Get, error message: ", err.Error())
			return data, err
		}
		involvedObjectName = cronjob.Name
		involvedObjectNamespace = cronjob.Namespace
		involvedObjectKind = ksc.KindNameCronJob
		involvedObjectUID = string(cronjob.UID)

	default:
		cfg.Mlog.Error("failed to list events, unsupported resource_kind")
		return data, err
	}

	fieldSelector := k8sCli.CoreV1().Events(reqObj.NamespaceName).GetFieldSelector(
		&involvedObjectName, &involvedObjectNamespace, &involvedObjectKind, &involvedObjectUID)

	data, err = k8sCli.CoreV1().Events(reqObj.NamespaceName).List(context.Background(), metav1.ListOptions{
		FieldSelector: fieldSelector.String(),
	})

	if err != nil {
		return data, err
	}

	return data, err
}
