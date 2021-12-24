package service

import (
	"context"
	"sort"

	"ftk8s/base/cfg"
	"ftk8s/ksc"
	"ftk8s/model"
	"ftk8s/util"

	jsoniter "github.com/json-iterator/go"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateJob(reqObj *model.ReqCreateJob) (*batchv1.Job, error) {
	var job batchv1.Job
	err := jsoniter.Unmarshal([]byte(reqObj.JobData), &job)
	if err != nil {
		cfg.Mlog.Error("failed to jsoniter.Unmarshal, error message: ", err.Error())
		return nil, err
	}

	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return nil, err
	}

	resultJob, err := k8sCli.BatchV1().Jobs(reqObj.NamespaceName).Create(context.Background(), &job, metav1.CreateOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli Create, error message: ", err.Error())
		return nil, err
	}

	return resultJob, err
}

func UpdateJob(reqObj *model.ReqUpdateJob) (*batchv1.Job, error) {
	var job batchv1.Job
	err := jsoniter.Unmarshal([]byte(reqObj.JobData), &job)
	if err != nil {
		cfg.Mlog.Error("failed to jsoniter.Unmarshal, error message: ", err.Error())
		return nil, err
	}

	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return nil, err
	}

	resultJob, err := k8sCli.BatchV1().Jobs(reqObj.NamespaceName).Update(context.Background(), &job, metav1.UpdateOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli Update, error message: ", err.Error())
		return nil, err
	}

	return resultJob, err
}

func ReadJob(reqObj *model.ReqReadJob) (*batchv1.Job, error) {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return nil, err
	}

	job, err := k8sCli.BatchV1().Jobs(reqObj.NamespaceName).Get(context.Background(), reqObj.JobName, metav1.GetOptions{})
	if err != nil {
		cfg.Mlog.Error("failed to k8sCli Get, error message: ", err.Error())
		return nil, err
	}

	return job, err
}

func DeleteJob(reqObj *model.ReqDeleteJob) error {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return err
	}

	deletionPropagation := metav1.DeletePropagationBackground
	err = k8sCli.BatchV1().Jobs(reqObj.NamespaceName).
		Delete(context.Background(), reqObj.JobName,
			metav1.DeleteOptions{PropagationPolicy: &deletionPropagation})
	return err
}

func ListJob(reqObj *model.ReqListJob) (extra util.Extra, data []batchv1.Job, err error) {
	k8sCli, err := ksc.GetK8sClientSet(reqObj.ClusterID)
	if err != nil {
		cfg.Mlog.Error("failed to GetK8sClientSet, error message: ", err.Error())
		return extra, data, err
	}

	jobs := new(batchv1.JobList)
	if reqObj.CronJobName == "" {
		jobs, err = k8sCli.BatchV1().Jobs(reqObj.NamespaceName).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			cfg.Mlog.Error("failed to k8sCli List, error message: ", err.Error())
			return extra, data, err
		}
	} else {
		tempJobs, err := k8sCli.BatchV1().Jobs(reqObj.NamespaceName).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			cfg.Mlog.Error("failed to k8sCli List, error message: ", err.Error())
			return extra, data, err
		}
		for _, v := range tempJobs.Items {
			for _, obj := range v.ObjectMeta.OwnerReferences {
				if obj.Name == reqObj.CronJobName {
					jobs.Items = append(jobs.Items, v)
				}
			}
		}
	}

	sort.SliceStable(jobs.Items, func(i, j int) bool {
		return jobs.Items[i].CreationTimestamp.Unix() > jobs.Items[j].CreationTimestamp.Unix()
	})

	dataCount := len(jobs.Items)

	reqObj.PageNum, reqObj.PageSize, reqObj.SkipNum, reqObj.SortField, reqObj.SortOrder, reqObj.SortOrderIsDesc = util.GetPageInfoAndSortWay(reqObj.PageNum, reqObj.PageSize, reqObj.SortField, reqObj.SortOrder)
	data = util.SlicePageForJob(jobs.Items, reqObj.SkipNum, reqObj.PageSize)

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
