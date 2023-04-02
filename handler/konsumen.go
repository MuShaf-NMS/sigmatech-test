package handler

import (
	"strconv"

	"github.com/MuShaf-NMS/sigmatech-test/dto"
	"github.com/MuShaf-NMS/sigmatech-test/helper"
	"github.com/MuShaf-NMS/sigmatech-test/service"
	"github.com/gin-gonic/gin"
)

type KonsumenHandler interface {
	GetAll(ctx *gin.Context)
	Create(ctx *gin.Context)
	GetOne(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type konsumenHandler struct {
	service service.KonsumenService
}

func (kh *konsumenHandler) GetAll(ctx *gin.Context) {
	konsumens, err := kh.service.GetAll()
	if err != nil {
		e := err.(*helper.CustomError)
		res := helper.ErrorResponseBuilder(e.Errors)
		ctx.JSON(e.Code, res)
		return
	}
	res := helper.ResponseBuilder(konsumens)
	ctx.JSON(200, res)
}
func (kh *konsumenHandler) Create(ctx *gin.Context) {
	var konsumen dto.KonsumenCreate
	ctx.BindJSON(&konsumen)
	if err := helper.Validate(konsumen); err != nil {
		errs := helper.ValidationError(err)
		res := helper.ErrorResponseBuilder(errs)
		ctx.JSON(400, res)
		return
	}
	err := kh.service.Create(konsumen)
	if err != nil {
		e := err.(*helper.CustomError)
		res := helper.ErrorResponseBuilder(e.Errors)
		ctx.JSON(e.Code, res)
		return
	}
	res := helper.ResponseBuilder("Berhasil membuat pegawai")
	ctx.JSON(200, res)
}
func (kh *konsumenHandler) GetOne(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		res := helper.ErrorResponseBuilder("Invalid Id")
		ctx.JSON(400, res)
		return
	}
	konsumen, err := kh.service.GetOne(uint(id))
	if err != nil {
		e := err.(*helper.CustomError)
		res := helper.ErrorResponseBuilder(e.Errors)
		ctx.JSON(e.Code, res)
		return
	}
	res := helper.ResponseBuilder(konsumen)
	ctx.JSON(200, res)
}
func (kh *konsumenHandler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		res := helper.ErrorResponseBuilder("Invalid Id")
		ctx.JSON(400, res)
		return
	}
	var konsumen dto.KonsumenUpdate
	ctx.BindJSON(&konsumen)
	if err := helper.Validate(konsumen); err != nil {
		errs := helper.ValidationError(err)
		res := helper.ErrorResponseBuilder(errs)
		ctx.JSON(400, res)
		return
	}
	err = kh.service.Update(konsumen, uint(id))
	if err != nil {
		e := err.(*helper.CustomError)
		res := helper.ErrorResponseBuilder(e.Errors)
		ctx.JSON(e.Code, res)
		return
	}
	res := helper.ResponseBuilder("Berhasil membuat pegawai")
	ctx.JSON(200, res)
}
func (kh *konsumenHandler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		res := helper.ErrorResponseBuilder("Invalid Id")
		ctx.JSON(400, res)
		return
	}
	err = kh.service.Delete(uint(id))
	if err != nil {
		e := err.(*helper.CustomError)
		res := helper.ErrorResponseBuilder(e.Errors)
		ctx.JSON(e.Code, res)
		return
	}
	res := helper.ResponseBuilder("Berhasil memuat pegawai")
	ctx.JSON(200, res)
}

func NewKonsumenHandler(service service.KonsumenService) KonsumenHandler {
	return &konsumenHandler{
		service: service,
	}
}
