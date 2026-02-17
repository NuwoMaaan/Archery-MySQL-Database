package models

import "fmt"

func (result AboveScoreValue) Print() {
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
