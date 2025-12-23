package handlers

import (
	"database/sql"

	"github.com/NuwoMaaan/Archery-MySQL-Database/internal/results"
)

func GetRoundAndCategoryDetails(db *sql.DB, sql string) (results.QueryResult, error) {
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result results.Details

	for rows.Next() {
		s := &results.DetailsRoundAndCategory{}
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
