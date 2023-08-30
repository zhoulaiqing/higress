package bodyprocessors

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core/url_util"
	"io"
	"strings"
)

type urlencodedBodyProcessor struct {
}

func (*urlencodedBodyProcessor) ProcessRequest(reader io.Reader, tx *core.Transaction, _ string) error {
	buf := new(strings.Builder)
	if _, err := io.Copy(buf, reader); err != nil {
		return err
	}

	b := buf.String()
	values := url_util.ParseQuery(b, '&')
	for k, vals := range values {
		for _, v := range vals {
			tx.Variables.ArgsPost[k] = v
		}
	}

	tx.Variables.RequestBody = b

	return nil
}

func (*urlencodedBodyProcessor) ProcessResponse(_ io.Reader, _ *core.Transaction, _ string) error {
	return nil
}

var (
	_ core.BodyProcessor = &urlencodedBodyProcessor{}
)

func init() {
	RegisterBodyProcessor("urlencoded", func() core.BodyProcessor {
		return &urlencodedBodyProcessor{}
	})
}
