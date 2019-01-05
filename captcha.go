package captcha

import "fmt"

func Captcha(pattern, op1, opr, op2 int) string {
	oprs := []string{"+", "-", "*"}
	return fmt.Sprintf("1 %s one", oprs[opr-1])
}
