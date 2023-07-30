package cc_tools

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZSetToString(t *testing.T) {
	data := &ZSet{}
	data.ZAdd("abc", 100)
	data.ZAdd("abc", 120)
	data.ZAdd("abc", 150)

	fmt.Println(data.ToString("abc"))

	assert.Equal(t, 2, data.ZCount("abc", 110, 150))

	data.ZRemByScore("abc", 0, 130)
	assert.Equal(t, 1, data.ZCount("abc", 110, 150))
}
