package rule_941

func matchXSS130(data []byte) bool {
%% machine xss130;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    var ts, te, act int
    _, _, _ = ts, te, act

    %%{

        word_bound = zlen | (space | punct)*;
        word_ele = [_0-9A-Za-z] ;
        not_word_ele = ^word_ele;

        identifier = [_A-Za-z][_0-9A-Za-z]*;
        # 941130
        pattern_tag = 'pattern' word_bound word_ele* '=';
        entity_tag = '!entity' space+ ('%' space+)? identifier+ space+ ('system' | 'public');
        x_tag = (not_word_ele)+ ('xlink:href' | 'xhtml' | 'xmlns' | 'data:text/html' | 'formaction'
            | '@import' | ';base64' | 'http' | pattern_tag | entity_tag);

        main := |*
            x_tag => {
                return true
            };

            any;
        *|;

        write init;
        write exec;
    }%%
        return false
}
