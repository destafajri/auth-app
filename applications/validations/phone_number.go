package validations

import (
	"fmt"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

func Phonenumber(input string) bool {
	if len([]rune(input)) < 9 {
		fmt.Println("phone number length")
		return false
	}
	
	regex := `^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`
	err := validation.Validate(input,
		validation.Required.Error("is required"),
		validation.Match(regexp.MustCompile(regex)).
		Error("must be a be phone number"),
	)

	if err != nil{
		fmt.Println(err)
		return false
	}

	return true
}