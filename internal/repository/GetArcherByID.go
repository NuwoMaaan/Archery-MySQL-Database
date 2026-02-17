package repository

import (
	"database/sql"
	"fmt"

	"github.com/NuwoMaaan/Archery-MySQL-Database/internal/models"
)

func GetArcherByID(db *sql.DB) (models.QueryResult, error) {
	var archer models.Archer
	query := "SELECT FirstName, LastName, Gender FROM archer WHERE ArcherId = 1"

	row := db.QueryRow(query)
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
