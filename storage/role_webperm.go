package storage

import (
	"ftk8s/base/cfg"
	"ftk8s/model"

	"gorm.io/gorm"
)

func UpdateWebpermByRole(reqObj *model.ReqUpdateWebpermTreeByRole) error {
	err := cfg.Gdb.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		// return any error will rollback

		// Delete role and webperm ass in table role_webperm_ass
		if err := tx.Delete(model.RoleWebpermAss{}, "role_id = ?", reqObj.RoleID).Error; err != nil {
			return err
		}

		// Add role and webperm ass in table role_webperm_ass
		roleWebpermAssSlice := make([]model.RoleWebpermAss, 0)
		for _, webpermID := range reqObj.WebpermIDSlice {
			roleWebpermAss := model.RoleWebpermAss{RoleID: reqObj.RoleID, WebpermID: webpermID}
			roleWebpermAssSlice = append(roleWebpermAssSlice, roleWebpermAss)
		}
		if len(roleWebpermAssSlice) != 0 {
			if err := tx.Create(&roleWebpermAssSlice).Error; err != nil {
				return err
			}
		}

		// return nil will commit the whole transaction
		return nil
	})
	return err
}

func ReadWebpermByRole(role model.Role) ([]model.Webperm, error) {
	webperms := make([]model.Webperm, 0)

	err := cfg.Gdb.Model(&role).Association("RoleWebperm").Find(&webperms)
	return webperms, err
}
