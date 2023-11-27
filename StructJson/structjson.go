package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Time int64  `json:"time"`
}

func main() {
	person := Person{"小明", 18, time.Now().Unix()}
	if result, err := json.Marshal(&person); err == nil {
		fmt.Println(string(result))
	}

	arr := "123 34 78 97 109 101 34 58 34 88 73 65 79 77 73 78 71 34 125"
	arr = strings.Replace(arr, " ", ",", -1)
	fmt.Println(arr)
	StringArr := string(arr)
	fmt.Println("StringArr", StringArr)
	if NumArr, err := strconv.Atoi(StringArr); err == nil {
		fmt.Println("NumArr=", NumArr)
	} else {
		fmt.Println("err=", err)
	}

	a := "123"
	fmt.Println(string(a))
}
