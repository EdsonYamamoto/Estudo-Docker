package main

import (
	"log"
	"net/http"
	"controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8082"
	}
	e := echo.New()
	e = middlewareApp(e)
	e = routerApp(e)

	log.Println("PORT: "+port)
	e.Logger.Fatal(e.Start(":"+port))
}

func middlewareApp(e *echo.Echo) *echo.Echo{

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	return e
}
func  routerApp(e *echo.Echo) *echo.Echo{

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/c", controllers.CreateUser)
	return e
}