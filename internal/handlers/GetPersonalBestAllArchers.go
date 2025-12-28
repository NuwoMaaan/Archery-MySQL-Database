package handlers

import (
	"database/sql"

	"github.com/NuwoMaaan/Archery-MySQL-Database/internal/results"
)

func GetPersonalBestAllArchers(db *sql.DB, sqlStatement string) (results.QueryResult, error) {
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result results.PersonalBestOfArchers

	for rows.Next() {
		s := &results.PersonalBest{}
		if err := rows.Scan(
			&s.ArcherID, &s.FirstName, &s.LastName, &s.PB_Total,
		); err != nil {
			return nil, err
		}
		result = append(result, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
