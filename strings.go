package main

import (
	"fmt"
	s "strings"
	// for the whole package go to
	// https://pkg.go.dev/strings
)

var p = fmt.Println

var str = "This is a string"

var ending = "ing"

var start = "Th"

var another = "This is another string"

var contained = "is"

func main() {

	p("Contains:", s.Contains(str, contained))
	p("Contains:", s.ContainsAny(str, "oiuyh"))
	p("Contains:", s.ContainsAny(str, "p")) // return false
	p("Contains:", s.ContainsAny("", "")) // returns false which is absurd
	p("ReplaceAll:", s.ReplaceAll(str, "i", "e")) // replaces i for e in all instances
}
