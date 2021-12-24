package model

// JwtInfo save user information to jwt
type JwtInfo struct {
	UserType    int
	TenantID    string
	UserAccount string
}
