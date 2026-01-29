package main

type OSData struct {
	OSDataCommon
	SpecificWindows SpecificWindows
}

type SpecificWindows struct {
	EvtxLogs string
}

func NewOSData() OSData {
	return OSData{
		OSDataCommon:    NewOSDataCommon(),
		SpecificWindows: NewSpecific(),
	}
}

func NewSpecific() SpecificWindows {
	return SpecificWindows{
		EvtxLogs: `C:\winevet\...`,
	}
}
