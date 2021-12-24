package service

import (
	"fmt"
	"sort"

	"ftk8s/base/enum"
	"ftk8s/model"
	"ftk8s/storage"
	"ftk8s/util"
)

func CreateWebperm(reqObj *model.ReqCreateWebperm) error {
	webperm := new(model.Webperm)

	webpermID, err := util.GenerateUUIDV4("")
	if err != nil {
		return err
	}

	webperms, err := storage.ListWebperm()
	if err != nil {
		return err
	}
	flag := false
	for _, webperm := range webperms {
		if reqObj.ParentID == webperm.ID {
			if reqObj.Display == 1 && webperm.Display == 2 {
				return fmt.Errorf("error display")
			}
			flag = true
			break
		} else if reqObj.ParentID == enum.WebpermTopLevelMenu {
			flag = true
			break
		}
	}

	if !flag {
		return fmt.Errorf("not found parent_id")
	}

	webperm.ID = webpermID
	webperm.ParentID = reqObj.ParentID
	webperm.Name = reqObj.Name
	webperm.Path = reqObj.Path
	webperm.ResourcesSort = reqObj.ResourcesSort
	webperm.ResourcesType = reqObj.ResourcesType
	webperm.Title = reqObj.Title
	webperm.Icon = reqObj.Icon
	webperm.Display = reqObj.Display
	webperm.OnlyBuiltinRoot = enum.WebpermForNormal
	err = storage.CreateWebperm(webperm)
	if err != nil {
		return err
	}

	return nil
}

func UpdateWebperm(reqObj *model.ReqUpdateWebperm) error {
	err := storage.UpdateWebperm(reqObj)
	if err != nil {
		return err
	}

	return nil
}

func DeleteWebperm(reqObj *model.ReqDeleteWebperm) error {
	err := storage.DeleteWebperm(reqObj.ID)
	return err
}

func ListWebpermTree() ([]*model.RespWebpermTree, error) {
	result := make([]*model.RespWebpermTree, 0)
	webperms, err := storage.ListWebperm()
	if err != nil {
		return result, err
	}

	respWebpermTrees := make([]*model.RespWebpermTree, 0)
	for _, v := range webperms {
		respWebpermTree := new(model.RespWebpermTree)
		respWebpermTree.ID = v.ID
		respWebpermTree.ParentID = v.ParentID
		respWebpermTree.Name = v.Name
		respWebpermTree.Path = v.Path
		respWebpermTree.ResourcesSort = v.ResourcesSort
		respWebpermTree.ResourcesType = v.ResourcesType
		respWebpermTree.Title = v.Title
		respWebpermTree.Icon = v.Icon
		respWebpermTree.Display = v.Display
		children := make([]*model.RespWebpermTree, 0)
		respWebpermTree.Children = children

		respWebpermTrees = append(respWebpermTrees, respWebpermTree)
	}
	sort.Slice(respWebpermTrees, func(i, j int) bool {
		return respWebpermTrees[i].ResourcesSort < respWebpermTrees[j].ResourcesSort
	})

	roots := make([]*model.RespWebpermTree, 0)
	for _, v := range respWebpermTrees {
		if v.ParentID == enum.WebpermTopLevelMenu {
			roots = append(roots, v)
		}
	}

	for _, root := range roots {
		util.GenerateTreeWebpermTree(root, respWebpermTrees)
		result = append(result, root)
	}
	sort.Slice(roots, func(i, j int) bool {
		return roots[i].ResourcesSort < roots[j].ResourcesSort
	})

	return result, err
}

func ListWebpermTreeBindRole() ([]*model.RespWebpermTree, error) {
	result := make([]*model.RespWebpermTree, 0)
	webperms, err := storage.ListWebperm()
	if err != nil {
		return result, err
	}

	respWebpermTrees := make([]*model.RespWebpermTree, 0)
	for _, v := range webperms {
		respWebpermTree := new(model.RespWebpermTree)
		respWebpermTree.ID = v.ID
		respWebpermTree.ParentID = v.ParentID
		respWebpermTree.Name = v.Name
		respWebpermTree.Path = v.Path
		respWebpermTree.ResourcesSort = v.ResourcesSort
		respWebpermTree.ResourcesType = v.ResourcesType
		respWebpermTree.Title = v.Title
		respWebpermTree.Icon = v.Icon
		respWebpermTree.Display = v.Display
		children := make([]*model.RespWebpermTree, 0)
		respWebpermTree.Children = children

		respWebpermTrees = append(respWebpermTrees, respWebpermTree)
	}
	sort.Slice(respWebpermTrees, func(i, j int) bool {
		return respWebpermTrees[i].ResourcesSort < respWebpermTrees[j].ResourcesSort
	})

	roots := make([]*model.RespWebpermTree, 0)
	for _, v := range respWebpermTrees {
		if v.ParentID == enum.WebpermTopLevelMenu {
			roots = append(roots, v)
		}
	}

	for _, root := range roots {
		if root.Name != "k8s_systemSetting" {
			util.GenerateTreeWebpermTree(root, respWebpermTrees)
			result = append(result, root)
		}
	}
	sort.Slice(roots, func(i, j int) bool {
		return roots[i].ResourcesSort < roots[j].ResourcesSort
	})

	return result, err
}
