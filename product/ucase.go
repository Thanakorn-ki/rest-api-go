package product

import "github.com/salapao2136/rest-api-go/models"

type ucaseProduct struct {
	repoProduct InterfaceRepo
}

func NewUcaseProduct(rp InterfaceRepo) InterfaceUcase {
	return &ucaseProduct{repoProduct: rp}
}

func (ucase ucaseProduct) Fetch() *[]models.Product {
	result := ucase.repoProduct.Fetch()
	return result
}

func (ucase ucaseProduct) FetchByID(id string) *models.Product {
	result := ucase.repoProduct.GetByID(id)
	return result
}
