package main

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// load file
	file, err := os.Open("input.txt")
	if err != nil {
		return
	}
	dump, err := io.ReadAll(file)
	if err != nil {
		return
	}
	splitCouple := strings.Split(dump, "\n")
	var f []int64
	var s []int64
	for _, ss := range splitCouple {
		t := strings.Split(ss, "   ")
		f1, _ := strconv.ParseInt(t[0], 10, 32)
		s1, _ := strconv.ParseInt(t[1], 10, 32)
		f = append(f, f1)
		s = append(s, s1)
	}
	slices.Sort(f)
	slices.Sort(s)
	var resp int64 = 0
	for i, _ := range f {
		g := f[i] - s[i]
		if g <= 0 {
			resp += int64(math.Abs(float64(g)))
		} else {
			resp += g
		}
	}
	println(resp)

	// part 2
	// extract set of first list
	var fMap map[int64]bool
	fMap = make(map[int64]bool)
	var sMap map[int64]int64
	sMap = make(map[int64]int64)
	for _, el := range f {
		fMap[el] = true
	}
	// extract count
	for _, el := range s {
		if val, ok := sMap[el]; ok {
			sMap[el] = val + 1
		} else {
			sMap[el] = 1
		}
	}
	// count similarity
	var similarity int64 = 0
	for k, _ := range fMap {
		similarity += sMap[k] * (k)
	}
	println(similarity)
}
