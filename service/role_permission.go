package service

import (
	"ftk8s/model"
	"ftk8s/storage"
)

func UpdateRoleByPermission(reqObj *model.ReqUpdateRoleByPermission) error {
	err := storage.UpdateRoleByPermission(reqObj)
	return err
}

func ReadRoleByPermission(reqObj *model.ReqReadRoleByPermission) (*model.DataOwnedOrNotRole, error) {
	data := new(model.DataOwnedOrNotRole)

	data.OwnedSlice = make([]model.Role, 0)
	data.NotOwnedSlice = make([]model.Role, 0)

	ownedSliceTemp, err := storage.ReadRoleByPermission(reqObj)
	if err != nil {
		return data, err
	}
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

func UpdatePermissionByRole(reqObj *model.ReqUpdatePermissionByRole) error {
	err := storage.UpdatePermissionByRole(reqObj)
	return err
}

func ReadPermissionByRole(reqObj *model.ReqReadPermissionByRole) *model.DataOwnedOrNotPermission {
	data := new(model.DataOwnedOrNotPermission)

	data.OwnedSlice = make([]model.Permission, 0)
	data.NotOwnedSlice = make([]model.Permission, 0)

	ownedSliceTemp := storage.ReadPermissionOwnedByRole(reqObj.RoleAccount, reqObj.TenantID)
	permissions, err := storage.ListPermission()
	if err != nil {
		return data
	}

	for _, permission := range permissions {
		flag := false
		for _, owned := range ownedSliceTemp {
			if owned[2]+owned[3] == permission.PermissionUrl+permission.PermissionAction {
				flag = true
			}
		}
		if flag {
			data.OwnedSlice = append(data.OwnedSlice, permission)
		} else {
			data.NotOwnedSlice = append(data.NotOwnedSlice, permission)
		}
	}

	return data
}
