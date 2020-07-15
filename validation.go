package strValid

import (
	"regexp"
	"sync"
)

const upperCasePattern = `.*[A-Z].*`
const numberPattern = `.*\d.*`
const specialCharacterPattern = `.*[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?].*`

var validUpperCaseRegex = regexp.MustCompile(upperCasePattern)
var validNumberRegex = regexp.MustCompile(numberPattern)
var validCharacterRegex = regexp.MustCompile(specialCharacterPattern)

func isValidNumberValidator(text string, wg *sync.WaitGroup, channel chan bool) {
	channel <- validNumberRegex.MatchString(text)
	wg.Done()
}

func isValidSpecialCharacterValidator(text string, wg *sync.WaitGroup, channel chan bool) {
	channel <- validCharacterRegex.MatchString(text)
	wg.Done()
}

func isValidUpperCaseValidator(text string, wg *sync.WaitGroup, channel chan bool) {
	channel <- validUpperCaseRegex.MatchString(text)
	wg.Done()
}
