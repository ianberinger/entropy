# entropy
[![Build Status](https://travis-ci.org/ianberinger/entropy.svg?branch=master)](https://travis-ci.org/ianberinger/entropy) [![GoDoc](https://godoc.org/github.com/ianberinger/entropy?status.svg)](https://godoc.org/github.com/ianberinger/entropy/entropy)

entropy calculation package & commandline tool

### Install
	$go get github.com/ianberinger/entropy

### Usage
Entropy takes a relative or absolute filepath as input e.g.,

	$entropy ../foo.txt

and outputs the entropy (between 0 and 1) of that file.

There is also a prettyprint flag:

	$entropy -pretty ../foo.txt

You can use entropy in your own go packages:

	import "github.com/ianberinger/entropy/entropy"

See [godoc](http://godoc.org/github.com/ianberinger/entropy/entropy) for documentation.
