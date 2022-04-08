package main

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func RouterSetUp() *gin.Engine {
	r := gin.Default()

	baseUrl := r.Group("api/heroes")
	{
		baseUrl.GET("/", controllers.GetHeroes)
		baseUrl.GET("/:id", controllers.GetHeroById)
		baseUrl.POST("/", controllers.AddHero)
		baseUrl.DELETE("/:id", controllers.DeleteHero)
		baseUrl.PUT("/", controllers.UpdateHero)
	}

	return r
}

func main() {
	r := RouterSetUp()
	r.Run()
}
