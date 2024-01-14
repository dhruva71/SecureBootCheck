package main

import (
	"fmt"
	"securebootcheck/utilities"
)

func main() {
	// Check if running as admin
	utilities.EnsureAdminAccess()

	// Check TPM status
	utilities.CheckTPMStatus()

	// Check Secure Boot status
	utilities.CheckSecureBoot()

	// wait for user input to exit
	var input string
	println("\n\nPress enter to exit")
	fmt.Scanln(&input)
}
