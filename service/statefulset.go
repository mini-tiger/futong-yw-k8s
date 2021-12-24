package service

import (
	"context"
	"sort"

	"ftk8s/base/cfg"
	"ftk8s/ksc"
	"ftk8s/model"
	"ftk8s/util"

	jsoniter "github.com/json-iterator/go"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

func CreateStatefulSet(reqObj *model.ReqCreateStatefulSet) (*appsv1.StatefulSet, error) {
	var statefulset appsv1.StatefulSet
	err := jsoniter.Unmarshal([]byte(reqObj.StatefulSetData), &statefulset)
	if err != nil {
		cfg.Mlog.Error("failed to jsoniter.Unmarshal, error message: ", err.Error())
		return nil, err
	}

	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return nil, err
	}

	resultStatefulSet, err := k8sCli.AppsV1().StatefulSets(reqObj.NamespaceName).Create(context.Background(), &statefulset, metav1.CreateOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli Create, error message: ", err.Error())
		return nil, err
	}

	return resultStatefulSet, err
}

func UpdateStatefulSet(reqObj *model.ReqUpdateStatefulSet) (*appsv1.StatefulSet, error) {
	var statefulset appsv1.StatefulSet
	err := jsoniter.Unmarshal([]byte(reqObj.StatefulSetData), &statefulset)
	if err != nil {
		cfg.Mlog.Error("failed to jsoniter.Unmarshal, error message: ", err.Error())
		return nil, err
	}

	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return nil, err
	}

	resultStatefulSet, err := k8sCli.AppsV1().StatefulSets(reqObj.NamespaceName).Update(context.Background(), &statefulset, metav1.UpdateOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli Update, error message: ", err.Error())
		return nil, err
	}

	return resultStatefulSet, err
}

func ReadStatefulSet(reqObj *model.ReqReadStatefulSet) (*appsv1.StatefulSet, error) {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return nil, err
	}

	statefulset, err := k8sCli.AppsV1().StatefulSets(reqObj.NamespaceName).Get(context.Background(), reqObj.StatefulSetName, metav1.GetOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli Get, error message: ", err.Error())
		return nil, err
	}

	return statefulset, err
}

func DeleteStatefulSet(reqObj *model.ReqDeleteStatefulSet) error {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return err
	}

	deletionPropagation := metav1.DeletePropagationBackground
	err = k8sCli.AppsV1().StatefulSets(reqObj.NamespaceName).
		Delete(context.Background(), reqObj.StatefulSetName,
			metav1.DeleteOptions{PropagationPolicy: &deletionPropagation})
	return err
}

func ListStatefulSet(reqObj *model.ReqListStatefulSet) (extra util.Extra, data []appsv1.StatefulSet, err error) {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return extra, data, err
	}

	statefulsets, err := k8sCli.AppsV1().StatefulSets(reqObj.NamespaceName).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli List, error message: ", err.Error())
		return extra, data, err
	}
	sort.SliceStable(statefulsets.Items, func(i, j int) bool {
		return statefulsets.Items[i].CreationTimestamp.Unix() > statefulsets.Items[j].CreationTimestamp.Unix()
	})

	dataCount := len(statefulsets.Items)

	reqObj.PageNum, reqObj.PageSize, reqObj.SkipNum, reqObj.SortField, reqObj.SortOrder, reqObj.SortOrderIsDesc = util.GetPageInfoAndSortWay(reqObj.PageNum, reqObj.PageSize, reqObj.SortField, reqObj.SortOrder)
	data = util.SlicePageForStatefulSet(statefulsets.Items, reqObj.SkipNum, reqObj.PageSize)

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

func ReadStatefulSetHistory(reqObj *model.ReqReadStatefulSetHistory) (data *appsv1.ReplicaSetList, err error) {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return data, err
	}

	statefulset, err := k8sCli.AppsV1().StatefulSets(reqObj.NamespaceName).Get(context.Background(), reqObj.StatefulSetName, metav1.GetOptions{})
	if err != nil {
		return data, err
	}

	labelsMap := statefulset.GetLabels()
	labelSelector := metav1.LabelSelector{MatchLabels: labelsMap}
	data, err = k8sCli.AppsV1().ReplicaSets(reqObj.NamespaceName).List(context.Background(), metav1.ListOptions{
		LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
	})
	if err != nil {
		return data, err
	}

	return data, err
}
