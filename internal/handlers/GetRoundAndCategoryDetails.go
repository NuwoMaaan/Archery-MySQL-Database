package handlers

import (
	"database/sql"
	"fmt"

	"github.com/NuwoMaaan/Archery-MySQL-Database/internal/results"
)

type Details []*results.DetailsRoundAndCategory

func (result Details) Print(sql string) {
	println("query:", sql)
	for _, r := range result {
		fmt.Printf(
			"RoundID: %d, RoundName: %s, AgeGroup: %s, GenderBracket: %s, BowType: %s\n"+
				"10m: %s, 20m: %s, 30m: %s, 40m: %s, 50m: %s, 60m: %s, 70m: %s, 90m: %s\n",
			r.RoundID, r.RoundName, r.AgeGroup, r.GenderBracket, r.BowType,
			r.Distance10m.String, r.Distance20m.String, r.Distance30m.String, r.Distance40m.String,
			r.Distance50m.String, r.Distance60m.String, r.Distance70m.String, r.Distance90m.String,
		)
	}
}

func GetRoundAndCategoryDetails(db *sql.DB, sql string) (results.QueryResult, error) {
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result Details

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
