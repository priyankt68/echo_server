package main

import(
	"fmt"
	"net/http"
	"github.com/labstack/echo"
)

func main() {
	fmt.Println("Yallo for the first line")

	e := echo.New()

	e.GET("/", func (c echo.Context) error {
		return c.String(http.StatusOK, "Yallow from the web side!")
		
	})

	e.Start(":8000")
}
