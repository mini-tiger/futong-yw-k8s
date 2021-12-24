package model

// The content of this file:
// the structure of the request parameters associated with API

type ReqCreateWebperm struct {
	ParentID      string `json:"parentId" binding:"required"`
	Name          string `json:"name" binding:"required"`
	Path          string `json:"path" binding:"required"`
	ResourcesSort int    `json:"resourcesSort" binding:"required,gt=0"`
	ResourcesType string `json:"resourcesType" binding:"oneof=M C F H"`
	Title         string `json:"title" binding:"required"`
	Icon          string `json:"icon"`
	Display       int    `json:"display" binding:"oneof=1 2"`
}

type ReqUpdateWebperm struct {
	ID            string `json:"id" binding:"required"`
	ParentID      string `json:"parentId" binding:"required"`
	Name          string `json:"name" binding:"required"`
	Path          string `json:"path" binding:"required"`
	ResourcesSort int    `json:"resourcesSort" binding:"required,gt=0"`
	ResourcesType string `json:"resourcesType" binding:"oneof=M C F H"`
	Title         string `json:"title" binding:"required"`
	Icon          string `json:"icon"`
	Display       int    `json:"display" binding:"oneof=1 2"`
}

type ReqDeleteWebperm struct {
	ID string `json:"id" binding:"required"`
}
