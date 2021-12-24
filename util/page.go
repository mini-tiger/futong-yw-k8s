package util

import (
	"ftk8s/base/enum"

	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	storagev1 "k8s.io/api/storage/v1"
)

// GetPageInfo get and correct page information, return: pageNum,pageSize,skipNum
func GetPageInfo(pageNum int, pageSize int) (int, int, int) {
	if pageNum < 1 {
		pageNum = 1
	}
	if pageSize < 1 {
		pageSize = 15
	}
	skipNum := (pageNum - 1) * pageSize
	return pageNum, pageSize, skipNum
}

// GetSortWay get and correct sort information
func GetSortWay(sortField string, sortOrder string) (string, string, bool) {
	sortOrderIsDesc := false
	if len(sortField) == 0 {
		sortField = "id"
	}
	if sortOrder != enum.SortOrderAsc {
		sortOrder = enum.SortOrderDesc
		sortOrderIsDesc = true
	}

	return sortField, sortOrder, sortOrderIsDesc
}

func GetPageInfoAndSortWay(pageNumTemp int, pageSizeTemp int, sortFieldTemp string, sortOrderTemp string) (int, int, int, string, string, bool) {
	pageNum, pageSize, skipNum := GetPageInfo(pageNumTemp, pageSizeTemp)
	sortField, sortOrder, sortOrderIsDesc := GetSortWay(sortFieldTemp, sortOrderTemp)
	return pageNum, pageSize, skipNum, sortField, sortOrder, sortOrderIsDesc
}

func GetPageCount(dataCount int, pageSize int) int {
	pageCount := (dataCount + pageSize - 1) / pageSize
	return pageCount
}

func SlicePageForNamespace(oriSlice []corev1.Namespace, skipNum int, pageSize int) []corev1.Namespace {
	if skipNum > len(oriSlice) {
		skipNum = len(oriSlice)
	}

	end := skipNum + pageSize
	if end > len(oriSlice) {
		end = len(oriSlice)
	}

	return oriSlice[skipNum:end]
}

func SlicePageForDeployment(oriSlice []*appsv1.Deployment, skipNum int, pageSize int) []*appsv1.Deployment {
	if skipNum > len(oriSlice) {
		skipNum = len(oriSlice)
	}

	end := skipNum + pageSize
	if end > len(oriSlice) {
		end = len(oriSlice)
	}

	return oriSlice[skipNum:end]
}

func SlicePageForPod(oriSlice []*corev1.Pod, skipNum int, pageSize int) []*corev1.Pod {
	if skipNum > len(oriSlice) {
		skipNum = len(oriSlice)
	}

	end := skipNum + pageSize
	if end > len(oriSlice) {
		end = len(oriSlice)
	}

	return oriSlice[skipNum:end]
}

func SlicePageForService(oriSlice []corev1.Service, skipNum int, pageSize int) []corev1.Service {
	if skipNum > len(oriSlice) {
		skipNum = len(oriSlice)
	}

	end := skipNum + pageSize
	if end > len(oriSlice) {
		end = len(oriSlice)
	}

	return oriSlice[skipNum:end]
}

func SlicePageForConfigMap(oriSlice []corev1.ConfigMap, skipNum int, pageSize int) []corev1.ConfigMap {
	if skipNum > len(oriSlice) {
		skipNum = len(oriSlice)
	}

	end := skipNum + pageSize
	if end > len(oriSlice) {
		end = len(oriSlice)
	}

	return oriSlice[skipNum:end]
}

func SlicePageForSecret(oriSlice []corev1.Secret, skipNum int, pageSize int) []corev1.Secret {
	if skipNum > len(oriSlice) {
		skipNum = len(oriSlice)
	}

	end := skipNum + pageSize
	if end > len(oriSlice) {
		end = len(oriSlice)
	}

	return oriSlice[skipNum:end]
}

func SlicePageForStatefulSet(oriSlice []appsv1.StatefulSet, skipNum int, pageSize int) []appsv1.StatefulSet {
	if skipNum > len(oriSlice) {
		skipNum = len(oriSlice)
	}

	end := skipNum + pageSize
	if end > len(oriSlice) {
		end = len(oriSlice)
	}

	return oriSlice[skipNum:end]
}

func SlicePageForDaemonSet(oriSlice []appsv1.DaemonSet, skipNum int, pageSize int) []appsv1.DaemonSet {
	if skipNum > len(oriSlice) {
		skipNum = len(oriSlice)
	}

	end := skipNum + pageSize
	if end > len(oriSlice) {
		end = len(oriSlice)
	}

	return oriSlice[skipNum:end]
}

func SlicePageForJob(oriSlice []batchv1.Job, skipNum int, pageSize int) []batchv1.Job {
	if skipNum > len(oriSlice) {
		skipNum = len(oriSlice)
	}

	end := skipNum + pageSize
	if end > len(oriSlice) {
		end = len(oriSlice)
	}

	return oriSlice[skipNum:end]
}

func SlicePageForCronJob(oriSlice []batchv1beta1.CronJob, skipNum int, pageSize int) []batchv1beta1.CronJob {
	if skipNum > len(oriSlice) {
		skipNum = len(oriSlice)
	}

	end := skipNum + pageSize
	if end > len(oriSlice) {
		end = len(oriSlice)
	}

	return oriSlice[skipNum:end]
}

func SlicePageForIngress(oriSlice []extensionsv1beta1.Ingress, skipNum int, pageSize int) []extensionsv1beta1.Ingress {
	if skipNum > len(oriSlice) {
		skipNum = len(oriSlice)
	}

	end := skipNum + pageSize
	if end > len(oriSlice) {
		end = len(oriSlice)
	}

	return oriSlice[skipNum:end]
}

func SlicePageForPVC(oriSlice []corev1.PersistentVolumeClaim, skipNum int, pageSize int) []corev1.PersistentVolumeClaim {
	if skipNum > len(oriSlice) {
		skipNum = len(oriSlice)
	}

	end := skipNum + pageSize
	if end > len(oriSlice) {
		end = len(oriSlice)
	}

	return oriSlice[skipNum:end]
}

func SlicePageForPV(oriSlice []corev1.PersistentVolume, skipNum int, pageSize int) []corev1.PersistentVolume {
	if skipNum > len(oriSlice) {
		skipNum = len(oriSlice)
	}

	end := skipNum + pageSize
	if end > len(oriSlice) {
		end = len(oriSlice)
	}

	return oriSlice[skipNum:end]
}

func SlicePageForStorageClass(oriSlice []storagev1.StorageClass, skipNum int, pageSize int) []storagev1.StorageClass {
	if skipNum > len(oriSlice) {
		skipNum = len(oriSlice)
	}

	end := skipNum + pageSize
	if end > len(oriSlice) {
		end = len(oriSlice)
	}

	return oriSlice[skipNum:end]
}
