package secondary

import "fmt"

func init() {
	fmt.Println("inited")
}

func Invoke() int {
	fmt.Println("invoked")
	return 10
}
