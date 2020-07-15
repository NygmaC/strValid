package strValid

import (
	"sync"
)

func ValidateAll(text string) bool {
	return isValidAll(text)
}

func IsValid(text string, optionSelect string) bool {

	switch optionSelect {
	case UpperCaseAndSpecialCharacter:
		return upperCaseSpecialCaracter(text)

	case UpperCaseAndNumber:
		return upperCaseNumber(text)
	case UpperCase:

		return isValidUpperCase(text)

	case Number:

		return isValidNumber(text)

	case NumberAndSpecialCharacter:

		return numberSpecialCharacter(text)

	case SpecialCharacter:

		return isValidSpecialCaracter(text)

	default:
		return false
	}

}

func isValidAll(text string) bool {
	channel := make(chan bool, 3)
	var wg sync.WaitGroup
	wg.Add(3)

	go isValidUpperCaseValidator(text, &wg, channel)
	go isValidSpecialCharacterValidator(text, &wg, channel)
	go isValidNumberValidator(text, &wg, channel)

	go func() {
		wg.Wait()
		close(channel)
	}()

	for resp := range channel {
		if !resp {
			return resp
		}
	}
	return true
}

func numberSpecialCharacter(text string) bool {
	channel := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(2)

	go isValidUpperCaseValidator(text, &wg, channel)
	go isValidSpecialCharacterValidator(text, &wg, channel)

	go func() {
		wg.Wait()
		close(channel)
	}()

	for resp := range channel {
		if !resp {
			return resp
		}
	}
	return true
}

func upperCaseNumber(text string) bool {
	channel := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(2)

	go isValidUpperCaseValidator(text, &wg, channel)
	go isValidNumberValidator(text, &wg, channel)

	go func() {
		wg.Wait()
		close(channel)
	}()

	for resp := range channel {
		if !resp {
			return resp
		}
	}
	return true
}

func upperCaseSpecialCaracter(text string) bool {
	channel := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(2)

	go isValidUpperCaseValidator(text, &wg, channel)
	go isValidSpecialCharacterValidator(text, &wg, channel)

	go func() {
		wg.Wait()
		close(channel)
	}()

	for resp := range channel {
		if !resp {
			return resp
		}
	}
	return true
}

func isValidNumber(text string) bool {
	channel := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(1)

	go isValidNumberValidator(text, &wg, channel)

	go func() {
		wg.Wait()
		close(channel)
	}()

	for resp := range channel {
		if !resp {
			return resp
		}
	}
	return true
}

func isValidSpecialCaracter(text string) bool {
	channel := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(1)

	go isValidSpecialCharacterValidator(text, &wg, channel)

	go func() {
		wg.Wait()
		close(channel)
	}()

	for resp := range channel {
		if !resp {
			return resp
		}
	}
	return true
}

func isValidUpperCase(text string) bool {
	channel := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(1)

	go isValidUpperCaseValidator(text, &wg, channel)

	go func() {
		wg.Wait()
		close(channel)
	}()

	for resp := range channel {
		if !resp {
			return resp
		}
	}
	return true
}
