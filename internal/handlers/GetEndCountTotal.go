package handlers

import (
	"database/sql"

	"github.com/NuwoMaaan/Archery-MySQL-Database/internal/results"
)

func GetEndCountTotal(db *sql.DB, sqlStatement string) (results.QueryResult, error) {
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result results.EndCounts

	for rows.Next() {
		s := &results.EndCount{}
		if err := rows.Scan(&s.ArcherID, &s.Count); err != nil {
			return nil, err
		}
		result = append(result, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
