package helper

type Context struct {
	intSlice []int
}

func (c *Context) getIntSlice(l int) []int {
	if cap(c.intSlice) < l {
		c.intSlice = make([]int, l)
	}
	return c.intSlice[:l]
}

func (c *Context) Distance(str1, str2 string) int {
	s1 := []rune(str1)
	s2 := []rune(str2)

	len_s1 := len(s1)
	len_s2 := len(s2)

	if len_s2 == 0 {
		return len_s1
	}

	column := c.getIntSlice(len_s1 + 1)

	for i := 1; i <= len_s1; i++ {
		column[i] = i
	}

	for x := 0; x < len_s2; x++ {
		s2Rune := s2[x]
		column[0] = x + 1
		lastdiag := x

		for y := 0; y < len_s1; y++ {
			olddiag := column[y+1]
			cost := 0
			if s1[y] != s2Rune {
				cost = 1
			}
			column[y+1] = min(
				column[y+1]+1,
				column[y]+1,
				lastdiag+cost,
			)
			lastdiag = olddiag
		}
	}

	return column[len_s1]
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}

func LevenshteinDistance(s1, s2 string) int {
	c := Context{}
	return c.Distance(s1, s2)
}
