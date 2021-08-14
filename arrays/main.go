package main

import (
	"fmt"
)

type names []string

func (n names) print() {
	for _, name := range n {
		fmt.Println("name:"+ name)
	}
}

func main() {
	namesS := names{"Ulises", "Mariangel", "Hanlet", "Isabella", "Teresa"}
	namesS.print()
}