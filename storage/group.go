package storage

import (
	"ftk8s/base/cfg"
	"ftk8s/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateGroup(group *model.Group) error {
	err := cfg.Gdb.Create(group).Error
	return err
}

func UpdateGroup(reqObj *model.ReqUpdateGroup) error {
	err := cfg.Gdb.Table(model.Group{}.TableName()).
		Where("id = ?", reqObj.GroupID).
		Updates(model.Group{GroupName: reqObj.GroupName, Description: reqObj.Description}).Error
	return err
}

func ReadGroup(groupID string) (model.Group, error) {
	group := model.Group{}
	err := cfg.Gdb.Table(model.Group{}.TableName()).
		Where("id = ?", groupID).
		First(&group).Error
	return group, err
}

func DeleteGroup(reqObj *model.ReqDeleteGroup) error {
	err := cfg.Gdb.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		// return any error will rollback

		// Delete group in table group
		if err := tx.Delete(model.Group{}, "id = ?", reqObj.GroupID).Error; err != nil {
			return err
		}

		// 由于开启了外键级联删除，故此处可以省略不写
		// // Delete group and user ass in table user_group_ass
		// if err := tx.Delete(model.UserGroupAss{}, "group_id = ?", reqObj.GroupID).Error; err != nil {
		// 	return err
		// }
		//
		// // Delete group and cluster ass in table group_cluster_ass
		// if err := tx.Delete(model.GroupClusterAss{}, "group_id = ?", reqObj.GroupID).Error; err != nil {
		// 	return err
		// }

		// Delete group and role ass in table casbin_rule
		if err := tx.Delete(model.CasbinRule{}, "p_type = ? and v0 = ? and v2 = ?",
			"g", reqObj.GroupAccount, reqObj.TenantID).Error; err != nil {
			return err
		}

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

func ListGroup(tenantID string) ([]model.Group, error) {
	groups := make([]model.Group, 0)

	err := cfg.Gdb.Table(model.Group{}.TableName()).
		Where("tenant_id = ?", tenantID).
		Find(&groups).Error
	return groups, err
}

func ListGroupWithPage(reqObj *model.ReqListGroup) (int64, []model.Group, error) {
	var dataCount int64
	groups := make([]model.Group, 0)

	err := cfg.Gdb.Table(model.Group{}.TableName()).
		Where("tenant_id = ?", reqObj.TenantID).
		Count(&dataCount).Error
	if err != nil {
		return dataCount, nil, err
	}

	err = cfg.Gdb.Table(model.Group{}.TableName()).
		Where("tenant_id = ?", reqObj.TenantID).
		Order(clause.OrderByColumn{Column: clause.Column{Name: reqObj.SortField}, Desc: reqObj.SortOrderIsDesc}).
		Limit(reqObj.PageSize).Offset(reqObj.SkipNum).
		Find(&groups).Error
	return dataCount, groups, err
}

func ListGroupByGroupAccountWithPage(reqObj *model.ReqListGroup) (int64, []model.Group, error) {
	var dataCount int64
	groups := make([]model.Group, 0)

	err := cfg.Gdb.Table(model.Group{}.TableName()).
		Where("tenant_id = ? and group_account = ?", reqObj.TenantID, reqObj.GroupAccount).
		Count(&dataCount).Error
	if err != nil {
		return dataCount, nil, err
	}

	err = cfg.Gdb.Table(model.Group{}.TableName()).
		Where("tenant_id = ? and group_account = ?", reqObj.TenantID, reqObj.GroupAccount).
		Order(clause.OrderByColumn{Column: clause.Column{Name: reqObj.SortField}, Desc: reqObj.SortOrderIsDesc}).
		Limit(reqObj.PageSize).Offset(reqObj.SkipNum).
		Find(&groups).Error
	return dataCount, groups, err
}
