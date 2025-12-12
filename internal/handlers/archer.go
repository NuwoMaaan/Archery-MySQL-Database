package handlers

import (
	"database/sql"
	"fmt"
)

type Archer struct {
	FirstName string
	LastName  string
	Gender    string
}

func GetArcherByID(db *sql.DB, sqlStatement string) (*Archer, error) {
	var archer Archer

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
