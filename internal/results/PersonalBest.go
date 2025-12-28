package results

import "fmt"

type PersonalBestIndividual struct {
	PB_TOTAL int
}

func (r *PersonalBestIndividual) Print(sql string) {
	fmt.Println("QUERY:", sql)
	fmt.Println("RESULT:")
	fmt.Println("Personal Best:", r.PB_TOTAL)
}
