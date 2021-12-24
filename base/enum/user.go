package enum

const BuiltinRoot = "builtin_root"
const WebpermTopLevelMenu = "0"

const (
	WebpermForOnlyBuiltinRoot = 1
	WebpermForNormal          = 2
	WebpermIsSelected         = 1
	WebpermNotSelected        = 2
)

const (
	UserTypeTenant = 1
	UserTypeNormal = 2
)

const (
	PrefixUser  = "FTUSER-"
	PrefixGroup = "FTGROUP-"
)

const (
	WebpermTypeDirectory = "M"
	WebpermTypeResource  = "C"
	WebpermTypeButton    = "F"
	WebpermTypeMixed     = "H"
)
