package utils

import "strings"


// Parses fields to be printed
func ParseFields(rawFields string, printableFields []*string) []*string {
	for _, field := range strings.Split(rawFields, ",") {
		// Trim whitespace and check if the field is not empty
		field = strings.TrimSpace(field)
		if field != "" {
			printableFields = append(printableFields, &field)
		}
	}
	return printableFields
}
