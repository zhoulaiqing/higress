package rule_941


func matchXSS200(data []byte) bool {
%% machine xss200;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    var ts, te, act int
        _, _, _ = ts, te, act

    step := 0

    %%{

        part1 = '<' any* ':'? 'vmlframe';
        part2 = 'src' (space | '/' | '+')* '=';

        main := |*
            part1 => {
                step = 1
            };

            part2 => {
                if step == 1 {
                    return true
                } else {
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
