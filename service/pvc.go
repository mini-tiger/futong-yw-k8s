package service

import (
	"context"
	"ftk8s/base/cfg"
	"ftk8s/ksc"
	"ftk8s/model"
	"ftk8s/util"
	"sort"

	jsoniter "github.com/json-iterator/go"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreatePVC(reqObj *model.ReqCreatePVC) (*corev1.PersistentVolumeClaim, error) {
	var pvc corev1.PersistentVolumeClaim
	err := jsoniter.Unmarshal([]byte(reqObj.PVCData), &pvc)
	if err != nil {
		cfg.Mlog.Error("failed to jsoniter.Unmarshal, error message: ", err.Error())
		return nil, err
	}

	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return nil, err
	}

	resultPVC, err := k8sCli.CoreV1().PersistentVolumeClaims(reqObj.NamespaceName).Create(context.Background(), &pvc, metav1.CreateOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli Create, error message: ", err.Error())
		return nil, err
	}

	return resultPVC, err
}

func UpdatePVC(reqObj *model.ReqUpdatePVC) (*corev1.PersistentVolumeClaim, error) {
	var pvc corev1.PersistentVolumeClaim
	err := jsoniter.Unmarshal([]byte(reqObj.PVCData), &pvc)
	if err != nil {
		cfg.Mlog.Error("failed to jsoniter.Unmarshal, error message: ", err.Error())
		return nil, err
	}

	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return nil, err
	}

	resultPVC, err := k8sCli.CoreV1().PersistentVolumeClaims(reqObj.NamespaceName).Update(context.Background(), &pvc, metav1.UpdateOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli Update, error message: ", err.Error())
		return nil, err
	}

	return resultPVC, err
}

func ReadPVC(reqObj *model.ReqReadPVC) (*corev1.PersistentVolumeClaim, error) {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return nil, err
	}

	pvc, err := k8sCli.CoreV1().PersistentVolumeClaims(reqObj.NamespaceName).Get(context.Background(), reqObj.PVCName, metav1.GetOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli Get, error message: ", err.Error())
		return nil, err
	}

	return pvc, err
}

func DeletePVC(reqObj *model.ReqDeletePVC) error {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return err
	}

	deletionPropagation := metav1.DeletePropagationBackground
	err = k8sCli.CoreV1().PersistentVolumeClaims(reqObj.NamespaceName).
		Delete(context.Background(), reqObj.PVCName,
			metav1.DeleteOptions{PropagationPolicy: &deletionPropagation})
	return err
}

func ListPVC(reqObj *model.ReqListPVC) (extra util.Extra, data []corev1.PersistentVolumeClaim, err error) {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return extra, data, err
	}

	pvcs, err := k8sCli.CoreV1().PersistentVolumeClaims(reqObj.NamespaceName).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli List, error message: ", err.Error())
		return extra, data, err
	}
	sort.SliceStable(pvcs.Items, func(i, j int) bool {
		return pvcs.Items[i].CreationTimestamp.Unix() > pvcs.Items[j].CreationTimestamp.Unix()
	})

	dataCount := len(pvcs.Items)

	reqObj.PageNum, reqObj.PageSize, reqObj.SkipNum, reqObj.SortField, reqObj.SortOrder, reqObj.SortOrderIsDesc = util.GetPageInfoAndSortWay(reqObj.PageNum, reqObj.PageSize, reqObj.SortField, reqObj.SortOrder)
	data = util.SlicePageForPVC(pvcs.Items, reqObj.SkipNum, reqObj.PageSize)

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
