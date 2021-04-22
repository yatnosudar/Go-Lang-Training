package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpEko NoKTP = "113124124124"
	var statusMarried = true

	fmt.Println(noKtpEko)
	fmt.Println(statusMarried)
}
