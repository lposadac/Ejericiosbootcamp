package router

import (
	"github.com/gin-gonic/gin"
	"class2/internal/Productos"
	"class2/cmd/server/handlers"
)

func Rutas(servidor *gin.Engine){

	ProductosRouter := servidor.Group("/Productos")

	Repo,_ := Productos.NuevoRepository()
	Serv := Productos.NuevoServicio(Repo)
	Contr := handlers.NuevoControlador(Serv)

	servidor.GET("/ping", handlers.Ping)
	ProductosRouter.GET("/", Contr.GetAllProductos)
	ProductosRouter.GET("/:id", Contr.GetByIdProducto)
	ProductosRouter.GET("/search", Contr.GetMayorQueProducto)
	ProductosRouter.POST("/", Contr.CrearNuevoProducto)
}