package service

import (
	"context"
	"sort"

	"ftk8s/base/cfg"
	"ftk8s/ksc"
	"ftk8s/model"
	"ftk8s/util"

	jsoniter "github.com/json-iterator/go"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateIngress(reqObj *model.ReqCreateIngress) (*extensionsv1beta1.Ingress, error) {
	var ingress extensionsv1beta1.Ingress
	err := jsoniter.Unmarshal([]byte(reqObj.IngressData), &ingress)
	if err != nil {
		cfg.Mlog.Error("failed to jsoniter.Unmarshal, error message: ", err.Error())
		return nil, err
	}

	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return nil, err
	}

	resultIngress, err := k8sCli.ExtensionsV1beta1().Ingresses(reqObj.NamespaceName).Create(context.Background(), &ingress, metav1.CreateOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli Create, error message: ", err.Error())
		return nil, err
	}

	return resultIngress, err
}

func UpdateIngress(reqObj *model.ReqUpdateIngress) (*extensionsv1beta1.Ingress, error) {
	var ingress extensionsv1beta1.Ingress
	err := jsoniter.Unmarshal([]byte(reqObj.IngressData), &ingress)
	if err != nil {
		cfg.Mlog.Error("failed to jsoniter.Unmarshal, error message: ", err.Error())
		return nil, err
	}

	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return nil, err
	}

	resultIngress, err := k8sCli.ExtensionsV1beta1().Ingresses(reqObj.NamespaceName).Update(context.Background(), &ingress, metav1.UpdateOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli Update, error message: ", err.Error())
		return nil, err
	}

	return resultIngress, err
}

func ReadIngress(reqObj *model.ReqReadIngress) (*extensionsv1beta1.Ingress, error) {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return nil, err
	}

	ingress, err := k8sCli.ExtensionsV1beta1().Ingresses(reqObj.NamespaceName).Get(context.Background(), reqObj.IngressName, metav1.GetOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli Get, error message: ", err.Error())
		return nil, err
	}

	return ingress, err
}

func DeleteIngress(reqObj *model.ReqDeleteIngress) error {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return err
	}

	deletionPropagation := metav1.DeletePropagationBackground
	err = k8sCli.ExtensionsV1beta1().Ingresses(reqObj.NamespaceName).
		Delete(context.Background(), reqObj.IngressName,
			metav1.DeleteOptions{PropagationPolicy: &deletionPropagation})
	return err
}

func ListIngress(reqObj *model.ReqListIngress) (extra util.Extra, data []extensionsv1beta1.Ingress, err error) {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return extra, data, err
	}

	ingresses := new(extensionsv1beta1.IngressList)
	if len(reqObj.ServiceNameList) == 0 {
		ingresses, err = k8sCli.ExtensionsV1beta1().Ingresses(reqObj.NamespaceName).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			cfg.Mlog.Error("failed to k8sCli List, error message: ", err.Error())
			return extra, data, err
		}
	} else {
		tempIngresses, err := k8sCli.ExtensionsV1beta1().Ingresses(reqObj.NamespaceName).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			cfg.Mlog.Error("failed to k8sCli List, error message: ", err.Error())
			return extra, data, err
		}

		for _, serviceName := range reqObj.ServiceNameList {
			for _, ingress := range tempIngresses.Items {
				for _, rule := range ingress.Spec.Rules {
					for _, path := range rule.HTTP.Paths {
						if path.Backend.ServiceName == serviceName {
							ingresses.Items = append(ingresses.Items, ingress)
						}
					}
				}
			}
		}
	}

	sort.SliceStable(ingresses.Items, func(i, j int) bool {
		return ingresses.Items[i].CreationTimestamp.Unix() > ingresses.Items[j].CreationTimestamp.Unix()
	})

	dataCount := len(ingresses.Items)

	reqObj.PageNum, reqObj.PageSize, reqObj.SkipNum, reqObj.SortField, reqObj.SortOrder, reqObj.SortOrderIsDesc = util.GetPageInfoAndSortWay(reqObj.PageNum, reqObj.PageSize, reqObj.SortField, reqObj.SortOrder)
	data = util.SlicePageForIngress(ingresses.Items, reqObj.SkipNum, reqObj.PageSize)

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
