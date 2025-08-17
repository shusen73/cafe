package router

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"

	"cafe/handlers"
)

func New(staticDir string) *gin.Engine {
	return NewRouter(staticDir)
}

func NewRouter(staticDir string) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Recovery())

	// --- APIs (prefix with /api)
	api := r.Group("/api")
	{
		api.GET("/health", handlers.Health)
		api.GET("/menu", handlers.MenuList)
	}

	// --- Static files for the built frontend
	// Serve asset folder (Vite build) at /assets
	r.Static("/assets", filepath.Join(staticDir, "assets"))
	// Serve index at "/"
	r.StaticFile("/", filepath.Join(staticDir, "index.html"))

	// --- SPA fallback (for client-side routing)
	r.NoRoute(func(c *gin.Context) {
		// Let API 404s be API 404s
		if strings.HasPrefix(c.Request.URL.Path, "/api/") {
			c.Status(http.StatusNotFound)
			return
		}
		c.File(filepath.Join(staticDir, "index.html"))
	})

	return r
}

