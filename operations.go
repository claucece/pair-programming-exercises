package operations

type number int

func (x number) Add(y number) number {
	return x + y
}

func substraction(a, b int) int {
	return a - b
}

func (x number) substr(b number) number {
	return x - b
}

func (x number) Multiplication(y number) number {
	return x * y
}

func (x number) Divide(divident number) (float32, string) {
	return float32(x / divident), ""
}
