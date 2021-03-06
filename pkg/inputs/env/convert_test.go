package env

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertName(t *testing.T) {
	tt := []struct {
		flag     string
		expected string
	}{
		{
			flag:     "help",
			expected: "ROOTLY_HELP",
		},
		{
			flag:     "HELP",
			expected: "ROOTLY_HELP",
		},
		{
			flag:     "old-help",
			expected: "ROOTLY_OLD_HELP",
		},
		{
			flag:     "OLD-HELP",
			expected: "ROOTLY_OLD_HELP",
		},
	}

	for _, test := range tt {
		envActionName := "ROOTLY_GH_ACTION"
		assert.Equal(t, test.expected, convertName(test.flag))

		err := os.Setenv(envActionName, "true")
		assert.NoError(t, err)

		assert.Equal(
			t,
			strings.Replace(test.expected, "ROOTLY", "INPUT", 1),
			convertName(test.flag),
		)

		err = os.Setenv(envActionName, "")
		assert.NoError(t, err)
	}
}
