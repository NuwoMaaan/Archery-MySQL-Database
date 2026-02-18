package repository

import (
	"database/sql"

	"github.com/NuwoMaaan/Archery-MySQL-Database/internal/models"
)

func GetArcherScoreAboveValue(db *sql.DB) (models.QueryResult, error) {
	query := "SELECT end.ArcherID, COUNT(TOTAL) FROM end LEFT JOIN archer ON archer.ArcherID = end.ArcherID " +
		"WHERE TOTAL > 50 AND Approved = 1 GROUP BY end.ArcherID"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result models.AboveScoreValue

	for rows.Next() {
		s := &models.ArcherScoreAboveValue{}
		if err := rows.Scan(
			&s.ArcherID, &s.Total,
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
