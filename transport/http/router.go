package http

import (
	"UEBA/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RouterImpl struct {
	UEBAController *UEBAController
	server         *gin.Engine
	logger         *zap.Logger
	url            string
}

func NewRouter(cfg *config.Config, logger *zap.Logger, UEBAController *UEBAController) *RouterImpl {
	return &RouterImpl{url: cfg.LocalURL, UEBAController: UEBAController, logger: logger}
}

func (r *RouterImpl) RegisterRoutes() {
	router := gin.Default()

	router.GET("/get-items", func(c *gin.Context) {
		r.UEBAController.GetItemsByID(c)
	})

	r.server = router
}

func (r *RouterImpl) Start() error {
	return r.server.Run(r.url)
}
