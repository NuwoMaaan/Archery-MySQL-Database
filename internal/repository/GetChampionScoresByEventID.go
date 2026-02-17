package repository

import (
	"database/sql"

	"github.com/NuwoMaaan/Archery-MySQL-Database/internal/models"
)

func GetChampionScoresByEventID(db *sql.DB) (models.QueryResult, error) {
	query := "SELECT archerevent.EventName, catergory.AgeGroup, catergory.GenderBracket, archer.FirstName, archer.LastName, " +
		"SUM(TOTAL) as Total FROM end INNER JOIN archer ON archer.ArcherID = end.ArcherID INNER JOIN archerevent on " +
		"archerevent.EventID = end.EventID INNER JOIN catergory on catergory.CatergoryID = end.CatergoryID " +
		"WHERE archerevent.EventName = 'December Championship' AND Approved = 1 GROUP BY end.ArcherID, catergory.AgeGroup, " +
		"catergory.GenderBracket;"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result models.ChampionScores

	for rows.Next() {
		s := &models.ChampionScoresTotal{}
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
