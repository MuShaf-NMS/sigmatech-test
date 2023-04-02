package middleware

import (
	"fmt"
	"strings"

	"github.com/MuShaf-NMS/sigmatech-test/helper"
	"github.com/MuShaf-NMS/sigmatech-test/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func KonsumenMiddleware(jwtService service.JwtService, blacklistTokenService service.BlacklistTokenService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			fmt.Println("1")
			res := helper.ErrorResponseBuilder("Unauthorized")
			ctx.AbortWithStatusJSON(401, res)
			return
		}
		splitAuth := strings.Split(auth, " ")
		if len(splitAuth) < 2 {
			fmt.Println("2")
			res := helper.ErrorResponseBuilder("Unauthorized")
			ctx.AbortWithStatusJSON(401, res)
			return
		}
		if splitAuth[0] != "Bearer" {
			fmt.Println("3")
			res := helper.ErrorResponseBuilder("Unauthorized")
			ctx.AbortWithStatusJSON(401, res)
			return
		}
		tokenString := splitAuth[1]
		token, err := jwtService.ExtractToken(tokenString)
		if err != nil || !token.Valid {
			fmt.Println("4")
			res := helper.ErrorResponseBuilder("Unauthorized")
			ctx.AbortWithStatusJSON(401, res)
			return
		}
		claim := token.Claims.(jwt.MapClaims)
		if blacklistTokenService.CheckBlacklistToken(claim["jti"].(string)) {
			fmt.Println("5")
			res := helper.ErrorResponseBuilder("Unauthorized")
			ctx.AbortWithStatusJSON(401, res)
			return
		}
		ctx.Set("konsumenId", claim["identity"])
		ctx.Set("tokenClaim", claim)
		ctx.Next()
	}
}
