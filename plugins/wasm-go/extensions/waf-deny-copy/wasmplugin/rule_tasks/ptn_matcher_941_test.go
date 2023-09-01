package rule_tasks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test941310Patterns(t *testing.T) {
	val := "/get"
	assert.False(t, Re9413101.MatchString(val))
	assert.False(t, Re9413102.MatchString(val))
}
