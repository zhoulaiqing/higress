package rule_941

func matchXSS350(data []byte) bool {
%% machine xss;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    // -1. No start; 0. Normal start; 1. Dangerous start
    startTag := -1

    var ts, te, act int
    _, _, _ = ts, te, act

    %%{
        main := |*
            '+adw-'  {
                startTag = 1
            };

            '<' => {
                startTag = 0
            };

            '+ad4-' => {
                if startTag >= 0 {
                    return true
                }
            };

            '>' => {
                if startTag > 0 {
                    return true
                }
            };

            any;

        *|;

        write init;
        write exec;
    }%%

    return false
}
