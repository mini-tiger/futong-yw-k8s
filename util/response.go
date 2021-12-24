package util

import (
	"ftk8s/base/cfg"
	"ftk8s/base/enum"
)

type Res struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type ResWithExtra struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`

	Extra Extra `json:"extra"`
}
type Extra struct {
	PageNum   int    `json:"page_num"`
	PageSize  int    `json:"page_size"`
	SortField string `json:"sort_field"`
	SortOrder string `json:"sort_order"`
	DataCount int    `json:"data_count"`
	PageCount int    `json:"page_count"`
}

// Re custom content of response
func Re(lang string, code int, data interface{}) Res {
	res := Res{
		Code:    code,
		Message: cfg.GetI18nMsg(lang, enum.GetStringOfMsgCode(code)),
		Data:    data,
	}
	return res
}

// ReOK general content of ok response
func ReOK(lang string, data interface{}) Res {
	res := Res{
		Code:    enum.MsgOK,
		Message: cfg.GetI18nMsg(lang, enum.GetStringOfMsgCode(enum.MsgOK)),
		Data:    data,
	}
	return res
}

// ReOKExtra general content with extra of ok response
func ReOKExtra(lang string, data interface{}, extra Extra) ResWithExtra {
	resWithExtra := ResWithExtra{
		Code:    enum.MsgOK,
		Message: cfg.GetI18nMsg(lang, enum.GetStringOfMsgCode(enum.MsgOK)),
		Data:    data,
		Extra:   extra,
	}
	return resWithExtra
}

// ReFail general content of fail response
func ReFail(lang string, data interface{}) Res {
	res := Res{
		Code:    enum.MsgFail,
		Message: cfg.GetI18nMsg(lang, enum.GetStringOfMsgCode(enum.MsgFail)),
		Data:    data,
	}
	return res
}
