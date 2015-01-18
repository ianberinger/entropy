//Package entropy calculates entropy values.
//Example usage:
//      file,_ := os.Open("foo.txt")
//      v, err := entropy.Entropy(file) // v is now the entropy value of "foo.txt" e.g., 0.38498
package entropy

import (
	"io"
	"math"
)

// Entropy takes any io.Reader and calculates its entropy
// It returns the entropy value (between 0 and 1) and any read error encountered.
func Entropy(r io.Reader) (entropy float32, err error) {
	var bytes [256]int
	var readSize int
	buf := make([]byte, 1)

	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return 0, err
		}
		if n == 0 {
			break
		}
		readSize++
		bytes[buf[0]]++
	}

	for _, e := range bytes {
		if e > 0 {
			p := float64(e) / float64(readSize)
			entropy = entropy - float32(p*math.Log2(p))
		}
	}
	entropy = entropy / 8
	return
}
