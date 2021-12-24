package storage

import (
	"ftk8s/base/cfg"
	"ftk8s/base/enum"
	"ftk8s/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func RegisterTenant(tenant *model.Tenant) error {
	err := cfg.Gdb.Create(tenant).Error
	return err
}

func GetTenantByTenantID(tenantID string) (*model.Tenant, error) {
	tenant := new(model.Tenant)
	err := cfg.Gdb.Where("id = ?", tenantID).First(&tenant).Error
	return tenant, err
}

func GetUserByUserAccount(tenantID string, userAccount string) (*model.User, error) {
	user := new(model.User)
	err := cfg.Gdb.Where("tenant_id = ? and user_account = ?", tenantID, userAccount).First(&user).Error
	return user, err
}

// GetGroupIDSliceByUserID get a slice of user group id based on user id
func GetGroupIDSliceByUserID(user *model.User) ([]string, error) {
	groups := make([]model.Group, 0)
	err := cfg.Gdb.Model(&user).Association("UserGroup").Find(&groups)
	if err != nil {
		return nil, err
	}
	groupIDSlice := make([]string, 0)
	for _, group := range groups {
		groupIDSlice = append(groupIDSlice, group.ID)
	}

	return groupIDSlice, err
}

// UpdateTenantLoginInfo update lastLoginTime and lastLoginIP after tenant login
func UpdateTenantLoginInfo(tenant *model.Tenant) error {
	err := cfg.Gdb.Model(tenant).Updates(map[string]interface{}{"last_login_time": tenant.LastLoginTime, "last_login_ip": tenant.LastLoginIP}).Error
	return err
}

// UpdateUserLoginInfo update lastLoginTime and lastLoginIP after user login
func UpdateUserLoginInfo(user *model.User) error {
	err := cfg.Gdb.Model(user).Updates(map[string]interface{}{"last_login_time": user.LastLoginTime, "last_login_ip": user.LastLoginIP}).Error
	return err
}

func CreateUser(user *model.User) error {
	err := cfg.Gdb.Create(user).Error
	return err
}

func UpdateUser(reqObj *model.ReqUpdateUser) error {
	err := cfg.Gdb.Table(model.User{}.TableName()).
		Where("id = ?", reqObj.UserID).
		Updates(model.User{Username: reqObj.Username, Email: reqObj.Email, Description: reqObj.Description}).Error
	return err
}

func ResetPassword(reqObj *model.ReqResetPassword) error {
	switch reqObj.UserType {
	case enum.UserTypeTenant:
		err := cfg.Gdb.Table(model.Tenant{}.TableName()).
			Where("id = ?", reqObj.Tenant.ID).
			Updates(model.Tenant{Salt: reqObj.Salt, Password: reqObj.Password}).Error
		return err

	default:
		err := cfg.Gdb.Table(model.User{}.TableName()).
			Where("id = ?", reqObj.UserID).
			Updates(model.User{Salt: reqObj.Salt, Password: reqObj.Password}).Error
		return err
	}
}

func ReadUser(userID string) (model.User, error) {
	user := model.User{}
	err := cfg.Gdb.Table(model.User{}.TableName()).
		Where("id = ?", userID).
		First(&user).Error
	return user, err
}

func DeleteUser(reqObj *model.ReqDeleteUser) error {
	err := cfg.Gdb.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		// return any error will rollback

		// Delete user in table user
		if err := tx.Delete(model.User{}, "id = ?", reqObj.UserID).Error; err != nil {
			return err
		}

		// 由于开启了外键级联删除，故此处可以省略不写
		// // Delete user and group ass in table user_group_ass
		// if err := tx.Delete(model.UserGroupAss{}, "user_id = ?", reqObj.UserID).Error; err != nil {
		// 	return err
		// }

		// Delete p_type=g match user data in table casbin_rule
		if err := tx.Delete(model.CasbinRule{}, "p_type = ? and v0 = ? and v2 = ?",
			"g", reqObj.UserID, reqObj.TenantID).Error; err != nil {
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

func ListUser(tenantID string) ([]model.User, error) {
	users := make([]model.User, 0)

	err := cfg.Gdb.Table(model.User{}.TableName()).
		Where("tenant_id = ?", tenantID).
		Find(&users).Error
	return users, err
}

func ListUserWithPage(reqObj *model.ReqListUser) (int64, []model.User, error) {
	var dataCount int64
	users := make([]model.User, 0)

	err := cfg.Gdb.Table(model.User{}.TableName()).
		Where("tenant_id = ?", reqObj.TenantID).
		Count(&dataCount).Error
	if err != nil {
		return dataCount, nil, err
	}

	err = cfg.Gdb.Table(model.User{}.TableName()).
		Where("tenant_id = ?", reqObj.TenantID).
		Order(clause.OrderByColumn{Column: clause.Column{Name: reqObj.SortField}, Desc: reqObj.SortOrderIsDesc}).
		Limit(reqObj.PageSize).Offset(reqObj.SkipNum).
		Find(&users).Error
	return dataCount, users, err
}

func ListUserByUserAccountWithPage(reqObj *model.ReqListUser) (int64, []model.User, error) {
	var dataCount int64
	users := make([]model.User, 0)

	err := cfg.Gdb.Table(model.User{}.TableName()).
		Where("tenant_id = ? and user_account = ?", reqObj.TenantID, reqObj.UserAccount).
		Count(&dataCount).Error
	if err != nil {
		return dataCount, nil, err
	}

	err = cfg.Gdb.Table(model.User{}.TableName()).
		Where("tenant_id = ? and user_account = ?", reqObj.TenantID, reqObj.UserAccount).
		Order(clause.OrderByColumn{Column: clause.Column{Name: reqObj.SortField}, Desc: reqObj.SortOrderIsDesc}).
		Limit(reqObj.PageSize).Offset(reqObj.SkipNum).
		Find(&users).Error
	return dataCount, users, err
}
