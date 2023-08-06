package sliding_hist_limiter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHistory(t *testing.T) {

	var hist = History([]int64{2, 5, 6, 8})

	hist = insert(hist, 7)
	hist = insert(hist, 10)

	assert.Equal(t, 6, len(hist))
	assert.Equal(t, int64(7), hist[3])
	assert.Equal(t, int64(10), hist[5])

	hist = removeBefore(hist, 5)
	assert.Equal(t, 5, len(hist))
	assert.Equal(t, 2, countAfter(hist, 8))

}
