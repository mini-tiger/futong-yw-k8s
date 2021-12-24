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

func CreateDeployment(reqObj *model.ReqCreateDeployment) (*appsv1.Deployment, error) {
	var deployment appsv1.Deployment
	err := jsoniter.Unmarshal([]byte(reqObj.DeploymentData), &deployment)
	if err != nil {
		cfg.Mlog.Error("failed to jsoniter.Unmarshal, error message: ", err.Error())
		return nil, err
	}

	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return nil, err
	}

	resultDeployment, err := k8sCli.AppsV1().Deployments(reqObj.NamespaceName).Create(context.Background(), &deployment, metav1.CreateOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli Create, error message: ", err.Error())
		return nil, err
	}

	return resultDeployment, err
}

func CreateDeploymentByUI(reqObj *model.ReqCreateDeploymentByUI) (deployment model.ReqCreateDeployment, err error) {
	deployment.ClusterID = reqObj.ClusterID
	deployment.NamespaceName = reqObj.NamespaceName
	deployment.DeploymentData = ""

	return deployment, err
}

func UpdateDeployment(reqObj *model.ReqUpdateDeployment) (*appsv1.Deployment, error) {
	var deployment appsv1.Deployment
	err := jsoniter.Unmarshal([]byte(reqObj.DeploymentData), &deployment)
	if err != nil {
		cfg.Mlog.Error("failed to jsoniter.Unmarshal, error message: ", err.Error())
		return nil, err
	}

	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return nil, err
	}

	resultDeployment, err := k8sCli.AppsV1().Deployments(reqObj.NamespaceName).Update(context.Background(), &deployment, metav1.UpdateOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli Update, error message: ", err.Error())
		return nil, err
	}

	return resultDeployment, err
}

func UpdateDeploymentByUI(reqObj *model.ReqUpdateDeploymentByUI) (deployment model.ReqUpdateDeployment, err error) {
	deployment.ClusterID = reqObj.ClusterID
	deployment.NamespaceName = reqObj.NamespaceName
	deployment.DeploymentData = ""

	return deployment, err
}

func ReadDeployment(reqObj *model.ReqReadDeployment) (*appsv1.Deployment, error) {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		return nil, err
	}

	deployment, err := k8sCli.AppsV1().Deployments(reqObj.NamespaceName).Get(context.Background(), reqObj.DeploymentName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return deployment, err
}

func DeleteDeployment(reqObj *model.ReqDeleteDeployment) error {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return err
	}

	deletionPropagation := metav1.DeletePropagationBackground
	err = k8sCli.AppsV1().Deployments(reqObj.NamespaceName).
		Delete(context.Background(), reqObj.DeploymentName,
			metav1.DeleteOptions{PropagationPolicy: &deletionPropagation})
	return err
}

func ListDeployment(reqObj *model.ReqListDeployment) (extra util.Extra, data []*appsv1.Deployment, err error) {
	ksCache, err := ksc.GetKsCache(reqObj.ClusterID)
	if err != nil {
		return extra, data, err
	}

	deployments, err := ksCache.DeploymentLister().Deployments(reqObj.NamespaceName).List(labels.Everything())
	if err != nil {
		return extra, data, err
	}
	sort.SliceStable(deployments, func(i, j int) bool {
		return deployments[i].CreationTimestamp.Unix() > deployments[j].CreationTimestamp.Unix()
	})

	dataCount := len(deployments)

	reqObj.PageNum, reqObj.PageSize, reqObj.SkipNum, reqObj.SortField, reqObj.SortOrder, reqObj.SortOrderIsDesc = util.GetPageInfoAndSortWay(reqObj.PageNum, reqObj.PageSize, reqObj.SortField, reqObj.SortOrder)
	data = util.SlicePageForDeployment(deployments, reqObj.SkipNum, reqObj.PageSize)

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

func ReadDeploymentHistory(reqObj *model.ReqReadDeploymentHistory) (data *appsv1.ReplicaSetList, err error) {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return data, err
	}

	deployment, err := k8sCli.AppsV1().Deployments(reqObj.NamespaceName).Get(context.Background(), reqObj.DeploymentName, metav1.GetOptions{})
	if err != nil {
		return data, err
	}

	labelsMap := deployment.GetLabels()
	labelSelector := metav1.LabelSelector{MatchLabels: labelsMap}
	data, err = k8sCli.AppsV1().ReplicaSets(reqObj.NamespaceName).List(context.Background(), metav1.ListOptions{
		LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
	})
	if err != nil {
		return data, err
	}

	return data, err
}
