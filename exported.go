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
type MgcNumber uint32

type Mgc struct {
	convert_to_int []uint32
	convert_to_mgc []MgcNumber
}

// Initialize a sequence of numbers, of length 'numBits'
// The tables are pre computed to speed up conversion
func New(numBits uint32) *Mgc {
	var ret Mgc
	l := 1 << numBits
	ret.convert_to_int = make([]uint32, l)
	ret.convert_to_mgc = make([]MgcNumber, l)
	m := monotonic(int(numBits))
	for i, mgc := range m {
		v := getValue(mgc)
		ret.convert_to_int[v] = uint32(i)
		ret.convert_to_mgc[i] = MgcNumber(v)
	}
	return &ret
}

// Convert a Monotonic Gray Code to the corresponding integer. Init() has to called first.
func (m *Mgc) GetInt(n MgcNumber) uint32 {
	return m.convert_to_int[n]
}

// Get the n:th Monotonic Gray Code. Init() has to called first.
func (m *Mgc) GetMgc(n uint32) MgcNumber {
	return m.convert_to_mgc[n]
}
