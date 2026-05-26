package main

import (
	"location/model"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	log.Println("Starting location service")

	r.GET("/locations/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, location := range MockLocations {
			if location.ID == id {
				c.JSON(200, location)
				return
			}
		}
		c.JSON(404, gin.H{"error": "Location not found"})
	})

	r.POST("/locations", func(c *gin.Context) {
		var location model.Location
		if err := c.ShouldBindJSON(&location); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		MockLocations = append(MockLocations, location)
		c.JSON(200, location)
	})

	r.Run(":8081")
}
