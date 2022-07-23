package handler

import (
	"net/http"
	"strconv"

	"github.com/dhevve/shop"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createItemInfo(c *gin.Context) {
	var item shop.ItemInfo

	if err := c.BindJSON(&item); err != nil {
		return
	}

	id, err := h.services.ItemInfo.Create(item)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getItemInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	item, err := h.services.ItemInfo.GetInfo(id)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) deleteItemInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	err = h.services.ItemInfo.DeleteInfo(id)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, "ok")
}

func (h *Handler) updateItemInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	var input shop.ItemInfoUpdateInput

	if err := c.BindJSON(&input); err != nil {
		return
	}

	err = h.services.ItemInfo.UpdateInfo(input, id)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, "ok")
}
