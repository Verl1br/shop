package handler

import (
	"net/http"
	"strconv"

	"github.com/dhevve/shop"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createItem(c *gin.Context) {
	var input shop.Item

	if err := c.BindJSON(&input); err != nil {
		return
	}

	id, err := h.services.Item.CreateItem(input)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllItemsResponse struct {
	Data []shop.Item `json:"data"`
}

func (h *Handler) getAllItems(c *gin.Context) {
	items, err := h.services.Item.GetAllItems()
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, getAllItemsResponse{
		Data: items,
	})
}

func (h *Handler) getItemById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	item, err := h.services.Item.GetById(id)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) deleteItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	err = h.services.Item.DeleteItem(id)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, "ok")
}

func (h *Handler) updateItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	var input shop.ItemUpdateInput

	if err := c.BindJSON(&input); err != nil {
		return
	}

	err = h.services.Item.UpdateItem(input, id)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, "ok")
}
