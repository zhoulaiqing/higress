package rule_941

import (
    "fmt"
    "strings"
)
var builder strings.Builder

func checkColon() bool {
    s := builder.String()
    fmt.Println(s)

    if s == "javascript" || s == "vbscript" {
        fmt.Println("here")
        builder.Reset()
        return true
    }
    builder.Reset()
    return false
}

func checkHtmlSpace(word string) bool {
    fmt.Println(word)

    if word == "&newline;" || word == "&tab;" {
        return false
    } else if word == "&colon" {
        return checkColon()
    }

    builder.Reset()
    return false
}

func appendWord(s string) {
    builder.WriteString(s)
}

func matchXSS210(data []byte) bool {
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

        main := |*

            '&' alpha+ ';' => {
                if checkHtmlSpace(string(data[ts:te])) {
                    return true
                }
            };

            alpha+ => {
                appendWord(string(data[ts:te]))
            };

            ':' => {
                if checkColon() {
                    return true
                }
            };

            space => {
            };

            any => {
                builder.Reset()
            };

        *|;


        write init;
        write exec;
    }%%

    return false
}
