package middle

import (
	"ftk8s/base/enum"

	"github.com/gin-gonic/gin"
)

const HeaderFieldLang = "lang"

// HandLang unify the lang field of the request
func HandLang() gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := c.Query("lang")
		if lang != enum.I18nEnUS && lang != enum.I18nZhCN && lang != enum.I18nZhTW {
			lang = enum.I18nZhCN
		}
		c.Request.Header.Set(HeaderFieldLang, lang)
		c.Next()
	}
}
