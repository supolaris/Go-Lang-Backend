package controller

import (
	"goProject/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthService
}

func InitController(authService *services.AuthService) *AuthController {
	return &AuthController{
		authService: *authService,
	}
}

func (n *AuthController) InitRoutes(router *gin.Engine){
	routes := router.Group("/auth")
	routes.POST("/login", n.loginFunc())
	routes.POST("/register", n.registerFunc())
}

func(n *AuthController) registerFunc() gin.HandlerFunc{
	type userData struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	return func(ctx *gin.Context) {
		var userData userData
		if err := ctx.BindJSON(&userData); err != nil{
			ctx.JSON(404, gin.H{
				"result": err.Error(),
			})
			return
		} 
		user,err := n.authService.Register(&userData.Email, &userData.Password)
		if err != nil {
			ctx.JSON(404, gin.H{
				"result": err.Error(),
			})
			return 
		} else {
		ctx.JSON(200, gin.H{
			"message": user,
		})
		return 
	}
	}
}

func(n *AuthController) loginFunc() gin.HandlerFunc{
	type userData struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	return func(ctx *gin.Context) {
		var userData userData
		if err := ctx.BindJSON(&userData); err != nil{
			ctx.JSON(404, gin.H{
				"result": err.Error(),
			})
			return
		} 
		user, err := n.authService.Login(&userData.Email, &userData.Password)
		if err != nil {
			ctx.JSON(404, gin.H{
				"result": err.Error(),
			})
			return 
		} else {
		ctx.JSON(200, gin.H{
			"message": user,
		})
		return 
	}
	}
}
