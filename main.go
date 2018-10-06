package main

import (
	"github.com/gin-gonic/gin"
	"github.com/salapao2136/rest-api-go/configs"
	"github.com/salapao2136/rest-api-go/models"
	"github.com/salapao2136/rest-api-go/product"
)

func main() {
	conn := configs.Init()
	conn.AutoMigrate(&models.Product{})
	defer conn.Close()
	repoProduct := product.NewRepoProduct(conn)
	ucaseProduct := product.NewUcaseProduct(repoProduct)
	r := gin.Default()
	product.ProductRegister(r, ucaseProduct)
	r.Run(":3000")
}
