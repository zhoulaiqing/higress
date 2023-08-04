package cc_tools

import "encoding/binary"

func Int64ArrayToByteArray(arr []int64) []byte {
	byteArr := make([]byte, len(arr)*8) // Each int64 takes 8 bytes
	for i, v := range arr {
		binary.BigEndian.PutUint64(byteArr[i*8:], uint64(v))
	}
	return byteArr
}

func ByteArrayToInt64Array(byteArr []byte) []int64 {
	if len(byteArr)%8 != 0 {
		panic("Invalid byte array length for conversion to int64 array.")
	}

	int64Arr := make([]int64, len(byteArr)/8)
	for i := 0; i < len(byteArr); i += 8 {
		int64Arr[i/8] = int64(binary.BigEndian.Uint64(byteArr[i:]))
	}
	return int64Arr
}
