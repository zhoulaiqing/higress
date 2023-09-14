package rule_941

import (
    "fmt"
    "strings"
    "golang.org/x/exp/slices"
)

type machine390 struct {
    Builder strings.Builder
    matchedWord bool
}

var keys390 = []string{"eval", "settimeout", "setinterval", "newfunction", "alert", "atob", "btoa"}

func (m *machine390) appendWord(s string) {
    m.Builder.WriteString(s)
}

func (m *machine390) reset() {
    m.Builder.Reset()
    m.matchedWord = false
}

func (m *machine390) checkSpace() {
    if m.matchedWord {
        return
    }

    if m.checkWord() >= 0 {
        return
    }

    m.reset()
}

func (m *machine390) checkWord() int {
    s := m.Builder.String()
    fmt.Println(s)

    if slices.Contains(keys390, s) {
        m.matchedWord = true
        return 1
    } else if s == "new" {
        return 0
    }

    return -1
}

func (m *machine390) checkFinish() bool {
    if m.matchedWord {
        return true
    }

    if m.checkWord() > 0 {
        return true
    }

    return false
}

func matchXSS390(data []byte) bool {
%% machine xss390;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    var ts, te, act int
        _ = act

    m := &machine390{}

    %%{
        action setMatched {
            return true
        }

        main :=  |*
            alpha+ => {
                m.appendWord(string(data[ts:te]))
            };

            '(' => {
                if m.checkFinish() {
                    return true
                }
            };

            space => {
                m.checkSpace()
            };

            any => {
                m.reset()
            };
        *|;

        write init;
        write exec;
    }%%

    return false
}
