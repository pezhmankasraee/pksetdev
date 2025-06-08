package help

import (
	"fmt"
	"os"
)

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
	fmt.Println("  -p, --path PATH     path to yaml config file")
	fmt.Println("  -h, --help          Show this help message")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("  ./pksetdev --path ./path/to/config.yaml")
	fmt.Println("  ./pksetdev -p ./path/to/config.yaml")
	fmt.Println("")
	fmt.Println("For more information, visit: https://github.com/pezhmankasraee/pksetdev")
}
