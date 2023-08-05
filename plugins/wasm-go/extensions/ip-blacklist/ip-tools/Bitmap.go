package ip_tools

// Bitmap represents a 2^32 bitmap.
type Bitmap struct {
	bits []byte
}

// NewBitmap creates a new 2^32 bitmap.
func NewBitmap() *Bitmap {
	// 2^32 bits can be represented by 2^29 bytes (2^29 * 8 = 2^32).
	bits := make([]byte, 1<<29)
	return &Bitmap{bits: bits}
}

// SetBit sets the bit at the given index to 1.
func (bm *Bitmap) SetBit(index uint32) {
	byteIndex := index / 8
	bitOffset := index % 8
	bm.bits[byteIndex] |= 1 << bitOffset
}

// SetRange sets all bits in the range [start, end] to 1 (closed interval).
func (bm *Bitmap) SetRange(start, end uint32) {
	if start > end {
		start, end = end, start
	}
	for i := start; i <= end; i++ {
		bm.SetBit(i)
	}
}

// ClearBit sets the bit at the given index to 0.
func (bm *Bitmap) ClearBit(index uint32) {
	byteIndex := index / 8
	bitOffset := index % 8
	bm.bits[byteIndex] &= ^(1 << bitOffset)
}

// GetBit checks if the bit at the given index is set to 1.
func (bm *Bitmap) GetBit(index uint32) bool {
	byteIndex := index / 8
	bitOffset := index % 8
	return (bm.bits[byteIndex]>>bitOffset)&1 == 1
}
