package handlers

import (
	"fmt"
	"gomysql-todo/config"
	"gomysql-todo/models"

	"github.com/gin-gonic/gin"
)

func ListTodo(c *gin.Context) {

	user_id := c.GetUint("user_id")
	var todos []models.Todo
	config.DB.Where("user_id= ?", user_id).Find(&todos)
	c.JSON(200, gin.H{"status": "suiccess", "data": todos})
}

func AddTodo(c *gin.Context) {
	user_id := c.GetUint("user_id")
	var todo models.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(400, gin.H{"message": "invalid input"})
		return
	}

	if todo.Title == "" {
		c.JSON(400, gin.H{"message": "title requried"})
		return
	}
	if todo.Description == "" {
		c.JSON(400, gin.H{"message": "description requried"})
		return
	}

	todo.UserID = user_id

	result := config.DB.Create(&todo)
	if result.Error != nil {
		c.JSON(400, gin.H{"message": "unable to save todo in database"})
		return
	}

	c.JSON(200, gin.H{"status": "success", "message": " todo added succesfully"})
}

func UpdateDoneToggle(c *gin.Context) {
	todo_id := c.Param("id")
	var todo models.Todo
	if result := config.DB.Where("id = ? ", todo_id).Find(&todo); result.Error != nil {
		c.JSON(400, gin.H{"message": "Data not available to update"})
		return
	}

	todo.Done = !todo.Done
	config.DB.Save(&todo)

	c.JSON(200, gin.H{
		"done": todo.Done,
	})
}

func DeleteTodo(c *gin.Context) {
	todo_id := c.Param("id")

	user_id := c.GetUint("user_id")

	fmt.Println(user_id)
	var todo models.Todo

	result := config.DB.Where("id = ? AND user_id = ?", todo_id, user_id).Delete(&todo)

	if result.Error != nil {
		c.JSON(500, gin.H{"message": "database error"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"message": "Data not available to delete"})
		return
	}

	c.JSON(200, gin.H{
		"message": "successfully deleted",
	})

}
