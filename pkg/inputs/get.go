package inputs

import (
	"github.com/rootly-io/cli/pkg/inputs/env"
	"github.com/rootly-io/cli/pkg/inputs/flags"
	"github.com/rootly-io/cli/pkg/inputs/parse"
	"github.com/rootly-io/cli/pkg/log"
	"github.com/spf13/cobra"
)

// Return an error if there is no value in the config
func errIfNoVal(name ConfigPiece) log.CtxErr {
	return log.NewErr("Please provide a value for " + string(name))
}

// Get a string configuration value
func GetString(name ConfigPiece, cmd *cobra.Command, required bool) (string, log.CtxErr) {
	// Getting the value from a command line flag if possible
	val, err := flags.GetString(string(name), cmd)
	if err.Error != nil {
		return "", err
	}
	if val != "" {
		return val, err
	}

	// No value from flag so falling back on possible env var
	val = env.GetString(string(name))
	if val == "" && required {
		return "", errIfNoVal(name)
	}

	return val, log.CtxErr{}
}

// Get a sting based array configuration value
func GetArray(
	name ConfigPiece,
	cmd *cobra.Command,
	required bool,
) ([]string, log.CtxErr) {
	// Getting the value from a command line flag if possible
	vals, err := flags.GetArray(string(name), cmd)
	if err.Error != nil {
		return []string{}, err
	}
	if len(vals) != 0 {
		return vals, err
	}

	// No value from flag so falling back on possible env var
	vals = env.GetArray(string(name))
	if len(vals) == 0 && required {
		return []string{}, errIfNoVal(name)
	}

	return vals, log.CtxErr{}
}

// Get a simple key value map based array configuration value
func GetStringSimpleMapArray(
	name ConfigPiece,
	cmd *cobra.Command,
	required bool,
) ([]map[string]string, log.CtxErr) {
	// Getting the value from a command line flag if possible
	str, err := flags.GetString(string(name), cmd)
	if err.Error != nil {
		return []map[string]string{}, err
	}

	vals := parse.Array(str)
	if len(vals) != 0 {
		return convertToSimpleMapArray(vals), err
	}

	// No value from flag so falling back on possible env var
	vals = env.GetArray(string(name))
	if len(vals) == 0 && required {
		return []map[string]string{}, errIfNoVal(name)
	}

	return convertToSimpleMapArray(vals), log.CtxErr{}
}
