package rule_941

func matchXSS(data []byte) bool {
%% machine xss;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof
    %%{
        main := any* '<script' any* '>' @{ return true } ;
        write init;
        write exec;
    }%%
        return false
}
