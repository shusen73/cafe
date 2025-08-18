package handlers

import (
	"net/http"
	"strconv"

	"cafe/db"
	"github.com/gin-gonic/gin"
)

// GetMenu returns active menu items (public).
func GetMenu(c *gin.Context) {
	items, err := db.GetAllMenuItems(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load menu"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": items})
}

type menuUpsert struct {
	Name        string `json:"name" binding:"required"`
	PriceCents  int64  `json:"priceCents" binding:"required"`
	Description string `json:"description"`
	Active      *bool  `json:"active"` // optional, default true on create
}

func CreateMenu(c *gin.Context) {
	var req menuUpsert
	if err := c.ShouldBindJSON(&req); err != nil || req.PriceCents < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	active := true
	if req.Active != nil {
		active = *req.Active
	}
	m := db.MenuItem{
		Name:        req.Name,
		PriceCents:  req.PriceCents,
		Description: req.Description,
		Active:      active,
	}
	if err := db.CreateMenuItem(c, &m); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "create failed"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"item": m})
}

func UpdateMenu(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad id"})
		return
	}
	var req menuUpsert
	if err := c.ShouldBindJSON(&req); err != nil || req.PriceCents < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	active := true
	if req.Active != nil {
		active = *req.Active
	}
	m := db.MenuItem{
		ID:          id,
		Name:        req.Name,
		PriceCents:  req.PriceCents,
		Description: req.Description,
		Active:      active,
	}
	if err := db.UpdateMenuItem(c, &m); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"item": m})
}

func DeleteMenu(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad id"})
		return
	}
	if err := db.DeleteMenuItem(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete failed"})
		return
	}
	c.Status(http.StatusNoContent)
}
