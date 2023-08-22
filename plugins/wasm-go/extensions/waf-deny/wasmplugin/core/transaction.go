package core

type Transaction struct {
	id        string
	Variables TransactionVariables
}

type TransactionVariables struct {
	RequestMethod   string
	RequestProtocol string
	RequestFileName string
	RequestLine     string
	RequestHeaders  map[string]string
	ArgsGet         map[string]string
	ArgsPost        map[string]string
	ArgsPath        map[string]string
	Args            []*map[string]string
	Files           map[string]map[string]string
	FilesNames      map[string]map[string]string

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
