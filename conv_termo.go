package main

import "fmt"

const ebulicaoK = 373

func main() {

	tempK := ebulicaoK

	tempC := tempK - 273

	fmt.Printf("O ponto de ebulição da água em Kelvin é de: %d e o ponto de ebulição da água em Celsius é de: %d.", tempK, tempC)

}
