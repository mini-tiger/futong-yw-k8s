package storage

import (
	"ftk8s/base/cfg"
	"ftk8s/model"

	"gorm.io/gorm"
)

func UpdateGroupByCluster(reqObj *model.ReqUpdateGroupByCluster) error {
	err := cfg.Gdb.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		// return any error will rollback

		// Delete all groups associated with this cluster
		if err := tx.Delete(model.GroupClusterAss{}, "cluster_id = ?", reqObj.ClusterID).Error; err != nil {
			return err
		}

		// Add all groups associated with this cluster
		groupClusterAssSlice := make([]model.GroupClusterAss, 0)
		for _, groupID := range reqObj.GroupIDSlice {
			groupClusterAss := model.GroupClusterAss{ClusterID: reqObj.ClusterID, GroupID: groupID}
			groupClusterAssSlice = append(groupClusterAssSlice, groupClusterAss)
		}
		if len(groupClusterAssSlice) != 0 {
			if err := tx.Create(&groupClusterAssSlice).Error; err != nil {
				return err
			}
		}

		// return nil will commit the whole transaction
		return nil
	})
	return err
}

func ReadGroupByCluster(cluster model.Cluster) ([]model.Group, error) {
	groups := make([]model.Group, 0)

	err := cfg.Gdb.Model(&cluster).Association("GroupCluster").Find(&groups)
	return groups, err
}

func UpdateClusterByGroup(reqObj *model.ReqUpdateClusterByGroup) error {
	err := cfg.Gdb.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		// return any error will rollback

		// Delete all clusters associated with this group
		if err := tx.Delete(model.GroupClusterAss{}, "group_id = ?", reqObj.GroupID).Error; err != nil {
			return err
		}

		// Add all clusters associated with this group
		groupClusterAssSlice := make([]model.GroupClusterAss, 0)
		for _, clusterID := range reqObj.ClusterIDSlice {
			groupClusterAss := model.GroupClusterAss{GroupID: reqObj.GroupID, ClusterID: clusterID}
			groupClusterAssSlice = append(groupClusterAssSlice, groupClusterAss)
		}
		if len(groupClusterAssSlice) != 0 {
			if err := tx.Create(&groupClusterAssSlice).Error; err != nil {
				return err
			}
		}

		// return nil will commit the whole transaction
		return nil
	})
	return err
}

func ReadClusterByGroup(group model.Group) ([]model.Cluster, error) {
	clusters := make([]model.Cluster, 0)

	err := cfg.Gdb.Model(&group).Association("GroupCluster").Find(&clusters)
	return clusters, err
}
