package cc_tools

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt64Arr2ByteArr2Int64Arr(test *testing.T) {
	int64Arr := []int64{123456789, 987654321, 9876543210}
	byteArr := Int64ArrayToByteArray(int64Arr)
	reversedInt64Arr := ByteArrayToInt64Array(byteArr)

	assert.Equal(test, 3, len(reversedInt64Arr))
	assert.Equal(test, int64(123456789), reversedInt64Arr[0])
	assert.Equal(test, int64(987654321), reversedInt64Arr[1])
	assert.Equal(test, int64(9876543210), reversedInt64Arr[2])
}
