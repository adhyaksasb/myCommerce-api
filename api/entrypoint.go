package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/adhyaksasb/mycommerce-api/controllers"
	"github.com/adhyaksasb/mycommerce-api/handler"
	"github.com/adhyaksasb/mycommerce-api/initializers"
	"github.com/gin-gonic/gin"
)

var(
	app *gin.Engine
)

func route(r *gin.RouterGroup) {
	// Write Service
	r.POST("/posts", controllers.CreatePosts)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	// Read Service
	r.GET("/posts", controllers.IndexPosts)
	r.GET("/posts/:id", controllers.ShowPost)

	r.GET("/ping", handler.Ping)

	r.GET("/err", handler.ErrRouter)


	r.GET("/hello", func(c *gin.Context) {
		name := c.Query("name")
		if name == "" {
			c.JSON(400, gin.H{
				"message": "name not found",
			})
		} else {
			c.JSON(200, gin.H{
				"data": fmt.Sprintf("Hello %s!", name),
			})
		}
	})
	r.GET("/env", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": os.Getenv("DB_URL"),
		})
	})
	r.GET("/user/:id", controllers.TestUser)
}

func init() {
	// handler.LoadEnvVariables()
	initializers.ConnectToDB()
	app = gin.New()
	r := app.Group("/api")
	route(r)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
