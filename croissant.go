package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Println("donnez le premier nombre")
	fmt.Scanln(&x)
	fmt.Println("donnez le deuxieme nombre")
	fmt.Scanln(&y)
	fmt.Println("donnez le troisieme nombre")
	fmt.Scanln(&z)
	if x <= y && y <= z {
		fmt.Println("ordonnés")
	} else {
		fmt.Println("non ordonnés")
	}
}
