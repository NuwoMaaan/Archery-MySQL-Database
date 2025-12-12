package handlers

import (
	"database/sql"
	"fmt"
)

type EndCount struct {
	ArcherID int
	Count    int
}

func GetEndCountTotal(db *sql.DB, sqlStatement string) (*EndCount, error) {
	var endcount EndCount

	row := db.QueryRow(sqlStatement)
	err := row.Scan(&endcount.ArcherID, &endcount.Count)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found.")
			return nil, nil
		}
		return nil, err
	}

	return &endcount, nil
}
