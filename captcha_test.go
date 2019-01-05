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

func TestCaptcha1PlusThree(t *testing.T) {
	expected := "1 + three"
	result := captcha.Captcha(1, 1, 1, 3)
	if expected != result {
		t.Errorf("It should say %s but get %s", expected, result)
	}
}

func TestCaptcha5PlusTwo(t *testing.T) {
	expected := "5 + two"
	result := captcha.Captcha(1, 5, 1, 2)
	if expected != result {
		t.Errorf("It should say %s but get %s", expected, result)
	}
}

func TestCaptcha4MinusTwo(t *testing.T) {
	expected := "4 - two"
	result := captcha.Captcha(1, 4, 2, 2)
	if expected != result {
		t.Errorf("It should say %s but get %s", expected, result)
	}
}

func TestCaptchaOnePlus1(t *testing.T) {
	expected := "one + 1"
	result := captcha.Captcha(2, 1, 1, 1)
	if expected != result {
		t.Errorf("It should say %s but get %s", expected, result)
	}
}
