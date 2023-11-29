package main

func head(str string) string {
	return string(str[0])
}

func tail(str string) string {
	return string(str[1:])
}

// https://en.m.wikipedia.org/wiki/Levenshtein_distance
func lev(a, b string) int {
	if len(a) == 0 {
		return len(b)
	}

	if len(b) == 0 {
		return len(a)
	}

	if head(a) == head(b) {
		return lev(tail(a), tail(b))
	}

	return 1 + min(lev(tail(a), b), lev(a, tail(b)), lev(tail(a), tail(b)))
}
