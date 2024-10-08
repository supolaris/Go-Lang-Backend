package controller

import (
	"goProject/services"

	"github.com/gin-gonic/gin"
)

type NotesController struct{
	notesService services.NotesService
}

func (n *NotesController) InitNotesControllerRoutes(router *gin.Engine){
	notes:= router.Group("/notes")
	notes.GET("/", n.GetNotes())
	notes.POST("/", n.PostNotes())
	notes.DELETE("/", n.DeleteNotes())
	notes.PUT("/", n.PutNotes())
	notes.PATCH("/", n.PatchNotes())
}

func (n *NotesController) GetNotes() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"notes": n.notesService.GetNotesService(),
		})
	}
}

func (n *NotesController) PostNotes() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"notes": n.notesService.PostNotesService(),
		})
	}
}

func (n *NotesController) DeleteNotes() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"notes": "Delete Request",
		})
	}
}

func (n *NotesController) PutNotes() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"notes": "Put Request",
		})
	}
}

func (n *NotesController) PatchNotes() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"notes": "Patch Request",
		})
	}
}