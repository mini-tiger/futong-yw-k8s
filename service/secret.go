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

func CreateSecret(reqObj *model.ReqCreateSecret) (*corev1.Secret, error) {
	var secret corev1.Secret
	err := jsoniter.Unmarshal([]byte(reqObj.SecretData), &secret)
	if err != nil {
		cfg.Mlog.Error("failed to jsoniter.Unmarshal, error message: ", err.Error())
		return nil, err
	}

	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return nil, err
	}

	resultSecret, err := k8sCli.CoreV1().Secrets(reqObj.NamespaceName).Create(context.Background(), &secret, metav1.CreateOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli Create, error message: ", err.Error())
		return nil, err
	}

	return resultSecret, err
}

func UpdateSecret(reqObj *model.ReqUpdateSecret) (*corev1.Secret, error) {
	var secret corev1.Secret
	err := jsoniter.Unmarshal([]byte(reqObj.SecretData), &secret)
	if err != nil {
		cfg.Mlog.Error("failed to jsoniter.Unmarshal, error message: ", err.Error())
		return nil, err
	}

	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return nil, err
	}

	resultSecret, err := k8sCli.CoreV1().Secrets(reqObj.NamespaceName).Update(context.Background(), &secret, metav1.UpdateOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli Update, error message: ", err.Error())
		return nil, err
	}

	return resultSecret, err
}

func ReadSecret(reqObj *model.ReqReadSecret) (*corev1.Secret, error) {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return nil, err
	}

	secret, err := k8sCli.CoreV1().Secrets(reqObj.NamespaceName).Get(context.Background(), reqObj.SecretName, metav1.GetOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli Get, error message: ", err.Error())
		return nil, err
	}

	return secret, err
}

func DeleteSecret(reqObj *model.ReqDeleteSecret) error {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return err
	}

	deletionPropagation := metav1.DeletePropagationBackground
	err = k8sCli.CoreV1().Secrets(reqObj.NamespaceName).
		Delete(context.Background(), reqObj.SecretName,
			metav1.DeleteOptions{PropagationPolicy: &deletionPropagation})
	return err
}

func ListSecret(reqObj *model.ReqListSecret) (extra util.Extra, data []corev1.Secret, err error) {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return extra, data, err
	}

	secrets, err := k8sCli.CoreV1().Secrets(reqObj.NamespaceName).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli List, error message: ", err.Error())
		return extra, data, err
	}
	sort.SliceStable(secrets.Items, func(i, j int) bool {
		return secrets.Items[i].CreationTimestamp.Unix() > secrets.Items[j].CreationTimestamp.Unix()
	})

	dataCount := len(secrets.Items)

	reqObj.PageNum, reqObj.PageSize, reqObj.SkipNum, reqObj.SortField, reqObj.SortOrder, reqObj.SortOrderIsDesc = util.GetPageInfoAndSortWay(reqObj.PageNum, reqObj.PageSize, reqObj.SortField, reqObj.SortOrder)
	data = util.SlicePageForSecret(secrets.Items, reqObj.SkipNum, reqObj.PageSize)

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
