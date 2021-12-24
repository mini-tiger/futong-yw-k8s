package service

import (
	"ftk8s/model"
	"ftk8s/storage"
)

func UpdateUserByGroup(reqObj *model.ReqUpdateUserByGroup) error {
	err := storage.UpdateUserByGroup(reqObj)
	if err != nil {
		return err
	}

	return nil
}

func ReadUserByGroup(reqObj *model.ReqReadUserByGroup) (*model.DataOwnedOrNotUser, error) {
	data := new(model.DataOwnedOrNotUser)

	data.OwnedSlice = make([]model.User, 0)
	data.NotOwnedSlice = make([]model.User, 0)

	group, err := storage.ReadGroup(reqObj.GroupID)
	if err != nil {
		return nil, err
	}
	ownedSliceTemp, err := storage.ReadUserByGroup(group)
	if err != nil {
		return nil, err
	}
	users, err := storage.ListUser(group.TenantID)
	if err != nil {
		return data, err
	}

	data.OwnedSlice = ownedSliceTemp
	for _, user := range users {
		flag := false
		for _, owned := range ownedSliceTemp {
			if owned.ID == user.ID {
				flag = true
				break
			}
		}
		if !flag {
			data.NotOwnedSlice = append(data.NotOwnedSlice, user)
		}
	}

	return data, err
}

func UpdateGroupByUser(reqObj *model.ReqUpdateGroupByUser) error {
	err := storage.UpdateGroupByUser(reqObj)
	if err != nil {
		return err
	}

	return nil
}

func ReadGroupByUser(reqObj *model.ReqReadGroupByUser) (*model.DataOwnedOrNotGroup, error) {
	data := new(model.DataOwnedOrNotGroup)

	data.OwnedSlice = make([]model.Group, 0)
	data.NotOwnedSlice = make([]model.Group, 0)

	user, err := storage.ReadUser(reqObj.UserID)
	if err != nil {
		return nil, err
	}
	ownedSliceTemp, err := storage.ReadGroupByUser(user)
	if err != nil {
		return data, err
	}
	groups, err := storage.ListGroup(user.TenantID)
	if err != nil {
		return data, err
	}

	data.OwnedSlice = ownedSliceTemp
	for _, group := range groups {
		flag := false
		for _, owned := range ownedSliceTemp {
			if owned.ID == group.ID {
				flag = true
				break
			}
		}
		if !flag {
			data.NotOwnedSlice = append(data.NotOwnedSlice, group)
		}
	}

	return data, err
}
