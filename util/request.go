package util

import (
	"fmt"

	"ftk8s/base/enum"
	"ftk8s/model"

	"github.com/gin-gonic/gin"
)

// GetTenantFromContext get tenant from the context of gin
func GetTenantFromContext(c *gin.Context) (*model.Tenant, error) {
	tenantTemp, exists := c.Get(enum.TenantInfo)
	if !exists {
		return nil, fmt.Errorf("failed to get tenant from the context of gin, not found key Tenant-Info")
	}
	tenant, ok := tenantTemp.(*model.Tenant)
	if !ok {
		return nil, fmt.Errorf("failed to get tenant from the context of gin, assertion of *model.Tenant type failed")
	}
	return tenant, nil
}

// GetUserFromContext get user from the context of gin
func GetUserFromContext(c *gin.Context) (*model.User, error) {
	userTemp, exists := c.Get(enum.UserInfo)
	if !exists {
		return nil, fmt.Errorf("failed to get user from the context of gin, not found key User-Info")
	}
	user, ok := userTemp.(*model.User)
	if !ok {
		return nil, fmt.Errorf("failed to get user from the context of gin, assertion of model.User type failed")
	}
	return user, nil
}

// GetTenantIDFromContext get tenantID from the context of gin
func GetTenantIDFromContext(c *gin.Context) (string, error) {
	tenantIDTemp, exists := c.Get(enum.TenantID)
	if !exists {
		return "", fmt.Errorf("failed to get tenantID from the context of gin, not found key(Tenant-ID)")
	}
	tenantID, ok := tenantIDTemp.(string)
	if !ok {
		return "", fmt.Errorf("failed to get tenantID from the context of gin, assertion of string type failed")
	}
	if tenantID == "" {
		return "", fmt.Errorf("failed to get tenantID from the context of gin, the value of key(Tenant-ID) is an empty string")
	}

	return tenantID, nil
}
