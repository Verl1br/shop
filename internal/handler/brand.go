package handler

import (
	"net/http"
	"strconv"

	"github.com/dhevve/shop"
	"github.com/gin-gonic/gin"
)

func (h *Handler) addBrand(c *gin.Context) {
	var input shop.Brand

	if err := c.BindJSON(&input); err != nil {
		return
	}

	id, err := h.services.Brand.AddBrand(input)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

func (h *Handler) deleteBrand(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	err = h.services.Brand.DeleteBrand(id)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, "ok")
}
