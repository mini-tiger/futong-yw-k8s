package storage

import (
	"fmt"

	"ftk8s/base/cfg"
	"ftk8s/base/enum"
	"ftk8s/model"

	"gorm.io/gorm"
)

func UpdateRoleByUserGroup(reqObj *model.ReqUpdateRoleByUserGroup) error {
	err := cfg.Gdb.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		// return any error will rollback

		// Delete role and user or group ass in table casbin_rule
		if err := tx.Delete(model.CasbinRule{}, "p_type = ? and v0 = ? and v2 = ?",
			"g", reqObj.RoleAssObj, reqObj.TenantID).Error; err != nil {
			return err
		}

		// Add role and user or group ass in table casbin_rule
		casbinRules := make([]model.CasbinRule, 0)
		for _, roleAccount := range reqObj.RoleAccountSlice {
			casbinRule := model.CasbinRule{
				PType: "g",
				V0:    reqObj.RoleAssObj,
				V1:    roleAccount,
				V2:    reqObj.TenantID,
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

func ReadRoleByUserGroup(roleAssObj string, tenantID string) []string {
	roles := cfg.CasbinSE.GetRolesForUserInDomain(roleAssObj, tenantID)
	return roles
}

func UpdateUserGroupByRole(reqObj *model.ReqUpdateUserGroupByRole) error {
	err := cfg.Gdb.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		// return any error will rollback

		// Delete role and user or group ass in table casbin_rule
		if reqObj.RoleAssObjPrefix == enum.PrefixUser {
			if err := tx.Delete(model.CasbinRule{}, "p_type = ? and v1 = ? and v2 = ? and v0 like ?",
				"g", reqObj.RoleAccount, reqObj.TenantID, enum.PrefixUser+"%").Error; err != nil {
				return err
			}

		} else if reqObj.RoleAssObjPrefix == enum.PrefixGroup {
			if err := tx.Delete(model.CasbinRule{}, "p_type = ? and v1 = ? and v2 = ? and v0 like ?",
				"g", reqObj.RoleAccount, reqObj.TenantID, enum.PrefixGroup+"%").Error; err != nil {
				return err
			}
		} else {
			return fmt.Errorf("invalid RoleAssObjPrefix")
		}

		// Add role and user or group ass in table casbin_rule
		casbinRules := make([]model.CasbinRule, 0)
		for _, roleAssObj := range reqObj.RoleAssObjSlice {
			casbinRule := model.CasbinRule{
				PType: "g",
				V0:    roleAssObj,
				V1:    reqObj.RoleAccount,
				V2:    reqObj.TenantID,
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

func ReadUserGroupByRole(roleAccount string, tenantID string) []string {
	roleAssObjSlice := cfg.CasbinSE.GetUsersForRoleInDomain(roleAccount, tenantID)
	return roleAssObjSlice
}
