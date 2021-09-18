package main

import (
	"fmt"
	"golang-echo-sql/db"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	d := db.New()
	db.AutoMigrate(d)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Success")
	})

	log.Fatal(e.Start(fmt.Sprintf(":%d", 3001)))
}
