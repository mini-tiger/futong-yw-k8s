package service

import (
	"ftk8s/model"
	"ftk8s/storage"
	"ftk8s/util"
)

func CreateRoleWithPermission(reqObj *model.ReqCreateRoleWithPermission) error {
	role := model.Role{
		TenantID:    reqObj.TenantID,
		RoleAccount: reqObj.RoleAccount,
		RoleName:    reqObj.RoleName,
		Description: reqObj.Description,
	}
	reqObj.Role = role

	err := storage.CreateRoleWithPermission(reqObj)
	return err
}

func UpdateRoleWithPermission(reqObj *model.ReqUpdateRoleWithPermission) error {
	role, err := storage.ReadRole(reqObj.RoleID)
	if err != nil {
		return err
	}
	reqObj.Role = role

	err = storage.UpdateRoleWithPermission(reqObj)
	return err
}

func ReadRole(reqObj *model.ReqReadRole) (model.Role, error) {
	role, err := storage.ReadRole(reqObj.RoleID)
	return role, err
}

func DeleteRole(reqObj *model.ReqDeleteRole) (deleteFlag int, err error) {
	role, err := storage.ReadRole(reqObj.RoleID)
	if err != nil {
		deleteFlag = 2
		return deleteFlag, err
	}
	reqObj.RoleAccount = role.RoleAccount

	// Check if there is associated data
	roleAssObjSlice := storage.ReadUserGroupByRole(reqObj.RoleAccount, reqObj.TenantID)
	permissions := storage.ReadPermissionOwnedByRole(reqObj.RoleAccount, reqObj.TenantID)

	if reqObj.Force != 1 {
		if len(roleAssObjSlice) != 0 || len(permissions) != 0 {
			deleteFlag = 2
			return deleteFlag, err
		}
	}

	err = storage.DeleteRole(reqObj)
	return deleteFlag, err
}

func ListRole(reqObj *model.ReqListRole) (extra util.Extra, roles []model.Role, err error) {
	reqObj.PageNum, reqObj.PageSize, reqObj.SkipNum, reqObj.SortField, reqObj.SortOrder, reqObj.SortOrderIsDesc = util.GetPageInfoAndSortWay(reqObj.PageNum, reqObj.PageSize, reqObj.SortField, reqObj.SortOrder)
	var dataCountTemp int64

	if reqObj.RoleAccount == "" {
		dataCountTemp, roles, err = storage.ListRoleWithPage(reqObj)
		if err != nil {
			return extra, roles, err
		}
	} else {
		dataCountTemp, roles, err = storage.ListRoleByRoleAccountWithPage(reqObj)
		if err != nil {
			return extra, roles, err
		}
	}

	dataCount := int(dataCountTemp)
	pageCount := util.GetPageCount(dataCount, reqObj.PageSize)
	extra = util.Extra{
		PageNum:   reqObj.PageNum,
		PageSize:  reqObj.PageSize,
		SortField: reqObj.SortField,
		SortOrder: reqObj.SortOrder,
		DataCount: dataCount,
		PageCount: pageCount,
	}

	return extra, roles, err
}
