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

// Server represents the HTTP server for the item store application.
type Server struct {
	config config.Config
	DB     db.Client
	router *gin.Engine
}

// NewServer initializes a new Server instance with the provided configuration.
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

	router.GET("/items/:key", server.GetItem)
	router.GET("/items", server.ListItem)
	router.POST("/items", server.CreateItem)
	router.DELETE("/items/:key", server.DeleteItem)
	router.PUT("/items/:key", server.UpdateItem)

	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
