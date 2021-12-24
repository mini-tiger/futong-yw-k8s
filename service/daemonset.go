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
)

func CreateDaemonSet(reqObj *model.ReqCreateDaemonSet) (*appsv1.DaemonSet, error) {
	var daemonset appsv1.DaemonSet
	err := jsoniter.Unmarshal([]byte(reqObj.DaemonSetData), &daemonset)
	if err != nil {
		cfg.Mlog.Error("failed to jsoniter.Unmarshal, error message: ", err.Error())
		return nil, err
	}

	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return nil, err
	}

	resultDaemonSet, err := k8sCli.AppsV1().DaemonSets(reqObj.NamespaceName).Create(context.Background(), &daemonset, metav1.CreateOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli Create, error message: ", err.Error())
		return nil, err
	}

	return resultDaemonSet, err
}

func UpdateDaemonSet(reqObj *model.ReqUpdateDaemonSet) (*appsv1.DaemonSet, error) {
	var daemonset appsv1.DaemonSet
	err := jsoniter.Unmarshal([]byte(reqObj.DaemonSetData), &daemonset)
	if err != nil {
		cfg.Mlog.Error("failed to jsoniter.Unmarshal, error message: ", err.Error())
		return nil, err
	}

	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return nil, err
	}

	resultDaemonSet, err := k8sCli.AppsV1().DaemonSets(reqObj.NamespaceName).Update(context.Background(), &daemonset, metav1.UpdateOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli Update, error message: ", err.Error())
		return nil, err
	}

	return resultDaemonSet, err
}

func ReadDaemonSet(reqObj *model.ReqReadDaemonSet) (*appsv1.DaemonSet, error) {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return nil, err
	}

	daemonset, err := k8sCli.AppsV1().DaemonSets(reqObj.NamespaceName).Get(context.Background(), reqObj.DaemonSetName, metav1.GetOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli Get, error message: ", err.Error())
		return nil, err
	}

	return daemonset, err
}

func DeleteDaemonSet(reqObj *model.ReqDeleteDaemonSet) error {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return err
	}

	deletionPropagation := metav1.DeletePropagationBackground
	err = k8sCli.AppsV1().DaemonSets(reqObj.NamespaceName).
		Delete(context.Background(), reqObj.DaemonSetName,
			metav1.DeleteOptions{PropagationPolicy: &deletionPropagation})
	return err
}

func ListDaemonSet(reqObj *model.ReqListDaemonSet) (extra util.Extra, data []appsv1.DaemonSet, err error) {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return extra, data, err
	}

	daemonsets, err := k8sCli.AppsV1().DaemonSets(reqObj.NamespaceName).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli List, error message: ", err.Error())
		return extra, data, err
	}
	sort.SliceStable(daemonsets.Items, func(i, j int) bool {
		return daemonsets.Items[i].CreationTimestamp.Unix() > daemonsets.Items[j].CreationTimestamp.Unix()
	})

	dataCount := len(daemonsets.Items)

	reqObj.PageNum, reqObj.PageSize, reqObj.SkipNum, reqObj.SortField, reqObj.SortOrder, reqObj.SortOrderIsDesc = util.GetPageInfoAndSortWay(reqObj.PageNum, reqObj.PageSize, reqObj.SortField, reqObj.SortOrder)
	data = util.SlicePageForDaemonSet(daemonsets.Items, reqObj.SkipNum, reqObj.PageSize)

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
