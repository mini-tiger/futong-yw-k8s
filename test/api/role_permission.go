package api

import (
	"net/http"

	"github.com/gavv/httpexpect/v2"
)

// UpdateRoleByPermission
func UpdateRoleByPermission(e *httpexpect.Expect) {}

// ReadRoleByPermission
func ReadRoleByPermission(e *httpexpect.Expect) {}

// UpdatePermissionByRole
func UpdatePermissionByRole(e *httpexpect.Expect) {}

// ReadPermissionByRole
func ReadPermissionByRole(e *httpexpect.Expect) {
	e.GET("/api/system/permission-role").
		WithHeader(headerAuthKey, headerAuthValue).
		WithQuery("role_account", "role-admin").
		Expect().Status(http.StatusOK).
		JSON().Object().ValueEqual("code", 200)
}
