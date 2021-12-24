package storage

import (
	"ftk8s/base/cfg"
	"ftk8s/model"
)

// GetPermissionByPermissionID
func GetPermissionByPermissionID(permissionID string) (model.Permission, error) {
	permission := model.Permission{}
	err := cfg.Gdb.Where("id = ?", permissionID).First(&permission).Error
	return permission, err
}

func ListPermission() ([]model.Permission, error) {
	permissions := make([]model.Permission, 0)
	err := cfg.Gdb.Table(model.Permission{}.TableName()).
		Find(&permissions).Error
	return permissions, err
}
