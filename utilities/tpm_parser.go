package utilities

import (
	"bufio"
	"strings"
)

// parseTPMStatus takes the output of the get-tpm command as a string
// and returns a map with the status of each property.
func parseTPMStatus(output string) (map[string]string, error) {
	status := make(map[string]string)
	scanner := bufio.NewScanner(strings.NewReader(output))

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			status[key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return status, nil
}

// checkTPMStatus takes the parsed status map and determines if TPM is enabled and active.
func checkTPMStatus(status map[string]string) bool {
	// Define the keys we are interested in
	keys := []string{"TpmPresent", "TpmReady", "TpmEnabled", "TpmActivated", "TpmOwned"}

	// Check if all the relevant keys have the value "True"
	for _, key := range keys {
		if status[key] != "True" {
			return false
		}
	}
	return true
}
