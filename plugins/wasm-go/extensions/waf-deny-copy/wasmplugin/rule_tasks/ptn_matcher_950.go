package rule_tasks

import (
	"github.com/wasilibs/go-re2"
)

const (
	Ptn950130 = `(?:<(?:TITLE>Index of.*?<H|title>Index of.*?<h)1>Index of|>\[To Parent Directory\]</[Aa]><br>)`
	Ptn950140 = `^#\!\s?/`
)

var Re950130, Re950140 *re2.Regexp

func init() {
	Re950130, _ = re2.Compile(Ptn950130)
	Re950140, _ = re2.Compile(Ptn950140)
}
