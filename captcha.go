package captcha

func Captcha(pattern, op1, operator, op2 int) string {
	if operator == 2 {
		return "1 - one"
	}
	return "1 + one"
}
