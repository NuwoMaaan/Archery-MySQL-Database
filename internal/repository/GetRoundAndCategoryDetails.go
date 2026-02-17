package repository

import (
	"database/sql"

	"github.com/NuwoMaaan/Archery-MySQL-Database/internal/models"
)

func GetRoundAndCategoryDetails(db *sql.DB) (models.QueryResult, error) {
	query := "SELECT round.RoundID, round.RoundName, catergory.AgeGroup, catergory.GenderBracket, " +
		"catergory.BowType, 10m,20m,30m,40m,50m,60m,70m,90m FROM catergory " +
		"INNER JOIN round on round.RoundID = catergory.RoundID WHERE round.RoundName = 'Perth' AND catergory.AgeGroup = 'Under 21'"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result models.Details

	for rows.Next() {
		s := &models.DetailsRoundAndCategory{}
		if err := rows.Scan(&s.RoundID, &s.RoundName, &s.AgeGroup, &s.GenderBracket, &s.BowType,
			&s.Distance10m, &s.Distance20m, &s.Distance30m, &s.Distance40m,
			&s.Distance50m, &s.Distance60m, &s.Distance70m, &s.Distance90m,
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
