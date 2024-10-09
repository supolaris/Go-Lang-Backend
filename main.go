package main

import (
	"fmt"
	"goProject/controller"

	"github.com/gin-gonic/gin"

	internal "goProject/internals/database"
)

func main() {
  router := gin.Default()

	db := internal.InitDb()

	if db == nil {
		fmt.Println("Error in db")
	}

//   router.GET("/ping", func(c *gin.Context) {
//     c.JSON(http.StatusOK, gin.H{
//       "message": "pong",
// 			"var": "hello world",
//     })
//   })

// 	router.GET("/me/:id/:newId", func(c *gin.Context) {
// 		var id=c.Param("id")
// 		var newId=c.Param(("newId"))
// 		c.JSON(http.StatusOK, gin.H{
// 		"id": id,
// 		"newId": newId,
// 		})
// 	})

// 	router.POST("/me", func(c *gin.Context){
// 		type MeRequest struct {
// 			Email string `json:"email" binding:"required"`
// 			Password string `json:"password"`
// 		}
// 		var meRequest MeRequest
// 	 if err := 	c.BindJSON(&meRequest); err !=nil{
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 			})
// 			return 		
// 	 } else {
// 		c.JSON(http.StatusOK, gin.H{
// 			"email": meRequest.Email,
// 			"password": meRequest.Password,
// 			})
// 	 }
// 	})

// 	router.PUT("/me", func(c *gin.Context) {
// 		type MeRequest struct {
// 			Email string `json:"email" binding:"required"`
// 			Password string`json:"password"`
// 		}
// 		var meRequest MeRequest
// 			if err := c.BindJSON(&meRequest); err !=nil{
// 			c.JSON(http.StatusBadRequest, gin.H {
// 			"error": err.Error(),
// 		})
// 		return 
// 		} else {
// 			c.JSON(http.StatusOK, gin.H{
// 			"email" : meRequest.Email,
// 			"password": meRequest.Password,
// 			})
// 		}
// 	})

// 	router.PATCH("/me", func(c *gin.Context) {
// 		type MeRequest struct {
// 			Email string `json:"email" binding:"required"`
// 			Password string `json:"password"`
// 		}
// 		var meRequest MeRequest
// 		if err := c.BindJSON(&meRequest); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H {
// 		"error": err.Error(),
// 		})
// 		return
// 		} else {
// 		c.JSON(http.StatusOK, gin.H {
// 		"email": meRequest.Email,
// 		"password": meRequest.Password,
// 		})
// 	}
// })

// router.DELETE("/me/:id", func(c *gin.Context) {
// 	var id = c.Param("id")
// 	c.JSON(http.StatusOK, gin.H{
// 		"id": id,
// 		"message": "Deleted",
// 	})
// })


notesController := &controller.NotesController{}
notesController.InitNotesControllerRoutes(router)

  router.Run(":8080")
}