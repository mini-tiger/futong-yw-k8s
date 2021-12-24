package model

// The content of this file:
// the structure of the request parameters associated with API

type ReqUpdateUserByGroup struct {
	GroupID     string   `json:"group_id" binding:"required"`
	UserIDSlice []string `json:"user_id_slice" binding:"required"`
}

type ReqReadUserByGroup struct {
	GroupID string `form:"group_id" binding:"required"`
}

type ReqUpdateGroupByUser struct {
	UserID       string   `json:"user_id" binding:"required"`
	GroupIDSlice []string `json:"group_id_slice" binding:"required"`
}

type ReqReadGroupByUser struct {
	UserID string `form:"user_id" binding:"required"`
}
