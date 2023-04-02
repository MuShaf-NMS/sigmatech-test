package handler

import (
	"fmt"
	"strings"

	"github.com/MuShaf-NMS/sigmatech-test/helper"
	"github.com/MuShaf-NMS/sigmatech-test/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FotoHandler interface {
	GetAll(ctx *gin.Context)
	CreateKTP(ctx *gin.Context)
	CreateSelfie(ctx *gin.Context)
	GetOneByKet(ctx *gin.Context)
}

type fotoHandler struct {
	service    service.FotoService
	assetsDir  string
	allowedExt []string
	maxSize    int64
}

func (fh *fotoHandler) GetAll(ctx *gin.Context) {
	if getKonsumenId, ok := ctx.Get("konsumenId"); ok {
		fKonsumenId, _ := getKonsumenId.(float64)
		konsumenId := uint(fKonsumenId)
		fotos, err := fh.service.GetAll(konsumenId)
		if err != nil {
			e := err.(*helper.CustomError)
			res := helper.ErrorResponseBuilder(e.Errors)
			ctx.JSON(e.Code, res)
			return
		}
		res := helper.ResponseBuilder(fotos)
		ctx.JSON(200, res)
		return
	}
	res := helper.ErrorResponseBuilder("Unauthorized")
	ctx.JSON(401, res)
}
func (fh *fotoHandler) CreateKTP(ctx *gin.Context) {
	if getKonsumenId, ok := ctx.Get("konsumenId"); ok {
		fKonsumenId, _ := getKonsumenId.(float64)
		konsumenId := uint(fKonsumenId)
		_, header, err := ctx.Request.FormFile("foto")
		if err != nil {
			res := helper.ErrorResponseBuilder("Failed to upload image")
			ctx.JSON(400, res)
			return
		}
		uuid, err := uuid.NewRandom()
		if err != nil {
			res := helper.ErrorResponseBuilder("Failed to upload image")
			ctx.JSON(400, res)
			return
		}
		if header.Size > fh.maxSize {
			res := helper.ErrorResponseBuilder("File size too big")
			ctx.JSON(400, res)
			return
		}
		splitName := strings.Split(header.Filename, ".")
		ext := splitName[len(splitName)-1]
		allowed := helper.Contains(fh.allowedExt, ext)
		if !allowed {
			res := helper.ErrorResponseBuilder("File not allowed")
			ctx.JSON(400, res)
			return
		}
		newname := fmt.Sprintf("%s/%s.%s", "ktp", uuid.String(), ext)
		path := fmt.Sprintf("%s/%s", fh.assetsDir, newname)
		if err := ctx.SaveUploadedFile(header, path); err != nil {
			res := helper.ErrorResponseBuilder("Failed to upload image")
			ctx.JSON(400, res)
			return
		}
		foto, err := fh.service.Create(konsumenId, newname, "ktp")
		if err != nil {
			e := err.(*helper.CustomError)
			res := helper.ErrorResponseBuilder(e.Errors)
			ctx.JSON(e.Code, res)
			return
		}
		res := helper.ResponseBuilder(gin.H{"url": foto.Url})
		ctx.JSON(200, res)
		return
	}
	fmt.Println("6")
	res := helper.ErrorResponseBuilder("Unauthorized")
	ctx.JSON(401, res)
}
func (fh *fotoHandler) CreateSelfie(ctx *gin.Context) {
	if getKonsumenId, ok := ctx.Get("konsumenId"); ok {
		fKonsumenId, _ := getKonsumenId.(float64)
		konsumenId := uint(fKonsumenId)
		_, header, err := ctx.Request.FormFile("foto")
		if err != nil {
			res := helper.ErrorResponseBuilder("Failed to upload image")
			ctx.JSON(400, res)
			return
		}
		uuid, err := uuid.NewRandom()
		if err != nil {
			res := helper.ErrorResponseBuilder("Failed to upload image")
			ctx.JSON(400, res)
			return
		}
		if header.Size > fh.maxSize {
			res := helper.ErrorResponseBuilder("File size too big")
			ctx.JSON(400, res)
			return
		}
		splitName := strings.Split(header.Filename, ".")
		ext := splitName[len(splitName)-1]
		allowed := helper.Contains(fh.allowedExt, ext)
		if !allowed {
			res := helper.ErrorResponseBuilder("File not allowed")
			ctx.JSON(400, res)
			return
		}
		newname := fmt.Sprintf("%s/%s.%s", "selfie", uuid.String(), ext)
		path := fmt.Sprintf("%s/%s", fh.assetsDir, newname)
		if err := ctx.SaveUploadedFile(header, path); err != nil {
			res := helper.ErrorResponseBuilder("Failed to upload image")
			ctx.JSON(400, res)
			return
		}
		foto, err := fh.service.Create(konsumenId, newname, "selfie")
		if err != nil {
			e := err.(*helper.CustomError)
			res := helper.ErrorResponseBuilder(e.Errors)
			ctx.JSON(e.Code, res)
			return
		}
		res := helper.ResponseBuilder(gin.H{"url": foto.Url})
		ctx.JSON(200, res)
		return
	}
	fmt.Println("yo")
	res := helper.ErrorResponseBuilder("Unauthorized")
	ctx.JSON(401, res)
}
func (fh *fotoHandler) GetOneByKet(ctx *gin.Context) {
	if getKonsumenId, ok := ctx.Get("konsumenId"); ok {
		fKonsumenId, _ := getKonsumenId.(float64)
		konsumenId := uint(fKonsumenId)
		ket := ctx.Param("ket")
		fotos, err := fh.service.GetOneByKet(konsumenId, ket)
		if err != nil {
			e := err.(*helper.CustomError)
			res := helper.ErrorResponseBuilder(e.Errors)
			ctx.JSON(e.Code, res)
			return
		}
		res := helper.ResponseBuilder(fotos)
		ctx.JSON(200, res)
		return
	}
	res := helper.ErrorResponseBuilder("Unauthorized")
	ctx.JSON(401, res)
}

func NewFotoHandler(service service.FotoService, assetsDir string, allowedExt []string, maxSize int64) FotoHandler {
	return &fotoHandler{
		service:    service,
		assetsDir:  assetsDir,
		allowedExt: allowedExt,
		maxSize:    maxSize,
	}
}
