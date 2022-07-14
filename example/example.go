package example

// Struct that starts with lowercase are not visible for consumers
type House struct {
	Color  string
	Number int
}

func PrintExample() {
	privateFunction()
}

// Functions that starts with lowercase are not visible for consumers
func privateFunction() {
	print("private function")
}
