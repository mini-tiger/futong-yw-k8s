package service

import (
	"ftk8s/model"
	"ftk8s/storage"
)

func UpdateRoleByUserGroup(reqObj *model.ReqUpdateRoleByUserGroup) error {
	err := storage.UpdateRoleByUserGroup(reqObj)
	return err
}

func ReadRoleByUserGroup(reqObj *model.ReqReadRoleByUserGroup) (*model.DataOwnedOrNotRole, error) {
	data := new(model.DataOwnedOrNotRole)

	data.OwnedSlice = make([]model.Role, 0)
	data.NotOwnedSlice = make([]model.Role, 0)

	ownedSliceTemp := storage.ReadRoleByUserGroup(reqObj.RoleAssObj, reqObj.TenantID)
	roles, err := storage.ListRole(reqObj.TenantID)
	if err != nil {
		return data, err
	}

	for _, role := range roles {
		flag := false
		for _, owned := range ownedSliceTemp {
			if owned == role.RoleAccount {
				flag = true
				break
			}
		}
		if flag {
			data.OwnedSlice = append(data.OwnedSlice, role)
		} else {
			data.NotOwnedSlice = append(data.NotOwnedSlice, role)
		}
	}

	return data, err
}

func UpdateUserGroupByRole(reqObj *model.ReqUpdateUserGroupByRole) error {
	err := storage.UpdateUserGroupByRole(reqObj)
	return err
}

func ReadUserGroupByRole(reqObj *model.ReqReadUserGroupByRole) (*model.DataOwnedOrNotUserGroup, error) {
	data := new(model.DataOwnedOrNotUserGroup)

	data.Users.OwnedSlice = make([]model.User, 0)
	data.Users.NotOwnedSlice = make([]model.User, 0)
	data.Groups.OwnedSlice = make([]model.Group, 0)
	data.Groups.NotOwnedSlice = make([]model.Group, 0)

	ownedSliceTemp := storage.ReadUserGroupByRole(reqObj.RoleAccount, reqObj.TenantID)
	users, err := storage.ListUser(reqObj.TenantID)
	if err != nil {
		return data, err
	}
	groups, err := storage.ListGroup(reqObj.TenantID)
	if err != nil {
		return data, err
	}

	for _, user := range users {
		flag := false
		for _, owned := range ownedSliceTemp {
			if owned == user.ID {
				flag = true
				break
			}
		}
		if flag {
			data.Users.OwnedSlice = append(data.Users.OwnedSlice, user)
		} else {
			data.Users.NotOwnedSlice = append(data.Users.NotOwnedSlice, user)
		}
	}

	for _, group := range groups {
		flag := false
		for _, owned := range ownedSliceTemp {
			if owned == group.ID {
				flag = true
				break
			}
		}
		if flag {
			data.Groups.OwnedSlice = append(data.Groups.OwnedSlice, group)
		} else {
			data.Groups.NotOwnedSlice = append(data.Groups.NotOwnedSlice, group)
		}
	}

	return data, err
}
