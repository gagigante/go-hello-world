package main

import (
	"fmt"
	"hello-go/example"
	"net/http"
)

func main() {
	const xpto int = 1

	const (
		x string = "aaaa"
		y string = "bbbb"
		z string = "cccc"
	)

	print(x)

	res, _ := http.Get("https://google.com.br")

	// if err != nil {
	// 	panic("Houve um erro na requisição")
	// }

	fmt.Println(res.Body)

	// var nome string = "hello world"
	// var nome = "hello world"

	a := "Gabriel"
	b := 10
	c := 10.44
	d := false
	e := 'a'
	f := `
		Gabriel
		Gigante
	`

	print(a + "\n")

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)

	example.PrintExample()
}
