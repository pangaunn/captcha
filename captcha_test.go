package captcha_test

import (
	"testing"

	"github.com/pangaunn/captcha"
)

func TestCaptcha1PlusOne(t *testing.T) {
	expected := "1 + one"
	result := captcha.Captcha(1, 1, 1, 1)
	if expected != result {
		t.Errorf("It should say %s but get %s", expected, result)
	}
}

func TestCaptcha1MinusOne(t *testing.T) {
	expected := "1 - one"
	result := captcha.Captcha(1, 1, 2, 1)
	if expected != result {
		t.Errorf("It should say %s but get %s", expected, result)
	}
}

func TestCaptcha1MultiplyOne(t *testing.T) {
	expected := "1 * one"
	result := captcha.Captcha(1, 1, 3, 1)
	if expected != result {
		t.Errorf("It should say %s but get %s", expected, result)
	}
}

func TestCaptcha1MultiplyTwo(t *testing.T) {
	expected := "1 * two"
	result := captcha.Captcha(1, 1, 3, 2)
	if expected != result {
		t.Errorf("It should say %s but get %s", expected, result)
	}
}
