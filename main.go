package main

import (
	"fmt"
)

func main() {
    info := SysInfo()
    fmt.Printf("%+v\n", info)
	fmt.Printf("%+s\n", info.toString())
	fmt.Printf("%+s\n", info.toHash())
}