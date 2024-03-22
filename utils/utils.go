package utils

func isLetter(ch byte) bool {
	if 'A' <= ch && 'Z' >= ch {
		return true
	}

	if 'a' <= ch && 'z' >= ch {
		return true
	}

	return false
}

func isSymbol(ch byte) bool {
	if isLetter(ch) {
		return true
	}

	if '0' <= ch && '9' >= ch {
		return true
	}

	if ch == '_' || ch == '-' {
		return true
	}

	return false
}

func isWord(s string) bool {
	if len(s) == 0 {
		return false
	}

	for i := 0; i < len(s); i++ {
		if !isSymbol(s[i]) {
			return false
		}
	}

	return true
}

func isPrefix(s string) bool {
	if len(s) == 0 {
		return false
	}

	var t string = ""
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			if !isWord(t) {
				return false
			}
			t = ""
		} else {
			t += string(s[i])
		}
	}

	return true
}

func isDomain(s string) bool {
	if !(len(s) == 2 || len(s) == 3) {
		return false
	}

	for i := 0; i < len(s); i++ {
		if !isLetter(s[i]) {
			return false
		}
	}

	return true
}

func findLastChar(s string, ch byte) int {
	index := -1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ch {
			index = i
			break
		}
	}
	return index
}

func prefStr(s string, index int) string {
	pref := ""
	for i := 0; i < index; i++ {
		pref += string(s[i])
	}
	return pref
}

func suffStr(s string, index int) string {
	suff := ""
	for i := index + 1; i < len(s); i++ {
		suff += string(s[i])
	}
	return suff
}

func isSuffix(s string) bool {
	index := findLastChar(s, '.')
	pref := prefStr(s, index)
	domain := suffStr(s, index)

	if !isPrefix(pref) {
		return false
	}

	if !isDomain(domain) {
		return false
	}

	return true
}

// IsEmail checks syntax of an email is valid or not based on some standards
func IsEmail(s string) bool {
	index := findLastChar(s, '@')
	pref := prefStr(s, index)
	suff := suffStr(s, index)

	if !isPrefix(pref) {
		return false
	}

	if !isSuffix(suff) {
		return false
	}

	return true
}
