package rule_941

import (
    "fmt"
    "strings"
    "golang.org/x/exp/slices"
)

var builder390 strings.Builder
var keys390 = []string{"eval", "settimeout", "setinterval", "newfunction", "alert", "atob", "btoa"}

func appendWord390(s string) {
    builder390.WriteString(s)
}

func reset390 () {
    builder390.Reset()
    matchedWord = false
}

func checkSpace390() {
    if matchedWord {
        return
    }

    if checkWord390() >= 0 {
        return
    }

    reset390()
}

func checkWord390() int {
    s := builder390.String()
    fmt.Println(s)

    if slices.Contains(keys390, s) {
        matchedWord = true
        return 1
    } else if s == "new" {
        return 0
    }

    return -1
}

func checkFinish390() bool {
    if matchedWord {
        return true
    }

    if checkWord390() > 0 {
        return true
    }

    return false
}

var matchedWord = false

func matchXSS390(data []byte) bool {
%% machine xss;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    var ts, te, act int
        _ = act

    %%{
        action setMatched {
            return true
        }

        main :=  |*
            alpha+ => {
                appendWord390(string(data[ts:te]))
            };

            '(' => {
                if checkFinish390() {
                    return true
                }
            };

            space => {
                checkSpace390()
            };

            any => {
                reset390()
            };
        *|;

        write init;
        write exec;
    }%%

    return false
}
