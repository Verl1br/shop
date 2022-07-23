package handler

import (
	"github.com/dhevve/shop/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{services: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sing-up", h.signUp)
		auth.POST("/sing-in", h.signIn)
	}

	item := router.Group("/item", h.userIdentity)
	{
		item.POST("/", h.createItem)
		item.GET("/", h.getAllItems)
		item.GET("/:id", h.getItemById)
		item.DELETE("/:id", h.deleteItem)
		item.PUT("/:id", h.updateItem)
	}

	basket := router.Group("/basket", h.userIdentity)
	{
		basket.POST("/:id", h.addToBasket)
		basket.GET("/", h.getBasketItems)
		basket.DELETE("/:id", h.deleteBasketItem)
	}

	brand := router.Group("/brand", h.userIdentity)
	{
		brand.POST("/", h.addBrand)
		brand.DELETE("/:id", h.deleteBrand)
	}

	info := router.Group("/item-info", h.userIdentity)
	{
		info.POST("/", h.createItemInfo)
		info.GET("/:id", h.getItemInfo)
		info.DELETE("/:id", h.deleteItemInfo)
		info.PUT("/:id", h.updateItemInfo)
	}

	return router
}
