package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/dckr-cmpse/internal/models"
	"github.com/ruziba3vich/dckr-cmpse/internal/repo"
)

type Handler struct {
	storage repo.IWokerService
}

func (h *Handler) WorkerHandler(c *gin.Context) {
	var req models.GetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	req.ReceivedAt = time.Now()
	response := h.storage.GetRequestAnsResend(&req)
	c.IndentedJSON(http.StatusOK, gin.H{"response": response})
}

func New(storage repo.IWokerService) *Handler {
	return &Handler{
		storage: storage,
	}
}
