package helper

func Calculator(a int, b int, operation string) int {
	switch operation {
	case "+":
		return a + b
	case "-":
		return a - b
	case "x":
		return a * b
	case "/":
		return a / b

	default:
		return 0
	}

}
