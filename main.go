package main

import (
	"fmt"
	"math"
)

func main() {

	var lado1, lado2 float64
	const precision = 2

	fmt.Print("ingrese lado 1:")
	fmt.Scanln(&lado1)
	fmt.Print("ingrese lado 2:")
	fmt.Scanln(&lado2)

	// procesos

	area := (lado1 * lado2) / 2

	//los numeros "2" son al cuadrado
	hipotenusa := math.Sqrt(math.Pow(lado1, 2) + math.Pow(lado2, 2))
	perimetro := lado1 + lado2 + hipotenusa

	fmt.Printf("\nArea: %.*f \n", precision, area)
	fmt.Printf("Perimetro: %.*f \n", precision, perimetro)

}
