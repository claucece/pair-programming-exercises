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

func (x number) Divide(divident number) float32 {
	defer func() {
		if err := recover(); err != nil {
			panic("integer divide by zero runtime error: integer divide by zero")
		}
	}()

	return (float32(x) / float32(divident))
}

func (x number) DivideInt(divident number) number {
	defer func() {
		if err := recover(); err != nil {
			panic("integer divide by zero runtime error: integer divide by zero")
		}
	}()

	return x / divident
}
