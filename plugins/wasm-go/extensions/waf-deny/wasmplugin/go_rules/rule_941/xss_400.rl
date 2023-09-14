package rule_941

import (
    "fmt"
)

func matchXSS400(data []byte) bool {
%% machine xss400;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    var ts, te, act int
        _, _, _ = ts, te, act

    step := 0

    %%{
        action setMatched {
            return true
        }

        part11 = '[' any* ']' [^.]* ;
        part12 = 'reflect' [^.]* ;

        part2 = '.' any* ('map' | 'sort' | 'apply') [^.]* '.';

        part3 = 'call' any* '`' any* '`';

        main := |*
            part11 | part12 => {
                fmt.Printf("match step1 %s \n", string(data[ts:te]))
                step = 1
            };

            part2 => {
                if step == 1 {
                    fmt.Printf("match step2 %s \n", string(data[ts:te]))
                    step = 2
                } else {
                    fmt.Println("back when match 2")
                    step = 0
                }
            };

            part3 => {
                if step == 2 {
                    fmt.Printf("match step3 %s \n", string(data[ts:te]))
                    return true
                } else {
                    fmt.Println("back when match 3")
                    step = 0
                }
            };

            any;
        *|;

        write init;
        write exec;
    }%%

    return false
}
