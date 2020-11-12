package controllers
import (
  "github.com/labstack/echo"
  "net/http"
)
//create user
func CreateUser(c echo.Context) error {  
  
  return c.JSON(http.StatusCreated, "teste")
}