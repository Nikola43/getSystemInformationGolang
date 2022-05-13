package sysinfo

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

// SysInfo saves the basic system information
type SysInfo struct {
    Hostname string `bson:hostname`
    Platform string `bson:platform`
    CPU      string `bson:cpu`
    RAM      uint64 `bson:ram`
    Disk     uint64 `bson:disk`
}

func newSysInfo() *SysInfo{
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
    return info
}

func (sysInfo *SysInfo) toString() string {
    return sysInfo.Hostname + ","+ sysInfo.Platform  + ","+ sysInfo.CPU + ","+ strconv.FormatUint(sysInfo.RAM, 10) +","+ strconv.FormatUint(sysInfo.Disk, 10)
}

func (sysInfo *SysInfo) toHash() string {
    Hostname := hashValue(sysInfo.Hostname)
    Platform :=  hashValue(sysInfo.Platform)
    CPU :=  hashValue(sysInfo.CPU)
    RAM :=  hashValue(strconv.FormatUint(sysInfo.RAM, 10))
    Disk :=  hashValue(strconv.FormatUint(sysInfo.Disk, 10))
    return hashValue(Hostname+Platform+CPU+RAM+Disk)
}

func hashValue(value string) string {
	hash := sha256.New()
	hash.Write([]byte(value))
    return hex.EncodeToString(hash.Sum(nil))
}

func decodeHex(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	
	return b
}