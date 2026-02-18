package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

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
	fmt.Println("a) Menu - CLI interaction with database")
	fmt.Println("b) API - Starts api server for http interaction with database")

	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := menu.GetInput("Enter option:", reader)

		switch strings.ToLower(input) {
		case "a":
			fmt.Println("Starting Menu...")
			menu.MainMenu(db, reader)
			return
		case "b":
			fmt.Println("Starting API Server...")
			app := app.New(db)
			_ = app.Start(context.TODO())
			return
		default:
			fmt.Println("Please choose a or b")
		}
	}

}
