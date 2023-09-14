package rule_941

func matchXSS170(data []byte) bool {
%% machine xss170;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    var ts, te, act int
    _, _, _ = ts, te, act

    %%{

        word = [_0-9a-z];
        word_plus = (word | '+' | '-');
        iden = [a-z]word*;

        data_suffix1 = iden '/' word word_plus+ word [;,];
        data_suffix2 = [^/;]* ';' any* ^word ('base64' | 'charset=');
        data_suffix3 = [^/,]* ',' any* '<' any* word any* '>';

        data = 'data:' ( data_suffix1 | data_suffix2 | data_suffix3 );

        temp = [!+ ];
        ptn360 = '!' temp '[' ']';

        lt = '<';
        alt = '+adw-';
        rt = '>';
        art = '+ad4-';

        main := |*
            lt any* art => {
                return true
            };

            alt any* (rt | art) => {
                return true
            };

            data => {
                return true
            };

            ptn360 => {
                return true
            };


            any;

        *|;

        write init;
        write exec;
    }%%

    return false
}
