package service

import (
	"fmt"

	"ftk8s/ksc"
	"ftk8s/model"
	"ftk8s/storage"
	"ftk8s/util"
)

func CreateTemplate(reqObj *model.ReqCreateTemplate) error {
	err := ksc.ValidateTemplate(reqObj.TemplateKind, reqObj.Content)
	if err != nil {
		return fmt.Errorf("TemplateKind=(%s), wrong resource template format: %w", reqObj.TemplateKind, err)
	}

	template := new(model.Template)

	template.TenantID = reqObj.TenantID
	template.TemplateAccount = reqObj.TemplateAccount
	template.TemplateName = reqObj.TemplateName
	template.TemplateKind = reqObj.TemplateKind
	template.Content = reqObj.Content
	template.Description = reqObj.Description

	err = storage.CreateTemplate(template)
	if err != nil {
		return err
	}

	return nil
}

func UpdateTemplate(reqObj *model.ReqUpdateTemplate) error {
	err := ksc.ValidateTemplate(reqObj.TemplateKind, reqObj.Content)
	if err != nil {
		return fmt.Errorf("TemplateKind=(%s), wrong resource template format: %w", reqObj.TemplateKind, err)
	}

	err = storage.UpdateTemplate(reqObj)
	if err != nil {
		return err
	}

	return nil
}

func ReadTemplate(reqObj *model.ReqReadTemplate) (model.Template, error) {
	template, err := storage.ReadTemplate(reqObj)
	if err != nil {
		return model.Template{}, err
	}

	return template, nil
}

func DeleteTemplate(reqObj *model.ReqDeleteTemplate) error {
	err := storage.DeleteTemplate(reqObj)
	if err != nil {
		return err
	}

	return nil
}

func ListTemplate(reqObj *model.ReqListTemplate) (extra util.Extra, templates []model.Template, err error) {
	reqObj.PageNum, reqObj.PageSize, reqObj.SkipNum, reqObj.SortField, reqObj.SortOrder, reqObj.SortOrderIsDesc = util.GetPageInfoAndSortWay(reqObj.PageNum, reqObj.PageSize, reqObj.SortField, reqObj.SortOrder)
	var dataCountTemp int64

	dataCountTemp, templates, err = storage.ListTemplateWithPage(reqObj)
	if err != nil {
		return extra, templates, err
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

	return extra, templates, err
}
