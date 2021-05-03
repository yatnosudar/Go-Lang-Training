package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("e([a-z])o")

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("ewo"))
	fmt.Println(regex.MatchString("eDo"))

	search := regex.FindAllString("eko eka edo ewo lili eto", -1)
	fmt.Println(search)
}
