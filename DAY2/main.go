package main

import (
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type levels struct {
	Val []int64
}

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
	// separate for \n
	var resp int64 = 0
	var respTolerant = 0
	repoList := strings.Split(string(dump), "\n")
	for _, repo := range repoList {
		// fetch single report
		stringLevels := strings.Split(repo, " ")
		l := makeLevels(stringLevels)

		if l.checkAdjTolerant() {
			respTolerant += 1
		}
	}
	println(resp, respTolerant)
}

// 0 1 2 3 4 5
// 6 5 4 2 1 1

func (l levels) checkAdj(index int) bool {

	for i := index; i < len(l.Val); i++ {
		// leftmost value
		if i == 0 {
			dif := l.Val[i+1] - l.Val[i]
			if !(math.Abs(float64(dif)) <= 3.0 && math.Abs(float64(dif)) >= 1.0) {
				return false
			}
		} else if i == len(l.Val)-1 {
			dif := l.Val[i] - l.Val[i-1]
			if !(math.Abs(float64(dif)) <= 3.0 && math.Abs(float64(dif)) >= 1.0) {
				return false
			}
		} else {
			leftDif := l.Val[i] - l.Val[i-1]
			rightDif := l.Val[i] - l.Val[i+1]
			if !((l.Val[i-1] < l.Val[i]) && (l.Val[i] < l.Val[i+1]) || (l.Val[i-1] > l.Val[i]) && (l.Val[i] > l.Val[i+1])) {
				return false
			} else if !(math.Abs(float64(leftDif)) <= 3.0 && math.Abs(float64(leftDif)) >= 1.0) ||
				!(math.Abs(float64(rightDif)) <= 3.0 && math.Abs(float64(rightDif)) >= 1.0) {
				return false
			}
		}
	}
	return true
}

func (l levels) checkAdjTolerant() bool {
	var tripWire = 0
	for i, _ := range l.Val {
		// leftmost value
		if i == 0 {
			dif := l.Val[i+1] - l.Val[i]
			if !(math.Abs(float64(dif)) <= 3.0 && math.Abs(float64(dif)) >= 1.0) {
				tripWire += 1
			}
		} else if i == len(l.Val)-1 {
			dif := l.Val[i] - l.Val[i-1]
			if !(math.Abs(float64(dif)) <= 3.0 && math.Abs(float64(dif)) >= 1.0) {
				tripWire += 1

			}
		} else {
			leftDif := l.Val[i] - l.Val[i-1]
			rightDif := l.Val[i] - l.Val[i+1]
			if !((l.Val[i-1] < l.Val[i]) && (l.Val[i] < l.Val[i+1]) || (l.Val[i-1] > l.Val[i]) && (l.Val[i] > l.Val[i+1])) {
				tripWire += 1
			} else if !(math.Abs(float64(leftDif)) <= 3.0 && math.Abs(float64(leftDif)) >= 1.0) ||
				!(math.Abs(float64(rightDif)) <= 3.0 && math.Abs(float64(rightDif)) >= 1.0) {
				tripWire += 1
			}
		}
		if tripWire == 1 {
			var innerLevel levels
			innerLevel.Val = append(l.Val[:i], l.Val[i+1:]...)
			return innerLevel.checkAdj(0)
		}
	}
	return true
}

func makeLevels(a []string) levels {
	var level levels
	level.Val = make([]int64, len(a))
	for i, j := range a {
		v, _ := strconv.ParseInt(j, 10, 64)
		level.Val[i] = v
	}
	return level
}
