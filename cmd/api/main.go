package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gomarket1.0/internal/config"
)

type Health struct {
	Status string `json:"status"`
}

func healthHandler(c *gin.Context) {
	health := Health{Status: "ok"}

	c.JSON(200, health)
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()

	r.GET("/health", healthHandler)

	if err := r.Run(":" + cfg.Port); err != nil {
		fmt.Println("Server failed:", err)
	}
}
