package main

import "fmt"

func main() {
	x := 10
	y := &x

	var z *int = &x // Sintaxe de ponteiro

	fmt.Println(&x)
	fmt.Println(*y) // Imprime o valor que está registrado no endereço de memória
	fmt.Println(*z)

	fmt.Println(abc(z))
}

func abc(a *int) int {
	return *a * 2
}
