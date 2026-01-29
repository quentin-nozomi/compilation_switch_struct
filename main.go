package main

import "fmt"

func main() {
	osData := NewOSData()
	fmt.Println(osData.Specific.EvtxLogs)
}
