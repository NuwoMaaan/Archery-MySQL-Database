package menu

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/NuwoMaaan/Archery-MySQL-Database/internal/queries"
)

func GetInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func MainMenu(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("==== MENU ====\n")
		for i, opt := range queries.Options {
			fmt.Printf("%d) %s\n", i+1, opt.Description)
		}
		fmt.Println("q) Quit")
		input, _ := GetInput("Enter option:", reader)

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

		if db == nil {
			fmt.Println("Database not connected; selected option is disabled.")
			continue
		}

		result, err := selected.Handler(db)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		result.Print()
		fmt.Print("\n")
	}
}
