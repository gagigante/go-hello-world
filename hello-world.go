package main

import (
	"encoding/json"
	"fmt"
)

type vehicle interface {
	start() string
}

type Car struct {
	CarName string `json:"car"`
	CarYear int    `json:"year"`
}

func (c Car) drive() {
	print(c.CarName, " vrum vrum")
}

func (c Car) start() string {
	return "**engine starts**"
}

func example(car vehicle) {
	fmt.Println(car.start())
}

func main() {
	j := []byte(`{"car": "bmw", "year": 2020 }`)

	var car Car

	json.Unmarshal(j, &car)

	fmt.Println(car)

	car1 := Car{
		CarName: "Palio",
		CarYear: 2008,
	}

	car1.drive()
	println(car1.CarName)

	result, _ := json.Marshal(car1)

	fmt.Println(string(result))

	example(car1)
}
