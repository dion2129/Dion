package main

import "github.com/gin-gonic/gin"

func main(){
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context){
		c.JSON(200	 gin.H{
			"massage": "HEllo, World!",
		})
	})

	r.Run(":8080")
}