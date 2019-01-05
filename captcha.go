package captcha

import "fmt"

func Captcha(pattern, op1, opr, op2 int) string {
	oprs := []string{"+", "-", "*"}
	ops := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	s2 := ops[op2-1]

	var s1 string
	if op1 == 1 {
		s1 = "1"
	} else {
		s1 = "5"
	}
	return fmt.Sprintf("%s %s %s", s1, oprs[opr-1], s2)
}
