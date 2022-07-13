package main

import (
	"errors"
	"fmt"
)

func main() {
	x := 10
	y := &x

	var z *int = &x // Sintaxe de ponteiro

	fmt.Println(&x)
	fmt.Println(*y) // Imprime o valor que está registrado no endereço de memória
	fmt.Println(*z)

	fmt.Println(abc(z))

	name, age, err := person("Gabriel", 21)

	if err != nil {
		panic(err.Error())
	}

	println(name, age)

	// Função anônima
	func() int {
		return 1
	}()

	// Arrays
	var arr [4]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4

	fmt.Println(arr)

	arr2 := [10]int{1, 2, 3, 4}
	fmt.Println(arr2)

	// Slices (Array "infinito)
	slice := make([]int, 3)
	slice[0] = 1
	slice[1] = 2

	slice = append(slice, 3, 4, 5, 6, 7)

	fmt.Println(slice)

	stringSlice := []string{
		"hello",
		"world",
	}

	stringSlice = append(stringSlice, "giga")

	fmt.Println(stringSlice)
}

func abc(a *int) int {
	return *a * 2
}

func person(name string, age int) (string, int, error) {
	if age == 20 {
		return "", 0, errors.New("Something went wrong")
	}

	return name, age, nil
}
