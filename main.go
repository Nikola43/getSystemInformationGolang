package main

import (
    "fmt"
)

func main() {
    info := newSysInfo()
    fmt.Printf("%+v\n", info)
	fmt.Printf("%+s\n", info.toString())
	fmt.Printf("%+s\n", info.toHash())
}