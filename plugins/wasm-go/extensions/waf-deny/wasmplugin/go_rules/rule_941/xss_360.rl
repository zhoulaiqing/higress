package rule_941

func matchXSS360(data []byte) bool {
%% machine xss;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof
    %%{
        action setMatched {
            return true
        }

        temp = [!+ ];
        main := any* '!' temp '[' ']' %/setMatched any* @setMatched;

        write init;
        write exec;
    }%%

    return false
}
