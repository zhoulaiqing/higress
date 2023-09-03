package rule_941

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatchValue(t *testing.T) {
	rule941 := Rule941{}
	rk := rule941.matchValue("xyz", false)
	assert.False(t, rk)

	rv := rule941.matchValue("<script>alert('1');<script>", false)
	assert.True(t, rv)
}
