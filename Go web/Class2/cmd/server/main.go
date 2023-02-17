package main

import (
	"github.com/gin-gonic/gin"
	router "class2/cmd/server/router"
)

func main() {
	servidor := gin.Default()
	router.Rutas(servidor)
	servidor.Run(":8080")
}