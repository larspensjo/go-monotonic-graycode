// Copyright 2014 Lars Pensjö
//
// monotonic_gray_code  is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, version 3.
//
// monotonic_gray_code is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with monotonic_gray_code.  If not, see <http://www.gnu.org/licenses/>.
//

// Create a sequence of Monotonic Gray Codes.
// Implementation is based the Savage-Winkler algorithm, see python script at http://sciyoshi.com/2010/12/gray-codes/.
package mgc

func pi(n int) []int {
	if n <= 1 {
		return []int{0}
	}
	x := append(pi(n-1), n-1)
	// Scramble the order
	x2 := make([]int, len(x))
	for i, k := range x {
		dest := i + 1
		if dest == len(x) {
			dest = 0
		}
		x2[dest] = x[k]
	}
	return x2
}

func p(n int, j int, reverse bool) (ret [][]int) {
	if n == 1 && j == 0 {
		if !reverse {
			ret = [][]int{{0}, {1}}
		} else {
			ret = [][]int{{1}, {0}}
		}
		return
	}
	if j < 0 || j >= n {
		return
	}
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
	return
}

func monotonic(n int) (ret [][]int) {
	if n < 0 {
		panic("Illegal argument")
	}
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
		power *= 2
	}
	return
}
