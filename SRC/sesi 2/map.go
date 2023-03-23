package main

import (
	"fmt"
	"strconv"
)

func main() {
	var nameObject = []map[string]string{
		{"name": "anwar Sahid", "job": strconv.Itoa(1)},
		{"name": "anwar Sahid2", "job": "on go2"},
		{"name": "anwar Sahid3", "job": "on go3"},
	}
	for _, nameObject := range nameObject {
		fmt.Println(nameObject["name"], nameObject["job"])
	}

	// var name = map[string]string{
	// 	"name": "cek",
	// 	"s":    "cek",
	// 	"c":    "cek",
	// }

	// fmt.Println(name["name"])

}
