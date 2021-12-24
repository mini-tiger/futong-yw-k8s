package storage

import (
	"ftk8s/base/cfg"
	"ftk8s/model"

	"gorm.io/gorm"
)

func UpdateRoleByPermission(reqObj *model.ReqUpdateRoleByPermission) error {
	permission, err := GetPermissionByPermissionID(reqObj.PermissionID)
	if err != nil {
		cfg.Mlog.Error("failed to GetPermissionByPermissionID, error message: ", err.Error())
		return err
	}

	err = cfg.Gdb.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		// return any error will rollback

		// Delete role and permission ass in table casbin_rule
		if err := tx.Delete(model.CasbinRule{}, "p_type = ? and v1 = ? and v2 = ? and v3 = ?",
			"p", reqObj.TenantID, permission.PermissionUrl, permission.PermissionAction).Error; err != nil {
			return err
		}

		// Add role and permission ass in table casbin_rule
		casbinRules := make([]model.CasbinRule, 0)
		for _, roleAccount := range reqObj.RoleAccountSlice {
			casbinRule := model.CasbinRule{
				PType: "p",
				V0:    roleAccount,
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

func ReadRoleByPermission(reqObj *model.ReqReadRoleByPermission) (roleAccountSlice []string, err error) {
	permission, err := GetPermissionByPermissionID(reqObj.PermissionID)
	if err != nil {
		return roleAccountSlice, err
	}

	casbinRules := make([]model.CasbinRule, 0)
	err = cfg.Gdb.Table(model.CasbinRule{}.TableName()).
		Where("p_type = ? and v1 = ? and v2 = ? and v3 = ?",
			"p", reqObj.TenantID, permission.PermissionUrl, permission.PermissionAction).
		Find(&casbinRules).Error
	if err != nil {
		return roleAccountSlice, err
	}

	for _, casbinRule := range casbinRules {
		roleAccountSlice = append(roleAccountSlice, casbinRule.V0)
	}

	return roleAccountSlice, err
}

func UpdatePermissionByRole(reqObj *model.ReqUpdatePermissionByRole) error {
	err := cfg.Gdb.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		// return any error will rollback

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

func ReadPermissionOwnedByRole(roleAccount string, tenantID string) (permissions [][]string) {
	permissions = cfg.CasbinSE.GetPermissionsForUserInDomain(roleAccount, tenantID)
	return permissions
}
