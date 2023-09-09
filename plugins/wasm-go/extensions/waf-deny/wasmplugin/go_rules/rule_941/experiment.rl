package rule_941

func matchExp(data []byte) bool {
%% machine xss;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof
    %%{

        action setMatched {
            return true
        }

        word_bound = zlen | (space | punct)*;
        word_ele = [_0-9A-Za-z];
        identifier = [_A-Za-z][_0-9A-Za-z]*;

        quotes = '"' | '\'';
        attribute_start = ('<' word_ele* (space | '/')) | (quotes (any* (space | '/'))?);

        html_attribute = attribute_start;


        main := html_attribute;
        write init;
        write exec;
    }%%
        return false
}
