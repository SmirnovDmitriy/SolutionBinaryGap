package Smirnov

import (
	"fmt"
	"strconv"
	"strings"
)

func SolutionBinaryGap(N int) int {
	x := strconv.FormatInt(int64(N), 2)
	array := strings.Split(x, "")
	fmt.Println(array)
	var zero []int
	var q int
	for i := 0; i < len(array); i++ {
		if array[i] == "0" {
			q++
		}
		if array[i] != "0" {
			zero = append(zero, q)
			q = 0
		}
	}
	zero = append(zero, q)
	fmt.Println(zero)
	var res int
	for i := 0; i < len(zero); i++ {
		if res < zero[i] {
			res = zero[i]
		}
	}
	return res
}
func main() {
	fmt.Println(SolutionBinaryGap(2352))
}
