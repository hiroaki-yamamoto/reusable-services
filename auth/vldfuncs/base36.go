package vldfuncs

import (
	"log"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var pattern *regexp.Regexp

func init() {
	var err error
	const patternTxt = "^[0-9,A-Z]*$"
	pattern, err = regexp.Compile("^[0-9,A-Z]*$")
	if err != nil {
		log.Println("Failed to compile the pattern:", patternTxt)
	}
}

// Base36 checks whther the text is following BASE36 astring.
func Base36(fl validator.FieldLevel) bool {
	return pattern.Match([]byte(fl.Field().String()))
}
