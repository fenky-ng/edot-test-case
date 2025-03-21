package regexp

import "regexp"

func IsEmail(input string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(input)
}

func IsPhone(input string) bool {
	phoneRegex := `^\+?\d{7,15}$`
	re := regexp.MustCompile(phoneRegex)
	return re.MatchString(input)
}
