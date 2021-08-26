package routes

import (
	"fmt"
	// "log"
	// "os"

	// "github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/myrachanto/testgo/data"
)

func API(){
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Erorr oading go enviroment files")
	// }
	// Port := os.Getenv("PORT")
	e := echo.New()
	fmt.Println("/////////////////")
	data.Getblog()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(":4321"))
}