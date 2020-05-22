package main

import (
	"./controllers"
	"./docs"
	"./repositories"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
	"net/http"
)

// @contact.name hejiangle
// @contact.email jiangle.he@foxmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	repositories.ConnectToDatabase()
	repositories.Migration()

	router := gin.New()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/healthyCheck", controllers.Healthy)
		v1.GET("/", notFound)

		todos := v1.Group("/todoItems")
		{
			todos.GET("", controllers.TodoItems)
			todos.POST("", controllers.CreateTodoItem)
			todos.GET(":id", controllers.GetToDoItem)
			todos.PATCH(":id", controllers.EditToDoItem)
			todos.DELETE(":id", controllers.DeleteToDoItem)
		}
	}

	docs.SwaggerInfo.Title = "TodoList API"
	docs.SwaggerInfo.Description = "This is a simple api demo for practice go language"
	docs.SwaggerInfo.Version = "v1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatal(router.Run(":8080"))
}

func notFound(c *gin.Context) {
	w := c.Writer

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusNotFound)
	_, _ = w.Write([]byte(`{"message": "not found"}`))
}
