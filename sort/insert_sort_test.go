package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertSort(t *testing.T) {
	cases := []struct {
		name   string
		source []int
		target []int
	}{
		{
			source: []int{4, 6, 1, 8, 11, 13, 3},
			target: []int{1, 3, 4, 6, 8, 11, 13},
		},
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			InsertSort(cc.source)
			assert.Equal(t, cc.source, cc.target)
		})
	}
}
