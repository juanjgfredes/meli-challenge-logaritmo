package main

import (
	"fmt"

	"github.com/juanjgfredes/meli-challenge-logaritmo.git/logaritmo"
)

func main() {
	var n int
	fmt.Print("Ingrese el tamaño del ADN que ingresara: ")
	fmt.Scanln(&n)

	var adn = make([]string, 4)

	fmt.Println("Acontinuación se le solicitara las secuencua de ADN, recuerde que solo podra ingresar",
		"la cantidad de letras que indico y que solo se permiten A, C, G y T")
	var secuencia string

	for i := 0; i < n; i++ {
		fmt.Print("Ingrese la secuencua nro ", i, " del ADN: ")
		fmt.Scanln(&secuencia)
		if i < 4 {
			adn[i] = secuencia
		} else {
			adn = append(adn, secuencia)
		}
	}

	esMutante, err := logaritmo.EsMutante(adn)

	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	if esMutante {
		fmt.Println("El adn ingresado es mutante")
	} else {
		fmt.Println("El adn ingresado no es mutante")
	}

}
