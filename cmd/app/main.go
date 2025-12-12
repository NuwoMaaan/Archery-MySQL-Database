package main

import (
	"fmt"
	"log"

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

	fmt.Println("Successfully connection to database:", cfg.DBName)

	menu.MainMenu(db)

}
