package main

import "fmt"

func main() {
	somaAteCem(1)

}

func somaAteCem(posicao uint) uint {
	if posicao >= 100 {
		return posicao
	} else {
		posicao++
	}

	fmt.Println(posicao)
	somaAteCem(posicao)
	return posicao

}