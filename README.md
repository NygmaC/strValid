# StrValid

## How works

Function Validate() expect a string parameter, this function validate all options
includes a special character, number and uppercase letters.

Function IsValid() expect two parameters, first parameter is a string,the second
is a option validation.

###### Validation Options

UpperCase                    --> only uppercase letters
Number                       --> only numbers
UpperCaseAndNumber           --> uppercase letters and numbers
SpecialCharacter             --> only special character
UpperCaseAndSpecialCharacter --> uppercase letters and special character
NumberAndSpecialCharacter    --> numbers and special character

## Example

import (
	"fmt"
	"github.com/NygmaC/strValid"
)

func main() {

	textExample := "TextMatch10$@"

	valid := strValid.ValidateAll(textExample)

	fmt.Println(valid)
	
	//Options
	/*UpperCase
	Number
	UpperCaseAndNumber
	SpecialCharacter
	UpperCaseAndSpecialCharacter
	NumberAndSpecialCharacter*/

	textExample = "Text2020"
	valid = strValid.IsValid(textExample, strValid.UpperCaseAndNumber)

	fmt.Println(valid)
}



