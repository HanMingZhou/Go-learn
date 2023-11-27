package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	s := "1,2,3,4,5"
	fmt.Println("s=", s)
	s = StringsReplace(s)
	fmt.Println("s=", s)
	fmt.Printf("s type = %T\n", s)

	num, err := strconv.Atoi(s)
	if err == nil {
		fmt.Println("strconv s=", num)
		fmt.Printf("s type = %T\n", num)
	} else {
		fmt.Println("Nil=", nil)
	}

	S1 := "1,2,3,4,5"
	fmt.Println("S1=", S1)
	S1 = strings.Replace(S1, ",", "", -1)
	fmt.Println("S1=", S1)
	if S1Num, err := strconv.Atoi(S1); err == nil {
		fmt.Println("s1_num =", S1Num)
	} else {
		fmt.Println(err)
	}

}

func StringsReplace(s string) string {
	s = strings.Replace(s, ",", "", -1)
	return s
}
