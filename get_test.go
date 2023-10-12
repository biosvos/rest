package rest

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	get, err := Get("https://www.google.com")
	require.NoError(t, err)
	require.NotEmpty(t, get)
}
