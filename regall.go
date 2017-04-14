package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"
)

var (
	prefix = flag.String("prefix", "", "リストの前置文字")
	suffix = flag.String("suffix", "", "リストの後置文字")
)

func init() {
	flag.Parse()
}

func main() {
	usable := ParseRune(flag.Arg(0))
	reg, err := regexp.Compile("^" + flag.Arg(1) + "$")
	if err != nil {
		fmt.Println("Regexp error arg:2")
		return
	}
	maxlen, err := strconv.Atoi(flag.Arg(2))
	if err != nil {
		fmt.Println("can't parse to int for max length arg:3.")
	}

	ind := make([]int, 1)
	for n := 0; ; n++ {
		s := IndexToStr(ind, usable)
		if len([]rune(s)) > maxlen {
			break
		}
		if reg.MatchString(s) {
			fmt.Println(*prefix + s + *suffix)
		}
		ind = NextIndex(ind, len(usable))
	}
}
func IndexToStr(i []int, rlist []rune) string {
	rv := make([]rune, len(i))
	for n, v := range i {
		rv[n] = rlist[v]
	}
	return string(rv)
}
func NextIndex(i []int, length int) []int {
	for n, v := range i {
		if v != length-1 {
			i[n]++
			return i
		}
		i[n] = 0
	}
	return append(i, 0)
}

var parsereg = regexp.MustCompile(".-.")

func ParseRune(s string) []rune {
	rv := []rune(parsereg.ReplaceAllString(s, ""))
	for _, v := range parsereg.FindAllString(s, -1) {
		r := []rune(v)
		rv = append(rv, RangeRunes(r[0], r[2])...)
	}
	return rv

}
func RangeRunes(start, end rune) []rune {
	if start > end {
		start, end = end, start
	}
	rv := make([]rune, 0, end-start)
	for c := start; c <= end; c++ {
		rv = append(rv, c)
	}
	return rv
}
