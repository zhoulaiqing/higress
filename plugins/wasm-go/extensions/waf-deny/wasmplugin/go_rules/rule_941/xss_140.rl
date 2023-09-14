package rule_941


func matchXSS140(data []byte) bool {
%% machine xss140;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof
    %%{
        action setMatched {
            return true
        }

        temp1 = [^:=];
        main := any* lower+ '=' (temp1+ ':' any+ ';')* temp1+ ':url(javascript' %setMatched any*;

        write init;
        write exec;
    }%%

    return false
}
