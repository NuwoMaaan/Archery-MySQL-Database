package handlers

import (
	"database/sql"

	"github.com/NuwoMaaan/Archery-MySQL-Database/internal/results"
)

func GetEndCountOrderTotalByDate(db *sql.DB, sqlStatement string) (results.QueryResult, error) {
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result results.ArcherEndScores

	for rows.Next() {
		s := &results.ArcherEndScore{}
		if err := rows.Scan(
			&s.RoundNum, &s.RoundName, &s.ArcherID, &s.FirstName, &s.LastName,
			&s.Date, &s.RoundTotalSum,
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
