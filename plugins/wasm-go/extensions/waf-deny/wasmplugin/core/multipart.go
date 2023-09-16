package core

import (
	"errors"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"strings"
)

type multipartBodyProcessor struct{}

func (mbp *multipartBodyProcessor) ProcessRequest(reader io.Reader, tx *Transaction, mimeType string) error {
	mediaType, params, err := mime.ParseMediaType(mimeType)
	if err != nil {
		return err
	}
	if !strings.HasPrefix(mediaType, "multipart/") {
		return errors.New("not a multipart body")
	}
	mr := multipart.NewReader(reader, params["boundary"])
	totalSize := int64(0)

	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		partName := p.FormName()

		if strings.ContainsAny(partName, "\r\n") {
			return ErrorIllegalCRLF
		}

		for key, values := range p.Header {
			for _, value := range values {
				tx.Variables.MultipartPartHeaders[partName] = append(tx.Variables.MultipartPartHeaders[partName],
					fmt.Sprintf("%s: %s", key, value))
			}
		}
		// if is a file
		filename := originFileName(p)
		if filename != "" {
			var size int64
			sz, err := io.Copy(io.Discard, p)
			if err != nil {
				return err
			}
			size = sz

			totalSize += size
			tx.Variables.Files[p.FormName()] = append(tx.Variables.Files[p.FormName()], filename)
		} else {
			// if is a field
			data, err := io.ReadAll(p)
			if err != nil {
				return err
			}
			totalSize += int64(len(data))
			tx.Variables.ArgsPost[p.FormName()] = string(data)
		}
		tx.Variables.FilesCombinedSize = fmt.Sprintf("%d", totalSize)
	}
	return nil
}

func (mbp *multipartBodyProcessor) ProcessResponse(_ io.Reader, _ *Transaction, _ string) error {
	return nil
}

var (
	_ BodyProcessor = (*multipartBodyProcessor)(nil)
)

// OriginFileName returns the filename parameter of the Part's Content-Disposition header.
// This function is based on (multipart.Part).parseContentDisposition,
// See https://go.googlesource.com/go/+/refs/tags/go1.17.9/src/mime/multipart/multipart.go#87
// for the current implementation and also notice this function hasn't change since go1.4, as in
// https://go.googlesource.com/go/+/refs/tags/go1.4/src/mime/multipart/multipart.go#75
func originFileName(p *multipart.Part) string {
	v := p.Header.Get("Content-Disposition")
	_, dispositionParams, err := mime.ParseMediaType(v)
	if err != nil {
		return ""
	}

	return dispositionParams["filename"]
}

func init() {
	RegisterBodyProcessor("multipart", func() BodyProcessor {
		return &multipartBodyProcessor{}
	})
}
