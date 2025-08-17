package handlers

import (
	"net/http"

	"cafe/db"
	"github.com/gin-gonic/gin"
)

// MenuList returns the public menu (read-only).
func MenuList(c *gin.Context) {
	items := db.Default().ListMenu()
	c.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}
