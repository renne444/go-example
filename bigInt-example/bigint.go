package main

import "math/big"

import "fmt"

func main() {
	bi := big.NewInt(1)
	bi.Lsh(bi, 2567)
	str := "d580a25969eb2c930bf7529f73cf493a13a46d42dad8ea08881a0a68bfb3df25"
	fmt.Println(len(str))
}

func fuck() {

}
