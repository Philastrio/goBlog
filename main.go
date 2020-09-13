package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Success")
}

func main() {
	fmt.Println("Hello")

	e := echo.New()
	e.GET("/", hello)

	e.Start(":7 000")
}
