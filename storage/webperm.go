package storage

import (
	"fmt"

	"ftk8s/base/cfg"
	"ftk8s/model"
	"ftk8s/util"

	"gorm.io/gorm"
)

func CreateWebperm(webperm *model.Webperm) error {
	err := cfg.Gdb.Create(webperm).Error
	return err
}

func UpdateWebperm(reqObj *model.ReqUpdateWebperm) error {
	err := cfg.Gdb.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		// return any error will rollback

		// Verify that the display field is correct
		webperms, err := ListWebperm()
		if err != nil {
			return err
		}
		for _, webperm := range webperms {
			if reqObj.ParentID == webperm.ID {
				if reqObj.Display == 1 && webperm.Display == 2 {
					return fmt.Errorf("error display")
				}
			}
		}

		// Update the specified object data
		err = tx.Table(model.Webperm{}.TableName()).
			Where("id = ?", reqObj.ID).
			Updates(map[string]interface{}{
				"parent_id":      reqObj.ParentID,
				"name":           reqObj.Name,
				"path":           reqObj.Path,
				"resources_sort": reqObj.ResourcesSort,
				"resources_type": reqObj.ResourcesType,
				"title":          reqObj.Title,
				"icon":           reqObj.Icon,
				"display":        reqObj.Display,
			}).Error
		if err != nil {
			return err
		}

		// Find all children
		allChildren := make([]*model.Webperm, 0)
		util.GenerateWebpermSlice(reqObj.ID, webperms, &allChildren)

		// Update the display field of all data
		for _, v := range allChildren {
			err := tx.Table(model.Webperm{}.TableName()).
				Where("id = ?", v.ID).
				Updates(model.Webperm{Display: reqObj.Display}).Error
			if err != nil {
				return err
			}
		}

		// return nil will commit the whole transaction
		return nil
	})
	return err
}

func DeleteWebperm(webpermID string) error {
	err := cfg.Gdb.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		// return any error will rollback

		// Find all children
		webperms, err := ListWebperm()
		if err != nil {
			return err
		}
		allChildren := make([]*model.Webperm, 0)
		util.GenerateWebpermSlice(webpermID, webperms, &allChildren)

		// Delete all eligible data
		if err := tx.Delete(model.Webperm{}, "id = ?", webpermID).Error; err != nil {
			return err
		}
		for _, v := range allChildren {
			if err := tx.Delete(model.Webperm{}, "id = ?", v.ID).Error; err != nil {
				return err
			}
		}

		// return nil will commit the whole transaction
		return nil
	})
	return err
}

func ListWebperm() ([]*model.Webperm, error) {
	webperms := make([]*model.Webperm, 0)

	err := cfg.Gdb.Table(model.Webperm{}.TableName()).Find(&webperms).Error
	return webperms, err
}

func ListWebpermForTenantOnlyDisplay() ([]*model.Webperm, error) {
	webperms := make([]*model.Webperm, 0)

	err := cfg.Gdb.Table(model.Webperm{}.TableName()).
		Where("display = 1 and only_builtin_root = 2").
		Find(&webperms).Error
	return webperms, err
}
