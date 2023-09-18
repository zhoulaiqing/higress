package rule_941

import (
    "fmt"
    "strings"
)

type machine210 struct {
    Builder strings.Builder
}

func (m *machine210) checkColon() bool {
    s := m.Builder.String()
    fmt.Println(s)

    if s == "javascript" || s == "vbscript" {
        m.Builder.Reset()
        return true
    }
    m.Builder.Reset()
    return false
}

func (m *machine210) checkHtmlSpace(word string) bool {
    fmt.Println(word)

    if word == "&newline;" || word == "&tab;" {
        return false
    } else if word == "&colon;" {
        return m.checkColon()
    }

    m.Builder.Reset()
    return false
}

func (m *machine210) appendWord(s string) {
    m.Builder.WriteString(s)
}

func matchXSS210(data []byte) bool {
%% machine xss210;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    var ts, te, act int
        _ = act

    m := &machine210{}

    %%{

        main := |*

            '&' alpha+ ';' => {
                if m.checkHtmlSpace(string(data[ts:te])) {
                    return true
                }
            };

            alpha+ => {
                m.appendWord(string(data[ts:te]))
            };

            ':' => {
                if m.checkColon() {
                    return true
                }
            };

            space => {
            };

            0 => {
                m.Builder.Reset()
            };

            any => {
                m.Builder.Reset()
            };

        *|;


        write init;
        write exec;
    }%%

    return false
}
