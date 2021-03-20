package main

type SkipString struct {
	// Counter for candidate for change
	c    int8
	skip int8
	in   string
	out  []rune
}

func (ss *SkipString) TransformRune(pos int) {
	r := rune(ss.in[pos])

	if !isAlphanumeric(r) {
		ss.out[pos] = r
		return
	}

	if isDigit(r) {
		ss.out[pos] = r
		ss.c++
		return
	}

	if ss.c%ss.skip == 0 && isLowercase(r) || ss.c%ss.skip != 0 && isUppercase(r) {
		ss.out[pos] = changeCase(r)
		ss.c++
		return
	}

	ss.out[pos] = r
	ss.c++
}

func (ss *SkipString) GetValueAsRuneSlice() []rune {
	return []rune(ss.in)
}

func (ss SkipString) String() string {
	// return fmt.Sprintf("%s transformed every %dth character is %s\n", ss.in, ss.skip, string(ss.out))
	return string(ss.out)
}

func NewSkipString(i int8, s string) SkipString {
	return SkipString{c: 1, skip: i, in: s, out: make([]rune, len(s))}
}

// changeCase changes the rune r from case if is between the ASCII range of [a-z][A-Z]
// otherwise returns the same rune
func changeCase(r rune) rune {
	// There is 32 positions between a and A in ASCII range
	const STEP = 32

	if isUppercase(r) {
		return r + STEP
	}
	if isLowercase(r) {
		return r - STEP
	}

	return r
}

// isAlphanumeric returns a bool if the rune r is between the correct ASCII ranges
func isAlphanumeric(r rune) bool {
	if isLetter(r) || isDigit(r) {
		return true
	}

	return false
}

// isUppercase verifies if the rune is between the [A-Z] ASCII range
func isUppercase(r rune) bool {
	if 'A' <= r && r <= 'Z' {
		return true
	}

	return false
}

// isLowercase verifies if the rune is between the [a-z] ASCII range
func isLowercase(r rune) bool {
	if 'a' <= r && r <= 'z' {
		return true
	}

	return false
}

// isLetter verifies if the rune is between the [A-Z] or [a-z] ASCII ranges
func isLetter(r rune) bool {
	if isUppercase(r) || isLowercase(r) {
		return true
	}

	return false
}

// isDigit verifies if the rune is between the [0-9] ASCII range
func isDigit(r rune) bool {
	if '0' <= r && r <= '9' {
		return true
	}

	return false
}
