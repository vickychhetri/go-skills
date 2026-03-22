package main

import (
	"fmt"
	"gomysql-todo/config"
	"gomysql-todo/handlers"
	"gomysql-todo/middleware"
	"gomysql-todo/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Server on Mysql Go Todo")

	//LOAD ENV FILE
	godotenv.Load(".env")

	//Conenct DB
	config.ConnectDB()

	//AUTO Migrate
	config.DB.AutoMigrate(&models.User{}, &models.Todo{})

	r := gin.Default()

	//Routes
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	todo := r.Group("/todos")

	todo.Use(middleware.JwtAuth())
	{
		todo.POST("", handlers.AddTodo)
		todo.POST("/:id", handlers.UpdateDoneToggle)
		todo.DELETE("/:id", handlers.DeleteTodo)
	}

	fmt.Println("server Running at 8080")
	r.Run(":8080")

}
