package controllers

import (
	"fmt"
	"net/http"
	"server/db"
	"server/models"

	"github.com/gin-gonic/gin"
)

func GetHeroes(c *gin.Context) {
	var (
		hero   models.Hero
		heroes []models.Hero
	)

	if name, ok := c.GetQuery("name"); ok {
		err := db.Init().QueryRow("select * from hero where name = ?", name).Scan(&hero.Id, &hero.Name)

		if err != nil {
			fmt.Println(err.Error())
			c.JSON(404, gin.H{"message": err.Error()})
			return
		}

		heroes = append(heroes, hero)
		c.JSON(http.StatusOK, heroes)
		return
	}

	rows, err := db.Init().Query("select * from hero")

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&hero.Id, &hero.Name)
		heroes = append(heroes, hero)

		if err != nil {
			fmt.Println(err.Error())
			c.JSON(400, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, heroes)
}

func GetHeroById(c *gin.Context) {
	var (
		hero models.Hero
	)

	id := c.Param("id")
	err := db.Init().QueryRow("select * from hero where id = ?", id).Scan(&hero.Id, &hero.Name)

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(404, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, hero)
}

func AddHero(c *gin.Context) {
	var (
		hero        models.Hero
		requestBody models.Hero
	)

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	stmt, err := db.Init().Prepare("insert into hero(name) values (?);")

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	_, err = stmt.Exec(requestBody.Name)

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	errWithName := db.Init().QueryRow("select * from hero where name = ? order by id desc limit 1", requestBody.Name).Scan(&hero.Id, &hero.Name)

	if errWithName != nil {
		fmt.Println(err.Error())
		c.JSON(404, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, hero)
}

func DeleteHero(c *gin.Context) {
	id := c.Param("id")
	stmt, err := db.Init().Prepare("delete from hero where id = ?;")

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	_, err = stmt.Exec(id)

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(404, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("successfully deleted hero with id %s", id)})
}

func UpdateHero(c *gin.Context) {
	var (
		hero        models.Hero
		requestBody models.Hero
	)

	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	stmt, err := db.Init().Prepare("update hero set name = ? where id = ?;")

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	_, err = stmt.Exec(requestBody.Name, requestBody.Id)

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(404, gin.H{"message": err.Error()})
		return
	}

	errWithName := db.Init().QueryRow("select * from hero where id = ?", requestBody.Id).Scan(&hero.Id, &hero.Name)

	if errWithName != nil {
		fmt.Println(err.Error())
		c.JSON(404, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, hero)
}
