package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bazeeko/mechta-test-task/data"
	"github.com/bazeeko/mechta-test-task/database"
	"github.com/bazeeko/mechta-test-task/route"
	"github.com/bazeeko/mechta-test-task/server"
)

func main() {
	addr := os.Getenv("HTTP_PORT")
	if addr == "" {
		addr = "localhost:8080"
	}

	db, err := database.ConnectDB("./database/sqlite.db")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	defer db.Close()

	cityRepository := data.NewCityRepository(db)

	handler := route.NewHandler(cityRepository)
	handler.Init()

	server := server.NewServer(addr, handler)
	log.Println(server.Run())
}
