package main

import (
	"fmt"
	"os"

	"github.com/rezmoss/go-cloudip"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: cloudip <IP>")
		os.Exit(1)
	}

	fmt.Println("Version:", cloudip.Version())
	fmt.Println("Providers:", cloudip.Providers())
	fmt.Println("Range count:", cloudip.RangeCount())

	ip := os.Args[1]

	result := cloudip.Lookup(ip)

	fmt.Println("Found:", result.Found)

	if !result.Found {
		fmt.Println("No match")
		return
	}

	fmt.Printf("Provider : %s\n", result.Provider)
	fmt.Printf("Region   : %s\n", result.Region)
	fmt.Printf("Service  : %s\n", result.Service)
	fmt.Printf("CIDR     : %s\n", result.CIDR)
}