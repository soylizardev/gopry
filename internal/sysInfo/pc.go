package sysinfo

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func (i *InfoPc) GetPc() {
	vendorRaw, errv := os.ReadFile("/sys/class/dmi/id/sys_vendor")
	nameRaw, errn := os.ReadFile("/sys/class/dmi/id/product_version")
	modelRaw, errm := os.ReadFile("/sys/class/dmi/id/product_name")

	var vendor, name, model string

	if errv != nil {
		vendor = "Unknown"
	} else {
		vendor = strings.TrimSpace(string(vendorRaw))
	}
	if errn != nil {
		name = "Unknown"
	} else {
		name = strings.TrimSpace(string(nameRaw))
	}
	if errm != nil {
		model = "Unknown"
	} else {
		model = strings.TrimSpace(string(modelRaw))
	}

	i.Pc = fmt.Sprintf("%s %s (%s)", vendor, name, model)
}

func (i *InfoPc) GetCPU() {
	cmd, err := os.ReadFile("/proc/cpuinfo")
	if err != nil {
		i.CPU = "Unknown"
		return
	}
	split := strings.Split(string(cmd), "\n")
	var cpuModel string
	var cpuCores string
	cpuThreads := 0

	for _, line := range split {
		if strings.Contains(line, "model name") {
			div := strings.Split(line, ":")
			cpuModel = strings.TrimSpace(div[1])
			cpuThreads++
		}
		if strings.Contains(line, "cpu cores") {
			div := strings.Split(line, ":")
			cpuCores = strings.TrimSpace(div[1])
		}
	}
	i.CPU = fmt.Sprintf("%s (%s Cores / %d Threads)", cpuModel, cpuCores, cpuThreads)
}

func (i *InfoPc) GetArch() {
	output, err := os.ReadFile("/proc/sys/kernel/arch")
	if err != nil {
		i.Arch = "Unknown"
	}
	i.Arch = strings.TrimSpace(string(output))
}

func (i *InfoPc) GetGraphics() {
	cmd := exec.Command("sh", "-c", "lspci | grep -i vga")
	out, err := cmd.Output()
	if err != nil {
		i.Graphic = "Unknown"
		return
	}
	split := strings.Split(string(out), ":")
	i.Graphic = strings.TrimSpace(split[2])
}

func (i *InfoPc) GetDisk() {
	cmd := exec.Command("df", "-h", "/")
	com, err := cmd.Output()
	if err != nil {
		i.Disk = "Unknown"
		return
	}
	lines := strings.Split(string(com), "\n")
	split := strings.Fields(lines[1])
	diskTotal := split[1]
	diskUsed := split[2]
	diskporc := split[4]

	i.Disk = fmt.Sprintf("%s / %s (used: %s)", diskUsed, diskTotal, diskporc)
}

func (i *InfoPc) GetRam() {
	content, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		i.Ram = "Unknown"
		return
	}
	lines := strings.Split(string(content), "\n")
	var memTotalRaw, memAvailRaw string
	for _, line := range lines {
		if strings.Contains(line, "MemTotal:") {
			fields := strings.Fields(line)
			memTotalRaw = fields[1]
		}
		if strings.Contains(line, "MemAvailable:") {
			fields := strings.Fields(line)
			memAvailRaw = fields[1]
		}
	}
	totalMemory, errTotal := strconv.ParseFloat(memTotalRaw, 64)
	if errTotal != nil {
		totalMemory = 0
	}
	const kbToGb = 1048576
	totalGB := totalMemory / kbToGb
	availableMemory, errAva := strconv.ParseFloat(memAvailRaw, 64)
	if errAva != nil {
		availableMemory = 0
	}

	availableGB := availableMemory / kbToGb
	memUsed := totalGB - availableGB

	i.Ram = fmt.Sprintf("%.2f GiB / %.2f GiB", memUsed, totalGB)
}

func (i *InfoPc) GetSwap() {
	content, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		i.Swap = "Unknown"
		return
	}
	lines := strings.Split(string(content), "\n")
	var swapTotalRaw, swapFreeRaw string
	for _, line := range lines {
		if strings.Contains(line, "SwapTotal:") {
			fields := strings.Fields(line)
			swapTotalRaw = fields[1]
		}
		if strings.Contains(line, "SwapFree:") {
			fields := strings.Fields(line)
			swapFreeRaw = fields[1]
		}
	}

	totalSwap, errTotal := strconv.ParseFloat(swapTotalRaw, 64)
	const kbToGb = 1048576.0
	if errTotal != nil {
		totalSwap = 0
	}
	totalGB := totalSwap / kbToGb
	freeSwap, errFree := strconv.ParseFloat(swapFreeRaw, 64)
	if errFree != nil {
		freeSwap = 0
	}
	freeGB := freeSwap / kbToGb
	swapUsed := totalGB - freeGB
	i.Swap = fmt.Sprintf("%.2f GiB / %.2f GiB", swapUsed, totalGB)
}
