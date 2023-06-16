package main

import (
	"github.com/SmagulLK/APIGateway/pkg/auth"
	"github.com/SmagulLK/APIGateway/pkg/config"
	"github.com/SmagulLK/APIGateway/pkg/product"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Error loading config", err)

	}
	r := gin.Default()
	authService := auth.RegisterRouter(r, &c)
	product.ProductRoutes(r, &c, authService)
	r.Run(c.Port)
}
