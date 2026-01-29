package main

type OSData struct {
	OSDataCommon
	Specific Specific
}

type Specific struct {
	SystemLog string
}

func NewSpecific() Specific {
	return Specific{
		SystemLog: "/var/log",
	}
}
