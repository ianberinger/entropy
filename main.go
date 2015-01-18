package main

import (
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()
	arg0 := flag.Arg(0)

	filePath, err := filepath.Abs(arg0)
	if err != nil {
		fmt.Println(err)
		return
	}
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	entropy, err := Entropy(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Entropy of", filePath, ":", entropy)
}

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
