package core

import (
	"errors"
	"fmt"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core/url_util"
	"math"
	"net/url"
	"strconv"
	"strings"
)

type Transaction struct {
	id                string
	RequestBodyBuffer *BodyBuffer
	Variables         TransactionVariables
}

func NewTransaction() Transaction {
	id := RandomString(16)
	variables := TransactionVariables{}
	variables.RequestHeaders = make(map[string]string)
	variables.RequestCookies = make(map[string]string)
	variables.ArgsGet = make(map[string]string)
	variables.ArgsPost = make(map[string]string)
	variables.ArgsPath = make(map[string]string)
	variables.Files = make(map[string][]string)
	variables.XML = make(map[string][]string)
	variables.MultipartPartHeaders = make(map[string][]string)
	variables.Args = append(variables.Args, &variables.ArgsGet, &variables.ArgsPost, &variables.ArgsPath)
	variables.TransMap = make(map[string][]string)

	return Transaction{id: id, Variables: variables, RequestBodyBuffer: NewBodyBuffer()}
}

type TransactionVariables struct {
	RemoteAddr           string
	RemotePort           string
	ServerAddr           string
	ServerPort           string
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
	Files                map[string][]string
	XML                  map[string][]string
	MultipartPartHeaders map[string][]string
	ResponseArgs         map[string]string
	FilesCombinedSize    string

	Skip941ForFileName bool
	TransMap           map[string][]string
	Interrupted        bool

	BlockingInboundAnomalyScore  int
	InboundAnomalyScorePl1       int
	InboundAnomalyScorePl2       int
	InboundAnomalyScorePl3       int
	InboundAnomalyScorePl4       int
	SqlInjectionScore            int
	XssScore                     int
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
	if len(key) == 0 {
		return
	}

	keyl := strings.ToLower(key)
	tx.Variables.RequestHeaders[keyl] = value

	switch keyl {
	case "content-type":
		val := strings.ToLower(value)
		if val == "application/x-www-form-urlencoded" {
			tx.Variables.ReqBodyProcessor = "URLENCODED"
		} else if strings.HasPrefix(val, "multipart/form-data") {
			tx.Variables.ReqBodyProcessor = "MULTIPART"
		}
	case "cookie":
		values := url_util.ParseQuery(value, ';')
		for k, vr := range values {
			for _, v := range vr {
				tx.Variables.RequestCookies[k] = v
			}
		}
	}

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

func (tx *Transaction) ProcessConnection(client string, cPort int, server string, sPort int) {
	p := strconv.Itoa(cPort)
	p2 := strconv.Itoa(sPort)

	tx.Variables.RemoteAddr = client
	tx.Variables.RemotePort = p
	tx.Variables.ServerAddr = server
	tx.Variables.ServerPort = p2
}

// WriteRequestBody /*
func (tx *Transaction) WriteRequestBody(b []byte) (bool, int, error) {
	writingBytes := int64(len(b))
	if tx.RequestBodyBuffer.length >= (math.MaxInt64 - writingBytes) {
		return false, 0, errors.New("overflow reached while writing request body")
	}

	w, err := tx.RequestBodyBuffer.Write(b[:writingBytes])
	if err != nil {
		return false, 0, err
	}

	return true, int(w), err
}

func (tx *Transaction) ProcessRequestHeaders() bool {
	return true
}

func (tx *Transaction) ProcessRequestBody() (bool, error) {

	if tx.RequestBodyBuffer.length == 0 {
		return true, nil
	}

	mime := tx.Variables.RequestHeaders["content-type"]
	reader, err := tx.RequestBodyBuffer.Reader()
	if err != nil {
		return false, err
	}

	// todo 注意这里暂时没有启用 XML 和 JSON 的解析逻辑，后续可以根据需要使用。参考 coraza.conf-recommended.conf
	rbp := strings.ToLower(tx.Variables.ReqBodyProcessor)
	if len(rbp) == 0 {
		return true, nil
	}

	bodyprocessor, err := GetBodyProcessor(rbp)
	if err != nil {
		return false, errors.New("invalid body processor")
	}

	if err := bodyprocessor.ProcessRequest(reader, tx, mime); err != nil {
		return false, errors.New("failed to process request body")
	}

	return true, nil
}

func (tx *Transaction) ProcessResponseHeader() bool {
	return false
}

func (tx *Transaction) ProcessResponseBody() bool {
	return false
}

func (tx *Transaction) CaptureField(idx int, value string) {

}

func (tx *Transaction) Close() error {

	err := tx.RequestBodyBuffer.Reset()
	if err != nil {
		return err
	}

	return nil
}
