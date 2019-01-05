package captcha

import (
	"fmt"
	"strconv"
)

func Captcha(pattern, op1, opr, op2 int) string {
	ops := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	oprs := [3]string{"+", "-", "*"}

	var s1 string
	if s1 = strconv.Itoa(op1); pattern == 2 {
		s1 = ops[op1-1]
	}

	var s2 string
	if s2 = strconv.Itoa(op2); pattern == 1 {
		s2 = ops[op2-1]
	}

	return fmt.Sprintf("%s %s %s", s1, oprs[opr-1], s2)
}
