package internal

import (
	"item-store/config"
	"item-store/internal/db"
	"item-store/internal/db/lvl"
	"item-store/internal/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	config config.Config
	DB     db.Client
	router *gin.Engine
}

func NewServer(cfg config.Config) (*Server, error) {

	db, err := lvl.NewDB(cfg.LvlConfig.Path)
	if err != nil {
		return nil, err
	}

	server := &Server{
		config: cfg,
		DB:     db,
	}

	server.setupRouter()
	return server, nil

}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.Use(gin.Recovery())

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "X-Requested-With"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60,
	}))

	//enable gzip compression
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(utils.HTTPLogger())

	router.GET("/:key", server.GetItem)
	router.GET("/list", server.ListItem)
	router.POST("/put", server.PutItem)
	router.DELETE("/:key", server.DeleteItem)
	router.PUT("/:key", server.UpdateItem)

	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
