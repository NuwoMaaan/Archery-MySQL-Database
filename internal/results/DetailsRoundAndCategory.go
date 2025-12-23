package results

import (
	"database/sql"
	"fmt"
)

func (result Details) Print(sql string) {
	println("QUERY:", sql)
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

type Details []*DetailsRoundAndCategory

type DetailsRoundAndCategory struct {
	RoundID       int
	RoundName     string
	AgeGroup      string
	GenderBracket string
	BowType       string
	Distance10m   sql.NullString
	Distance20m   sql.NullString
	Distance30m   sql.NullString
	Distance40m   sql.NullString
	Distance50m   sql.NullString
	Distance60m   sql.NullString
	Distance70m   sql.NullString
	Distance90m   sql.NullString
}
