package rule_941

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestXss(t *testing.T) {
	assert.True(t, matchXSS([]byte("< x on=")))
}
