package config

// A certain piece of config
type ConfigPiece string

const (
	// Name for the summary name related to pulses
	PulseSummaryName ConfigPiece = "summary"
	// Name for the labels name related to pulses
	PulseLabelsName ConfigPiece = "labels"
	// Name for the API key
	ApiKeyName ConfigPiece = "api-key"
)
