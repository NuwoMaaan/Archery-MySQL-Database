package results

import "fmt"

type EndCount struct {
	ArcherID int
	Count    int
}

type EndCounts []*EndCount

func (result EndCounts) Print(sql string) {
	fmt.Println("Query:", sql)
	fmt.Println("Result:")
	for _, r := range result {
		fmt.Printf("ArcherID: %d, EndsCount: %d\n", r.ArcherID, r.Count)
	}
}
