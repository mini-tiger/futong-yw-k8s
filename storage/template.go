package storage

import (
	"ftk8s/base/cfg"
	"ftk8s/model"

	"gorm.io/gorm/clause"
)

func CreateTemplate(template *model.Template) error {
	err := cfg.Gdb.Create(template).Error
	return err
}

func UpdateTemplate(reqObj *model.ReqUpdateTemplate) error {
	err := cfg.Gdb.Table(model.Template{}.TableName()).
		Where("id = ?", reqObj.TemplateID).
		Updates(model.Template{TemplateName: reqObj.TemplateName, TemplateKind: reqObj.TemplateKind, Content: reqObj.Content, Description: reqObj.Description}).Error
	return err
}

func ReadTemplate(reqObj *model.ReqReadTemplate) (model.Template, error) {
	template := model.Template{}
	err := cfg.Gdb.Table(model.Template{}.TableName()).
		Where("id = ?", reqObj.TemplateID).
		First(&template).Error
	return template, err
}

func DeleteTemplate(reqObj *model.ReqDeleteTemplate) error {
	err := cfg.Gdb.Delete(model.Template{}, "id = ?", reqObj.TemplateID).Error
	return err
}

func ListTemplate(tenantID string) ([]model.Template, error) {
	templates := make([]model.Template, 0)

	err := cfg.Gdb.Table(model.Template{}.TableName()).
		Where("tenant_id = ?", tenantID).Find(&templates).Error
	return templates, err
}

func ListTemplateWithPage(reqObj *model.ReqListTemplate) (int64, []model.Template, error) {
	var dataCount int64
	templates := make([]model.Template, 0)

	qc := cfg.Gdb.Table(model.Template{}.TableName()).Where("tenant_id = ?", reqObj.TenantID)
	err := qc.Count(&dataCount).Error
	if err != nil {
		return dataCount, nil, err
	}

	for k, v := range reqObj.UrlQueryPara {
		if len(v[0]) != 0 {
			if k != "page_num" && k != "page_size" &&
				k != "sort_field" && k != "sort_order" {
				qc = qc.Where(map[string]interface{}{k: v[0]})
			}
		}
	}
	err = qc.
		Order(clause.OrderByColumn{Column: clause.Column{Name: reqObj.SortField}, Desc: reqObj.SortOrderIsDesc}).
		Limit(reqObj.PageSize).Offset(reqObj.SkipNum).
		Find(&templates).Error
	return dataCount, templates, err
}
