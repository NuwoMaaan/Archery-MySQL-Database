package models

import "fmt"

type EndCount struct {
	ArcherID int
	Count    int
}

type EndCounts []*EndCount

func (result EndCounts) Print() {
	fmt.Println("Result:")
	for _, r := range result {
		fmt.Printf("ArcherID: %d, EndsCount: %d\n", r.ArcherID, r.Count)
	}
}
