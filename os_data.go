package main

type OSDataCommon struct {
	Common string
}

func NewOSDataCommon() OSDataCommon {
	return OSDataCommon{}
}

func NewOSData() OSData {
	return OSData{
		OSDataCommon: NewOSDataCommon(),
		Specific:     NewSpecific(),
	}
}
