package storage

import (
	"ftk8s/base/cfg"
	"ftk8s/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func ImportCluster(cluster *model.Cluster) error {
	err := cfg.Gdb.Create(cluster).Error
	return err
}

func UpdateCluster(reqObj *model.ReqUpdateCluster) error {
	err := cfg.Gdb.Table(model.Cluster{}.TableName()).
		Where("id = ?", reqObj.ClusterID).
		Updates(model.Cluster{ClusterName: reqObj.ClusterName, ClusterAPI: reqObj.ClusterAPI, K8sConfig: reqObj.K8sConfig, Description: reqObj.Description}).Error
	return err
}

func ReadCluster(clusterID int) (model.Cluster, error) {
	cluster := model.Cluster{}
	err := cfg.Gdb.Table(model.Cluster{}.TableName()).
		Where("id = ?", clusterID).
		First(&cluster).Error
	return cluster, err
}

func DeleteCluster(clusterID int) error {
	err := cfg.Gdb.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		// return any error will rollback

		// Delete cluster in table cluster
		if err := tx.Delete(model.Cluster{}, "id = ?", clusterID).Error; err != nil {
			return err
		}

		// 由于开启了外键级联删除，故此处可以省略不写
		// // Delete cluster and group ass in table group_cluster_ass
		// if err := tx.Delete(model.GroupClusterAss{}, "cluster_id = ?", clusterID).Error; err != nil {
		// 	return err
		// }

		// return nil will commit the whole transaction
		return nil
	})

	if err == nil {
		// Load policy
		if err := cfg.CasbinSE.LoadPolicy(); err != nil {
			cfg.Mlog.Error("failed to CasbinSE.LoadPolicy, error message: ", err.Error())
			return err
		}
		return err

	} else {
		return err
	}
}

func ListAllCluster() ([]model.Cluster, error) {
	clusters := make([]model.Cluster, 0)

	err := cfg.Gdb.Table(model.Cluster{}.TableName()).Find(&clusters).Error
	return clusters, err
}

func ListClusterByTenantID(tenantID string) ([]model.Cluster, error) {
	clusters := make([]model.Cluster, 0)

	err := cfg.Gdb.Table(model.Cluster{}.TableName()).
		Where("tenant_id = ?", tenantID).Find(&clusters).Error
	return clusters, err
}

func ListClusterByUserID(userID string) ([]model.Cluster, error) {
	clusters := make([]model.Cluster, 0)

	err := cfg.Gdb.Table(model.User{}.TableName()).
		Joins("join user_group_ass as t1 on t1.user_id = ?", userID).
		Joins("join group_cluster_ass as t2 on t2.group_id = t1.group_id").
		Joins("join cluster on cluster.id = t2.cluster_id").
		Find(&clusters).Error
	return clusters, err
}

func ListClusterIDByUserID(userID string) ([]model.Cluster, error) {
	clusters := make([]model.Cluster, 0)

	err := cfg.Gdb.Table(model.User{}.TableName()).Select("t2.cluster_id").
		Joins("join user_group_ass as t1 on t1.user_id = ?", userID).
		Joins("join group_cluster_ass as t2 on t2.group_id = t1.group_id").
		Find(&clusters).Error
	return clusters, err
}

func ListClusterByTenantIDWithPage(reqObj *model.ReqListCluster) (int64, []model.Cluster, error) {
	var dataCount int64
	clusters := make([]model.Cluster, 0)

	err := cfg.Gdb.Table(model.Cluster{}.TableName()).
		Where("tenant_id = ?", reqObj.TenantID).
		Count(&dataCount).Error
	if err != nil {
		return dataCount, nil, err
	}

	err = cfg.Gdb.Table(model.Cluster{}.TableName()).
		Where("tenant_id = ?", reqObj.TenantID).
		Order(clause.OrderByColumn{Column: clause.Column{Name: reqObj.SortField}, Desc: reqObj.SortOrderIsDesc}).
		Limit(reqObj.PageSize).Offset(reqObj.SkipNum).
		Find(&clusters).Error
	return dataCount, clusters, err
}

func ListClusterByUserIDWithPage(reqObj *model.ReqListCluster) (int64, []model.Cluster, error) {
	var dataCount int64
	clusters := make([]model.Cluster, 0)

	err := cfg.Gdb.Table(model.User{}.TableName()).
		Joins("join user_group_ass as t1 on t1.user_id = ?", reqObj.UserID).
		Joins("join group_cluster_ass as t2 on t2.group_id = t1.group_id").
		Joins("join cluster on cluster.id = t2.cluster_id").
		Count(&dataCount).Error
	if err != nil {
		return 0, nil, err
	}

	err = cfg.Gdb.Table(model.User{}.TableName()).
		Joins("join user_group_ass as t1 on t1.user_id = ?", reqObj.UserID).
		Joins("join group_cluster_ass as t2 on t2.group_id = t1.group_id").
		Joins("join cluster on cluster.id = t2.cluster_id").
		Order(clause.OrderByColumn{Column: clause.Column{Name: reqObj.SortField}, Desc: reqObj.SortOrderIsDesc}).
		Limit(reqObj.PageSize).Offset(reqObj.SkipNum).
		Find(&clusters).Error
	return dataCount, clusters, err
}

func ListClusterByTenantIDByClusterAccountWithPage(reqObj *model.ReqListCluster) (int64, []model.Cluster, error) {
	var dataCount int64
	clusters := make([]model.Cluster, 0)

	err := cfg.Gdb.Table(model.Cluster{}.TableName()).
		Where("tenant_id = ? and cluster_account = ?", reqObj.TenantID, reqObj.ClusterAccount).
		Count(&dataCount).Error
	if err != nil {
		return dataCount, nil, err
	}

	err = cfg.Gdb.Table(model.Cluster{}.TableName()).
		Where("tenant_id = ? and cluster_account = ?", reqObj.TenantID, reqObj.ClusterAccount).
		Order(clause.OrderByColumn{Column: clause.Column{Name: reqObj.SortField}, Desc: reqObj.SortOrderIsDesc}).
		Limit(reqObj.PageSize).Offset(reqObj.SkipNum).
		Find(&clusters).Error
	return dataCount, clusters, err
}

func ListClusterByUserIDByClusterAccountWithPage(reqObj *model.ReqListCluster) (int64, []model.Cluster, error) {
	var dataCount int64
	clusters := make([]model.Cluster, 0)

	err := cfg.Gdb.Table(model.User{}.TableName()).
		Joins("join user_group_ass as t1 on t1.user_id = ?", reqObj.UserID).
		Joins("join group_cluster_ass as t2 on t2.group_id = t1.group_id").
		Joins("join cluster on cluster.id = t2.cluster_id and cluster.cluster_account = ?", reqObj.ClusterAccount).
		Count(&dataCount).Error
	if err != nil {
		return 0, nil, err
	}

	err = cfg.Gdb.Table(model.User{}.TableName()).
		Joins("join user_group_ass as t1 on t1.user_id = ?", reqObj.UserID).
		Joins("join group_cluster_ass as t2 on t2.group_id = t1.group_id").
		Joins("join cluster on cluster.id = t2.cluster_id and cluster.cluster_account = ?", reqObj.ClusterAccount).
		Order(clause.OrderByColumn{Column: clause.Column{Name: reqObj.SortField}, Desc: reqObj.SortOrderIsDesc}).
		Limit(reqObj.PageSize).Offset(reqObj.SkipNum).
		Find(&clusters).Error
	return dataCount, clusters, err
}
