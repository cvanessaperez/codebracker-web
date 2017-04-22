package CodeBreaker

import "testing"

func TestAllX(t *testing.T) {
	expected := "xxxx"
	actual := guess("1234")
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got: '%s'", expected, actual)
	}
}
func TestNot(t *testing.T) {
	expected := ""
	actual := guess("5678")
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got: '%s'", expected, actual)
	}
}

func TestThreeX(t *testing.T) {
	expected := "xxx"
	actual := guess("1235")
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got: '%s'", expected, actual)
	}
}

func TestTwoX(t *testing.T) {
	expected := "xx"
	actual := guess("1635")
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got: '%s'", expected, actual)
	}
}

func TestOneX(t *testing.T) {
	expected := "x"
	actual := guess("8635")
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got: '%s'", expected, actual)
	}
}
func TestFour_(t *testing.T) {
	expected := "----"
	actual := guess("4321")
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got: '%s'", expected, actual)
	}
}
func TestX_(t *testing.T) {
	expected := "x-"
	actual := guess("5248")
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got: '%s'", expected, actual)
	}
}
