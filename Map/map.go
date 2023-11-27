package main

import (
	"fmt"
)

func main() {

	siteMap := make(map[string]string)

	siteMap["wiki"] = "维基"
	siteMap["google"] = "谷歌"
	siteMap["baidu"] = "百度"
	siteMap["runoob"] = "菜鸟教程"

	for k, v := range siteMap {
		fmt.Println("siteMap=", k, v)
	}

	v, ok := siteMap["facebook"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("not in siteMap")
	}

	delete(siteMap, "baidu")
	for k, v := range siteMap {
		fmt.Println("siteMap=", k, v)
	}

}
