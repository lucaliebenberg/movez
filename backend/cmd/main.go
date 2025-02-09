package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"

	"github.com/lucaliebenberg/movez/backend/api"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")

    // Handle initialization
    var (
        userStore = db.NewSqlUserStore(client)
    //     store = &db.Store{
    //         User: userStore,
    //     }
    //     userHandler = api.NewUserHandler(userStore)
    // var authHandler = api.NewUserHandler(userStore)
    authHandler = api.NewUserHandler(userStore)

    //     app = fiber.New(config)
    auth = app.Group("/api")
    //     apiv1 = app.Group("/api/v1", api.JWTAuthentication(userStore))
    )

    // Auth Handler


    // Auth Handlers
    auth.Post("/auth", authHandler.HandleAuthenticate)

    // User Handlers

    // Geolocation Handlers

    listenAddr := os.Getenv("HTTP_LISTEN_ADDRESS")
    if err := app.Listen(listenAddr); err != nil {
        panic(err)
    }
}

func init() {
    if err := godotenv.Load(); err != nil {
        log.Fatal(err)
    }
}