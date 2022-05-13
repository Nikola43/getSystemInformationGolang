package main

import (
    "fmt"
    "github.com/nikola43/getsysteminformationgolang/sysinfo"

)

func main() {
    info := newSysInfo()
    fmt.Printf("%+v\n", info)
	fmt.Printf("%+s\n", info.toString())
	fmt.Printf("%+s\n", info.toHash())
}