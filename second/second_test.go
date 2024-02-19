package second

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSecond(test *testing.T) {
	require.Equal(test, "second module, version 0: first function, release 0", Second())
}
