package routers

import (
	"github.com/OtavioValadao/api-main-go.git/models"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	basePath := "/api/v1"
	v1 := router.Group(basePath)
	//handlers.InitializeHandler()
	{
		//v1.GET("/animal_guardian", )
		v1.GET("/animal_guardian", func(ctx *gin.Context) {

			model := &models.AnimalGuardianResponse{
				Animals: []models.AnimalResponse{
					models.AnimalResponse{},
					models.AnimalResponse{},
				},
			}

			ctx.JSON(200, gin.H{
				"message": "ta rodando parça",
				"value:": models.BasicResponse{
					Status:  200,
					Message: "Objeto vazio",
					Data:    model,
				},
			})
		})
	}

}
