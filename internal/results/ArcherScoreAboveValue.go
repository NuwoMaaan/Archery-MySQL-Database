package results

import "fmt"

func (result AboveScoreValue) Print(sql string) {
	println("QUERY:", sql)
	for _, r := range result {
		fmt.Printf(
			"ArcherID: %d, Count(Total): %d\n", r.ArcherID, r.Total,
		)
	}
}

type AboveScoreValue []*ArcherScoreAboveValue

type ArcherScoreAboveValue struct {
	ArcherID int
	Total    int
}
