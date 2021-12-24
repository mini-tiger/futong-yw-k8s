package enum

// Token header key
const (
	HeaderAuthKey = "X-Auth-Token"
)

// Application running mode
const (
	AppRunModeDebug   = "debug"
	AppRunModeRelease = "release"
	AppRunModeTest    = "test"
)

// Select the user service
const (
	TypeUserServiceSelf = "self"
	TypeUserServiceCmp  = "cmp"
)

// i18n
const (
	I18nEnUS = "en_US"
	I18nZhCN = "zh_CN"
	I18nZhTW = "zh_TW"
)
