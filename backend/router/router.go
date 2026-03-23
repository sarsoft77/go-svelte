package router

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/example/myapp/handlers"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	workDir, _ := os.Getwd()
	absPath := filepath.Join(workDir, "../frontend/dist")
	r.Static("/assets", absPath+"/assets")
	r.LoadHTMLGlob(absPath + "/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/api/hello", handlers.Hello)
	r.GET("/api/products", handlers.GetProducts)
	r.POST("/api/products", handlers.CreateProduct)
	r.PUT("/api/products/:id", handlers.UpdateProduct)
	r.DELETE("/api/products/:id", handlers.DeleteProduct)

	return r
}
