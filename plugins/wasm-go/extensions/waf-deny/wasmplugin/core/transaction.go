package core

import (
	"fmt"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core/url_util"
	"net/url"
	"strings"
)

type Transaction struct {
	id        string
	Variables TransactionVariables
}

func NewTransaction() Transaction {
	id := RandomString(16)
	variables := TransactionVariables{}
	variables.Args = append(variables.Args, &variables.ArgsGet, &variables.ArgsPost, &variables.ArgsPath)

	return Transaction{id: id, Variables: variables}
}

type TransactionVariables struct {
	RequestMethod        string
	RequestProtocol      string
	RequestUriRaw        string
	RequestUri           string
	RequestBaseName      string
	RequestFileName      string
	RequestLine          string
	RequestHeaders       map[string]string
	RequestCookies       map[string]string
	QueryString          string
	RequestBody          string
	ReqBodyProcessor     string
	ArgsGet              map[string]string
	ArgsPost             map[string]string
	ArgsPath             map[string]string
	Args                 []*map[string]string
	ArgsNames            []*[]string
	Files                map[string][]string
	FilesNames           map[string][]string
	XML                  map[string][]string
	MultipartPartHeaders map[string][]string

	BlockingInboundAnomalyScore  int
	InboundAnomalyScorePl1       int
	InboundAnomalyScorePl2       int
	InboundAnomalyScorePl3       int
	InboundAnomalyScorePl4       int
	SqlInjectionScore            int
	SssScore                     int
	RfiScore                     int
	LfiScore                     int
	RceScore                     int
	PhpInjectionScore            int
	HttpViolationScore           int
	SessionFixationScore         int
	BlockingOutboundAnomalyScore int
	OutboundAnomalyScorePl1      int
	OutboundAnomalyScorePl2      int
	OutboundAnomalyScorePl3      int
	OutboundAnomalyScorePl4      int
	AnomalyScore                 int
}

const ARG_LIMIT = 1000

func (tx *Transaction) AddRequestHeader(key string, value string) {
	tx.Variables.RequestHeaders[key] = value
}

func (tx *Transaction) ExtractGetArguments(uri string) {
	data := url_util.ParseQuery(uri, '&')
	for k, vs := range data {
		for _, v := range vs {
			tx.AddGetRequestArgument(k, v)
		}
	}
}

func (tx *Transaction) AddGetRequestArgument(key string, value string) {
	if !tx.checkArgumentLimit() {
		return
	}
	tx.Variables.ArgsGet[key] = value
}

func (tx *Transaction) AddPostRequestArgument(key string, value string) {
	if !tx.checkArgumentLimit() {
		return
	}
	tx.Variables.ArgsPost[key] = value
}

func (tx *Transaction) AddPathRequestArgument(key string, value string) {
	if !tx.checkArgumentLimit() {
		return
	}
	tx.Variables.ArgsPath[key] = value
}

func (tx *Transaction) checkArgumentLimit() bool {
	argLen := len(tx.Variables.ArgsGet) + len(tx.Variables.ArgsPost) + len(tx.Variables.ArgsPath)
	return argLen <= ARG_LIMIT
}

func (tx *Transaction) ProcessURI(uri string, method string, httpVersion string) {
	tx.Variables.RequestMethod = method
	tx.Variables.RequestProtocol = httpVersion
	tx.Variables.RequestUriRaw = uri
	tx.Variables.RequestLine = fmt.Sprintf("%s %s %s", method, uri, httpVersion)

	var err error
	// we remove anchors
	if in := strings.Index(uri, "#"); in != -1 {
		uri = uri[:in]
	}
	path := ""
	parsedURL, err := url.Parse(uri)
	query := ""

	if err != nil {
		path = uri
		tx.Variables.RequestUri = uri
	} else {
		tx.ExtractGetArguments(parsedURL.RawQuery)
		tx.Variables.RequestUri = parsedURL.String()
		path = parsedURL.Path
		query = parsedURL.RawQuery
	}

	offset := strings.LastIndexAny(path, "/\\")
	if offset >= 0 && len(path) > offset+1 {
		tx.Variables.RequestBaseName = path[offset+1:]
	} else {
		tx.Variables.RequestBaseName = path
	}

	tx.Variables.RequestFileName = path
	tx.Variables.QueryString = query
}

func (tx *Transaction) ProcessRequestHeaders() bool {

	for _, rule := range RULES {
		if rule.Phase() == 1 || rule.Phase() == 100 {
			r := rule.Evaluate(*tx)
			if !r {
				return false
			}
		}
	}

	return true
}

func (tx *Transaction) ProcessRequestBody() bool {
	return false
}

func (tx *Transaction) ProcessResponseHeader() bool {
	return false
}

func (tx *Transaction) ProcessResponseBody() bool {
	return false
}

func (tx *Transaction) CaptureField(idx int, value string) {

}
