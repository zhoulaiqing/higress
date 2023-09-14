package rule_941

import (
    "fmt"
    "strings"
    "golang.org/x/exp/slices"
)

type machine370 struct {
    Builder strings.Builder
    // 0. None; 1. Passed word; 2. Passed left;
    step int
}

var keys370 = []string{"self", "document", "this", "top", "window"}

func (m *machine370) appendWord(s string) {
    m.Builder.WriteString(s)
}

func (m *machine370) checkWord() {
    if m.step >= 1 {
        return
    }

    s := m.Builder.String()
    fmt.Println(s)

    if slices.Contains(keys370, s) {
        m.step = 1
    } else {
        m.Builder.Reset()
    }
}

func (m *machine370) checkLeft() {
    if m.step >= 1 {
        m.step = 2
    } else {
        m.checkWord()
        if m.step >= 1 {
            m.step = 2
        } else {
            m.Builder.Reset()
        }
    }
}

func (m *machine370) checkRight() bool {
    if m.step == 2 {
        return true
    } else {
        m.Builder.Reset()
        return false
    }
}

func matchXSS370(data []byte) bool {
%% machine xss370;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    var ts, te, act int
            _ = act

    m := &machine370{}

    %%{

        main := |*
            alpha+ => {
                m.appendWord(string(data[ts:te]))
            };

            '/*' | '[' | ')' => {
                m.checkLeft()
            };

            ']' | '*/' => {
                if m.checkRight() {
                    return true
                }
            };

            space => {
                m.checkWord()
            };

            any => {
                if m.step < 2 {
                    m.Builder.Reset()
                }

            };
        *|;

        write init;
        write exec;
    }%%

    return false
}
