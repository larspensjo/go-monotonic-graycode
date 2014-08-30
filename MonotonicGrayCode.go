// Copyright 2014 Lars Pensjö
//
// Monotonic Gray Code  is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, version 3.
//
// Ephenation is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Ephenation.  If not, see <http://www.gnu.org/licenses/>.
//

// Implementation based on python script http://sciyoshi.com/2010/12/gray-codes/

package monotonic_gray_code

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

// Convert a bit array to a number
// Strictly speaking, the bits are reversed, but that shouldn't matter.
func getValue(number []int) (ret int) {
	power := 1
	for _, bit := range number {
		ret += power * bit
	}
	return
}
