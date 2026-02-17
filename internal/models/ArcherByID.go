package models

import "fmt"

type Archer struct {
	FirstName string
	LastName  string
	Gender    string
}

func (archer *Archer) Print() {
	fmt.Println("RESULT:")
	fmt.Println(" First Name:", archer.FirstName)
	fmt.Println(" Last Name :", archer.LastName)
	fmt.Println(" Gender    :", archer.Gender)
}
