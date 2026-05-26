package main

import (
	"damage/model"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	log.Println("Starting damage service")

	r.GET("/damage/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, damage := range MockDamage {
			if damage.PedelecId == id {
				c.JSON(200, damage)
				return
			}
		}
		c.JSON(404, gin.H{"error": "Damage not found"})
	})

	r.POST("/damage", func(c *gin.Context) {
		var damage model.Damage
		if err := c.ShouldBindJSON(&damage); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		MockDamage = append(MockDamage, damage)
		c.JSON(200, damage)
	})

	r.Run(":8082")
}
