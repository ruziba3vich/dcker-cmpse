package repo

import "github.com/ruziba3vich/dckr-cmpse/internal/models"

type IWokerService interface {
	GetRequestAnsResend(*models.GetRequest) *models.Response
}
