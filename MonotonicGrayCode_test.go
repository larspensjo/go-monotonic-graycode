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

package monotonic_gray_code

import (
	"fmt"
	"testing"
)

func main() {
	for i := 1; i < 17; i++ {
		fmt.Println("monotonic(", i, "):", len(monotonic(i)))
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
	power := 2
	for iter := 1; iter < 10; iter++ {
		m := monotonic(iter)
		// The number of elements shall double for every next iteration
		if len(m) != power {
			t.Error("Unexpected length monotonic", iter, ":", len(m))
		}
		power *= 2

		// An important property is "monotonic". In this case, it means that we do not want
		// two consequitive numbers to decrease number of bits, except at most by 1.
		highest := 0
		for _, e := range m {
			count := SummarizeBits(e)
			delta := count - highest
			if delta > 1 || delta < -1 {
				t.Error("Not monotonic number for", iter)
			}
			if count > highest {
				highest = count
			}
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
