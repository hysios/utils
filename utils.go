package utils

import "strings"

func SplitLast(s string, c rune, n int) (out []string) {
	var (
		ss = strings.Split(s, string(c))
		l  = len(ss)
	)
	if l <= n {
		return ss
	}

	r := l - n
	out = []string{strings.Join(ss[:r+1], string(c))}
	out = append(out, ss[r+1:]...)
	return
}
