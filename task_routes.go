package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func TaskRoutes(e *echo.Echo) {
	e.POST("/api/tasks", func(c echo.Context) error {
		req := TaskRequest{}
		if err := c.Bind(&req); err != nil {
			return err
		}

		ae := req.Validate()
		if ae != nil {
			return ae.ToApiResponse(c)
		}

		ae = taskService.Create(c.Request().Context(), req)
		if ae != nil {
			return ae.ToApiResponse(c)
		}

		return c.NoContent(http.StatusCreated)
	})

	e.GET("/api/tasks", func(c echo.Context) error {
		res, ae := taskService.GetAll(c.Request().Context())
		if ae != nil {
			return ae.ToApiResponse(c)
		}

		return c.JSON(http.StatusOK, ApiResponse{
			Data: res,
		})
	})
}
