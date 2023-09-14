package rule_941


func matchXSS170(data []byte) bool {
%% machine xss;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof



    %%{

        action setMatched {
            return true
        }

        word_ele = [_0-9A-Za-z] ;
        not_word_ele = ^word_ele;
        identifier = [_a-z][_0-9a-z]+;

        word_bound = (any* not_word_ele)?;
        not_word = not_word_ele*;

        data_suffix1 = (identifier '/' word_ele ( word_ele | '+' | '-' )+ word_ele)? [;,];
        data_suffix2 = any* ';' word_bound ('base64'|'charset=');
        data_suffix3 = any* ',' any* '<' any* word_ele any* '>';

        data = 'data:' (data_suffix1 | data_suffix2 | data_suffix3) %/setMatched not_word_ele @setMatched;

        main := word_bound data ;

        write init;
        write exec;
    }%%

    return false
}
