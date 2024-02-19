package third

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestThird(test *testing.T) {
	require.Equal(test, "third version 0: second module, version 0: first function, release 0", Third())
}
