package core

import (
	"bytes"
	"io"
)

type BodyBuffer struct {
	buffer  *bytes.Buffer
	length  int64
	readers []*bodyBufferReader
}

var (
	_ io.Writer = (*BodyBuffer)(nil)
)

// Write appends data to the body buffer by chunks
// You may dump io.Readers using io.Copy(br, reader)
func (br *BodyBuffer) Write(data []byte) (n int, err error) {
	if len(data) == 0 {
		return 0, nil
	}

	targetLen := br.length + int64(len(data))
	br.length = targetLen
	return br.buffer.Write(data)
}

type bodyBufferReader struct {
	pos int
	br  *BodyBuffer
}

func (b *bodyBufferReader) Read(p []byte) (n int, err error) {
	if b.br == nil {
		// reader has been closed and hence we don't attempt to do anymore read
		return 0, io.EOF
	}

	buf := b.br.buffer.Bytes()
	n = len(p)
	if b.pos+n > len(buf) {
		n = len(buf) - b.pos
	}
	if n == 0 {
		return 0, io.EOF
	}

	an := copy(p, buf[b.pos:b.pos+n])
	b.pos += an
	return an, nil
}

// Close closes the reader
func (b *bodyBufferReader) Close() {
	b.br = nil
	b.pos = 0
}

func (br *BodyBuffer) Reader() (io.Reader, error) {
	r := &bodyBufferReader{
		br: br,
	}
	br.readers = append(br.readers, r)
	return r, nil
}

// Size returns the current size of the body buffer
func (br *BodyBuffer) Size() int64 {
	return br.length
}

// Reset will reset buffers
func (br *BodyBuffer) Reset() error {
	br.buffer.Reset()
	br.length = 0

	// close all readers, this is important because connectors may have
	// a reference to a reader but the transaction is already closed and
	// hence the reader is not valid anymore.
	for _, r := range br.readers {
		r.Close()
	}
	br.readers = nil
	return nil
}

// NewBodyBuffer Initializes a body reader
// After writing memLimit bytes to the memory buffer, data will be
// written to a temporary file
// Temporary files will be written to tmpDir
func NewBodyBuffer() *BodyBuffer {
	return &BodyBuffer{
		buffer: &bytes.Buffer{},
	}
}
