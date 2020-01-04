package pkg

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Start() {
	g := gin.New()
	g.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "this is a module test",
		})
	})
	err := g.Run(":9999")
	if err != nil {
		log.Fatal(err)
	}
}
