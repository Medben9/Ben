package main

import "fmt"

func main() {
	var nb, rest, total, montant, arendre, nb10, nb5, nb2, nb1 int
	const ph = 2
	fmt.Printf("ecrivez le nombre d heure de stationnement :")
	fmt.Scanln(&nb)
	total = nb * ph
	fmt.Println("votre somme est :", total)
	fmt.Println("c'est quoi votre montant :")
	fmt.Scanln(&montant)
	arendre = montant - total
	nb10 = arendre / 10
	rest = arendre % 10
	nb5 = rest / 10
	rest = arendre % 5
	nb2 = arendre / 2
	fmt.Println("le reste est ", arendre)

	fmt.Println("piece de 10 : ", nb10)
	fmt.Println("piece de 5 : ", nb5)
	fmt.Println("piece de 2 : ", nb2)
	fmt.Println("piece de 1 : ", nb1)
}
