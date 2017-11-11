# Membuat Library
## Membuat direktory library
- > mkdir $GOPATH/src/github.com/user/stringutil

## Membuat file library
- > cd $GOPATH/src/github.com/user/stringutil
- > touch reverse.go

## Isi file nya
`  
// Package stringutil contains utility functions for working with strings.
package stringutil

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
`

## Build library
- > go build github.com/user/stringutil


## Menggunakan library
` package main

import (
	"fmt"

	"github.com/user/stringutil" //this same with path file from #build
)

func main() {
	fmt.Printf(stringutil.Reverse("!oG ,olleH"))
} `

## Install aplikasi 
- > go install github.com/user/hello

## Hasil
_ Hasil install akan membuat suatu file di pkg -> username -> repo -> file.a _

` bin/
    hello                 # command executable
pkg/
    darwin_amd64/       # this will reflect your OS and architecture
        github.com/user/
            stringutil.a  # package object
src/
    github.com/user/
        hello/
            hello.go      # command source
        stringutil/
            reverse.go    # package source
`
