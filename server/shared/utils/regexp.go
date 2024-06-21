package utils

import "regexp"

func IsValidRegexp(regexpStr, str string) bool {
	re, err := regexp.Compile(regexpStr)
	if err != nil {
		return false
	}
	return re.MatchString(str)
}
