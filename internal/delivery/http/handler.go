package http

import (
	"github.com/gin-gonic/gin"
	"github.com/onemgvv/storage-service.git/internal/config"
	"github.com/onemgvv/storage-service.git/internal/service"
	"github.com/onemgvv/storage-service.git/pkg/storage"
	"net/http"
)

type Handler struct {
	services *service.Services
	storage  *storage.Storage
}

func NewHandler(services *service.Services, storage *storage.Storage) *Handler {
	return &Handler{
		services: services,
		storage:  storage,
	}
}

func (h *Handler) Init(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
		// limiter
		// cors
	)

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		h.FileHandlerInit(api)
	}
}
