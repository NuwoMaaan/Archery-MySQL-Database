package models

import (
	"fmt"
)

func (result PersonalBestOfArchers) Print() {
	for _, r := range result {
		fmt.Printf(
			"ArcherID: %d, FirstName: %s, LastName: %s, PersonalBest: %d\n",
			r.ArcherID, r.FirstName, r.LastName, r.PB_Total,
		)
	}
}

type PersonalBestOfArchers []*PersonalBest
type PersonalBest struct {
	ArcherID  int
	FirstName string
	LastName  string
	PB_Total  int
}
