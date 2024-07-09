package handlers
import (
	"github.com/ABC10946/todo-go/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

// CreateTodo - Create todo
func (c *Container) CreateTodo(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld {
		Message: "Hello World",
	})
}


// DeleteTodo - Delete todo
func (c *Container) DeleteTodo(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld {
		Message: "Hello World",
	})
}


// GetTodo - Get todo
func (c *Container) GetTodo(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld {
		Message: "Hello World",
	})
}


// GetTodoList - Get All todo
func (c *Container) GetTodoList(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld {
		Message: "Hello World",
	})
}


// UpdateTodo - Update todo
func (c *Container) UpdateTodo(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld {
		Message: "Hello World",
	})
}

