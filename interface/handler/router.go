package handler

import (
	"github.com/labstack/echo"
)

func InitRouting(e *echo.Echo, taskHandler TaskHandler) {

	e.POST("/task", taskHandler.Post())
	e.GET("/task/:id", taskHandler.Get())
	e.PUT("/task/:id", taskHandler.Put())
	e.DELETE("/task/:id", taskHandler.Delete())

}
