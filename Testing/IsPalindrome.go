package test

func removeSpace(s string) string {
	temp := ""
	for _, c := range s {
		if c != ' ' {
			temp = temp + string(c)
		}
	}
	return temp
}

func lowerCase(s string) string {
	temp := ""
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			temp = temp + string(c)
		} else if c >= 'A' && c <= 'Z' {
			temp = temp + string(c+32)
		}
	}
	return temp
}

func IsPalindrome(s string) bool {
	s = removeSpace(s)
	s = lowerCase(s)
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
