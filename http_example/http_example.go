package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	jReq, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("j_req = %s", string(jReq))
	res.Write([]byte("hello"))
}

func main() {
	http.HandleFunc("/", hello)

	http.ListenAndServe("0.0.0.0:40404", nil)
}
