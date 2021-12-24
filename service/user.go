package service

import (
	"fmt"
	"time"

	"ftk8s/base/cfg"
	"ftk8s/base/enum"
	"ftk8s/model"
	"ftk8s/storage"
	"ftk8s/util"

	"github.com/dgrijalva/jwt-go"
)

func RegisterTenant(reqObj *model.ReqRegisterTenant) error {
	tenant := new(model.Tenant)

	tenant.ID = reqObj.TenantID
	tenant.TenantName = reqObj.TenantName
	tenant.Email = reqObj.Email
	tenant.Description = reqObj.Description

	salt := util.GetRandomString(10)
	passwordHashed := util.EncodePassword(reqObj.Password, salt)
	tenant.Salt = salt
	tenant.Password = passwordHashed

	err := storage.RegisterTenant(tenant)
	if err != nil {
		return err
	}

	return nil
}

func Login(reqObj *model.ReqLogin) (string, error) {

	nowTime := time.Now()
	var token string
	jwtInfoObj := new(model.JwtInfo)

	if reqObj.UserType == enum.UserTypeTenant {
		tenant, err := storage.GetTenantByTenantID(reqObj.TenantID)
		if err != nil {
			return "", err
		}
		if tenant.Password == "" || tenant.Salt == "" {
			return "", fmt.Errorf("this tenant has no password or salt")
		}
		passwordHashed := util.EncodePassword(reqObj.Password, tenant.Salt)
		if passwordHashed != tenant.Password {
			return "", fmt.Errorf("failed to login, tenant_id or password error")
		}

		tenant.LastLoginIP = reqObj.ClientIP
		tenant.LastLoginTime = &nowTime

		// generate JWT token
		jwtInfoObj.UserType = enum.UserTypeTenant
		jwtInfoObj.TenantID = tenant.ID
		jwtInfoObj.UserAccount = ""
		tokenObj := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
			"iss": "ftk8s",
			"iat": nowTime.Unix(),
			"exp": nowTime.Add(time.Duration(cfg.AppConfObj.TokenExpTime) * time.Second).Unix(),
			"ft":  jwtInfoObj,
		})
		token, err = tokenObj.SignedString(cfg.RsaPrivateKey)
		if err != nil {
			return "", fmt.Errorf("failed to generate JWT token, error message: %s", err.Error())
		}

		// Update tenant login info
		err = storage.UpdateTenantLoginInfo(tenant)
		if err != nil {
			return "", err
		}

	} else if reqObj.UserType == enum.UserTypeNormal {
		user, err := storage.GetUserByUserAccount(reqObj.TenantID, reqObj.UserAccount)
		if err != nil {
			return "", err
		}
		if user.Password == "" || user.Salt == "" {
			return "", fmt.Errorf("this user has no password or salt")
		}
		passwordHashed := util.EncodePassword(reqObj.Password, user.Salt)
		if passwordHashed != user.Password {
			return "", fmt.Errorf("failed to login, user_id or password error")
		}

		user.LastLoginIP = reqObj.ClientIP
		user.LastLoginTime = &nowTime

		// generate JWT token
		jwtInfoObj.UserType = enum.UserTypeNormal
		jwtInfoObj.TenantID = user.TenantID
		jwtInfoObj.UserAccount = user.UserAccount
		tokenObj := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
			"iss": "ftk8s",
			"iat": nowTime.Unix(),
			"exp": nowTime.Add(time.Duration(cfg.AppConfObj.TokenExpTime) * time.Second).Unix(),
			"ft":  jwtInfoObj,
		})
		token, err = tokenObj.SignedString(cfg.RsaPrivateKey)
		if err != nil {
			return "", fmt.Errorf("failed to generate JWT token, error message: %s", err.Error())
		}

		// Update user login info
		err = storage.UpdateUserLoginInfo(user)
		if err != nil {
			return "", err
		}
	}

	return token, nil
}

func CreateUser(reqObj *model.ReqCreateUser) error {
	user := new(model.User)

	userID, err := util.GenerateUUIDV4(enum.PrefixUser)
	if err != nil {
		return err
	}

	user.ID = userID
	user.TenantID = reqObj.TenantID
	user.UserAccount = reqObj.UserAccount
	user.Username = reqObj.Username
	user.Email = reqObj.Email
	user.Description = reqObj.Description

	salt := util.GetRandomString(10)
	passwordHashed := util.EncodePassword(reqObj.Password, salt)
	user.Salt = salt
	user.Password = passwordHashed

	err = storage.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(reqObj *model.ReqUpdateUser) error {
	err := storage.UpdateUser(reqObj)
	if err != nil {
		return err
	}

	return nil
}

func ResetPassword(reqObj *model.ReqResetPassword) error {
	salt := util.GetRandomString(10)
	passwordHashed := util.EncodePassword(reqObj.Password, salt)
	reqObj.Salt = salt
	reqObj.Password = passwordHashed

	err := storage.ResetPassword(reqObj)
	if err != nil {
		return err
	}

	return nil
}

func ReadUser(reqObj *model.ReqReadUser) (model.User, error) {
	user, err := storage.ReadUser(reqObj.UserID)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func DeleteUser(reqObj *model.ReqDeleteUser) error {
	err := storage.DeleteUser(reqObj)
	if err != nil {
		return err
	}

	return nil
}

func ListUser(reqObj *model.ReqListUser) (extra util.Extra, users []model.User, err error) {
	reqObj.PageNum, reqObj.PageSize, reqObj.SkipNum, reqObj.SortField, reqObj.SortOrder, reqObj.SortOrderIsDesc = util.GetPageInfoAndSortWay(reqObj.PageNum, reqObj.PageSize, reqObj.SortField, reqObj.SortOrder)
	var dataCountTemp int64

	if reqObj.UserAccount == "" {
		dataCountTemp, users, err = storage.ListUserWithPage(reqObj)
		if err != nil {
			return extra, users, err
		}
	} else {
		dataCountTemp, users, err = storage.ListUserByUserAccountWithPage(reqObj)
		if err != nil {
			return extra, users, err
		}
	}

	dataCount := int(dataCountTemp)
	pageCount := util.GetPageCount(dataCount, reqObj.PageSize)
	extra = util.Extra{
		PageNum:   reqObj.PageNum,
		PageSize:  reqObj.PageSize,
		SortField: reqObj.SortField,
		SortOrder: reqObj.SortOrder,
		DataCount: dataCount,
		PageCount: pageCount,
	}

	return extra, users, err
}
