package sysinfo

type InfoOs struct {
	Host     string
	OS       string
	Kernel   string
	User     string
	Terminal string
	Shell    string
	Uptime   string
}

type InfoPc struct {
	Pc      string
	Arch    string
	CPU     string
	Graphic string
	Disk    string
	Ram     string
	Swap    string
}

type SystemReport struct {
	OS InfoOs
	PC InfoPc
}
