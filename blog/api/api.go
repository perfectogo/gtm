package api

import (
    "monolith-app/api/handlers"
    "github.com/labstack/echo/v4"
)

type Options struct {

}

func RegisterRoutes(router *echo.Echo, opt Options) {

    h := handlers.NewHandler()

    v1 := router.Group("/api/v1")
    {
        // check service
        v1.Get("/ping", h.Ping)
    }
}