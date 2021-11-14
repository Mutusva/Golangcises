package algorithms

func reverse(s string) string {
	ss := []rune(s)
	if len(ss) == 0 {
		return ""
	}

	if len(ss) == 1 {
		return s
	}

	return reverse(string(ss[1:])) + string(ss[0])
}

// cat
/*
cat
  at c
    t
*/
