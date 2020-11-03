package main

import (
	"encoding/json"
	"fmt"
)

type Base struct {
	Number int    `json:"number-base"`
	Str    string `json:"string-base"`
}

type SonExtend struct {
	Base
	Num2 int    `json:"number-implement"`
	Str2 string `json:"string-implement"`
}

func extendsTest() {
	var son SonExtend
	son.Num2 = 2
	son.Number = 1
	son.Str = "str-base"
	son.Str2 = "str-son"

	jsonM := func(e interface{}) string {
		j, _ := json.Marshal(e)
		return string(j)
	}

	fmt.Println(jsonM(son))
}

type SonCombine struct {
	Father interface{} `json:"father"`

	Num2 int    `json:"number-implement"`
	Str2 string `json:"string-implement"`
}

func combineTest() {
	var son SonCombine
	var father Base
	father.Number = 1
	father.Str = "str-base"
	son.Num2 = 2
	son.Str2   = "str-son"
	son.Father = father

	jsonM := func(e interface{}) string {
		j, _ := json.Marshal(e)
		return string(j)
	}

	fmt.Println(jsonM(father))
	fmt.Println(jsonM(son))

}

func unmarshalTest(){
	var unmarshalString string
	unmarshalString = `{"father":{"number-base":1,"string-base":"str-base"},"number-implement":2,"string-implement":"str-son"}`
	fmt.Println(unmarshalString)

	var sonDeserialize SonCombine
	err := json.Unmarshal([]byte(unmarshalString), &sonDeserialize)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sonDeserialize)

}

func main() {
	extendsTest()
	combineTest()
	unmarshalTest()
}
