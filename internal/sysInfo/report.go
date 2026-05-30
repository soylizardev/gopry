package sysinfo

import "fmt"

func (sr *SystemReport) Render() {
	//Software
	sr.OS.GetUser()
	sr.OS.GetHost()
	sr.OS.GetOS()
	sr.OS.GetKernel()
	sr.OS.GetShell()
	sr.OS.GetTerm()
	sr.OS.GetTime()

	//Hardware
	sr.PC.GetPc()
	sr.PC.GetCPU()
	sr.PC.GetArch()
	sr.PC.GetGraphics()
	sr.PC.GetDisk()
	sr.PC.GetRam()
	sr.PC.GetSwap()

	//Format
	fmt.Println()
	fmt.Printf("   \033[1;35m%s\033[0m@\033[1;35m%s\033[0m\n", sr.OS.User, sr.OS.Host)
	fmt.Println("   ------------------OS--------------------")

	// Bloque de Software (OS)
	fmt.Printf("   \033[1;36mOS:\033[0m       %s\n", sr.OS.OS)
	fmt.Printf("   \033[1;36mKernel:\033[0m   %s\n", sr.OS.Kernel)
	fmt.Printf("   \033[1;36mUptime:\033[0m   %s\n", sr.OS.Uptime)
	fmt.Printf("   \033[1;36mShell:\033[0m    %s\n", sr.OS.Shell)
	fmt.Printf("   \033[1;36mTerminal:\033[0m %s\n", sr.OS.Terminal)

	fmt.Println("   ------------------PC--------------------")

	// Bloque de Hardware (PC)
	fmt.Printf("   \033[1;33mHost:\033[0m     %s\n", sr.PC.Pc)
	fmt.Printf("   \033[1;33mCPU:\033[0m      %s\n", sr.PC.CPU)
	fmt.Printf("   \033[1;33mArch:\033[0m     %s\n", sr.PC.Arch)
	fmt.Printf("   \033[1;33mGraphics:\033[0m %s\n", sr.PC.Graphic)
	fmt.Printf("   \033[1;33mDisk:\033[0m     %s\n", sr.PC.Disk)
	fmt.Printf("   \033[1;33mMemory:\033[0m   %s\n", sr.PC.Ram)
	fmt.Printf("   \033[1;33mSwap:\033[0m     %s\n", sr.PC.Swap)

	fmt.Println("   ----------------------------------------")
	fmt.Println()
}
