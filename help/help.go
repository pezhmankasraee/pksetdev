package help

import (
	"fmt"
	"os"
)

var Version string

func ShowHelp(isHelp bool) {
	if isHelp {
		generalHelp()
		os.Exit(0)
	}
}

func generalHelp() {
	fmt.Println("Usage: pksetdev [OPTIONS]")
	fmt.Println("")
	fmt.Println("Options:")
	fmt.Println("  -i, --install       Download and install the applications")
	fmt.Println("  -p, --path PATH     Path to yaml config file")
	fmt.Println("  -h, --help          Show this help message")
	fmt.Println("  -v, --Version       Show the version of the application")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("  ./pksetdev --path ./path/to/config.yaml")
	fmt.Println("  ./pksetdev -p ./path/to/config.yaml")
	fmt.Println("")
	fmt.Println("For more information, visit: https://github.com/pezhmankasraee/pksetdev")
}

func ShowVersion(isVersion bool) {

	if isVersion {
		fmt.Println("pksetdev " + Version)
		fmt.Println("Written by Pezhman Kasraee <github@pezhmankasraee.com>")
		os.Exit(0)
	}
}
