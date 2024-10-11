package controller

import (
	"goProject/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NotesController struct{
	notesService services.NotesService
}

func (n *NotesController) InitController(notesService services.NotesService) *NotesController{
	n.notesService = notesService;
	return n;
}

func (n *NotesController) InitRoutes(router *gin.Engine){
	notes:= router.Group("/notes")
	notes.GET("/", n.GetNotes())
	notes.GET("/:id", n.GetSingleNote())
	notes.POST("/", n.PostNotes())
	notes.DELETE("/:id", n.DeleteNotes())
	notes.PUT("/", n.PutNotes())
	notes.PATCH("/", n.PatchNotes())
}

func (n *NotesController) GetSingleNote() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		convertedId,err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{
				"result": err.Error(),
			})
			return 
		}
		note,err := n.notesService.GetSingleNotesService(convertedId)
		if err != nil {
			ctx.JSON(400, gin.H{
				"result": err.Error(),
			})
			return 
		} else {
			ctx.JSON(200, gin.H{
				"note": note,
			})
		}
	}
}

func (n *NotesController) GetNotes() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		status := ctx.Query("status")
		var actualStatus *bool
		if status != ""{
			convertedStatus,err := strconv.ParseBool(status)
			actualStatus = &convertedStatus
			if err != nil {
				ctx.JSON(400, gin.H{
					"result": err.Error(),
				})
				return 
			}

		}
		
		notes,err := n.notesService.GetNotesService(actualStatus)
		if err != nil {
			ctx.JSON(400, gin.H{
				"result": err.Error(),
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
				"result": err.Error(),
			})
			return 
		} else {
		note,err :=	n.notesService.CreateNotesService(notesBody.Title, notesBody.Status) 
		if err != nil{
			ctx.JSON(400, gin.H{
				"result": err,
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

func (n *NotesController) PutNotes() gin.HandlerFunc{
	type NotesBody struct {
		Title string `json:"title" binding:"required"`
		Status bool `json:"status"`
		Id int `json:"id" binding:"required"`
	}
	var notesBody NotesBody
	return func(ctx *gin.Context) {
		if err :=	ctx.BindJSON(&notesBody); err != nil{
			ctx.JSON(400, gin.H{
				"result": err.Error(),
			})
			return 
		} else {
		note,err :=	n.notesService.UpdateNotesService(notesBody.Title, notesBody.Status, notesBody.Id) 
		if err != nil{
			ctx.JSON(400, gin.H{
				"result": err,
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
		id := ctx.Param("id")
		convertedId,err := strconv.ParseInt(id, 10, 64)
		if err != nil{
			ctx.JSON(400, gin.H{
				"result": err,
			})
			return
		}
		err = n.notesService.DeleteNotesService(convertedId)
		if err != nil{
			ctx.JSON(400, gin.H{
				"result": err,
			})
			return
		} else {
			ctx.JSON(200, gin.H{
				"message": "Note deleted successfully",
			})
		}
	}
}

func (n *NotesController) PatchNotes() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"notes": "Patch Request",
		})
	}
}