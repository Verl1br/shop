package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) addToBasket(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	id, err := h.services.Basket.AddToBasket(userId, itemId)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) deleteBasketItem(c *gin.Context) {
	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	err = h.services.Basket.DeleteBasketItem(itemId)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, "ok")
}

func (h *Handler) getBasketItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	items, err := h.services.Basket.GetBasketItems(userId)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, getAllItemsResponse{
		Data: items,
	})
}
