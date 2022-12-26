package main

import (
	"log"
	"os"

	"github.com/budhalantara/filebag/pkg"
	"github.com/budhalantara/filebag/task"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func setupDir() (string, string) {
	tmpPath := "./tmp"
	resultPath := "./result"

	if _, err := os.Stat(tmpPath); os.IsNotExist(err) {
		if err := os.Mkdir("./tmp", 0744); err != nil {
			log.Fatal("Failed to create tmp dir", err)
		}
	}

	if _, err := os.Stat(resultPath); os.IsNotExist(err) {
		if err := os.Mkdir("./result", 0744); err != nil {
			log.Fatal("Failed to create result dir", err)
		}
	}

	return tmpPath, resultPath
}

func main() {
	pkg.SetupDB()

	e := echo.New()

	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173", "http://127.0.0.1:5173"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	task.Routes(e)

	e.Logger.Fatal(e.Start(":4321"))
}
