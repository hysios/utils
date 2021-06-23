package utils

import (
	"math"
	"strconv"
	"strings"
	"time"
)

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

func Timestamp(ts int64) time.Time {
	return time.Unix(ts/1000, (ts%1000)*1000*1000)
}

// ParseDatestr 解析数字日期格式，n 决定毫秒的位数
func ParseDatestr(s string, n int) (time.Time, error) {
	var (
		ms string
		dt string
	)
	if len(s) > 14 {
		dt = s[:14]
		ms = s[14:]
	} else {
		dt = s
	}

	t, err := time.Parse("20060102150405", dt)
	if err != nil {
		return time.Time{}, err
	}

	i, err := strconv.Atoi(ms)
	if err != nil {
		return time.Time{}, err
	}

	return t.Add(time.Duration(i) * time.Duration(math.Pow(10, 9-float64(n)))), nil
}
