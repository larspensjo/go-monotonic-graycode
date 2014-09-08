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

package mgc

import ()

// The type representation of a Monotonic Gray Code number
type Mgc int

var convert_to_int []int32
var convert_to_mgc []Mgc

// Initialize a sequence of numbers, of length 'count'
// The tables are pre computed to speed up conversion
func Init(count int) {
	convert_to_int = make([]int32, count)
	convert_to_mgc = make([]Mgc, count)
	m := monotonic(count)
	for i, mgc := range m {
		convert_to_int[getValue(mgc)] = int32(i)
		convert_to_mgc[i] = Mgc(getValue(mgc))
	}
}

// Convert a Monotonic Gray Code to the corresponding integer. Init() has to called first.
func GetInt(n Mgc) int32 {
	return convert_to_int[n]
}

// Get the n:th Monotonic Gray Code. Init() has to called first.
func GetMgc(n int32) Mgc {
	return convert_to_mgc[n]
}
