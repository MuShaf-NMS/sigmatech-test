package handler

import (
	"github.com/MuShaf-NMS/sigmatech-test/dto"
	"github.com/MuShaf-NMS/sigmatech-test/helper"
	"github.com/MuShaf-NMS/sigmatech-test/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type AuthHandler interface {
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type authHandler struct {
	service               service.AuthService
	jwtService            service.JwtService
	blacklistTokenService service.BlacklistTokenService
}

func (ah *authHandler) Login(ctx *gin.Context) {
	var login dto.Auth
	ctx.BindJSON(&login)
	if err := helper.Validate(login); err != nil {
		errs := helper.ValidationError(err)
		res := helper.ErrorResponseBuilder(errs)
		ctx.JSON(400, res)
		return
	}
	auth, err := ah.service.Login(login)
	if err != nil {
		e := err.(*helper.CustomError)
		res := helper.ErrorResponseBuilder(e.Errors)
		ctx.JSON(e.Code, res)
		return
	}
	token := ah.jwtService.GenerateToken(auth.KonsumenId)
	res := gin.H{
		"token": token,
	}
	ctx.JSON(200, res)
}
func (ah *authHandler) Logout(ctx *gin.Context) {
	if getTokenClaim, ok := ctx.Get("tokenClaim"); ok {
		if tokenClaim, ok := getTokenClaim.(jwt.MapClaims); ok {
			ah.blacklistTokenService.CreateBlacklistToken(tokenClaim["jti"].(string))
			res := helper.ResponseBuilder(nil)
			ctx.JSON(200, res)
			return
		}
		res := helper.ErrorResponseBuilder("Unauthorized")
		ctx.JSON(401, res)
		return
	}
	res := helper.ErrorResponseBuilder("Unauthorized")
	ctx.JSON(401, res)
}
func (ah *authHandler) Update(ctx *gin.Context) {
	if getKonsumenId, ok := ctx.Get("konsumenId"); ok {
		fKonsumenId, _ := getKonsumenId.(float64)
		konsumenId := uint(fKonsumenId)
		var auth dto.UpdateAuth
		ctx.BindJSON(auth)
		if err := helper.Validate(auth); err != nil {
			errs := helper.ValidationError(err)
			res := helper.ErrorResponseBuilder(errs)
			ctx.JSON(400, res)
			return
		}
		err := ah.service.Update(auth, konsumenId)
		if err != nil {
			e := err.(*helper.CustomError)
			res := helper.ErrorResponseBuilder(e.Errors)
			ctx.JSON(e.Code, res)
			return
		}
		res := helper.ResponseBuilder(nil)
		ctx.JSON(200, res)
		return
	}
	res := helper.ErrorResponseBuilder("Unauthorized")
	ctx.JSON(401, res)
}

func NewAuthHandler(service service.AuthService, jwtService service.JwtService, blacklistTokenService service.BlacklistTokenService) AuthHandler {
	return &authHandler{
		service:               service,
		jwtService:            jwtService,
		blacklistTokenService: blacklistTokenService,
	}
}
