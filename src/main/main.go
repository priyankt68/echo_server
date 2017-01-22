package main

import(
	"fmt"
	"net/http"
	"github.com/labstack/echo"
)

func main() {
	fmt.Println("Yallo for the first line")

	e := echo.New()

	e.GET("/", index)
	e.GET("/users/:id", path_paramters)

	e.Start(":8000")
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Universe!")
	
}

func path_paramters(c echo.Context) error {
	
	id := c.Param("id")
	return c.String(http.StatusOK, id)
	
}