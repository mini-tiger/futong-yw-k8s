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

func CreateConfigMap(reqObj *model.ReqCreateConfigMap) (*corev1.ConfigMap, error) {
	var configMap corev1.ConfigMap
	err := jsoniter.Unmarshal([]byte(reqObj.ConfigMapData), &configMap)
	if err != nil {
		cfg.Mlog.Error("failed to jsoniter.Unmarshal, error message: ", err.Error())
		return nil, err
	}

	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return nil, err
	}

	resultConfigMap, err := k8sCli.CoreV1().ConfigMaps(reqObj.NamespaceName).Create(context.Background(), &configMap, metav1.CreateOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli Create, error message: ", err.Error())
		return nil, err
	}

	return resultConfigMap, err
}

func UpdateConfigMap(reqObj *model.ReqUpdateConfigMap) (*corev1.ConfigMap, error) {
	var configMap corev1.ConfigMap
	err := jsoniter.Unmarshal([]byte(reqObj.ConfigMapData), &configMap)
	if err != nil {
		cfg.Mlog.Error("failed to jsoniter.Unmarshal, error message: ", err.Error())
		return nil, err
	}

	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return nil, err
	}

	resultConfigMap, err := k8sCli.CoreV1().ConfigMaps(reqObj.NamespaceName).Update(context.Background(), &configMap, metav1.UpdateOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli Update, error message: ", err.Error())
		return nil, err
	}

	return resultConfigMap, err
}

func ReadConfigMap(reqObj *model.ReqReadConfigMap) (*corev1.ConfigMap, error) {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return nil, err
	}

	configMap, err := k8sCli.CoreV1().ConfigMaps(reqObj.NamespaceName).Get(context.Background(), reqObj.ConfigMapName, metav1.GetOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli Get, error message: ", err.Error())
		return nil, err
	}

	return configMap, err
}

func DeleteConfigMap(reqObj *model.ReqDeleteConfigMap) error {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return err
	}

	deletionPropagation := metav1.DeletePropagationBackground
	err = k8sCli.CoreV1().ConfigMaps(reqObj.NamespaceName).
		Delete(context.Background(), reqObj.ConfigMapName,
			metav1.DeleteOptions{PropagationPolicy: &deletionPropagation})
	return err
}

func ListConfigMap(reqObj *model.ReqListConfigMap) (extra util.Extra, data []corev1.ConfigMap, err error) {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return extra, data, err
	}

	configMaps, err := k8sCli.CoreV1().ConfigMaps(reqObj.NamespaceName).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli List, error message: ", err.Error())
		return extra, data, err
	}
	sort.SliceStable(configMaps.Items, func(i, j int) bool {
		return configMaps.Items[i].CreationTimestamp.Unix() > configMaps.Items[j].CreationTimestamp.Unix()
	})

	dataCount := len(configMaps.Items)

	reqObj.PageNum, reqObj.PageSize, reqObj.SkipNum, reqObj.SortField, reqObj.SortOrder, reqObj.SortOrderIsDesc = util.GetPageInfoAndSortWay(reqObj.PageNum, reqObj.PageSize, reqObj.SortField, reqObj.SortOrder)
	data = util.SlicePageForConfigMap(configMaps.Items, reqObj.SkipNum, reqObj.PageSize)

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
