package handlers

import (
	"database/sql"

	"github.com/NuwoMaaan/Archery-MySQL-Database/internal/results"
)

func GetEndCountOrderTotalByRoundNum(db *sql.DB, sqlStatement string) (results.QueryResult, error) {
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result results.ArcherEndScores

	for rows.Next() {
		s := &results.ArcherEndScore{}
		if err := rows.Scan(
			&s.EndID, &s.RoundNum, &s.FirstName, &s.LastName,
			&s.RoundName, &s.Date, &s.Distance, &s.TargetType, &s.TOTAL,
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
