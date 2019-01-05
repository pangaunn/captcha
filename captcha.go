package captcha

import "fmt"

func Captcha(pattern, op1, opr, op2 int) string {
	oprs := []string{"+", "-", "*"}

	var ops2 string
	if op2 == 1 {
		ops2 = "one"
	} else {
		ops2 = "two"
	}

	return fmt.Sprintf("1 %s %s", oprs[opr-1], ops2)
}
