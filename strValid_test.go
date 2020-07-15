package strValid

import "testing"

const errorString = "text not match: "

func TestNumberAndSpecialCharacter(t *testing.T) {
	correctString := "TextMatch00and@$"
	wrongString := "TextNotMatch0"

	if !IsValid(correctString, NumberAndSpecialCharacter) {
		t.Errorf(errorString + "text: " + correctString)
	}

	if IsValid(wrongString, NumberAndSpecialCharacter) {
		t.Errorf(errorString + "text: " + wrongString)
	}
}

func TestUpperCaseAndSpecialCharacter(t *testing.T) {
	correctString := "TextCont@ins@"
	wrongString := "textnotmatch"

	if !IsValid(correctString, UpperCaseAndSpecialCharacter) {
		t.Errorf(errorString + "text: " + correctString)
	}

	if IsValid(wrongString, UpperCaseAndSpecialCharacter) {
		t.Errorf(errorString + "text: " + wrongString)
	}
}

func TestUpperCaseAndNumber(t *testing.T) {
	correctString := "TextContainsNumber10"
	wrongString := "textnotcontainsnumber"

	if !IsValid(correctString, UpperCaseAndNumber) {
		t.Errorf(errorString + "text: " + correctString)
	}

	if IsValid(wrongString, UpperCaseAndNumber) {
		t.Errorf(errorString + "text: " + wrongString)
	}
}
