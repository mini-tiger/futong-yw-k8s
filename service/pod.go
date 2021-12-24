package service

import (
	"context"
	"io/ioutil"
	"sort"

	"ftk8s/base/cfg"
	"ftk8s/ksc"
	"ftk8s/model"
	"ftk8s/util"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

func ReadPod(reqObj *model.ReqReadPod) (pod *corev1.Pod, err error) {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return pod, err
	}

	pod, err = k8sCli.CoreV1().Pods(reqObj.NamespaceName).Get(context.Background(), reqObj.PodName, metav1.GetOptions{})
	return pod, err
}

func ListPod(reqObj *model.ReqListPod) (extra util.Extra, data []*corev1.Pod, err error) {
	ksCache, err := ksc.GetKsCache(reqObj.ClusterID)
	if err != nil {
		return extra, data, err
	}

	pods, err := ksCache.PodLister().Pods(reqObj.NamespaceName).List(labels.Everything())
	sort.SliceStable(pods, func(i, j int) bool {
		return pods[i].CreationTimestamp.Unix() > pods[j].CreationTimestamp.Unix()
	})

	dataCount := len(pods)

	reqObj.PageNum, reqObj.PageSize, reqObj.SkipNum, reqObj.SortField, reqObj.SortOrder, reqObj.SortOrderIsDesc = util.GetPageInfoAndSortWay(reqObj.PageNum, reqObj.PageSize, reqObj.SortField, reqObj.SortOrder)
	data = util.SlicePageForPod(pods, reqObj.SkipNum, reqObj.PageSize)

	pageCount := util.GetPageCount(dataCount, reqObj.PageSize)
	extra = util.Extra{
		PageNum:   reqObj.PageNum,
		PageSize:  reqObj.PageSize,
		SortField: reqObj.SortField,
		SortOrder: reqObj.SortOrder,
		DataCount: dataCount,
		PageCount: pageCount,
	}

	return extra, data, err
}

func ListPodByResource(reqObj *model.ReqListPodByResource) (data []*corev1.Pod, err error) {
	ksCache, err := ksc.GetKsCache(reqObj.ClusterID)
	if err != nil {
		return data, err
	}

	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		return nil, err
	}

	var labelsMap map[string]string
	switch reqObj.ResourceKind {
	case ksc.KindNameDeployment:
		deployment, err := k8sCli.AppsV1().Deployments(reqObj.NamespaceName).Get(context.Background(), reqObj.ResourceName, metav1.GetOptions{})
		if err != nil {
			return data, err
		}
		labelsMap = deployment.Spec.Template.GetLabels()

	case ksc.KindNameStatefulSet:
		statefulset, err := k8sCli.AppsV1().StatefulSets(reqObj.NamespaceName).Get(context.Background(), reqObj.ResourceName, metav1.GetOptions{})
		if err != nil {
			cfg.Mlog.Error("failed to k8sCli Get, error message: ", err.Error())
			return data, err
		}
		labelsMap = statefulset.Spec.Template.GetLabels()

	case ksc.KindNameDaemonSet:
		daemonset, err := k8sCli.AppsV1().DaemonSets(reqObj.NamespaceName).Get(context.Background(), reqObj.ResourceName, metav1.GetOptions{})
		if err != nil {
			cfg.Mlog.Error("failed to k8sCli Get, error message: ", err.Error())
			return data, err
		}
		labelsMap = daemonset.Spec.Template.GetLabels()

	case ksc.KindNameJob:
		job, err := k8sCli.BatchV1().Jobs(reqObj.NamespaceName).Get(context.Background(), reqObj.ResourceName, metav1.GetOptions{})
		if err != nil {
			cfg.Mlog.Error("failed to k8sCli Get, error message: ", err.Error())
			return data, err
		}
		labelsMap = job.Spec.Template.GetLabels()

	default:
		cfg.Mlog.Error("failed to ListPodByResource, unsupported resource_kind: %s", reqObj.ResourceKind)
		return data, err
	}

	data, err = ksCache.PodLister().Pods(reqObj.NamespaceName).List(labels.SelectorFromSet(labelsMap))
	sort.SliceStable(data, func(i, j int) bool {
		return data[i].CreationTimestamp.Unix() > data[j].CreationTimestamp.Unix()
	})
	if err != nil {
		return data, err
	}

	return data, err
}

func ReadPodLog(reqObj *model.ReqReadPodLog) (data string, err error) {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return data, err
	}

	readCloser, err := k8sCli.CoreV1().Pods(reqObj.NamespaceName).GetLogs(
		reqObj.PodName,
		&corev1.PodLogOptions{
			Container: reqObj.ContainerName,
			TailLines: &reqObj.LogLineNum,
		}).
		Stream(context.Background())
	if err != nil {
		return data, err
	}
	defer readCloser.Close()

	dataTemp, err := ioutil.ReadAll(readCloser)
	data = string(dataTemp)

	return data, err
}
