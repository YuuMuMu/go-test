package test

import "fmt"

func Testmyfunc(args ...interface{}) {
	for _, arg := range args {
		fmt.Printf("%d", arg)
	}
}
func main() {
	Testmyfunc(1, 1, "1231")
}
