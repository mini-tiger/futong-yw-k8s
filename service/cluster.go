package service

import (
	"ftk8s/base/enum"
	"ftk8s/ksc"
	"ftk8s/model"
	"ftk8s/storage"
	"ftk8s/util"
)

func ImportCluster(reqObj *model.ReqImportCluster) error {
	cluster := new(model.Cluster)

	cluster.TenantID = reqObj.TenantID
	cluster.ClusterAccount = reqObj.ClusterAccount
	cluster.ClusterName = reqObj.ClusterName
	cluster.ClusterAPI = reqObj.ClusterAPI
	cluster.K8sConfig = reqObj.K8sConfig
	cluster.Description = reqObj.Description

	err := ksc.CheckConnectCluster(reqObj.ClusterAPI, reqObj.K8sConfig)
	if err != nil {
		return err
	}

	err = storage.ImportCluster(cluster)
	if err != nil {
		return err
	}

	return nil
}

func UpdateCluster(reqObj *model.ReqUpdateCluster) error {
	err := ksc.CheckConnectCluster(reqObj.ClusterAPI, reqObj.K8sConfig)
	if err != nil {
		return err
	}

	err = storage.UpdateCluster(reqObj)
	if err != nil {
		return err
	}

	return nil
}

func ReadCluster(reqObj *model.ReqReadCluster) (model.Cluster, error) {
	cluster, err := storage.ReadCluster(reqObj.ClusterID)
	if err != nil {
		return model.Cluster{}, err
	}

	return cluster, nil
}

func DeleteCluster(reqObj *model.ReqDeleteCluster) error {
	err := storage.DeleteCluster(reqObj.ClusterID)
	return err
}

func ListCluster(reqObj *model.ReqListCluster) (extra util.Extra, clusters []model.Cluster, err error) {
	reqObj.PageNum, reqObj.PageSize, reqObj.SkipNum, reqObj.SortField, reqObj.SortOrder, reqObj.SortOrderIsDesc = util.GetPageInfoAndSortWay(reqObj.PageNum, reqObj.PageSize, reqObj.SortField, reqObj.SortOrder)
	var dataCountTemp int64

	if reqObj.ClusterAccount == "" {
		if reqObj.UserType == enum.UserTypeTenant {
			dataCountTemp, clusters, err = storage.ListClusterByTenantIDWithPage(reqObj)
			if err != nil {
				return extra, clusters, err
			}
		} else {
			dataCountTemp, clusters, err = storage.ListClusterByUserIDWithPage(reqObj)
			if err != nil {
				return extra, clusters, err
			}
		}
	} else {
		if reqObj.UserType == enum.UserTypeTenant {
			dataCountTemp, clusters, err = storage.ListClusterByTenantIDByClusterAccountWithPage(reqObj)
			if err != nil {
				return extra, clusters, err
			}
		} else {
			dataCountTemp, clusters, err = storage.ListClusterByUserIDByClusterAccountWithPage(reqObj)
			if err != nil {
				return extra, clusters, err
			}
		}
	}

	dataCount := int(dataCountTemp)
	pageCount := util.GetPageCount(dataCount, reqObj.PageSize)
	extra = util.Extra{
		PageNum:   reqObj.PageNum,
		PageSize:  reqObj.PageSize,
		SortField: reqObj.SortField,
		SortOrder: reqObj.SortOrder,
		DataCount: dataCount,
		PageCount: pageCount,
	}

	return extra, clusters, err
}
