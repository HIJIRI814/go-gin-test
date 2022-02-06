package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) { GetView(c) })
	r.POST("/", func(c *gin.Context) { CreateView(c) })
	r.GET("/:id", func(c *gin.Context) { DetailView(c) })
	r.POST("/:id", func(c *gin.Context) { UpdateView(c) })
	r.POST("/:id/delete", func(c *gin.Context) { DeleteView(c) })

	r.Run(":8000")
}
