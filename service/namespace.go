package service

import (
	"context"
	"sort"

	"ftk8s/base/cfg"
	"ftk8s/ksc"
	"ftk8s/model"
	"ftk8s/util"

	jsoniter "github.com/json-iterator/go"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateNamespace(reqObj *model.ReqCreateNamespace) (*corev1.Namespace, error) {
	var namespace corev1.Namespace
	err := jsoniter.Unmarshal([]byte(reqObj.NamespaceData), &namespace)
	if err != nil {
		cfg.Mlog.Error("failed to jsoniter.Unmarshal, error message: ", err.Error())
		return nil, err
	}

	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return nil, err
	}

	resultNamespace, err := k8sCli.CoreV1().Namespaces().Create(context.Background(), &namespace, metav1.CreateOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli Create, error message: ", err.Error())
		return nil, err
	}

	return resultNamespace, err
}

func UpdateNamespace(reqObj *model.ReqUpdateNamespace) (*corev1.Namespace, error) {
	var namespace corev1.Namespace
	err := jsoniter.Unmarshal([]byte(reqObj.NamespaceData), &namespace)
	if err != nil {
		cfg.Mlog.Error("failed to jsoniter.Unmarshal, error message: ", err.Error())
		return nil, err
	}

	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return nil, err
	}

	resultNamespace, err := k8sCli.CoreV1().Namespaces().Update(context.Background(), &namespace, metav1.UpdateOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli Update, error message: ", err.Error())
		return nil, err
	}

	return resultNamespace, err
}

func ReadNamespace(reqObj *model.ReqReadNamespace) (*corev1.Namespace, error) {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return nil, err
	}

	namespace, err := k8sCli.CoreV1().Namespaces().Get(context.Background(), reqObj.NamespaceName, metav1.GetOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli Get, error message: ", err.Error())
		return nil, err
	}

	return namespace, err
}

func DeleteNamespace(reqObj *model.ReqDeleteNamespace) error {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return err
	}

	deletionPropagation := metav1.DeletePropagationBackground
	err = k8sCli.CoreV1().Namespaces().
		Delete(context.Background(), reqObj.NamespaceName,
			metav1.DeleteOptions{PropagationPolicy: &deletionPropagation})
	return err
}

func ListNamespace(reqObj *model.ReqListNamespace) (extra util.Extra, data []corev1.Namespace, err error) {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return extra, data, err
	}

	namespaces, err := k8sCli.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli List, error message: ", err.Error())
		return extra, data, err
	}
	OriSlice := namespaces.Items

	dataCount := len(OriSlice)

	reqObj.PageNum, reqObj.PageSize, reqObj.SkipNum, reqObj.SortField, reqObj.SortOrder, reqObj.SortOrderIsDesc = util.GetPageInfoAndSortWay(reqObj.PageNum, reqObj.PageSize, reqObj.SortField, reqObj.SortOrder)
	data = util.SlicePageForNamespace(OriSlice, reqObj.SkipNum, reqObj.PageSize)

	sort.SliceStable(data, func(i, j int) bool {
		return data[i].CreationTimestamp.Unix() > data[j].CreationTimestamp.Unix()
	})

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
