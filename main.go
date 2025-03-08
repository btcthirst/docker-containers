package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	// init port
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "4000"
	}
	addr := fmt.Sprintf(":%s", port)
	e := echo.New()

	e.GET("", health)

	e.Logger.Fatal(e.Start(addr))
}

func health(c echo.Context) error {
	now := time.Now()

	return c.JSON(http.StatusOK, now)
}
