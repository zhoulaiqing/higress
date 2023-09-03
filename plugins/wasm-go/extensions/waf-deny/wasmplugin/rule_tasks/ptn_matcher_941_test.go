package rule_tasks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test941110Patterns(t *testing.T) {
	val := "xyz=<script>alert(1);</script>"
	assert.True(t, Re941110.MatchString(val))
}

func Test941310Patterns(t *testing.T) {
	val := "/get"
	assert.False(t, Re9413101.MatchString(val))
	assert.False(t, Re9413102.MatchString(val))
}
