package first

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFirst(test *testing.T) {
	require.Equal(test, "first function, release 0", First())
}
