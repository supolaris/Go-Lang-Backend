package controller

import (
	"goProject/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NotesController struct{
	notesService services.NotesService
}

func (n *NotesController) InitNotesControllerRoutes(router *gin.Engine, notesService services.NotesService){
	notes:= router.Group("/notes")
	notes.GET("/", n.GetNotes())
	notes.POST("/", n.PostNotes())
	notes.DELETE("/", n.DeleteNotes())
	notes.PUT("/", n.PutNotes())
	notes.PATCH("/", n.PatchNotes())
	n.notesService = notesService 
}


func (n *NotesController) GetNotes() gin.HandlerFunc{
	return func(ctx *gin.Context) {


		status := ctx.Query("status")
		convertedStatus,err := strconv.ParseBool(status)
		if err != nil {
			ctx.JSON(400, gin.H{
				"Error message": err.Error(),
			})
			return 
		}


		notes,err := n.notesService.GetNotesService(convertedStatus)
		if err != nil {
			ctx.JSON(400, gin.H{
				"Error message": err.Error(),
			})
			return 
		} else {
			ctx.JSON(200, gin.H{
				"notes": notes,
			})
		}
	}
}

func (n *NotesController) PostNotes() gin.HandlerFunc{
	type NotesBody struct {
		Title string `json:"title" binding:"required"`
		Status bool `json:"status"`
	}
	var notesBody NotesBody
	return func(ctx *gin.Context) {
		if err :=	ctx.BindJSON(&notesBody); err != nil{
			ctx.JSON(400, gin.H{
				"Error message": err.Error(),
			})
			return 
		} else {
		note,err :=	n.notesService.CreateNotesService(notesBody.Title, notesBody.Status) 
		if err != nil{
			ctx.JSON(400, gin.H{
				"Error message": err,
			})
			return
		} else {
			ctx.JSON(200, gin.H{
				"note": note,
			}) 
		}
		}
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

