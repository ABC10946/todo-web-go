package main

import (
	"github.com/ABC10946/todo-go/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	//todo: handle the error!
	c, _ := handlers.NewContainer()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


	// CreateTodo - Create todo
	e.POST("/todo", c.CreateTodo)

	// DeleteTodo - Delete todo
	e.DELETE("/todo/:todoId", c.DeleteTodo)

	// GetTodo - Get todo
	e.GET("/todo/:todoId", c.GetTodo)

	// GetTodoList - Get All todo
	e.GET("/todo", c.GetTodoList)

	// UpdateTodo - Update todo
	e.PUT("/todo/:todoId", c.UpdateTodo)


	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
