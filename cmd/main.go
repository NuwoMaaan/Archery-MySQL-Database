package main

import (
	"context"
	"fmt"
	"log"

	"github.com/NuwoMaaan/Archery-MySQL-Database/cmd/app"
	"github.com/NuwoMaaan/Archery-MySQL-Database/internal/connection"
	"github.com/NuwoMaaan/Archery-MySQL-Database/internal/menu"
)

func main() {

	cfg, err := connection.LoadConfig()
	if err != nil {
		log.Fatalf("Connection Error To MySQL: %s ", err)
	}

	db, err := connection.ConnectDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Successful connection to database:", cfg.DBName)

	app := app.New(db)
	err = app.Start(context.TODO())
	if err != nil {
		fmt.Println("Application start failure:", err)
	}

	menu.MainMenu(db)

}
