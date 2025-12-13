package menu

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/NuwoMaaan/Archery-MySQL-Database/internal/handlers"
	"github.com/NuwoMaaan/Archery-MySQL-Database/internal/queries"
)

func MainMenu(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("==== MENU ====\n")
		for i, opt := range queries.Options {
			fmt.Printf("%d) %s\n", i+1, opt.Description)
		}
		fmt.Println("q) Quit")
		fmt.Print("Enter option: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.EqualFold(input, "q") {
			fmt.Println("Exiting menu.")
			return
		}

		// Check input and convert to int for slice
		choice, err := strconv.Atoi(input)
		if err != nil || choice < 1 || choice > len(queries.Options) {
			fmt.Println("Invalid option — enter a valid number or 'q' to quit.")
			continue
		}
		selected := queries.Options[choice-1]

		handler, ok := handlers.Registry[selected.HandlerKey]
		if !ok {
			fmt.Println("Handler not implemented.")
			continue
		}

		result, err := handler(db, selected.SQL)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		result.Print(selected.SQL)
		fmt.Print("\n\n")
	}
}
