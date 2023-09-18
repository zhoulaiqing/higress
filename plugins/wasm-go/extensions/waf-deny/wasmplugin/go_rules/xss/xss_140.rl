package rule_941


func matchXSS140(data []byte) bool {
%% machine xss140;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)
       _ = eof
    var ts, te, act int
    _, _, _ = ts, te, act

    %%{
        temp1 = [^:=];
        main := |*
            lower+ '=' (temp1+ ':' any+ ';')* temp1+ ':url(javascript' => {
                return true
            };

            any;
        *|;

        write init;
        write exec;
    }%%

    return false
}
