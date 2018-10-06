package product

import (
	"github.com/jinzhu/gorm"
	"github.com/salapao2136/rest-api-go/models"
)

type repoProduct struct {
	Conn *gorm.DB
}

func NewRepoProduct(Conn *gorm.DB) InterfaceRepo {
	return &repoProduct{Conn}
}

func (repo repoProduct) Fetch() *[]models.Product {
	product := &[]models.Product{}
	repo.Conn.Model(&models.Product{}).Find(product)
	return product
}

func (repo repoProduct) GetByID(id string) *models.Product {
	product := &models.Product{}
	repo.Conn.Model(&models.Product{}).Find(product, id)
	return product
}
