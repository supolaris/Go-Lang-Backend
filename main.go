package main

import (
	"fmt"
	"goProject/controller"
	"goProject/services"

	"github.com/gin-gonic/gin"

	internal "goProject/internals/database"
)

func main() {
	
  router := gin.Default()
	db := internal.InitDb()

	if db != nil { 
		fmt.Println("Connected to db")
	} else {
		fmt.Println("Not connected to db")
	}

	notesService := &services.NotesService{}
	notesService.InitService(db)
	notesController := &controller.NotesController{}
	notesController.InitNotesControllerRoutes(router, *notesService)

  router.Run(":8080")
}