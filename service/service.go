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
	"k8s.io/apimachinery/pkg/labels"
)

func CreateService(reqObj *model.ReqCreateService) (*corev1.Service, error) {
	var service corev1.Service
	err := jsoniter.Unmarshal([]byte(reqObj.ServiceData), &service)
	if err != nil {
		cfg.Mlog.Error("failed to jsoniter.Unmarshal, error message: ", err.Error())
		return nil, err
	}

	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return nil, err
	}

	resultService, err := k8sCli.CoreV1().Services(reqObj.NamespaceName).Create(context.Background(), &service, metav1.CreateOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli Create, error message: ", err.Error())
		return nil, err
	}

	return resultService, err
}

func UpdateService(reqObj *model.ReqUpdateService) (*corev1.Service, error) {
	var service corev1.Service
	err := jsoniter.Unmarshal([]byte(reqObj.ServiceData), &service)
	if err != nil {
		cfg.Mlog.Error("failed to jsoniter.Unmarshal, error message: ", err.Error())
		return nil, err
	}

	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return nil, err
	}

	resultService, err := k8sCli.CoreV1().Services(reqObj.NamespaceName).Update(context.Background(), &service, metav1.UpdateOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli Update, error message: ", err.Error())
		return nil, err
	}

	return resultService, err
}

func ReadService(reqObj *model.ReqReadService) (*corev1.Service, error) {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return nil, err
	}

	service, err := k8sCli.CoreV1().Services(reqObj.NamespaceName).Get(context.Background(), reqObj.ServiceName, metav1.GetOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli Get, error message: ", err.Error())
		return nil, err
	}

	return service, err
}

func DeleteService(reqObj *model.ReqDeleteService) error {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return err
	}

	deletionPropagation := metav1.DeletePropagationBackground
	err = k8sCli.CoreV1().Services(reqObj.NamespaceName).
		Delete(context.Background(), reqObj.ServiceName,
			metav1.DeleteOptions{PropagationPolicy: &deletionPropagation})
	return err
}

func ListService(reqObj *model.ReqListService) (extra util.Extra, data []corev1.Service, err error) {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return extra, data, err
	}

	var services *corev1.ServiceList
	var labelsMap map[string]string

	if reqObj.ResourceKind == "" {
		services, err = k8sCli.CoreV1().Services(reqObj.NamespaceName).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			cfg.Mlog.Error("failed to k8sCli List, error message: ", err.Error())
			return extra, data, err
		}
	} else {
		switch reqObj.ResourceKind {
		case ksc.KindNameDeployment:
			deployment, err := k8sCli.AppsV1().Deployments(reqObj.NamespaceName).Get(context.Background(), reqObj.ResourceName, metav1.GetOptions{})
			if err != nil {
				return extra, data, err
			}
			labelsMap = deployment.Spec.Template.GetLabels()

		case ksc.KindNameStatefulSet:
			statefulset, err := k8sCli.AppsV1().StatefulSets(reqObj.NamespaceName).Get(context.Background(), reqObj.ResourceName, metav1.GetOptions{})
			if err != nil {
				cfg.Mlog.Error("failed to k8sCli Get, error message: ", err.Error())
				return extra, data, err
			}
			labelsMap = statefulset.Spec.Template.GetLabels()

		case ksc.KindNameDaemonSet:
			daemonset, err := k8sCli.AppsV1().DaemonSets(reqObj.NamespaceName).Get(context.Background(), reqObj.ResourceName, metav1.GetOptions{})
			if err != nil {
				cfg.Mlog.Error("failed to k8sCli Get, error message: ", err.Error())
				return extra, data, err
			}
			labelsMap = daemonset.Spec.Template.GetLabels()

		default:
			cfg.Mlog.Error("failed to ListPodByResource, unsupported resource_kind: %s", reqObj.ResourceKind)
			return extra, data, err
		}

		labelSelector := metav1.LabelSelector{MatchLabels: labelsMap}
		services, err = k8sCli.CoreV1().Services(reqObj.NamespaceName).List(context.Background(), metav1.ListOptions{
			LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
		})
		if err != nil {
			cfg.Mlog.Error("failed to k8sCli List, error message: ", err.Error())
			return extra, data, err
		}
	}

	sort.SliceStable(services.Items, func(i, j int) bool {
		return services.Items[i].CreationTimestamp.Unix() > services.Items[j].CreationTimestamp.Unix()
	})

	dataCount := len(services.Items)

	reqObj.PageNum, reqObj.PageSize, reqObj.SkipNum, reqObj.SortField, reqObj.SortOrder, reqObj.SortOrderIsDesc = util.GetPageInfoAndSortWay(reqObj.PageNum, reqObj.PageSize, reqObj.SortField, reqObj.SortOrder)
	data = util.SlicePageForService(services.Items, reqObj.SkipNum, reqObj.PageSize)

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
