package service

import (
	"ftk8s/model"
	"ftk8s/storage"
)

func ListPermission() ([]model.Permission, error) {
	permissions, err := storage.ListPermission()
	return permissions, err
}
