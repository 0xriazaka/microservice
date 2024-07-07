package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type champion struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Role           string `json:"role"`
	Region         string `json:"region"`
	Ready_To_Fight bool   `json:"ready_to_fight"`
}

var champions = []champion{
	{ID: "1", Name: "clog", Role: "assassin", Region: "aurelia", Ready_To_Fight: false},
	{ID: "2", Name: "jimbi", Role: "tank", Region: "eldoria", Ready_To_Fight: true},
}

// get reqeust
func getChampion(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, champions)
}

// post reqeust
func addChampion(c *gin.Context) {
	var newChampion champion

	// new champion
	if err := c.BindJSON(&newChampion); err != nil {
		return
	}

	champions = append(champions, newChampion)
	c.IndentedJSON(http.StatusCreated, newChampion)
}

// get reqeust by id
func getChampionID(c *gin.Context) {
	id := c.Param("id")

	// champion by id
	for _, a := range champions {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "champion not found"})
}

func main() {
	router := gin.Default()
	router.GET("/champions", getChampion)
	router.GET("/champions/:id", getChampionID)
	router.POST("/champions", addChampion)

	router.Run("localhost:7070")
}
