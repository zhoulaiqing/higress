package rule_941

import (
    "fmt"
)

func matchXSS130(data []byte) bool {
%% machine xss;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    lastIsWordEle := false

    %%{

        action markWordEle {
            fmt.Printf("Word Ele: %s \n", string(data[p]))
            lastIsWordEle = true
        }

        action markNotWordEle {
            fmt.Printf("Not word Ele: %s \n", string(data[p]))
            lastIsWordEle = false
        }

        action enterMatch {
            if !lastIsWordEle {
                fmt.Println(lastIsWordEle)
            } else {
                fmt.Println("no")
            }
        }

        action setMatched {
            return true
        }

        word_bound = zlen | (space | punct)*;
        word_ele = [_0-9A-Za-z] ;
        not_word_ele = ^word_ele;

        identifier = [_A-Za-z][_0-9A-Za-z]*;
        # 941130
        pattern_tag = 'pattern' word_bound word_ele* '=';
        entity_tag = '!entity' space+ ('%' space+)? identifier+ space+ ('system' | 'public');
        x_tag = any+ (not_word_ele)+ ('xlink:href' | 'xhtml' | 'xmlns' | 'data:text/html' | 'formaction'
            | '@import' | ';base64' | pattern_tag | entity_tag) >enterMatch  %/setMatched word_bound @setMatched;

        main := x_tag;
        write init;
        write exec;
    }%%
        return false
}
