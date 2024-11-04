package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Member7 struct{
	Id string `json:"id"`
	name string `json:"name"`
	year string `json:"year"`
	batch string `json:"batch"`
}

var member = []Member7{
	Id: "1", name: "daffa", year:"2005", batch: "Astrosphere"
	Id: "2", name: "lila", year:"2005", batch: "Astrosphere"
}

func main(){
	PORT := (":8080")
	r := gin.Default()

	r.POST("/member", createMember)
	r.GET("/member", getAllMember)
	r.GET("/member/:id", getByIdMember)
	r.PUT("/member/:id", updateMember)
	r.DELETE("/member/:id", deleteMember)
}

func getAllMember(c *gin.Context){
	c.JSON(http.StatusOK, member)
}

func getByIdMember(c *gin.Context){
	id := c.Param("id")
	for _, v := range member{
		if v.Id == id{
			c.JSON(http.StatusOK,v)
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "member not found"})
}

func createMember(c *gin.Context){
	var newMember Member7

	if err := c.BindJSON(&newMember); err != nil
	return
}