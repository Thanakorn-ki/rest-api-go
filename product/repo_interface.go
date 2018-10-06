package product

import "github.com/salapao2136/rest-api-go/models"

type InterfaceRepo interface {
	Fetch() *[]models.Product
	GetByID(id string) *models.Product
}
