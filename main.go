package main

import (
	conf "github.com/PKL-Angkasa-Pura-I/backend-pkl/config"
	handler "github.com/PKL-Angkasa-Pura-I/backend-pkl/controller/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	config := conf.InitConfiguration()
	e := echo.New()

	handler.RegisterGroupAPI(e, config)

	e.Logger.Fatal(e.Start((config.SERVER_ADDRESS)))
}
