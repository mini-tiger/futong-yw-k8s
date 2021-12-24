package enum

// Message code
const (
	// General
	MsgOK   = 200
	MsgFail = -1

	MsgUnauthorized      = 401
	MsgLoginFailed       = 1100
	MsgPermissionDenied  = 1101
	MsgHasAssociatedData = 1102

	// Request parameters
	MsgInvalidParameter = 1001

	// Database
	MsgFailDataInsert = 2001
	MsgFailDataUpdate = 2002
	MsgFailDataQuery  = 2003
	MsgFailDataDelete = 2004

	// Other
	MsgBuiltInCannotBeModified = 3001
	MsgBuiltInCannotBeDeleted  = 3002
	MsgFailProcessFile         = 3011
	MsgFailCreateOpLog         = 3012
)

// Text information corresponding to the Message code
var stringOfMsgCode = map[int]string{
	// General
	MsgOK:   "MsgOK",
	MsgFail: "MsgFail",

	MsgUnauthorized:      "MsgUnauthorized",
	MsgLoginFailed:       "MsgLoginFailed",
	MsgPermissionDenied:  "MsgPermissionDenied",
	MsgHasAssociatedData: "MsgHasAssociatedData",

	// Request parameters
	MsgInvalidParameter: "MsgInvalidParameter",

	// Database
	MsgFailDataInsert: "MsgFailDataInsert",
	MsgFailDataUpdate: "MsgFailDataUpdate",
	MsgFailDataQuery:  "MsgFailDataQuery",
	MsgFailDataDelete: "MsgFailDataDelete",

	// Other
	MsgBuiltInCannotBeModified: "MsgBuiltInCannotBeModified",
	MsgBuiltInCannotBeDeleted:  "MsgBuiltInCannotBeDeleted",
	MsgFailProcessFile:         "MsgFailProcessFile",
	MsgFailCreateOpLog:         "MsgFailCreateOpLog",
}

func GetStringOfMsgCode(code int) string {
	return stringOfMsgCode[code]
}
