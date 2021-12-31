package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/thsanan/idolist/dbs"
	"github.com/thsanan/idolist/handlers"
	"github.com/thsanan/idolist/repositories"
	"github.com/thsanan/idolist/services"
)

func main() {

	app := fiber.New()

	db, err := sqlx.Open("mysql", "sandland:IntelliP24.X@tcp(203.159.94.79:3306)/ananyadb")
	if err != nil {
		panic(err)
	}

	dbAlphabet := dbs.NewAlphabetNumbers()
	repo := repositories.NewAlphabetDb(*dbAlphabet)

	service := services.NewName(repo)

	repoDb := repositories.NewMiracleNumber(db)

	miracleService := services.NewMiracle(repoDb)

	handlerNameMaster := handlers.NewNameHandler(service, miracleService)

	app.Get("/", handlerNameMaster.NameMasterHandler)

	log.Fatal(app.Listen(":3000"))

}
