package main

import (
	"github.com/gin-gonic/gin"
)

//lets understand go gin basics for interview
// Server   - done
//Routes
//Response JSON
//Decode json

func main() {
	/*
		##################START BASIC SERVER #######################
	*/
	//BASIC SERVER
	// router := gin.Default()

	// router.GET("/info", GinHandlerFunc)
	// router.POST("/info-submit", GinHandlerFunc)
	// err := router.Run(":8080")
	// if err != nil {
	// 	return
	// }

	/*
		##################END BASIC SERVER #######################
	*/

	/*
		##################START  includes Logger + Recovery #######################
	*/
	// gin.Default() includes Logger + Recovery
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// r.GET("/panic", func(c *gin.Context) {
	// 	panic("something went wrong")
	// })

	// r.Run(":8080")
	/*
		##################END  includes Logger + Recovery #######################
	*/

	/*
		################## START gin.New #######################
	*/
	/*
		gin.New mannual without default
	*/

	// r := gin.New()
	// r.Use(gin.Logger())
	// r.Use(gin.Recovery())

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// r.GET("/panic", func(c *gin.Context) {
	// 	panic("something went wrong")
	// })

	// r.Run(":8080")
	/*
		##################END gin.New #######################
	*/

}

type APIResponse struct {
	Message string      `json:"message,omitempty"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
}

func GinHandlerFunc(c *gin.Context) {

	response := APIResponse{
		Success: true,
		Message: "success",
		Data: gin.H{
			"status":  "success",
			"message": "Message Succesfully Passed from Go Gin.",
		},
	}

	c.JSON(200, response)
}

/*
Gin Framework Core
2.1 Basic Server
r := gin.Default()
r.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{"message": "pong"})
})
r.Run(":8080")


Interview question
What does gin.Default() include?
✅ Logger + Recovery middleware

*/
