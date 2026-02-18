package repository

import (
	"database/sql"

	"github.com/NuwoMaaan/Archery-MySQL-Database/internal/models"
)

func GetPersonalBestByRound(db *sql.DB) (models.QueryResult, error) {
	query := "SELECT MAX(round_total_sum) AS PB_total FROM (SELECT RoundNum, SUM(total) " +
		"AS round_total_sum FROM end INNER JOIN round on round.RoundID = end.RoundID " +
		"WHERE round.RoundName = 'Hobart' AND Approved = 1 AND ArcherID = 1  AND round.DateAdded " +
		"IS NOT NULL and round.DateChange IS NULL GROUP BY RoundNum) AS round_sums"

	var result models.PersonalBestIndividual

	row := db.QueryRow(query)
	err := row.Scan(&result.PB_TOTAL)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}
