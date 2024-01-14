package utilities

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func CheckTPMStatus() bool {
	tpmStatus, err := checkStatus("Get-Tpm", "TPM")
	if err != nil {
		fmt.Println("Error checking TPM status:", err)
		return true
	}
	fmt.Println("TPM Status:", tpmStatus)
	return false
}

func CheckSecureBoot() bool {
	secureBootStatus, err := checkStatus("Confirm-SecureBootUEFI", "Secure Boot")
	if err != nil {
		fmt.Println("Error checking Secure Boot status:", err)
		return true
	}
	fmt.Println("Secure Boot Status:", secureBootStatus)
	return false
}

// / checkStatus takes a powershell command and a name and returns the status of the command as a string.
// / The name is used to determine how to parse the output of the command.
func checkStatus(command, name string) (string, error) {
	cmd := exec.Command("powershell", "-Command", command)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("%v: %v", err, stderr.String())
	}

	var returnString string

	output := out.String()
	// if name is TPM
	if name == "TPM" {
		// Parse the TPM status from the command output
		status, err := parseTPMStatus(output)
		if err != nil {
			fmt.Println("Error parsing TPM status:", err)
			return "", err
		}

		// Check if TPM is enabled and active
		if checkTPMStatus(status) {
			returnString = "TPM is enabled and active."
		} else {
			returnString = "TPM is not fully enabled and active."
		}
	} else if name == "Secure Boot" {
		// output should be "True" or "False"
		if strings.Contains(output, "True") {
			returnString = "Secure Boot is enabled."
		} else {
			returnString = "Secure Boot is disabled."
		}
	}

	// check if returnString is empty
	if returnString == "" {
		return "", fmt.Errorf("Error checking %v status: %v", name, err)
	}

	return returnString, nil
}
