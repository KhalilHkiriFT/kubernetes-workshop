package main

import (
	"log"
	"reservation/model"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	log.Println("Starting reservation service")

	r.GET("/pedelec", func(c *gin.Context) {
		c.JSON(200, MockPedelecs)
	})

	r.GET("/pedelec/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, pedelec := range MockPedelecs {
			if pedelec.ID == id {
				c.JSON(200, pedelec)
				return
			}
		}
		c.JSON(404, gin.H{"error": "Pedelec not found"})
	})

	r.POST("/reservation", func(c *gin.Context) {
		var reservation model.Reservation
		if err := c.ShouldBindJSON(&reservation); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		MockReservations = append(MockReservations, reservation)
		c.JSON(200, reservation)
	})

	r.GET("/reservation/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, reservation := range MockReservations {
			if reservation.ID == id {
				c.JSON(200, reservation)
				return
			}
		}
		c.JSON(404, gin.H{"error": "Reservation not found"})
	})

	r.DELETE("/reservation/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, reservation := range MockReservations {
			if reservation.ID == id {
				MockReservations = append(MockReservations[:i], MockReservations[i+1:]...)
				c.JSON(200, gin.H{"message": "Reservation deleted"})
				return
			}
		}
		c.JSON(404, gin.H{"error": "Reservation not found"})
	})

	r.Run(":8080")
}
