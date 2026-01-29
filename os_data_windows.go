package main

type OSData struct {
	OSDataCommon
	Specific Specific
}

type Specific struct {
	EvtxLogs string
}

func NewSpecific() Specific {
	return Specific{
		EvtxLogs: `C:\winevet\...`,
	}
}
