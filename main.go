package main

import "fmt"

func main() {
	osData := NewOSData()
	fmt.Println(osData.SpecificWindows.EvtxLogs)
}
