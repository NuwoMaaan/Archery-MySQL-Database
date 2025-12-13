package handlers

import (
	"database/sql"
	"fmt"
)

type EndCount struct {
	ArcherID int
	Count    int
}

type EndCounts []*EndCount

func (result EndCounts) Print(sql string) {
	fmt.Println("Query:", sql)
	fmt.Println("Result:")
	for _, r := range result {
		fmt.Printf("ArcherID: %d, Ends Count: %d\n", r.ArcherID, r.Count)
	}
}

func GetEndCountTotal(db *sql.DB, sqlStatement string) (QueryResult, error) {
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results EndCounts

	for rows.Next() {
		s := &EndCount{}
		if err := rows.Scan(&s.ArcherID, &s.Count); err != nil {
			return nil, err
		}
		results = append(results, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
