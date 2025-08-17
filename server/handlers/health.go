package handlers

import "github.com/gin-gonic/gin"

// Health responds with a simple JSON payload to indicate the server is alive.
func Health(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}
