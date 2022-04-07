package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Hero struct {
	Id   int    `json:"id"`
	Name string ` json:"name"`
}

var fakeHeroesDataArr = []Hero{
	{Id: 11, Name: "Dr Nice"},
	{Id: 12, Name: "Narco"},
	{Id: 13, Name: "Bombasto"},
	{Id: 14, Name: "Celeritas"},
	{Id: 15, Name: "Magneta"},
	{Id: 16, Name: "RubberMan"},
	{Id: 17, Name: "Dynama"},
	{Id: 18, Name: "Dr IQ"},
	{Id: 19, Name: "Magma"},
	{Id: 20, Name: "Tornado"},
}

func getHeroesHandler(c *gin.Context) {
	if name, ok := c.GetQuery("name"); ok {
		for _, v := range fakeHeroesDataArr {
			if v.Name == name {
				var heroNameArr []Hero
				heroNameArr = append(heroNameArr, v)
				c.JSON(http.StatusOK, heroNameArr)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "hero with name not found"})
		return
	} else {
		c.JSON(http.StatusOK, fakeHeroesDataArr)
	}

}

func getHeroByIdHandler(c *gin.Context) {
	// change str type to int
	heroId, _ := strconv.Atoi(c.Param("id"))

	for _, v := range fakeHeroesDataArr {
		if v.Id == heroId {
			c.JSON(http.StatusOK, v)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "hero can not be found"})
}

func addHeroHandler(c *gin.Context) {
	var requestBody Hero

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "add hero failed"})
	}
	var newId int = fakeHeroesDataArr[len(fakeHeroesDataArr)-1].Id + 1

	fakeHeroesDataArr = append(fakeHeroesDataArr, Hero{Id: newId, Name: requestBody.Name})
	c.JSON(http.StatusOK, Hero{Id: newId, Name: requestBody.Name})
}

func remove(slice []Hero, s int) []Hero {
	return append(slice[:s], slice[s+1:]...)
}

func deleteHeroHandler(c *gin.Context) {
	// change str type to int
	heroId, _ := strconv.Atoi(c.Param("id"))

	for i, v := range fakeHeroesDataArr {
		if v.Id == heroId {
			fakeHeroesDataArr = remove(fakeHeroesDataArr, i)
			c.JSON(http.StatusOK, gin.H{"message": "successfully deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "delete failed"})
}

func updateHeroHandler(c *gin.Context) {
	var requestBody Hero

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "update hero failed"})
	}

	for i, v := range fakeHeroesDataArr {
		if v.Id == requestBody.Id {
			fakeHeroesDataArr[i].Name = requestBody.Name
			c.JSON(http.StatusOK, v)
			return
		}
	}

}

func main() {

	r := gin.Default()
	r.GET("/api/heroes", getHeroesHandler)
	r.GET("/api/heroes/:id", getHeroByIdHandler)
	r.POST("/api/heroes", addHeroHandler)
	r.DELETE("/api/heroes/:id", deleteHeroHandler)
	r.PUT("/api/heroes", updateHeroHandler)
	r.Run()
}
