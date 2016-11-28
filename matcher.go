package errors

type Matcher interface {
	Match(required ...*Kind) bool
}

func Is(err error, matchers ...Matcher) bool {
	if err == nil {
		return false
	}

	e, ok := err.(*Error)
	if !ok {
		return false
	}

	if e.Child != nil && Is(e.Child, matchers...) {
		return true
	}

	for _, m := range matchers {
		if m.Match(e.Kind) {
			return true
		}
	}

	return false
}

func (k *Kind) matchChild(required ...*Kind) bool {

	return false
}
