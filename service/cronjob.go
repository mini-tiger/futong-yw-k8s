package service

import (
	"context"
	"sort"

	"ftk8s/base/cfg"
	"ftk8s/ksc"
	"ftk8s/model"
	"ftk8s/util"

	jsoniter "github.com/json-iterator/go"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateCronJob(reqObj *model.ReqCreateCronJob) (*batchv1beta1.CronJob, error) {
	var cronjob batchv1beta1.CronJob
	err := jsoniter.Unmarshal([]byte(reqObj.CronJobData), &cronjob)
	if err != nil {
		cfg.Mlog.Error("failed to jsoniter.Unmarshal, error message: ", err.Error())
		return nil, err
	}

	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return nil, err
	}

	resultCronJob, err := k8sCli.BatchV1beta1().CronJobs(reqObj.NamespaceName).Create(context.Background(), &cronjob, metav1.CreateOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli Create, error message: ", err.Error())
		return nil, err
	}

	return resultCronJob, err
}

func UpdateCronJob(reqObj *model.ReqUpdateCronJob) (*batchv1beta1.CronJob, error) {
	var cronjob batchv1beta1.CronJob
	err := jsoniter.Unmarshal([]byte(reqObj.CronJobData), &cronjob)
	if err != nil {
		cfg.Mlog.Error("failed to jsoniter.Unmarshal, error message: ", err.Error())
		return nil, err
	}

	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return nil, err
	}

	resultCronJob, err := k8sCli.BatchV1beta1().CronJobs(reqObj.NamespaceName).Update(context.Background(), &cronjob, metav1.UpdateOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli Update, error message: ", err.Error())
		return nil, err
	}

	return resultCronJob, err
}

func ReadCronJob(reqObj *model.ReqReadCronJob) (*batchv1beta1.CronJob, error) {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return nil, err
	}

	cronjob, err := k8sCli.BatchV1beta1().CronJobs(reqObj.NamespaceName).Get(context.Background(), reqObj.CronJobName, metav1.GetOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli Get, error message: ", err.Error())
		return nil, err
	}

	return cronjob, err
}

func DeleteCronJob(reqObj *model.ReqDeleteCronJob) error {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return err
	}

	deletionPropagation := metav1.DeletePropagationBackground
	err = k8sCli.BatchV1beta1().CronJobs(reqObj.NamespaceName).
		Delete(context.Background(), reqObj.CronJobName,
			metav1.DeleteOptions{PropagationPolicy: &deletionPropagation})
	return err
}

func ListCronJob(reqObj *model.ReqListCronJob) (extra util.Extra, data []batchv1beta1.CronJob, err error) {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return extra, data, err
	}

	cronjobs, err := k8sCli.BatchV1beta1().CronJobs(reqObj.NamespaceName).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli List, error message: ", err.Error())
		return extra, data, err
	}
	sort.SliceStable(cronjobs.Items, func(i, j int) bool {
		return cronjobs.Items[i].CreationTimestamp.Unix() > cronjobs.Items[j].CreationTimestamp.Unix()
	})

	dataCount := len(cronjobs.Items)

	reqObj.PageNum, reqObj.PageSize, reqObj.SkipNum, reqObj.SortField, reqObj.SortOrder, reqObj.SortOrderIsDesc = util.GetPageInfoAndSortWay(reqObj.PageNum, reqObj.PageSize, reqObj.SortField, reqObj.SortOrder)
	data = util.SlicePageForCronJob(cronjobs.Items, reqObj.SkipNum, reqObj.PageSize)

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
