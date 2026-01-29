package main

type OSData struct {
	OSDataCommon
	SpecificLinux SpecificLinux
}

type SpecificLinux struct {
	SystemLog string
}

func NewOSData() OSData {
	return OSData{
		OSDataCommon:  NewOSDataCommon(),
		SpecificLinux: NewSpecific(),
	}
}

func NewSpecific() SpecificLinux {
	return SpecificLinux{
		SystemLog: "/var/log",
	}
}
