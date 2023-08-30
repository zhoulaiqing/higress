package bodyprocessors

import (
	"encoding/xml"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"io"
	"strings"
)

type xmlBodyProcessor struct {
}

func (*xmlBodyProcessor) ProcessRequest(reader io.Reader, tx *core.Transaction, _ string) error {
	values, contents, err := readXML(reader)
	if err != nil {
		return err
	}
	tx.Variables.XML["//@*"] = values
	tx.Variables.XML["/*"] = contents
	return nil
}

func (*xmlBodyProcessor) ProcessResponse(reader io.Reader, tx *core.Transaction, _ string) error {
	return nil
}

func readXML(reader io.Reader) ([]string, []string, error) {
	var attrs []string
	var content []string
	dec := xml.NewDecoder(reader)
	dec.Strict = false
	dec.AutoClose = xml.HTMLAutoClose
	dec.Entity = xml.HTMLEntity
	for {
		token, err := dec.Token()
		if err != nil && err != io.EOF {
			return nil, nil, err
		}
		if token == nil {
			break
		}
		switch tok := token.(type) {
		case xml.StartElement:
			for _, attr := range tok.Attr {
				attrs = append(attrs, attr.Value)
			}
		case xml.CharData:
			if c := strings.TrimSpace(string(tok)); c != "" {
				content = append(content, c)
			}
		}
	}
	return attrs, content, nil
}

var (
	_ core.BodyProcessor = &xmlBodyProcessor{}
)

func init() {
	RegisterBodyProcessor("xml", func() core.BodyProcessor {
		return &xmlBodyProcessor{}
	})
}
