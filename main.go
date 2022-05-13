package main

import (
    "fmt"

    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/disk"
    "github.com/shirou/gopsutil/host"
    "github.com/shirou/gopsutil/mem"
)

func main() {
    hostStat, _ := host.Info()
    cpuStat, _ := cpu.Info()
    vmStat, _ := mem.VirtualMemory()
    diskStat, _ := disk.Usage("\\") // If you're in Unix change this "\\" for "/"

    info := new(SysInfo)

    info.Hostname = hostStat.Hostname
    info.Platform = hostStat.Platform
    info.CPU = cpuStat[0].ModelName
    info.RAM = vmStat.Total / 1024 / 1024
    info.Disk = diskStat.Total / 1024 / 1024

    fmt.Printf("%+v\n", info)

}