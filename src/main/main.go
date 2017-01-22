package main

import(
	"fmt"
<<<<<<< HEAD
	"net/http"
	"github.com/labstack/echo"
)

type (

	user struct {

		ID int 'json:"id"'
		Name string 'json:"name"'
	}

)

var (

	users = map[int]*user{}
	seq   = 1
=======
>>>>>>> a3c6e7ca332890cc2bb7d5570cd5a7e9e6b6f1ab
)

func main() {
	fmt.Println("Yallo for the first line")
<<<<<<< HEAD

	e := echo.New()

	e.GET("/", index)
	e.GET("/users/:id", path_paramters)

	e.POST("/users", createUser)


	e.Start(":8000")
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Universe!")
	
}

// Handlers

func path_paramters(c echo.Context) error {
	
	id := c.Param("id")
	return c.String(http.StatusOK, id)
	
}

func createUser(c echo.Context) error {
	u := &user{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}
=======
}
>>>>>>> a3c6e7ca332890cc2bb7d5570cd5a7e9e6b6f1ab
