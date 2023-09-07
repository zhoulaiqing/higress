package rule_941

func matchXSS(data []byte) bool {
%% machine xss;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof
    %%{

        action setMatched {
            return true
        }

        word_bound = zlen | (space | punct)*;
        word_ele = [_0-9A-Za-z]*;

        # 941110
        script_tag = any* '<script' any* '>' @setMatched;

        # 941130
        x_tag = word_bound ('xlink:href' | 'xhtml' | 'xmlns' | 'data:text/html' | 'formaction')  %/setMatched word_bound @setMatched;
        pattern_tag = word_bound 'pattern' word_bound word_ele '=' word_ele %/setMatched word_bound @setMatched;

        main := script_tag | x_tag | pattern_tag;
        write init;
        write exec;
    }%%
        return false
}
