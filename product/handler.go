package product

import (
	"github.com/gin-gonic/gin"
)

type HttpProductHandler struct {
	InterfaceUcase
}

func (h *HttpProductHandler) FetchProducts(c *gin.Context) {
	result := h.Fetch()
	c.JSON(200, result)
}

func (h *HttpProductHandler) FetchProduct(c *gin.Context) {
	productId := c.Param("id")
	result := h.FetchByID(productId)
	c.JSON(200, result)
}
func ProductRegister(router *gin.Engine, ucaseProduct InterfaceUcase) {
	handler := &HttpProductHandler{ucaseProduct}
	router.GET("/product", handler.FetchProducts)
	router.GET("/product/:id", handler.FetchProduct)
}
