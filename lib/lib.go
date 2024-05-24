package lib

import "strings"

func IsCSVFile(filename string) bool {
	// Convert filename to lowercase to handle case-insensitive matching
	filename = strings.ToLower(filename)
	// Check if the filename ends with ".csv"
	return strings.HasSuffix(filename, ".csv")
}
