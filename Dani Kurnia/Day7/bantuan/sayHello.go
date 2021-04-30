package bantuan

import "fmt"

var version = 1
var Application = "Belajar golang"

func SayHello(name string) {
	fmt.Println("Hello ", name)
}

func sayGoodBye(name string) {
	fmt.Println("Good bye ", name)
}
