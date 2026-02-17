package repository

import (
	"database/sql"

	"github.com/NuwoMaaan/Archery-MySQL-Database/internal/models"
)

func GetEndCountOrderTotalByRoundNum(db *sql.DB) (models.QueryResult, error) {
	query := "SELECT RoundNum, round.RoundName, archer.ArcherID, archer.FirstName, Date, archer.LastName, SUM(total) " +
		"AS RoundTotalSum FROM `end` INNER JOIN round on round.RoundID = `end`.RoundID " +
		"INNER JOIN archer on archer.ArcherID = `end`.ArcherID WHERE archer.ArcherID = 1 " +
		"AND Approved = 1 AND round.RoundName = 'Hobart' GROUP BY RoundNum, archer.ArcherID, " +
		"archer.FirstName, archer.LastName, round.RoundName, Date ORDER BY RoundNum DESC LIMIT 10"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result models.ArcherEndScores

	for rows.Next() {
		s := &models.ArcherEndScore{}
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
