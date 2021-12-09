package main

import (
	"github.com/labstack/echo/v4"
)

func main() {

	start()

}

func start() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Olá, Astro!")
	})

	e.POST("/users", func(c echo.Context) error {
		user := new(User)
		if err := c.Bind(user); err != nil {
			return c.JSON(400, err)
		}
		err := addUser(*user)
		if err != nil {
			return c.JSON(500, "Não foi possível adicionar o usuário.")
		}
		return c.JSON(201, user)
	})

	e.GET("/users/:email", func(c echo.Context) error {
		email := c.Param("email")
		user, err := getUser(email)
		if err != nil {
			return c.JSON(404, "Usuário não encontrado.")
		}
		return c.JSON(200, user)
	})

	e.GET("/users", func(c echo.Context) error {
		users, err := getAllUsers()
		if err != nil {
			return c.JSON(500, "Erro interno.")
		}
		return c.JSON(200, users)
	})

	e.DELETE("/users/:email", func(c echo.Context) error {
		email := c.Param("email")
		err := deleteUser(email)
		if err != nil {
			return c.JSON(500, "Não foi possível remover o usuário.")
		}
		return c.JSON(200, "Usuário removido.")
	})

	e.Logger.Fatal(e.Start(":5000"))

}
