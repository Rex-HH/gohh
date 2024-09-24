package main

import (
	"fmt"
)

func shortestToChar(s string, c byte) []int {
	n := len(s)
	res := make([]int, n)
	for i := range res {
		res[i] = n
	}
	for i := 0; i < n; i++ {
		if s[i] == c {
			res[i] = 0
		} else if i > 0 {
			res[i] = res[i-1] + 1
		}
	}
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && res[i+1]+1 < res[i] {
			res[i] = res[i+1] + 1
		}
	}

	return res
}

func main() {
	s := "loveleetcode"
	c := byte('e')
	fmt.Println(shortestToChar(s, c))

}
