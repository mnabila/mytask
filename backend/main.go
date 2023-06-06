package main

import (
	"embed"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/mnabila/mytask/api/routers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	RELEASE = "NO"
)

//go:embed static/*
var frontend embed.FS

func main() {
	if RELEASE == "NO" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("POSTGRESQL_URI")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "mytask_",
			SingularTable: true,
			NameReplacer:  nil,
			NoLowerCase:   false,
		},
	})

	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	if err != nil {
		log.Fatal("failed to connect database")
	}

	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:  "http://127.0.0.1:5173, http://127.0.0.1:8080",
		AllowHeaders:  "Origin, Content-Type, Accept, Authorization",
		ExposeHeaders: "Authorization",
	}))

	v1 := app.Group("/api/v1")
	routers.UseUserRouter(db, v1)
	routers.UseTodoRouter(db, v1)

	app.Use("/", filesystem.New(filesystem.Config{
		Root:         http.FS(frontend),
		PathPrefix:   "static",
		Browse:       false,
		NotFoundFile: "static/index.html",
	}))

	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err.Error())
	}
}
