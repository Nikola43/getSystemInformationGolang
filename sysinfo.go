package main

// SysInfo saves the basic system information
type SysInfo struct {
    Hostname string `bson:hostname`
    Platform string `bson:platform`
    CPU      string `bson:cpu`
    RAM      uint64 `bson:ram`
    Disk     uint64 `bson:disk`
}
