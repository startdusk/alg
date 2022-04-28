package recursion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFib(t *testing.T) {
	cases := []struct {
		name   string
		n      int
		target int
	}{
		{
			name:   "fib",
			n:      9,
			target: 55,
		},
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			assert.Equal(t, Fib(cc.n), cc.target)
			assert.Equal(t, FibWithLoop(cc.n), cc.target)
			assert.Equal(t, FibWithDynamic(cc.n), cc.target)
			assert.Equal(t, FibWithTailRecursice(cc.n), cc.target)
		})
	}
}
