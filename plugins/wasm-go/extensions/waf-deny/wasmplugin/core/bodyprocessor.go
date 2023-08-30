package core

import (
	"io"
)

type BodyProcessor interface {
	ProcessRequest(reader io.Reader, tx *Transaction, mime string) error
	ProcessResponse(reader io.Reader, tx *Transaction, mime string) error
}
