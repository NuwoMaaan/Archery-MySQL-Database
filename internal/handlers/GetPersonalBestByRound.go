package handlers

import (
	"database/sql"

	"github.com/NuwoMaaan/Archery-MySQL-Database/internal/results"
)

func GetPersonalBestByRound(db *sql.DB, sqlStatement string) (results.QueryResult, error) {
	var result results.PersonalBestIndividual

	row := db.QueryRow(sqlStatement)
	err := row.Scan(&result.PB_TOTAL)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}
