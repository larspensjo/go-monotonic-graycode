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

import (
	"fmt"
	"testing"
)

func ExamplePi() {
	for i := 1; i < 17; i++ {
		fmt.Println("pi(", i, "):", pi(i))
	}
}

func ExampleMonotonic() {
	m := monotonic(5)
	for _, n := range m {
		fmt.Println(n)
	}
}

// Count number of 1:s
func SummarizeBits(vector []int) (count int) {
	for _, element := range vector {
		if element == 1 {
			count++
		}
	}
	return
}

func TestInitial(t *testing.T) {
	for numBits := 1; numBits < 10; numBits++ {
		m := monotonic(numBits)
		if len(m) != 1<<uint32(numBits) {
			t.Error("Unexpected length monotonic", numBits, ":", len(m))
		}

		// An important property is "monotonic". In this case, it means that we do not want
		// two consequitive numbers to decrease number of bits, except at most by 1.
		highest := 0
		for _, e := range m {
			count := SummarizeBits(e)
			delta := count - highest
			if delta > 1 || delta < -1 {
				t.Error("Not monotonic number for", numBits)
			}
			if count > highest {
				highest = count
			}
		}
	}
}

// Test that converting an integer to a MGC number, and back, gives back the original number
func TestConversion(t *testing.T) {
	const bits = 16
	m := New(bits)
	const max = 1 << bits
	for i := int32(0); i < max; i++ {
		a := m.GetMgc(i)
		b := m.GetInt(a)
		if b != i {
			t.Error("Converting int", i, "gives mgc", a, "and back, gave", b)
			t.FailNow() // No need to go through all the others
		}
	}
}

// Measure a couple of use cases
func BenchmarkGenerate10(t *testing.B) {
	for i := 0; i < t.N; i++ {
		monotonic(10)
	}
}
func BenchmarkGenerate14(t *testing.B) {
	for i := 0; i < t.N; i++ {
		monotonic(14)
	}
}
func BenchmarkGenerate16(t *testing.B) {
	for i := 0; i < t.N; i++ {
		monotonic(16)
	}
}
func BenchmarkGenerate18(t *testing.B) {
	for i := 0; i < t.N; i++ {
		monotonic(18)
	}
}
