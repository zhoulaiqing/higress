package core

const (
	INBOUND_ANOMALY_SCORE_THRESHOLD  = 5
	OUTBOUND_ANOMALY_SCORE_THRESHOLD = 4
	REPORTING_LEVEL                  = 4
	EARLY_BLOCKING                   = 0
	BLOCKING_PARANONIA_LEVEL         = 1
	CRITICAL_ANOMALY_SCORE           = 5
	ENFORCE_BODYPROC_URLENCODED      = 0
	CRS_VALIDATE_UTF8_ENCODING       = 0
)

var (
	ALLOWED_METHODS              = []string{"GET", "HEAD", "OPTIONS", "POST"}
	ALLOWED_REQUEST_CONTENT_TYPE = [9]string{
		"application/x-www-form-urlencoded",
		"multipart/form-data",
		"multipart/related",
		"text/xml",
		"application/xml",
		"application/soap+xml",
		"application/json",
		"application/cloudevents+json",
		"application/cloudevents-batch+json",
	}
	ALLOWED_REQUEST_CONTENT_TYPE_CHARSET = [4]string{"utf-8", "iso-8859", "iso-8859-15", "windows-1252"}
	ALLOWED_HTTP_VERSIONS                = [4]string{"HTTP/1.0", "HTTP/1.1", "HTTP/2", "HTTP/2.0"}
	RESTRICTED_EXTENSIONS                = [51]string{
		".asa", ".asax", ".ascx",
		".backup", ".bak", ".bat",
		".cdx", ".cer", ".cfg", ".cmd", ".com", ".config", ".conf", ".cs", ".csproj", ".csr",
		".dat", ".db", ".dbf", ".dll", ".dos",
		".htr", ".htw",
		".ida", ".idc", ".idq", ".inc", ".ini",
		".key",
		".licx", ".lnk", ".log",
		".mdb",
		".old",
		".pass", ".pdb", ".pol", ".printer", ".pwd", ".rdb",
		".resources", ".resx",
		".sql", ".swp", ".sys", ".vb", ".vbs", ".vbproj", ".vsdisco",
		".xsd", ".xsx",
	}
	RESTRICTED_HEADERS = []string{"accept-charset", "content-encoding", "proxy", "lock-token", "content-range", "if",
		"x-http-method-override", "x-http-method", "x-method-override"}
)

type RuleEngine struct {
}
