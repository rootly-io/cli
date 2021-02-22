package inputs

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// Add a flag for the API key
func AddKeyFlag(cmd *cobra.Command) {
	cmd.Flags().String(
		string(ApiKeyName),
		"",
		"api key generated from rootly.io. See https://rootly.io/api#section/How-to-generate-an-API-Key for more info.",
	)
}

// Add a flag for a pulse summary
func AddPulseSummaryFlag(cmd *cobra.Command) {
	cmd.Flags().StringP(
		string(PulseSummaryName),
		"p",
		"",
		"Summary of the pulse",
	)
}

// Consistent message telling the user how to use array string flags
func arrayUsage(name string) string {
	return fmt.Sprintf(
		"%v associated with the pulse. Give a new flag for each %v and separate the key from the value with an equals sign.",
		strings.Title(name),
		name,
	)
}

// Add a flag for pulse labels
func AddPulseLabelsFlag(cmd *cobra.Command) {
	name := string(PulseLabelsName)
	cmd.Flags().StringP(
		name,
		"l",
		"",
		arrayUsage(name),
	)
}

// Add a flag for pulse services
func AddPulseServicesFlag(cmd *cobra.Command) {
	name := string(PulseServicesName)
	cmd.Flags().StringP(
		name,
		"s",
		"",
		arrayUsage(name),
	)
}

// Add a flag for pulse services
func AddPulseEnvironmentsFlag(cmd *cobra.Command) {
	name := string(PulseEnvironmentsName)
	cmd.Flags().StringP(
		name,
		"e",
		"",
		arrayUsage(name),
	)
}
