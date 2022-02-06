package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Text string `form:"text"`
}

func GetView(c *gin.Context) {
	d := Database{}
	err := d.Connect()
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	todos := d.FetchTodos()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"todos": todos,
	})
}

func CreateView(c *gin.Context) {
	d := Database{}
	err := d.Connect()
	if err != nil {
		panic("failed to connect database")
	}
	var request Request
	err = c.Bind(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	d.Create(request.Text)
	c.Redirect(http.StatusFound, "/")
}

func DetailView(c *gin.Context) {
	id := c.Param("id")
	d := Database{}
	err := d.Connect()
	if err != nil {
		panic("failed to connect database")
	}
	todo := d.FetchTodo(id)
	c.HTML(http.StatusOK, "detail.html", gin.H{
		"todo": todo,
	})
}

func UpdateView(c *gin.Context) {
	id := c.Param("id")
	d := Database{}
	err := d.Connect()
	if err != nil {
		panic("failed to connect database")
	}
	var request Request
	err = c.Bind(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	d.UpdateTodo(id, request.Text)
	c.Redirect(http.StatusFound, "/")
}

func DeleteView(c *gin.Context) {
	id := c.Param("id")
	d := Database{}
	err := d.Connect()
	if err != nil {
		panic("failed to connect database")
	}
	var request Request
	err = c.Bind(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	d.DeleteTodo(id)
	c.Redirect(http.StatusFound, "/")
}
