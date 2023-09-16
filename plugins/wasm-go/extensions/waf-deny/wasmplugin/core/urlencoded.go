package core

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core/url_util"
	"io"
	"strings"
)

type urlencodedBodyProcessor struct {
}

func (*urlencodedBodyProcessor) ProcessRequest(reader io.Reader, tx *Transaction, _ string) error {
	buf := new(strings.Builder)
	if _, err := io.Copy(buf, reader); err != nil {
		return err
	}

	b := buf.String()
	values := url_util.ParseQuery(b, '&')
	for k, vals := range values {

		if strings.ContainsAny(k, "\r\n") {
			return ErrorIllegalCRLF
		}

		for _, v := range vals {
			tx.Variables.ArgsPost[k] = v
		}
	}

	tx.Variables.RequestBody = b

	return nil
}

func (*urlencodedBodyProcessor) ProcessResponse(_ io.Reader, _ *Transaction, _ string) error {
	return nil
}

var (
	_ BodyProcessor = &urlencodedBodyProcessor{}
)

func init() {
	RegisterBodyProcessor("urlencoded", func() BodyProcessor {
		return &urlencodedBodyProcessor{}
	})
}
