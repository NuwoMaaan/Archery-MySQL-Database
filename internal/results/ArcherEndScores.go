package results

import "fmt"

type ArcherEndScores []*ArcherEndScore
type ArcherEndScore struct {
	RoundNum      int
	ArcherID      int
	FirstName     string
	LastName      string
	RoundName     string
	Date          string
	RoundTotalSum int
}

func (result ArcherEndScores) Print(sql string) {
	println("QUERY:", sql)
	for _, r := range result {
		fmt.Printf("RoundNum: %d RoundName: %s ArcherID: %d FirstName: %s LastName: %s Date: %s, RoundTotalSum: %d\n",
			r.RoundNum, r.RoundName, r.ArcherID, r.FirstName, r.LastName, r.Date, r.RoundTotalSum)
	}
}
