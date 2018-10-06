package product

import "github.com/salapao2136/rest-api-go/models"

type InterfaceUcase interface {
	Fetch() *[]models.Product
	FetchByID(id string) *models.Product
}
