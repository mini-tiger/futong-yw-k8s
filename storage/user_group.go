package storage

import (
	"ftk8s/base/cfg"
	"ftk8s/model"

	"gorm.io/gorm"
)

func UpdateUserByGroup(reqObj *model.ReqUpdateUserByGroup) error {
	err := cfg.Gdb.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		// return any error will rollback

		// Delete all users associated with this group
		if err := tx.Delete(model.UserGroupAss{}, "group_id = ?", reqObj.GroupID).Error; err != nil {
			return err
		}

		// Add all users associated with this group
		userGroupAssSlice := make([]model.UserGroupAss, 0)
		for _, userID := range reqObj.UserIDSlice {
			userGroupAss := model.UserGroupAss{UserID: userID, GroupID: reqObj.GroupID}
			userGroupAssSlice = append(userGroupAssSlice, userGroupAss)
		}
		if len(userGroupAssSlice) != 0 {
			if err := tx.Create(&userGroupAssSlice).Error; err != nil {
				return err
			}
		}

		// return nil will commit the whole transaction
		return nil
	})
	return err
}

func ReadUserByGroup(group model.Group) ([]model.User, error) {
	users := make([]model.User, 0)

	err := cfg.Gdb.Model(&group).Association("UserGroup").Find(&users)
	return users, err
}

func UpdateGroupByUser(reqObj *model.ReqUpdateGroupByUser) error {
	err := cfg.Gdb.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		// return any error will rollback

		// Delete all groups associated with this user
		if err := tx.Delete(model.UserGroupAss{}, "user_id = ?", reqObj.UserID).Error; err != nil {
			return err
		}

		// Add all groups associated with this user
		userGroupAssSlice := make([]model.UserGroupAss, 0)
		for _, groupID := range reqObj.GroupIDSlice {
			userGroupAss := model.UserGroupAss{UserID: reqObj.UserID, GroupID: groupID}
			userGroupAssSlice = append(userGroupAssSlice, userGroupAss)
		}
		if len(userGroupAssSlice) != 0 {
			if err := tx.Create(&userGroupAssSlice).Error; err != nil {
				return err
			}
		}

		// return nil will commit the whole transaction
		return nil
	})
	return err
}

func ReadGroupByUser(user model.User) ([]model.Group, error) {
	groups := make([]model.Group, 0)

	err := cfg.Gdb.Model(&user).Association("UserGroup").Find(&groups)
	return groups, err
}
