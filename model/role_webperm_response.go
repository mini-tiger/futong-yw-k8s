package model

type RespWebpermTreeByRoleOnlyDisplay struct {
	ID            string                              `json:"id"`
	ParentID      string                              `json:"parentId"`
	Name          string                              `json:"name"`
	Path          string                              `json:"path"`
	ResourcesSort int                                 `json:"resourcesSort"`
	ResourcesType string                              `json:"resourcesType"`
	Title         string                              `json:"title"`
	Icon          string                              `json:"icon"`
	Display       int                                 `json:"display"`
	Selected      int                                 `json:"selected"`
	Children      []*RespWebpermTreeByRoleOnlyDisplay `json:"children"`
}

type RespWebpermFunsAndRights struct {
	Funs   []string             `json:"funs"`
	Rights []*RespWebpermRights `json:"rights"`
}
type RespWebpermRights struct {
	ID            string                `json:"id"`
	ParentID      string                `json:"parentId"`
	Name          string                `json:"name"`
	Path          string                `json:"path"`
	ResourcesSort int                   `json:"resourcesSort"`
	ResourcesType string                `json:"resourcesType"`
	Meta          RespWebpermRightsMeta `json:"meta"`
	Children      []*RespWebpermRights  `json:"children"`
}
type RespWebpermRightsMeta struct {
	Title string `json:"title"`
	Icon  string `json:"icon"`
}
