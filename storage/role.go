package storage

import (
	"ftk8s/base/cfg"
	"ftk8s/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateRoleWithPermission(reqObj *model.ReqCreateRoleWithPermission) error {
	err := cfg.Gdb.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		// return any error will rollback

		// Add role in table role
		if err := tx.Create(&reqObj.Role).Error; err != nil {
			return err
		}

		// Add role and webperm ass in table role_webperm_ass
		roleWebpermAssSlice := make([]model.RoleWebpermAss, 0)
		for _, webpermID := range reqObj.WebpermIDSlice {
			roleWebpermAss := model.RoleWebpermAss{RoleID: reqObj.Role.ID, WebpermID: webpermID}
			roleWebpermAssSlice = append(roleWebpermAssSlice, roleWebpermAss)
		}
		if len(roleWebpermAssSlice) != 0 {
			if err := tx.Create(&roleWebpermAssSlice).Error; err != nil {
				return err
			}
		}

		// Delete role and permission ass in table casbin_rule
		if err := tx.Delete(model.CasbinRule{}, "p_type = ? and v0 = ? and v1 = ?",
			"p", reqObj.RoleAccount, reqObj.TenantID).Error; err != nil {
			return err
		}

		// Add role and permission ass in table casbin_rule
		casbinRules := make([]model.CasbinRule, 0)
		for _, permissionID := range reqObj.PermissionIDSlice {
			permission, err := GetPermissionByPermissionID(permissionID)
			if err != nil {
				cfg.Mlog.Error("failed to GetPermissionByPermissionID, error message: ", err.Error())
				return err
			}
			casbinRule := model.CasbinRule{
				PType: "p",
				V0:    reqObj.RoleAccount,
				V1:    reqObj.TenantID,
				V2:    permission.PermissionUrl,
				V3:    permission.PermissionAction,
			}
			casbinRules = append(casbinRules, casbinRule)
		}
		if len(casbinRules) != 0 {
			if err := tx.Create(&casbinRules).Error; err != nil {
				return err
			}
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

func UpdateRoleWithPermission(reqObj *model.ReqUpdateRoleWithPermission) error {
	err := cfg.Gdb.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		// return any error will rollback

		// Update role in table role
		if err := tx.Table(model.Role{}.TableName()).
			Where("id = ?", reqObj.RoleID).
			Updates(model.Role{RoleName: reqObj.RoleName, Description: reqObj.Description}).Error; err != nil {
			return err
		}

		// Delete role and webperm ass in table role_webperm_ass
		if err := tx.Delete(model.RoleWebpermAss{}, "role_id = ?", reqObj.Role.ID).Error; err != nil {
			return err
		}

		// Add role and webperm ass in table role_webperm_ass
		roleWebpermAssSlice := make([]model.RoleWebpermAss, 0)
		for _, webpermID := range reqObj.WebpermIDSlice {
			roleWebpermAss := model.RoleWebpermAss{RoleID: reqObj.Role.ID, WebpermID: webpermID}
			roleWebpermAssSlice = append(roleWebpermAssSlice, roleWebpermAss)
		}
		if len(roleWebpermAssSlice) != 0 {
			if err := tx.Create(&roleWebpermAssSlice).Error; err != nil {
				return err
			}
		}

		// Delete role and permission ass in table casbin_rule
		if err := tx.Delete(model.CasbinRule{}, "p_type = ? and v0 = ? and v1 = ?",
			"p", reqObj.Role.RoleAccount, reqObj.Role.TenantID).Error; err != nil {
			return err
		}

		// Add role and permission ass in table casbin_rule
		casbinRules := make([]model.CasbinRule, 0)
		for _, permissionID := range reqObj.PermissionIDSlice {
			permission, err := GetPermissionByPermissionID(permissionID)
			if err != nil {
				cfg.Mlog.Error("failed to GetPermissionByPermissionID, error message: ", err.Error())
				return err
			}
			casbinRule := model.CasbinRule{
				PType: "p",
				V0:    reqObj.Role.RoleAccount,
				V1:    reqObj.Role.TenantID,
				V2:    permission.PermissionUrl,
				V3:    permission.PermissionAction,
			}
			casbinRules = append(casbinRules, casbinRule)
		}
		if len(casbinRules) != 0 {
			if err := tx.Create(&casbinRules).Error; err != nil {
				return err
			}
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

func ReadRole(roleID int) (model.Role, error) {
	role := model.Role{}
	err := cfg.Gdb.Table(model.Role{}.TableName()).
		Where("id = ?", roleID).
		First(&role).Error
	return role, err
}

func ReadRoleByRoleAccount(tenantID string, roleAccount string) (model.Role, error) {
	role := model.Role{}
	err := cfg.Gdb.Table(model.Role{}.TableName()).
		Where("tenant_id = ? and role_account = ?", tenantID, roleAccount).
		First(&role).Error
	return role, err
}

func DeleteRole(reqObj *model.ReqDeleteRole) error {
	err := cfg.Gdb.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		// return any error will rollback

		// Delete role in table role
		if err := tx.Delete(model.Role{}, "id = ?", reqObj.RoleID).Error; err != nil {
			return err
		}

		// Delete role and permission ass in table casbin_rule
		if err := tx.Delete(model.CasbinRule{}, "p_type = ? and v0 = ? and v1 = ?",
			"p", reqObj.RoleAccount, reqObj.TenantID).Error; err != nil {
			return err
		}

		// Delete role and user or group ass in table casbin_rule
		if err := tx.Delete(model.CasbinRule{}, "p_type = ? and v1 = ? and v2 = ?",
			"g", reqObj.RoleAccount, reqObj.TenantID).Error; err != nil {
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

func ListRole(tenantID string) ([]model.Role, error) {
	roles := make([]model.Role, 0)

	err := cfg.Gdb.Table(model.Role{}.TableName()).
		Where("tenant_id = ?", tenantID).
		Find(&roles).Error
	return roles, err
}

func ListRoleWithPage(reqObj *model.ReqListRole) (int64, []model.Role, error) {
	var dataCount int64
	roles := make([]model.Role, 0)

	err := cfg.Gdb.Table(model.Role{}.TableName()).
		Where("tenant_id = ?", reqObj.TenantID).
		Count(&dataCount).Error
	if err != nil {
		return dataCount, nil, err
	}

	err = cfg.Gdb.Table(model.Role{}.TableName()).
		Where("tenant_id = ?", reqObj.TenantID).
		Order(clause.OrderByColumn{Column: clause.Column{Name: reqObj.SortField}, Desc: reqObj.SortOrderIsDesc}).
		Limit(reqObj.PageSize).Offset(reqObj.SkipNum).
		Find(&roles).Error
	return dataCount, roles, err
}

func ListRoleByRoleAccountWithPage(reqObj *model.ReqListRole) (int64, []model.Role, error) {
	var dataCount int64
	roles := make([]model.Role, 0)

	err := cfg.Gdb.Table(model.Role{}.TableName()).
		Where("tenant_id = ? and role_account = ?", reqObj.TenantID, reqObj.RoleAccount).
		Count(&dataCount).Error
	if err != nil {
		return dataCount, nil, err
	}

	err = cfg.Gdb.Table(model.Role{}.TableName()).
		Where("tenant_id = ? and role_account = ?", reqObj.TenantID, reqObj.RoleAccount).
		Order(clause.OrderByColumn{Column: clause.Column{Name: reqObj.SortField}, Desc: reqObj.SortOrderIsDesc}).
		Limit(reqObj.PageSize).Offset(reqObj.SkipNum).
		Find(&roles).Error
	return dataCount, roles, err
}
