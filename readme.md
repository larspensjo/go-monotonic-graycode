Monotonic Gray Codes
====================
[Gray codes](http://en.wikipedia.org/wiki/Gray_code) are sequences of binary numbers where only one bit change between every subsequent number. E.g. a 4 bit number can be:
```
0000 0001 0011 0010 0110 0111 0101 0100 1100 1101 1111 1110 1010 1011 1001 1000
```
[Monotonic Gray codes](http://en.wikipedia.org/wiki/Gray_code#Monotonic_Gray_codes) also have the
property that the total number of bits increase with the value of the number. E.g. a 4 bit bit numbers
can be:
```
0000 0001 0011 0010 0110 0100 1100 1000 1010 1011 1001 1101 0101 0111 1111 1110
```
Implementation is based on http://sciyoshi.com/2010/12/gray-codes/.

## Copyright and licensing
Distributed under GNU Lesser General Public License.

## Usage
Initialize with the number of bits. This must be done before any other functions are used.
```
	import 	"github.com/larspensjo/go-monotonic-graycode"
	numbers := mgc.New(numBits)
```
To get the Monotonic Gray Code for the integer 'n':
```
	numbers.GetMgc(n)
```
To convert a Monotonic Gray Code 'c' back to an integer:
```
	numbers.GetInt(c)
```
