package rule_tasks

import (
	ahocorasick "github.com/wasilibs/go-aho-corasick"
)

var SSRF = []string{
	"http://instance-data/latest/",
	"http://169.254.169.254/latest/",
	"http://2852039166/latest/",
	"http://025177524776/latest/",
	"http://0251.0376.0251.0376/latest/",
	"http://[::ffff:a9fe:a9fe]/latest/",
	"http://[0:0:0:0:0:ffff:a9fe:a9fe]/latest/",
	"http://[0:0:0:0:0:ffff:169.254.169.254]/latest/",
	"http://169.254.169.254.nip.io/latest/",
	"http://nicob.net/redir-http-169.254.169.254:80-",
	"http://2130706433/",
	"http://3232235521/",
	"http://3232235777/",
	"http://2852039166/",
	"http://[::]:",
	"http://169.254.170.2/v2",
	"http://169.254.169.254/computeMetadata/v1/",
	"http://metadata.google.internal/computeMetadata/v1/",
	"http://metadata/computeMetadata/v1/",
	"http://2852039166/computeMetadata/v1/",
	"http://[::ffff:a9fe:a9fe]/computeMetadata/v1/",
	"http://[0:0:0:0:0:ffff:a9fe:a9fe]/computeMetadata/v1/",
	"http://[0:0:0:0:0:ffff:169.254.169.254]/computeMetadata/v1/",
	"http://169.254.169.254.nip.io/computeMetadata/v1/",
	"http://metadata.google.internal/computeMetadata/v1/instance/disks/?recursive=true",
	"http://metadata.google.internal/computeMetadata/v1beta1/",
	"http://169.254.169.254/metadata/v1.json",
	"https://metadata.packet.net/userdata",
	"http://169.254.169.254/metadata/v1/",
	"http://169.254.169.254/metadata/instance?api-version=2017-04-02",
	"http://169.254.169.254/metadata/instance/network/interface/0/ipv4/ipAddress/0/publicIpAddress?api-version=2017-04-02&format=text",
	"http://2852039166/metadata/v1/",
	"http://[::ffff:a9fe:a9fe]/metadata/v1/",
	"http://[0:0:0:0:0:ffff:a9fe:a9fe]/metadata/v1/",
	"http://[0:0:0:0:0:ffff:169.254.169.254]/metadata/v1/",
	"http://169.254.169.254.nip.io/metadata/v1/",
	"http://169.254.169.254/openstack",
	"http://169.254.169.254/2009-04-04/meta-data/",
	"http://192.0.0.192/latest/",
	"http://100.100.100.200/latest/meta-data/",
	"http://rancher-metadata/",
	"http://127.0.0.1:2375",
	"http://2130706433:2375/",
	"http://[::]:2375/",
	"http://[0000::1]:2375/",
	"http://[0:0:0:0:0:ffff:127.0.0.1]:2375/",
	"http://2130706433:2375/",
	"http://017700000001:2375/",
	"http://0x7f000001:2375/",
	"http://0xc0a80014:2375/",
	"http://127.0.0.1:2379",
	"http://169。254。169。254",
	"http://169｡254｡169｡254",
	"http://⑯⑨。②⑤④。⑯⑨｡②⑤④",
	"http://⓪ⓧⓐ⑨｡⓪ⓧⓕⓔ｡⓪ⓧⓐ⑨｡⓪ⓧⓕⓔ",
	"http://⓪ⓧⓐ⑨ⓕⓔⓐ⑨ⓕⓔ",
	"http://②⑧⑤②⓪③⑨①⑥⑥",
	"http://④②⑤｡⑤①⓪｡④②⑤｡⑤①⓪",
	"http://⓪②⑤①。⓪③⑦⑥。⓪②⑤①。⓪③⑦⑥",
	"http://⓪⓪②⑤①｡⓪⓪⓪③⑦⑥｡⓪⓪⓪⓪②⑤①｡⓪⓪⓪⓪⓪③⑦⑥",
	"http://[::①⑥⑨｡②⑤④｡⑯⑨｡②⑤④]",
	"http://[::ⓕⓕⓕⓕ:①⑥⑨。②⑤④。⑯⑨。②⑤④]",
	"http://⓪ⓧⓐ⑨。⓪③⑦⑥。④③⑤①⑧",
	"http://⓪ⓧⓐ⑨｡⑯⑥⑧⑨⑥⑥②",
	"http://⓪⓪②⑤①。⑯⑥⑧⑨⑥⑥②",
	"http://⓪⓪②⑤①｡⓪ⓧⓕⓔ｡④③⑤①⑧",
	"jar:http://127.0.0.1!/",
	"jar:https://127.0.0.1!/",
	"jar:ftp://127.0.0.1!/",
	"gopher://127.0.0.1",
	"gopher://localhost",
	"http://localhost:9001/2018-06-01/runtime/",
}
var Rule934110Matcher ahocorasick.AhoCorasick

func init() {

	Rule934110Matcher = AHO_CORASICK_BUILDER.Build(SSRF)
}
