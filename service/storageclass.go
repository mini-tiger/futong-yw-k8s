package service

import (
	"context"
	"sort"

	"ftk8s/base/cfg"
	"ftk8s/ksc"
	"ftk8s/model"
	"ftk8s/util"

	jsoniter "github.com/json-iterator/go"
	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateStorageClass(reqObj *model.ReqCreateStorageClass) (*storagev1.StorageClass, error) {
	var storageclass storagev1.StorageClass
	err := jsoniter.Unmarshal([]byte(reqObj.StorageClassData), &storageclass)
	if err != nil {
		cfg.Mlog.Error("failed to jsoniter.Unmarshal, error message: ", err.Error())
		return nil, err
	}

	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return nil, err
	}

	resultStorageClass, err := k8sCli.StorageV1().StorageClasses().Create(context.Background(), &storageclass, metav1.CreateOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli Create, error message: ", err.Error())
		return nil, err
	}

	return resultStorageClass, err
}

func UpdateStorageClass(reqObj *model.ReqUpdateStorageClass) (*storagev1.StorageClass, error) {
	var storageclass storagev1.StorageClass
	err := jsoniter.Unmarshal([]byte(reqObj.StorageClassData), &storageclass)
	if err != nil {
		cfg.Mlog.Error("failed to jsoniter.Unmarshal, error message: ", err.Error())
		return nil, err
	}

	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return nil, err
	}

	resultStorageClass, err := k8sCli.StorageV1().StorageClasses().Update(context.Background(), &storageclass, metav1.UpdateOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli Update, error message: ", err.Error())
		return nil, err
	}

	return resultStorageClass, err
}

func ReadStorageClass(reqObj *model.ReqReadStorageClass) (*storagev1.StorageClass, error) {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return nil, err
	}

	storageclass, err := k8sCli.StorageV1().StorageClasses().Get(context.Background(), reqObj.StorageClassName, metav1.GetOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli Get, error message: ", err.Error())
		return nil, err
	}

	return storageclass, err
}

func DeleteStorageClass(reqObj *model.ReqDeleteStorageClass) error {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return err
	}

	deletionPropagation := metav1.DeletePropagationBackground
	err = k8sCli.StorageV1().StorageClasses().
		Delete(context.Background(), reqObj.StorageClassName,
			metav1.DeleteOptions{PropagationPolicy: &deletionPropagation})
	return err
}

func ListStorageClass(reqObj *model.ReqListStorageClass) (extra util.Extra, data []storagev1.StorageClass, err error) {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return extra, data, err
	}

	storageclasss, err := k8sCli.StorageV1().StorageClasses().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli List, error message: ", err.Error())
		return extra, data, err
	}
	sort.SliceStable(storageclasss.Items, func(i, j int) bool {
		return storageclasss.Items[i].CreationTimestamp.Unix() > storageclasss.Items[j].CreationTimestamp.Unix()
	})

	dataCount := len(storageclasss.Items)

	reqObj.PageNum, reqObj.PageSize, reqObj.SkipNum, reqObj.SortField, reqObj.SortOrder, reqObj.SortOrderIsDesc = util.GetPageInfoAndSortWay(reqObj.PageNum, reqObj.PageSize, reqObj.SortField, reqObj.SortOrder)
	data = util.SlicePageForStorageClass(storageclasss.Items, reqObj.SkipNum, reqObj.PageSize)

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
