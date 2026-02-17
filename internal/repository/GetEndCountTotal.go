package repository

import (
	"database/sql"

	"github.com/NuwoMaaan/Archery-MySQL-Database/internal/models"
)

func GetEndCountTotal(db *sql.DB) (models.QueryResult, error) {
	query := "SELECT ArcherID, COUNT(ArcherID) FROM end GROUP BY ArcherID"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result models.EndCounts

	for rows.Next() {
		s := &models.EndCount{}
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
