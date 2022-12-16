package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"url-shortner-go/api/route"
	"log"
	"os"
)

func setupRoutes(app *fiber.App){
	app.Get("/:url",route.ResolveUrl)
	app.Post("/api/v1",route.ShortenUrl)
}

func main(){
	err := godotenv.Load()
	if err != nil{
		fmt.Println(err)
	}
	app := fiber.New()
	app.Use(logger.New())
	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
