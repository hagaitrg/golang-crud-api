package main

import (
	"basic-golang-rest-api/internal/config"
	"basic-golang-rest-api/internal/connection"
	"basic-golang-rest-api/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func main() {

	cnf := config.Get()	
	dbConnection := connection.GetDatabase(cnf.Database)

	app := fiber.New();

	customerRepository := repository.NewCustomer(dbConnection)

	_ = app.Listen(cnf.Server.Host +":" +cnf.Server.Port)
}

func developers(ctx *fiber.Ctx)  error{
	return ctx.Status(200).JSON("data");
}