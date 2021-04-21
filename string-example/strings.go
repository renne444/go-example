package main

import (
	"fmt"
	"strings"
)
func interpret(command string) string {

	command = strings.ReplaceAll(command,"(al)", "al")
	command = strings.ReplaceAll(command,"()", "o")
	return command
}
func main(){
	command := "G()(al)"
	fmt.Println(interpret(command))
}
