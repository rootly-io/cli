package flags

import (
	"fmt"
	"strings"
)

// Consistent message telling the user how to use array string flags
func arrayUsage(name string) string {
	return fmt.Sprintf(
		"%v associated with the pulse. Separate each item with a comma.",
		strings.Title(name),
	)
}

// Consistent message telling the user how to use map flags
const mapUsage = "Key-value pair separated with an equal sign (= with no surrounding spaces)"
