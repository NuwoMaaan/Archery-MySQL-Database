package handlers

import (
	"database/sql"
	"fmt"
)

type ArcherEndScoreDATE struct {
	EndID      int
	RoundNum   int
	FirstName  string
	LastName   string
	RoundName  string
	Date       string
	Distance   string
	TargetType string
	TOTAL      int
}

type ArcherEndScoresD []*ArcherEndScoreDATE

func (result ArcherEndScoresD) Print(sql string) {
	println("query:", sql)
	for _, r := range result {
		fmt.Printf("EndID: %d Round: %d, FirstName: %s LastName: %s RoundName: %s, Date: %s, Distance: %s TargetType: %s, Total: %d\n",
			r.EndID, r.RoundNum, r.FirstName, r.LastName, r.RoundName, r.Date, r.Distance, r.TargetType, r.TOTAL)
	}
}

func GetEndCountOrderTotalByDate(db *sql.DB, sqlStatement string) (QueryResult, error) {
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results ArcherEndScoresD

	for rows.Next() {
		s := &ArcherEndScoreDATE{}
		if err := rows.Scan(
			&s.EndID, &s.RoundNum, &s.FirstName, &s.LastName,
			&s.RoundName, &s.Date, &s.Distance, &s.TargetType, &s.TOTAL,
		); err != nil {
			return nil, err
		}
		results = append(results, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
