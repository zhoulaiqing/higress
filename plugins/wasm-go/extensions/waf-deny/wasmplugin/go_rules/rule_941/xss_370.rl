package rule_941

import (
    "fmt"
    "strings"
    "golang.org/x/exp/slices"
)
var builder370 strings.Builder

var keys370 = []string{"self", "document", "this", "top", "window"}

// 0. None; 1. Passed word; 2. Passed left;
var step370 = 0

func appendWord370(s string) {
    builder370.WriteString(s)
}

func checkWord370() {
    if step370 >= 1 {
        return
    }

    s := builder370.String()
    fmt.Println(s)

    if slices.Contains(keys370, s) {
        step370 = 1
    } else {
        builder370.Reset()
    }
}

func checkLeft370() {
    if step370 >= 1 {
        step370 = 2
    } else {
        checkWord370()
        if step370 >= 1 {
            step370 = 2
        } else {
            builder370.Reset()
        }
    }
}

func checkRight370() bool {
    if step370 == 2 {
        return true
    } else {
        builder370.Reset()
        return false
    }
}

func matchXSS370(data []byte) bool {
%% machine xss;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    var ts, te, act int
            _ = act

    %%{

        main := |*
            alpha+ => {
                appendWord370(string(data[ts:te]))
            };

            '/*' | '[' | ')' => {
                checkLeft370()
            };

            ']' | '*/' => {
                if checkRight370() {
                    return true
                }
            };

            space => {
                checkWord370()
            };

            any => {
                if step370 < 2 {
                    builder370.Reset()
                }

            };
        *|;

        write init;
        write exec;
    }%%

    return false
}
