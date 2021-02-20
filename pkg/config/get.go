package config

import (
	"github.com/rootly-io/cli/pkg/config/env"
	"github.com/rootly-io/cli/pkg/config/flags"
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
		return val, errIfNoVal(name)
	}

	return val, log.CtxErr{}
}

// Get a string array configuration value
func GetStringArray(name ConfigPiece, cmd *cobra.Command, required bool) ([]string, log.CtxErr) {
	// Getting the value from a command line flag if possible
	val, err := flags.GetStringArray(string(name), cmd)
	if err.Error != nil {
		return []string{}, err
	}
	if len(val) == 0 {
		return []string{}, err
	}

	// No value from flag so falling back on possible env var
	val, err = env.GetStringArray(string(name))
	if err.Error != nil {
		return []string{}, err
	}
	if len(val) == 0 && required {
		return []string{}, errIfNoVal(name)
	}

	return []string{}, log.CtxErr{}
}
