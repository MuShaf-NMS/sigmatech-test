package handler

import (
	"strconv"

	"github.com/MuShaf-NMS/sigmatech-test/dto"
	"github.com/MuShaf-NMS/sigmatech-test/helper"
	"github.com/MuShaf-NMS/sigmatech-test/service"
	"github.com/gin-gonic/gin"
)

type TransaksiHandler interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetOne(ctx *gin.Context)
}

type transaksiHandler struct {
	service service.TransaksiService
}

func (th *transaksiHandler) Create(ctx *gin.Context) {
	if getKonsumenId, ok := ctx.Get("konsumenId"); ok {
		fKonsumenId, _ := getKonsumenId.(float64)
		konsumenId := uint(fKonsumenId)
		var transaksi dto.Transaksi
		ctx.BindJSON(&transaksi)
		if err := helper.Validate(transaksi); err != nil {
			errs := helper.ValidationError(err)
			res := helper.ErrorResponseBuilder(errs)
			ctx.JSON(400, res)
			return
		}
		err := th.service.Create(transaksi, konsumenId)
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
func (th *transaksiHandler) GetAll(ctx *gin.Context) {
	if getKonsumenId, ok := ctx.Get("konsumenId"); ok {
		fKonsumenId, _ := getKonsumenId.(float64)
		konsumenId := uint(fKonsumenId)
		var transaksi dto.Transaksi
		ctx.BindJSON(&transaksi)
		if err := helper.Validate(transaksi); err != nil {
			errs := helper.ValidationError(err)
			res := helper.ErrorResponseBuilder(errs)
			ctx.JSON(400, res)
			return
		}
		transaksis, err := th.service.GetAll(konsumenId)
		if err != nil {
			e := err.(*helper.CustomError)
			res := helper.ErrorResponseBuilder(e.Errors)
			ctx.JSON(e.Code, res)
			return
		}
		res := helper.ResponseBuilder(transaksis)
		ctx.JSON(200, res)
		return
	}
	res := helper.ErrorResponseBuilder("Unauthorized")
	ctx.JSON(401, res)
}
func (th *transaksiHandler) GetOne(ctx *gin.Context) {
	if getKonsumenId, ok := ctx.Get("konsumenId"); ok {
		fKonsumenId, _ := getKonsumenId.(float64)
		konsumenId := uint(fKonsumenId)
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			res := helper.ErrorResponseBuilder("invalid id")
			ctx.JSON(400, res)
			return
		}
		var transaksi dto.Transaksi
		ctx.BindJSON(&transaksi)
		if err := helper.Validate(transaksi); err != nil {
			errs := helper.ValidationError(err)
			res := helper.ErrorResponseBuilder(errs)
			ctx.JSON(400, res)
			return
		}
		transaksis, err := th.service.GetOne(konsumenId, uint(id))
		if err != nil {
			e := err.(*helper.CustomError)
			res := helper.ErrorResponseBuilder(e.Errors)
			ctx.JSON(e.Code, res)
			return
		}
		res := helper.ResponseBuilder(transaksis)
		ctx.JSON(200, res)
		return
	}
	res := helper.ErrorResponseBuilder("Unauthorized")
	ctx.JSON(401, res)
}

func NewTransaksiHandler(service service.TransaksiService) TransaksiHandler {
	return &transaksiHandler{
		service: service,
	}
}
