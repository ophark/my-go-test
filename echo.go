package main

import (
	"fmt"
	"os"
	"strings"
)

func echo() {
	var sep = " "
	var s = ""
	var l = len(os.Args)

	for i := 1; i < l-1; i++ {
		s += os.Args[i] + sep
	}
	s += os.Args[l-1]
	fmt.Println(s)
}

func echo1() {
	var sep = " "
	var s = ""

	for _, arg := range os.Args[1:] {
		s += arg + sep
	}
	fmt.Println(s)
}

func echo2() {
	fmt.Println(os.Args[1:])
	fmt.Println(strings.Join(os.Args[1:], " "))
}
