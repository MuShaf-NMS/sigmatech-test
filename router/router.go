package router

import (
	"net/http"

	"github.com/MuShaf-NMS/sigmatech-test/config"
	"github.com/MuShaf-NMS/sigmatech-test/handler"
	"github.com/MuShaf-NMS/sigmatech-test/middleware"
	"github.com/MuShaf-NMS/sigmatech-test/repository"
	"github.com/MuShaf-NMS/sigmatech-test/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IntializeRouter(server *gin.Engine, db *gorm.DB, config *config.Config) {
	// Define Repositories
	konsumenRepository := repository.NewKonsumenRepository(db)
	authRepository := repository.NewAuthRepository(db)
	limitRepository := repository.NewLimitRepository(db)
	limitTerpakaiRepository := repository.NewLimitTerpakaiRepository(db)
	trsansaksiRepository := repository.NewTransaksiRepository(db)
	fotoRepository := repository.NewFotoRepository(db)
	blacklistTokenRepository := repository.NewBlacklistTokenRepository(db)
	// Define Services
	konsumenService := service.NewKonsumenService(konsumenRepository, authRepository, limitRepository, limitTerpakaiRepository)
	authService := service.NewAuthService(authRepository)
	transaksiService := service.NewTransaksiService(trsansaksiRepository, limitRepository, limitTerpakaiRepository)
	fotoService := service.NewFotoService(fotoRepository)
	jwtService := service.NewJwtService(config.Secret_Key)
	blacklistTokenService := service.NewBlacklistTokenService(blacklistTokenRepository)
	// Define Handlers
	konsumenHandler := handler.NewKonsumenHandler(konsumenService)
	authHandler := handler.NewAuthHandler(authService, jwtService, blacklistTokenService)
	transaksihandler := handler.NewTransaksiHandler(transaksiService)
	fotoHandler := handler.NewFotoHandler(fotoService, config.Assets_Dir, config.Allowed_Ext, 2*1024*1024)
	// Define Middleware
	konsumenMiddleware := middleware.KonsumenMiddleware(jwtService, blacklistTokenService)
	// Define Routes
	konsumenRouter := server.Group("/konsumen")
	{
		konsumenRouter.GET("", konsumenHandler.GetAll)
		konsumenRouter.POST("", konsumenHandler.Create)
		konsumenRouter.GET("/:id", konsumenHandler.GetOne)
		konsumenRouter.PUT("/:id", konsumenHandler.Update)
		konsumenRouter.DELETE("/:id", konsumenHandler.Delete)
	}
	authRouter := server.Group("/auth")
	{
		authRouter.POST("/login", authHandler.Login)
		authRouter.GET("/logout", konsumenMiddleware, authHandler.Logout)
		authRouter.PUT("", konsumenMiddleware, authHandler.Update)
	}
	transaksiRouter := server.Group("/transaksi", konsumenMiddleware)
	{
		transaksiRouter.POST("", transaksihandler.Create)
		transaksiRouter.GET("", transaksihandler.GetAll)
		transaksiRouter.GET("/:id", transaksihandler.GetOne)
	}
	fotoRouter := server.Group("/foto", konsumenMiddleware)
	{
		fotoRouter.GET("", fotoHandler.GetAll)
		fotoRouter.GET("/:ket", fotoHandler.GetOneByKet)
		fotoRouter.POST("/ktp", fotoHandler.CreateKTP)
		fotoRouter.POST("/selfie", fotoHandler.CreateSelfie)
	}
	server.StaticFS("/assets", http.Dir("./assets"))
}
