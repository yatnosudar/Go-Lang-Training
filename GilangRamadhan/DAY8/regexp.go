package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("e([a-z])o")

	fmt.Print(regex.MatchString("eko"))
	fmt.Print(regex.MatchString("eio"))
	fmt.Print(regex.MatchString("eOo"))

	search := regex.FindAllString("eko eka edo eta eyo", -1)
	fmt.Print(search)
}
