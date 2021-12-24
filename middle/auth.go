package middle

import (
	"net/http"
	"strings"

	"ftk8s/base/cfg"
	"ftk8s/base/enum"
	"ftk8s/model"
	"ftk8s/storage"
	"ftk8s/util"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

const HeaderFieldUserType = "UserType"

// Auth verify that the user is logged in
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		langReq := c.Request.Header.Get(HeaderFieldLang)

		authString := c.Request.Header.Get(enum.HeaderAuthKey)
		kv := strings.Split(authString, " ")
		if len(kv) != 2 || kv[0] != "Bearer" {
			cfg.Mlog.Error("invalid Authorization: ", authString)
			c.JSON(http.StatusUnauthorized, util.Re(langReq, enum.MsgUnauthorized, "invalid Authorization"))
			c.Abort()
			return
		}

		token := kv[1]
		tokenObj, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			// since we only use the one private key to sign the tokens,
			// we also only use its public counter part to verify
			return cfg.RsaPublicKey, nil
		})
		if err != nil || !tokenObj.Valid {
			cfg.Mlog.Error("invalid token: ", token)
			c.JSON(http.StatusUnauthorized, util.Re(langReq, enum.MsgUnauthorized, "invalid token"))
			c.Abort()
			return
		}

		// 由于tokenObj是jwt.Parse生成的，已经确保了类型，故下面这个断言没有判断ok
		// 如果无法确定类型的时候，务必写ok来判断，否则会panic
		claimObj := tokenObj.Claims.(jwt.MapClaims)
		jwtInfoTemp, ok := claimObj["ft"].(map[string]interface{})
		if !ok {
			cfg.Mlog.Error("invalid [ft] field of token: ", claimObj)
			c.JSON(http.StatusUnauthorized, util.Re(langReq, enum.MsgUnauthorized, "invalid [ft] field of token"))
			c.Abort()
			return
		}

		jwtInfoObj := new(model.JwtInfo)
		jwtInfoByte, err := jsoniter.Marshal(jwtInfoTemp)
		if err != nil {
			cfg.Mlog.Error("failed to json marshal, error message: ", err.Error())
			c.JSON(http.StatusUnauthorized, util.Re(langReq, enum.MsgUnauthorized, "invalid [ft] field of token"))
			c.Abort()
			return
		}
		err = jsoniter.Unmarshal(jwtInfoByte, jwtInfoObj)
		if err != nil {
			cfg.Mlog.Error("failed to json unmarshal, error message: ", err.Error())
			c.JSON(http.StatusUnauthorized, util.Re(langReq, enum.MsgUnauthorized, "invalid [ft] field of token"))
			c.Abort()
			return
		}

		if jwtInfoObj.UserType == enum.UserTypeTenant {
			tenant, err := storage.GetTenantByTenantID(jwtInfoObj.TenantID)
			if err != nil {
				c.JSON(http.StatusUnauthorized, util.Re(langReq, enum.MsgUnauthorized, "the tenant does not exist"))
				c.Abort()
				return
			}
			c.Set(HeaderFieldUserType, enum.UserTypeTenant)
			c.Set(enum.TenantInfo, tenant)
			c.Set(enum.TenantID, tenant.ID)

		} else if jwtInfoObj.UserType == enum.UserTypeNormal {
			user, err := storage.GetUserByUserAccount(jwtInfoObj.TenantID, jwtInfoObj.UserAccount)
			if err != nil {
				c.JSON(http.StatusUnauthorized, util.Re(langReq, enum.MsgUnauthorized, "the user does not exist"))
				c.Abort()
				return
			}
			c.Set(HeaderFieldUserType, enum.UserTypeNormal)
			c.Set(enum.UserInfo, user)
			c.Set(enum.TenantID, user.TenantID)
		}

		c.Next()
	}
}
