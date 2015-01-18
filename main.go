package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ianberinger/entropy/entropy"
)

func main() {
	pretty := flag.Bool("pretty", false, "activate prettyprint")
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
	entropy, err := entropy.Entropy(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	if *pretty {
		fmt.Printf("Entropy of %s: %f\n", filePath, entropy)
	} else {
		fmt.Println(entropy)
	}

}
