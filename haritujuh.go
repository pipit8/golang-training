package main

import (
	"net/http"

	"github.com/labstack/echo"
)

//BASIC ROUTING
/*
func main() {
	e := echo.New()

	e.GET("/", HelloController)

	e.Start(":8000")
}

func HelloController(d echo.Context) error {
	return d.String(http.StatusOK, "Hai guysss")
}
*/

//RENDERING
/*
type User struct {
	Name  string
	Email string
}

func GetUser(c echo.Context) error {
	user := User{"pipit", "pipit@alterra.id"}
	return c.JSON(http.StatusOK, user)
}

func main() {
	e := echo.New()

	e.GET("/user", GetUser)

	e.Start(":8000")
}
*/

//RENDERING
/*
type User struct {
	Name  string
	Email string
}

func GetUser(c echo.Context) error {
	user := User{"pipit", "pipit@alterra.id"}
	return c.JSON(http.StatusOK, user)
}

func main() {
	e := echo.New()

	e.GET("/user", GetUser)

	e.Start(":8000")
}
*/

//RENDERING
/*
type User struct {
	Name  string
	Email string
}

func GetUser(c echo.Context) error {
	user := User{"pipit", "pipit@alterra.id"}
	return c.JSON(http.StatusOK, user)
}

func main() {
	e := echo.New()

	e.GET("/user", GetUser)

	e.Start(":8000")
}
*/

//URL PARAMS
/*
type User struct {
	Id    int
	Name  string
	Email string
}

func GetUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := User{id, "pipit", "pipit@alterra.id"}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}

func main() {
	e := echo.New()

	e.GET("/users/:id", GetUserController)

	e.Start(":8000")
}
*/

//QUERY PARAMS tanpa if
/*
type User struct {
	Id    int
	Name  string
	Email string
}

func UserSearchController(c echo.Context) error {
	match := c.QueryParam("match")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"match":  match,
		"result": []string{"lala", "dipsi", "pooo"},
	})
}

func main() {
	e := echo.New()

	e.GET("/users", UserSearchController)

	e.Start(":8000")
}
*/

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

//controller

//get all user
func GetUserController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "succes get data all users",
		"users":    users,
	})
}

//create new users
func CreateUserController() error {
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}

	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "succes create all users",
		"user":     user,
	})
}

func main() {
	e := echo.New()

	e.GET("/users", GetUserController)
	e.POST("/users", CreateUserController)

	e.Logger.Fatal(e.Start(":8000"))
}

//nambah buat commit git
