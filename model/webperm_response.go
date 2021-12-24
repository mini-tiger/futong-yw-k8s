package model

type RespWebpermTree struct {
	ID            string             `json:"id"`
	ParentID      string             `json:"parentId"`
	Name          string             `json:"name"`
	Path          string             `json:"path"`
	ResourcesSort int                `json:"resourcesSort"`
	ResourcesType string             `json:"resourcesType"`
	Title         string             `json:"title"`
	Icon          string             `json:"icon"`
	Display       int                `json:"display"`
	Children      []*RespWebpermTree `json:"children"`
}
