package storage

import (
	"time"

	"github.com/ruziba3vich/dckr-cmpse/internal/models"
)

type Storage struct{}

func (s Storage) GetRequestAnsResend(req *models.GetRequest) *models.Response {
	response := req.Convert()
	response.ResentAt = time.Now()
	return response
}

func New() Storage {
	return Storage{}
}
