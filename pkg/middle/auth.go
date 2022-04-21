package middle

import (
	"strings"

	"github.com/ztalab/ZAManager/pconst"
	jwtGet "github.com/ztalab/ZAManager/pkg/jwt"
	"github.com/ztalab/ZAManager/pkg/response"
	"github.com/ztalab/ZAManager/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code = pconst.CODE_COMMON_OK
		var token string
		Authorization := ctx.Request.Header.Get("Authorization")
		if Authorization == "" {
			jwtToken := util.GetCookieFromGin(ctx, "jwt-token")
			if len(jwtToken) > 0 {
				token = jwtToken
			}
		} else {
			tokenSli := strings.Split(Authorization, " ")
			if len(tokenSli) != 2 {
				code = pconst.CODE_COMMON_USER_NO_LOGIN
			} else {
				token = tokenSli[1]
			}
		}
		if len(token) != 0 {
			tokenClaims, err := jwtGet.ParseTokenV2(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = pconst.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
					goto out
				default:
					code = pconst.ERROR_AUTH_CHECK_TOKEN_FAIL
					goto out
				}
			} else if !tokenClaims.Valid {
				code = pconst.ERROR_AUTH
				goto out
			}
			ctx.Set("claims.username", jwtGet.GetIDFromClaims("username", tokenClaims.Claims))
			ctx.Set("claims.id", jwtGet.GetIDFromClaims("id", tokenClaims.Claims))
			ctx.Set("claims.uuid", jwtGet.GetIDFromClaims("uuid", tokenClaims.Claims))
		} else {
			code = pconst.CODE_COMMON_USER_NO_LOGIN
		}
	out:
		if code != pconst.CODE_COMMON_OK {
			res := map[string]interface{}{
				"code":    code,
				"data":    nil,
				"message": response.GetResponseMsg(ctx, code),
			}
			ctx.JSON(response.GetResponseCode(code), res)
			ctx.Abort()
		}
		ctx.Next()
	}
}
