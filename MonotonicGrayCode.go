// Implementation based on python script http://sciyoshi.com/2010/12/gray-codes/

package main

import "fmt"

func rotate_right(x []int, n int) []int {
	l := len(x)
	return append(x[l-n:], x[0:l-n]...)
}

func pi(n int) []int {
	if n <= 1 {
		return []int{0}
	}
	x := append(pi(n-1), n-1)
	x2 := make([]int, len(x))
	for i, k := range x {
		x2[i] = x[k]
	}
	return rotate_right(x2, 1)
}

func p(n int, j int, reverse bool) (ret [][]int) {
	if n == 1 && j == 0 {
		if !reverse {
			ret = [][]int{{0}, {1}}
		} else {
			ret = [][]int{{1}, {0}}
		}
	} else if j >= 0 && j < n {
		perm := pi(n - 1)
		if !reverse {
			for _, x := range p(n-1, j-1, false) {
				t := []int{1}
				for _, k := range perm {
					t = append(t, x[k])
				}
				ret = append(ret, t)
			}
			for _, x := range p(n-1, j, false) {
				t := append([]int{0}, x...)
				ret = append(ret, t)
			}
		} else {
			for _, x := range p(n-1, j, true) {
				t := append([]int{0}, x...)
				ret = append(ret, t)
			}
			for _, x := range p(n-1, j-1, true) {
				t := []int{1}
				for _, k := range perm {
					t = append(t, x[k])
				}
				ret = append(ret, t)
			}
		}
	}
	return ret
}

func monotonic(n int) (ret [][]int) {
	for i := 0; i < n; i++ {
		var p2 [][]int
		if i%2 == 0 {
			p2 = p(n, i, false)
		} else {
			p2 = p(n, i, true)
		}
		ret = append(ret, p2...)
	}
	return
}

func main() {
	for i := 1; i < 5; i++ {
		fmt.Println("monotonic(", i, "):", monotonic(i))
	}
}
