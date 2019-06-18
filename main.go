package main

import (
	"./tool"
	"fmt"
)

func main() {
	rs := tool.Reveser("abcedfg")
	fmt.Println(rs)
}
