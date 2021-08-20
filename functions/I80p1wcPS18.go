package main

import (
	"fmt"
)

func getMessage() (a string, b string) {
  return "Hello", "World"
}

func rNames() {
  names := []string{"Ulises", "Mariangel", "Hanlet", "Teresa"}
  for _, name := range names {
	fmt.Println(name)	
  }
}

func main() {
	a, b := getMessage()
	fmt.Println(a, b)
	rNames()
}
