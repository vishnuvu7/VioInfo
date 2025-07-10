package main

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
)

func printBanner() {
	red := "\033[31m"
	reset := "\033[0m"
	banner := `
 ___      ___ ___  ________          ___  ________   ________ ________     
|\  \    /  /|\  \|\   __  \        |\  \|\   ___  \|\  _____\\   __  \    
\ \  \  /  / | \  \ \  \|\  \       \ \  \ \  \\ \  \ \  \__/\ \  \|\  \   
 \ \  \/  / / \ \  \ \  \\\  \       \ \  \ \  \\ \  \ \   __\\ \  \\\  \  
  \ \    / /   \ \  \ \  \\\  \       \ \  \ \  \\ \  \ \  \_| \ \  \\\  \ 
   \ \__/ /     \ \__\ \_______\       \ \__\ \__\\ \__\ \__\   \ \_______\
    \|__|/       \|__|\|_______|        \|__|\|__| \|__|\|__|    \|_______|
                                                                           
                                                                           
`
	fmt.Print(red + banner + reset)
}

func main() {
	printBanner()
	fmt.Println("=== System Information Tool ===")
	fmt.Println()

	// Display CPU information
	displayCPUInfo()

	// Display memory information
	displayMemoryInfo()

	// Display disk usage
	displayDiskUsage()

	// Display running processes
	displayRunningProcesses()

	// Display network interface information
	displayNetworkInfo()

	// Display system uptime and OS details
	displaySystemInfo()
}

func displayCPUInfo() {
	fmt.Println("=== CPU Information ===")

	// Get CPU info
	cpuInfo, err := cpu.Info()
	if err != nil {
		log.Printf("Error getting CPU info: %v", err)
		return
	}

	// Get CPU usage percentage
	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		log.Printf("Error getting CPU usage: %v", err)
		return
	}

	for i, cpu := range cpuInfo {
		fmt.Printf("CPU %d:\n", i+1)
		fmt.Printf("  Model: %s\n", cpu.ModelName)
		fmt.Printf("  Cores: %d\n", cpu.Cores)
		fmt.Printf("  Usage: %.2f%%\n", cpuPercent[i])
		fmt.Printf("  Architecture: %s\n", runtime.GOARCH)
		fmt.Println()
	}
}

func displayMemoryInfo() {
	fmt.Println("=== Memory Information ===")

	memInfo, err := mem.VirtualMemory()
	if err != nil {
		log.Printf("Error getting memory info: %v", err)
		return
	}

	fmt.Printf("Total Memory: %s\n", formatBytes(memInfo.Total))
	fmt.Printf("Used Memory: %s (%.2f%%)\n", formatBytes(memInfo.Used), memInfo.UsedPercent)
	fmt.Printf("Available Memory: %s\n", formatBytes(memInfo.Available))
	fmt.Printf("Free Memory: %s\n", formatBytes(memInfo.Free))
	fmt.Println()
}

func displayDiskUsage() {
	fmt.Println("=== Disk Usage ===")

	partitions, err := disk.Partitions(false)
	if err != nil {
		log.Printf("Error getting disk partitions: %v", err)
		return
	}

	for _, partition := range partitions {
		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			continue
		}

		fmt.Printf("Mount Point: %s\n", partition.Mountpoint)
		fmt.Printf("  Device: %s\n", partition.Device)
		fmt.Printf("  File System: %s\n", partition.Fstype)
		fmt.Printf("  Total: %s\n", formatBytes(usage.Total))
		fmt.Printf("  Used: %s (%.2f%%) ", formatBytes(usage.Used), usage.UsedPercent)
		printUsageBar(usage.UsedPercent)
		fmt.Printf("  Free: %s\n", formatBytes(usage.Free))
		fmt.Println()
	}
}

func printUsageBar(percent float64) {
	barWidth := 20
	usedBars := int(percent / 100 * float64(barWidth))
	freeBars := barWidth - usedBars
	fmt.Print("[")
	for i := 0; i < usedBars; i++ {
		fmt.Print("#")
	}
	for i := 0; i < freeBars; i++ {
		fmt.Print(".")
	}
	fmt.Printf("] %.2f%%\n", percent)
}

func displayRunningProcesses() {
	fmt.Println("=== Running Processes ===")

	processes, err := process.Processes()
	if err != nil {
		log.Printf("Error getting processes: %v", err)
		return
	}

	fmt.Printf("Total Processes: %d\n", len(processes))
	fmt.Println()

	// Show top 10 processes by CPU usage
	fmt.Println("Top 10 Processes by CPU Usage:")
	fmt.Printf("%-8s %-20s %-10s %-10s %s\n", "PID", "Name", "CPU%", "Memory%", "Status")
	fmt.Println(string(make([]byte, 80, 80)))

	count := 0
	for _, p := range processes {
		if count >= 10 {
			break
		}

		name, err := p.Name()
		if err != nil {
			continue
		}

		cpuPercent, err := p.CPUPercent()
		if err != nil {
			continue
		}

		memPercent, err := p.MemoryPercent()
		if err != nil {
			continue
		}

		status, err := p.Status()
		if err != nil {
			status = []string{"unknown"}
		}

		pid := p.Pid
		statusStr := "unknown"
		if len(status) > 0 {
			statusStr = status[0]
		}
		fmt.Printf("%-8d %-20s %-10.2f %-10.2f %s\n", pid, truncateString(name, 18), cpuPercent, memPercent, statusStr)
		count++
	}
	fmt.Println()
}

func displayNetworkInfo() {
	fmt.Println("=== Network Interface Information ===")

	interfaces, err := net.Interfaces()
	if err != nil {
		log.Printf("Error getting network interfaces: %v", err)
		return
	}

	for _, iface := range interfaces {
		if len(iface.Addrs) == 0 {
			continue
		}

		fmt.Printf("Interface: %s\n", iface.Name)
		fmt.Printf("  Hardware Address: %s\n", iface.HardwareAddr)
		fmt.Printf("  MTU: %d\n", iface.MTU)
		fmt.Printf("  Flags: %s\n", iface.Flags)

		for _, addr := range iface.Addrs {
			fmt.Printf("  Address: %s\n", addr.Addr)
		}
		fmt.Println()
	}

	// Get network I/O statistics
	ioCounters, err := net.IOCounters(false)
	if err != nil {
		log.Printf("Error getting network I/O: %v", err)
		return
	}

	if len(ioCounters) > 0 {
		io := ioCounters[0]
		fmt.Printf("Network I/O Statistics:\n")
		fmt.Printf("  Bytes Sent: %s\n", formatBytes(io.BytesSent))
		fmt.Printf("  Bytes Received: %s\n", formatBytes(io.BytesRecv))
		fmt.Printf("  Packets Sent: %d\n", io.PacketsSent)
		fmt.Printf("  Packets Received: %d\n", io.PacketsRecv)
		fmt.Printf("  Errors In: %d\n", io.Errin)
		fmt.Printf("  Errors Out: %d\n", io.Errout)
		fmt.Printf("  Drops In: %d\n", io.Dropin)
		fmt.Printf("  Drops Out: %d\n", io.Dropout)
		fmt.Println()
	}
}

func displaySystemInfo() {
	fmt.Println("=== System Information ===")

	hostInfo, err := host.Info()
	if err != nil {
		log.Printf("Error getting host info: %v", err)
		return
	}

	fmt.Printf("Hostname: %s\n", hostInfo.Hostname)
	fmt.Printf("OS: %s %s\n", hostInfo.OS, hostInfo.Platform)
	fmt.Printf("Platform Family: %s\n", hostInfo.PlatformFamily)
	fmt.Printf("Platform Version: %s\n", hostInfo.PlatformVersion)
	fmt.Printf("Kernel Version: %s\n", hostInfo.KernelVersion)
	fmt.Printf("Architecture: %s\n", hostInfo.KernelArch)
	fmt.Printf("Uptime: %s\n", formatUptime(hostInfo.Uptime))
	fmt.Printf("Boot Time: %s\n", time.Unix(int64(hostInfo.BootTime), 0).Format("2006-01-02 15:04:05"))
	fmt.Printf("Procs: %d\n", hostInfo.Procs)
	fmt.Printf("Go Version: %s\n", runtime.Version())
	fmt.Printf("Go OS: %s\n", runtime.GOOS)
	fmt.Printf("Go Architecture: %s\n", runtime.GOARCH)
	fmt.Println()
}

func formatBytes(bytes uint64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := uint64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

func formatUptime(uptime uint64) string {
	duration := time.Duration(uptime) * time.Second
	days := int(duration.Hours() / 24)
	hours := int(duration.Hours()) % 24
	minutes := int(duration.Minutes()) % 60

	if days > 0 {
		return fmt.Sprintf("%d days, %d hours, %d minutes", days, hours, minutes)
	} else if hours > 0 {
		return fmt.Sprintf("%d hours, %d minutes", hours, minutes)
	} else {
		return fmt.Sprintf("%d minutes", minutes)
	}
}

func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}
