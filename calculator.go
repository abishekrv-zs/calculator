package calculator

type myError struct {
	message string
}

func (e myError) Error() string {
	return e.message
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, myError{"Division by zero occurred"}
	}
	return a / b, nil
}

func mod(a, b int) (int, error) {
	if b == 0 {
		return 0, myError{"Modulo by zero occurred"}
	}

	if b < 0 {
		return -(a % b), nil
	}
	return a % b, nil
}

func calculator(a, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return add(a, b), nil
	case "-":
		return sub(a, b), nil
	case "*":
		return mul(a, b), nil
	case "/":
		return div(a, b)
	case "%":
		return mod(a, b)
	default:
		return 0, myError{"Invalid operator"}
	}
}
