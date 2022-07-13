package main

import (
	"fmt"
)

func main() {
	mapx := make(map[string]int)

	mapx["gabriel"] = 1
	mapx["gigante"] = 2

	fmt.Println(mapx["gabriel"])
}
