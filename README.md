# entropy
[![GoDoc](https://godoc.org/github.com/ianberinger/entropy?status.svg)](https://godoc.org/github.com/ianberinger/entropy)

entropy calculation package & commandline tool

### Install
	$go get github.com/ianberinger/entropy
	
### Usage
Entropy takes a relative or absolute filepath as input e.g.,
	
	$entropy ../foo.txt
	
and outputs the entropy (0,1) of that file.

You can use entropy in your own go packages:
	
	import "github.com/ianberinger/entropy"

See [godoc](http://godoc.org/github.com/ianberinger/entropy) for documentation.