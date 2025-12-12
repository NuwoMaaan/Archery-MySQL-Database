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
		selectedQuery := queries.Options[choice-1]

		switch choice {
		case 1:
			archer, err := handlers.GetArcherByID(db, selectedQuery.SQL)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Println("Executed Query:", selectedQuery.SQL)
			fmt.Println("QUERY RESULT:", archer.FirstName, archer.LastName, archer.Gender)

		case 2:
			endcount, err := handlers.GetEndCountTotal(db, selectedQuery.SQL)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Println("Executed Query:", selectedQuery.SQL)
			fmt.Println("QUERY RESULT:", endcount.ArcherID, endcount.Count)
		default:
			fmt.Println("Option not implemented yet.")
		}
	}
}
