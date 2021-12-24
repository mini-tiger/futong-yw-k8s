package service

import (
	"ftk8s/model"
	"ftk8s/storage"
)

func UpdateGroupByCluster(reqObj *model.ReqUpdateGroupByCluster) error {
	err := storage.UpdateGroupByCluster(reqObj)
	if err != nil {
		return err
	}

	return nil
}

func ReadGroupByCluster(reqObj *model.ReqReadGroupByCluster) (*model.DataOwnedOrNotGroup, error) {
	data := new(model.DataOwnedOrNotGroup)

	data.OwnedSlice = make([]model.Group, 0)
	data.NotOwnedSlice = make([]model.Group, 0)

	cluster, err := storage.ReadCluster(reqObj.ClusterID)
	if err != nil {
		return data, err
	}
	ownedSliceTemp, err := storage.ReadGroupByCluster(cluster)
	if err != nil {
		return data, err
	}
	groups, err := storage.ListGroup(cluster.TenantID)
	if err != nil {
		return data, err
	}

	data.OwnedSlice = ownedSliceTemp
	for _, group := range groups {
		flag := false
		for _, owned := range ownedSliceTemp {
			if owned.ID == group.ID {
				flag = true
				break
			}
		}
		if !flag {
			data.NotOwnedSlice = append(data.NotOwnedSlice, group)
		}
	}

	return data, err
}

func UpdateClusterByGroup(reqObj *model.ReqUpdateClusterByGroup) error {
	err := storage.UpdateClusterByGroup(reqObj)
	if err != nil {
		return err
	}

	return nil
}

func ReadClusterByGroup(reqObj *model.ReqReadClusterByGroup) (*model.DataOwnedOrNotCluster, error) {
	data := new(model.DataOwnedOrNotCluster)

	data.OwnedSlice = make([]model.Cluster, 0)
	data.NotOwnedSlice = make([]model.Cluster, 0)

	group, err := storage.ReadGroup(reqObj.GroupID)
	if err != nil {
		return nil, err
	}
	ownedSliceTemp, err := storage.ReadClusterByGroup(group)
	if err != nil {
		return data, err
	}
	clusters, err := storage.ListClusterByTenantID(group.TenantID)
	if err != nil {
		return data, err
	}

	data.OwnedSlice = ownedSliceTemp
	for _, cluster := range clusters {
		flag := false
		for _, owned := range ownedSliceTemp {
			if owned.ID == cluster.ID {
				flag = true
				break
			}
		}
		if !flag {
			data.NotOwnedSlice = append(data.NotOwnedSlice, cluster)
		}
	}

	return data, err
}
