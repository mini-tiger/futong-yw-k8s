package service

import (
	"ftk8s/base/enum"
	"ftk8s/model"
	"ftk8s/storage"
	"ftk8s/util"
)

func CreateGroup(reqObj *model.ReqCreateGroup) error {
	group := new(model.Group)

	groupID, err := util.GenerateUUIDV4(enum.PrefixGroup)
	if err != nil {
		return err
	}

	group.ID = groupID
	group.TenantID = reqObj.TenantID
	group.GroupAccount = reqObj.GroupAccount
	group.GroupName = reqObj.GroupName
	group.Description = reqObj.Description

	err = storage.CreateGroup(group)
	if err != nil {
		return err
	}

	return nil
}

func UpdateGroup(reqObj *model.ReqUpdateGroup) error {
	err := storage.UpdateGroup(reqObj)
	if err != nil {
		return err
	}

	return nil
}

func ReadGroup(reqObj *model.ReqReadGroup) (model.Group, error) {
	group, err := storage.ReadGroup(reqObj.GroupID)
	if err != nil {
		return model.Group{}, err
	}

	return group, nil
}

func DeleteGroup(reqObj *model.ReqDeleteGroup) (deleteFlag int, err error) {
	group, err := storage.ReadGroup(reqObj.GroupID)
	if err != nil {
		return deleteFlag, err
	}
	reqObj.GroupAccount = group.GroupAccount

	// Check if there is associated data
	users, err := storage.ReadUserByGroup(group)
	if err != nil {
		deleteFlag = 2
		return deleteFlag, err
	}
	clusters, err := storage.ReadClusterByGroup(group)
	if err != nil {
		deleteFlag = 2
		return deleteFlag, err
	}

	if reqObj.Force != 1 {
		if len(users) != 0 || len(clusters) != 0 {
			deleteFlag = 2
			return deleteFlag, err
		}
	}

	err = storage.DeleteGroup(reqObj)
	return deleteFlag, err
}

func ListGroup(reqObj *model.ReqListGroup) (extra util.Extra, groups []model.Group, err error) {
	reqObj.PageNum, reqObj.PageSize, reqObj.SkipNum, reqObj.SortField, reqObj.SortOrder, reqObj.SortOrderIsDesc = util.GetPageInfoAndSortWay(reqObj.PageNum, reqObj.PageSize, reqObj.SortField, reqObj.SortOrder)
	var dataCountTemp int64

	if reqObj.GroupAccount == "" {
		dataCountTemp, groups, err = storage.ListGroupWithPage(reqObj)
		if err != nil {
			return extra, groups, err
		}
	} else {
		dataCountTemp, groups, err = storage.ListGroupByGroupAccountWithPage(reqObj)
		if err != nil {
			return extra, groups, err
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

	return extra, groups, err
}
