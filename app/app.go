package app

import (
	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/dckr-cmpse/app/handler"
)

type APP struct {
	handler *handler.Handler
}

func (a *APP) RUN(port string) error {
	router := gin.Default()
	router.POST("/send", a.handler.WorkerHandler)
	return router.Run(":" + port)
}

func New(handler *handler.Handler) *APP {
	return &APP{
		handler: handler,
	}
}
