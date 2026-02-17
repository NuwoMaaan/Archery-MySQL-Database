package repository

import (
	"database/sql"

	"github.com/NuwoMaaan/Archery-MySQL-Database/internal/models"
)

func GetPersonalBestAllArchers(db *sql.DB) (models.QueryResult, error) {
	query := "SELECT archer.ArcherID, archer.FirstName, archer.LastName, MAX(round_total_sum) " +
		"AS PB_total FROM (SELECT RoundNum, archer.ArcherID, archer.FirstName, archer.LastName, SUM(total) " +
		"AS round_total_sum FROM end INNER JOIN round on round.RoundID = end.RoundID INNER JOIN archer on " +
		"archer.ArcherID = end.ArcherID WHERE round.RoundName = 'Hobart' AND Approved = 1 AND round.DateAdded " +
		"IS NOT NULL AND round.DateChange IS NULL GROUP BY RoundNum, archer.ArcherID, archer.FirstName, archer.LastName) " +
		"AS round_sums INNER JOIN  archer on archer.ArcherID = round_sums.ArcherID GROUP BY archer.ArcherID"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result models.PersonalBestOfArchers

	for rows.Next() {
		s := &models.PersonalBest{}
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
