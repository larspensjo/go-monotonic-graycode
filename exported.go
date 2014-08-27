package monotonic_gray_code

import ()

var convert_to_int []int
var convert_to_mgc []int

// Initialize a sequence of numbers, of length 'count'
// The tables are pre computed to speed up conversion
func Init(count int) {
	convert_to_int = make([]int, count)
	convert_to_mgc = make([]int, count)
	m := monotonic(count)
	for i, mgc := range m {
		convert_to_int[getValue(mgc)] = i
		convert_to_mgc[i] = getValue(mgc)
	}
}
