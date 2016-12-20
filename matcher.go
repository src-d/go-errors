package errors

// Matcher match a given error
type Matcher interface {
	// Is return true if the err match
	Is(err error) bool
}

// Is check if err match all matchers
func Is(err error, matchers ...Matcher) bool {
	var are int
	for _, m := range matchers {
		if m.Is(err) {
			are++
		}
	}

	wanted := len(matchers)
	return wanted == are
}

// Any check if err match any matchers
func Any(err error, matchers ...Matcher) bool {
	for _, m := range matchers {
		if m.Is(err) {
			return true
		}
	}

	return false
}
