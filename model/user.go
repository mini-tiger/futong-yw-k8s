package model

import "time"

// UserType: Tenant=1
type Tenant struct {
	ID          string `json:"id" gorm:"column:id; type:varchar(150); primaryKey; comment:主键,租户ID"`
	TenantName  string `json:"tenant_name" gorm:"column:tenant_name; type:varchar(150); not null; comment:租户名称"`
	Password    string `json:"password" gorm:"column:password; type:varchar(150); comment:密码"`
	Salt        string `json:"salt" gorm:"column:salt; type:varchar(200); comment:密码加密盐"`
	Email       string `json:"email" gorm:"column:email; type:varchar(150); default:''; comment:电子邮箱"`
	Description string `json:"description" gorm:"column:description; type:text; comment:用户介绍"`

	LastLoginTime *time.Time `json:"last_login_time" gorm:"column:last_login_time; comment:最近一次登录时间"`
	LastLoginIP   string     `json:"last_login_ip" gorm:"column:last_login_ip; type:varchar(100); comment:最近一次登录IP"`

	CreateTime *time.Time `json:"create_time" gorm:"column:create_time; autoCreateTime; comment:创建时间"`
	UpdateTime *time.Time `json:"update_time" gorm:"column:update_time; autoUpdateTime; comment:更新时间"`
}

func (Tenant) TableName() string {
	return "tenant"
}

// UserType: User=2
type User struct {
	ID          string `json:"id" gorm:"column:id; type:varchar(150); primaryKey; comment:主键,用户ID"`
	TenantID    string `json:"tenant_id" gorm:"column:tenant_id; type:varchar(150); uniqueIndex:ui_tenant_id_user_account; not null; comment:租户ID"`
	UserAccount string `json:"user_account"  gorm:"column:user_account; type:varchar(150); uniqueIndex:ui_tenant_id_user_account; not null; comment:用户账号"`
	Username    string `json:"username" gorm:"column:username; type:varchar(150); not null; comment:用户名称"`
	Password    string `json:"password" gorm:"column:password; type:varchar(150); comment:密码"`
	Salt        string `json:"salt" gorm:"column:salt; type:varchar(200); comment:密码加密盐"`
	Email       string `json:"email" gorm:"column:email; type:varchar(150); default:''; comment:电子邮箱"`
	Description string `json:"description" gorm:"column:description; type:text; comment:用户介绍"`

	LastLoginTime *time.Time `json:"last_login_time" gorm:"column:last_login_time; comment:最近一次登录时间"`
	LastLoginIP   string     `json:"last_login_ip" gorm:"column:last_login_ip; type:varchar(100); comment:最近一次登录IP"`

	CreateTime *time.Time `json:"create_time" gorm:"column:create_time; autoCreateTime; comment:创建时间"`
	UpdateTime *time.Time `json:"update_time" gorm:"column:update_time; autoUpdateTime; comment:更新时间"`

	// User and Tenant association
	UserTenant Tenant `gorm:"foreignKey:TenantID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// User and Group association
	UserGroup []*Group `gorm:"many2many:user_group_ass; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (User) TableName() string {
	return "user"
}
