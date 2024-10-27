package middleware

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"

)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        // Authentication logic
        return next(c)
    }
}

func SetMiddlewares(router *echo.Echo) {
    // set echo's middlewares
    router.Use(middleware.RemoveTrailingSlash())
	router.Use(middleware.RequestID())
	router.Use(middleware.Recover())
	router.Use(middleware.CORS())
    // router.Use(middleware.Logger())

    // set custom middlewares
}

