 ___      ___ ___  ________          ___  ________   ________ ________     
|\  \    /  /|\  \|\   __  \        |\  \|\   ___  \|\  _____\\   __  \    
\ \  \  /  / | \  \ \  \|\  \       \ \  \ \  \\ \  \ \  \__/\ \  \|\  \   
 \ \  \/  / / \ \  \ \  \\\  \       \ \  \ \  \\ \  \ \   __\\ \  \\\  \  
  \ \    / /   \ \  \ \  \\\  \       \ \  \ \  \\ \  \ \  \_| \ \  \\\  \ 
   \ \__/ /     \ \__\ \_______\       \ \__\ \__\\ \__\ \__\   \ \_______\
    \|__|/       \|__|\|_______|        \|__|\|__| \|__|\|__|    \|_______|
                                                                           
                                                                           
# Vio Info

A stylish and comprehensive terminal-based system information tool written in Go.

## Features
- **CPU Information**: Model, cores, usage, and architecture
- **Memory Information**: Total, used, available, and free memory
- **Disk Usage**: Partition details with graphical usage bars
- **Running Processes**: Top 10 processes by CPU usage
- **Network Interfaces**: Interface details and network I/O stats
- **System Details**: Hostname, OS, uptime, boot time, and more
- **ASCII Art Banner**: Eye-catching "Vio Info" banner at startup

## Requirements
- Go 1.18 or higher
- Works on macOS, Linux, and Windows

## Installation
1. Clone this repository or download the source code.
2. Open a terminal and navigate to the project directory.
3. Download dependencies:
   ```sh
   go mod tidy
   ```

## Build
To create an executable named `vioinfo`:
```sh
go build -o vioinfo
```

## Run
You can run the tool directly:
```sh
./vioinfo
```
Or, run without building:
```sh
go run main.go
```

## Example Output
```
 ___      ___ ___  ________          ___  ________   ________ ________     
|\  \    /  /|\  \|\   __  \        |\  \|\   ___  \|\  _____\\   __  \    
\ \  \  /  / | \  \ \  \|\  \       \ \  \ \  \\ \  \ \  \__/\ \  \|\  \   
 \ \  \/  / / \ \  \ \  \\\  \       \ \  \ \  \\ \  \ \   __\\ \  \\\  \  
  \ \    / /   \ \  \ \  \\\  \       \ \  \ \  \\ \  \ \  \_| \ \  \\\  \ 
   \ \__/ /     \ \__\ \_______\       \ \__\ \__\\ \__\ \__\   \ \_______\
    \|__|/       \|__|\|_______|        \|__|\|__| \|__|\|__|    \|_______|

=== System Information Tool ===

=== CPU Information ===
CPU 1:
  Model: Apple M4
  Cores: 10
  Usage: 9.49%
  Architecture: arm64

=== Memory Information ===
Total Memory: 24.0 GB
Used Memory: 13.5 GB (56.24%)
Available Memory: 10.5 GB
Free Memory: 140.8 MB

=== Disk Usage ===
Mount Point: /
  Device: /dev/disk3s1s1
  File System: apfs
  Total: 926.4 GB
  Used: 274.4 GB (29.62%) [######..............] 29.62%
  Free: 651.9 GB
...

## License
MIT 