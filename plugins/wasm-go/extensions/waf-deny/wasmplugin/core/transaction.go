package core

type Transaction struct {
	id        string
	Variables TransactionVariables
}

type TransactionVariables struct {
	RequestMethod        string
	RequestProtocol      string
	RequestUri           string
	RequestBaseName      string
	RequestFileName      string
	RequestLine          string
	RequestHeaders       map[string]string
	RequestCookies       map[string]string
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

func (tx *Transaction) ProcessRequestHeaders() bool {
	return false
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
