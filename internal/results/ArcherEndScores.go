package results

import "fmt"

type ArcherEndScores []*ArcherEndScore
type ArcherEndScore struct {
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

func (result ArcherEndScores) Print(sql string) {
	println("QUERY:", sql)
	for _, r := range result {
		fmt.Printf("EndID: %d Round: %d, FirstName: %s LastName: %s RoundName: %s, Date: %s, Distance: %s TargetType: %s, Total: %d\n",
			r.EndID, r.RoundNum, r.FirstName, r.LastName, r.RoundName, r.Date, r.Distance, r.TargetType, r.TOTAL)
	}
}
