package handlers

import (
	"database/sql"
	"fmt"

	"github.com/NuwoMaaan/Archery-MySQL-Database/internal/results"
)

func GetArcherByID(db *sql.DB, sqlStatement string) (results.QueryResult, error) {
	var archer results.Archer

	row := db.QueryRow(sqlStatement)
	err := row.Scan(&archer.FirstName, &archer.LastName, &archer.Gender)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No archers found.")
			return nil, nil
		}
		return nil, err
	}

	return &archer, nil
}
