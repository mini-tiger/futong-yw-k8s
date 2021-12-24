package service

import (
	"sort"

	"ftk8s/base/enum"
	"ftk8s/model"
	"ftk8s/storage"
	"ftk8s/util"
)

func UpdateWebpermTreeByRole(reqObj *model.ReqUpdateWebpermTreeByRole) error {
	err := storage.UpdateWebpermByRole(reqObj)
	return err
}

func ReadWebpermTreeByRole(reqObj *model.ReqReadWebpermTreeByRole) ([]*model.RespWebpermTreeByRoleOnlyDisplay, error) {
	respWebpermTrees := make([]*model.RespWebpermTreeByRoleOnlyDisplay, 0)

	switch reqObj.RoleID {
	// create role
	case -1:
		webperms, err := storage.ListWebpermForTenantOnlyDisplay()
		if err != nil {
			return respWebpermTrees, err
		}

		for _, v := range webperms {
			respWebpermTree := new(model.RespWebpermTreeByRoleOnlyDisplay)
			respWebpermTree.ID = v.ID
			respWebpermTree.ParentID = v.ParentID
			respWebpermTree.Name = v.Name
			respWebpermTree.Path = v.Path
			respWebpermTree.ResourcesSort = v.ResourcesSort
			respWebpermTree.ResourcesType = v.ResourcesType
			respWebpermTree.Title = v.Title
			respWebpermTree.Icon = v.Icon
			respWebpermTree.Display = v.Display
			respWebpermTree.Selected = enum.WebpermNotSelected
			children := make([]*model.RespWebpermTreeByRoleOnlyDisplay, 0)
			respWebpermTree.Children = children

			respWebpermTrees = append(respWebpermTrees, respWebpermTree)
		}
		sort.Slice(respWebpermTrees, func(i, j int) bool {
			return respWebpermTrees[i].ResourcesSort < respWebpermTrees[j].ResourcesSort
		})

		roots := make([]*model.RespWebpermTreeByRoleOnlyDisplay, 0)
		for _, v := range respWebpermTrees {
			if v.ParentID == enum.WebpermTopLevelMenu {
				roots = append(roots, v)
			}
		}

		result := make([]*model.RespWebpermTreeByRoleOnlyDisplay, 0)
		for _, root := range roots {
			if root.Name != "k8s_systemSetting" {
				util.GenerateTreeWebpermTreeByRoleOnlyDisplay(root, respWebpermTrees)
				result = append(result, root)
			}
		}
		sort.Slice(result, func(i, j int) bool {
			return result[i].ResourcesSort < result[j].ResourcesSort
		})

		return result, err

	default:
		role, err := storage.ReadRole(reqObj.RoleID)
		if err != nil {
			return respWebpermTrees, err
		}

		ownedWebperms, err := storage.ReadWebpermByRole(role)
		if err != nil {
			return respWebpermTrees, err
		}
		allWebperms, err := storage.ListWebpermForTenantOnlyDisplay()
		if err != nil {
			return respWebpermTrees, err
		}

		for _, a := range allWebperms {
			selected := enum.WebpermNotSelected
			for _, o := range ownedWebperms {
				if a.ID == o.ID {
					selected = enum.WebpermIsSelected
					break
				}
			}

			respWebpermTree := new(model.RespWebpermTreeByRoleOnlyDisplay)
			respWebpermTree.ID = a.ID
			respWebpermTree.ParentID = a.ParentID
			respWebpermTree.Name = a.Name
			respWebpermTree.Path = a.Path
			respWebpermTree.ResourcesSort = a.ResourcesSort
			respWebpermTree.ResourcesType = a.ResourcesType
			respWebpermTree.Title = a.Title
			respWebpermTree.Icon = a.Icon
			respWebpermTree.Display = a.Display
			respWebpermTree.Selected = selected
			children := make([]*model.RespWebpermTreeByRoleOnlyDisplay, 0)
			respWebpermTree.Children = children

			respWebpermTrees = append(respWebpermTrees, respWebpermTree)
		}
		sort.Slice(respWebpermTrees, func(i, j int) bool {
			return respWebpermTrees[i].ResourcesSort < respWebpermTrees[j].ResourcesSort
		})

		roots := make([]*model.RespWebpermTreeByRoleOnlyDisplay, 0)
		for _, v := range respWebpermTrees {
			if v.ParentID == enum.WebpermTopLevelMenu {
				roots = append(roots, v)
			}
		}

		result := make([]*model.RespWebpermTreeByRoleOnlyDisplay, 0)
		for _, root := range roots {
			if root.Name != "k8s_systemSetting" {
				util.GenerateTreeWebpermTreeByRoleOnlyDisplay(root, respWebpermTrees)
				result = append(result, root)
			}
		}
		sort.Slice(result, func(i, j int) bool {
			return result[i].ResourcesSort < result[j].ResourcesSort
		})

		return result, err
	}
}

func ListWebpermFunsAndRights(reqObj *model.ReqListWebpermFunsAndRights) (*model.RespWebpermFunsAndRights, error) {
	result := new(model.RespWebpermFunsAndRights)
	result.Funs = make([]string, 0)
	nodeSliceTemp := make([]*model.RespWebpermRights, 0)

	switch reqObj.UserType {
	case enum.UserTypeTenant:
		if reqObj.TenantID == enum.BuiltinRoot {
			webperms, err := storage.ListWebperm()
			if err != nil {
				return nil, err
			}
			for _, webperm := range webperms {
				if webperm.ResourcesType != enum.WebpermTypeButton {
					node := new(model.RespWebpermRights)
					node.ID = webperm.ID
					node.ParentID = webperm.ParentID
					node.Name = webperm.Name
					node.Path = webperm.Path
					node.ResourcesSort = webperm.ResourcesSort
					node.ResourcesType = webperm.ResourcesType
					node.Meta.Title = webperm.Title
					node.Meta.Icon = webperm.Icon
					children := make([]*model.RespWebpermRights, 0)
					node.Children = children
					nodeSliceTemp = append(nodeSliceTemp, node)
				}
			}

		} else {
			webperms, err := storage.ListWebpermForTenantOnlyDisplay()
			if err != nil {
				return nil, err
			}
			for _, webperm := range webperms {
				if webperm.ResourcesType != enum.WebpermTypeButton {
					node := new(model.RespWebpermRights)
					node.ID = webperm.ID
					node.ParentID = webperm.ParentID
					node.Name = webperm.Name
					node.Path = webperm.Path
					node.ResourcesSort = webperm.ResourcesSort
					node.ResourcesType = webperm.ResourcesType
					node.Meta.Title = webperm.Title
					node.Meta.Icon = webperm.Icon
					children := make([]*model.RespWebpermRights, 0)
					node.Children = children
					nodeSliceTemp = append(nodeSliceTemp, node)
				}
			}
		}
	case enum.UserTypeNormal:
		groups, err := storage.ReadGroupByUser(reqObj.User)
		if err != nil {
			return nil, err
		}

		roleAccountsTemp := storage.ReadRoleByUserGroup(reqObj.User.ID, reqObj.TenantID)
		for _, group := range groups {
			tempData := storage.ReadRoleByUserGroup(group.ID, reqObj.TenantID)
			roleAccountsTemp = append(roleAccountsTemp, tempData...)
		}
		roleAccounts := make([]string, 0)
		tempMap := map[string]struct{}{}
		for _, v := range roleAccountsTemp {
			if _, ok := tempMap[v]; !ok {
				tempMap[v] = struct{}{}
				roleAccounts = append(roleAccounts, v)
			}
		}

		for _, roleAccount := range roleAccounts {
			role, err := storage.ReadRoleByRoleAccount(reqObj.TenantID, roleAccount)
			if err != nil {
				return nil, err
			}
			webperms, err := storage.ReadWebpermByRole(role)
			if err != nil {
				return nil, err
			}

			for _, webperm := range webperms {
				if webperm.ResourcesType != enum.WebpermTypeButton {
					node := new(model.RespWebpermRights)
					node.ID = webperm.ID
					node.ParentID = webperm.ParentID
					node.Name = webperm.Name
					node.Path = webperm.Path
					node.ResourcesSort = webperm.ResourcesSort
					node.ResourcesType = webperm.ResourcesType
					node.Meta.Title = webperm.Title
					node.Meta.Icon = webperm.Icon
					children := make([]*model.RespWebpermRights, 0)
					node.Children = children
					nodeSliceTemp = append(nodeSliceTemp, node)
				}
			}
		}
	}

	nodeMap := make(map[string]*model.RespWebpermRights)
	nodeSlice := make([]*model.RespWebpermRights, 0)
	for _, node := range nodeSliceTemp {
		nodeMap[node.ID] = node
	}
	for _, node := range nodeMap {
		nodeSlice = append(nodeSlice, node)
	}
	sort.Slice(nodeSlice, func(i, j int) bool {
		return nodeSlice[i].ResourcesSort < nodeSlice[j].ResourcesSort
	})

	roots := make([]*model.RespWebpermRights, 0)
	for _, v := range nodeSlice {
		if v.ParentID == enum.WebpermTopLevelMenu {
			roots = append(roots, v)
		}
		if v.ResourcesType == enum.WebpermTypeButton || v.ResourcesType == enum.WebpermTypeMixed {
			result.Funs = append(result.Funs, v.Name)
		}
	}

	rights := make([]*model.RespWebpermRights, 0)
	for _, root := range roots {
		util.GenerateTreeRoleWebpermRightsTree(root, nodeSlice)
		rights = append(rights, root)
	}
	sort.Slice(rights, func(i, j int) bool {
		return rights[i].ResourcesSort < rights[j].ResourcesSort
	})

	result.Rights = rights
	return result, nil
}
