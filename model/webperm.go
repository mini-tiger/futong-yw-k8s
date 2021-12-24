package model

import (
	"time"
)

type Webperm struct {
	ID              string     `gorm:"column:id; type:varchar(150); primaryKey; comment:主键,前端权限ID"`
	ParentID        string     `gorm:"column:parent_id; type:varchar(150); comment:前端权限父级ID"`
	Name            string     `gorm:"column:name; type:varchar(191); uniqueIndex; not null; comment:前端权限名称"`
	Path            string     `gorm:"column:path; type:varchar(191); not null; comment:前端权限路径"`
	ResourcesSort   int        `gorm:"column:resources_sort; not null; comment:排序顺序,1最前,越大越往后"`
	ResourcesType   string     `gorm:"column:resources_type; type:varchar(10); not null; comment:权限类型,M:目录,C:资源,F:按钮,H:混合"`
	Title           string     `gorm:"column:title; type:varchar(150); not null; comment:前端权限显示名称"`
	Icon            string     `gorm:"column:icon; type:varchar(150); not null; comment:前端权限显示图片"`
	Display         int        `gorm:"column:display; not null; comment:是否显示,1:显示,2:隐藏"`
	OnlyBuiltinRoot int        `gorm:"column:only_builtin_root; not null; comment:是否只是内置后台管理用户操作,1:是,2:否"`
	CreateTime      *time.Time `gorm:"column:create_time; autoCreateTime; comment:创建时间"`
	UpdateTime      *time.Time `gorm:"column:update_time; autoUpdateTime; comment:更新时间"`

	// Role and Webperm association
	RoleWebperm []*Role `gorm:"many2many:role_webperm_ass; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Webperm) TableName() string {
	return "webperm"
}

type RoleWebpermAss struct {
	RoleID    int    `gorm:"column:role_id;"`
	WebpermID string `gorm:"column:webperm_id;"`
}

func (RoleWebpermAss) TableName() string {
	return "role_webperm_ass"
}
