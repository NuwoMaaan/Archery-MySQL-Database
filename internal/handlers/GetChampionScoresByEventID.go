package handlers

import (
	"database/sql"

	"github.com/NuwoMaaan/Archery-MySQL-Database/internal/results"
)

func GetChampionScoresByEventID(db *sql.DB, sql string) (results.QueryResult, error) {
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result results.ChampionScores

	for rows.Next() {
		s := &results.ChampionScoresTotal{}
		if err := rows.Scan(&s.EventName, &s.AgeGroup, &s.GenderBracket, &s.FirstName, &s.LastName, &s.Total); err != nil {
			return nil, err
		}
		result = append(result, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil

}
