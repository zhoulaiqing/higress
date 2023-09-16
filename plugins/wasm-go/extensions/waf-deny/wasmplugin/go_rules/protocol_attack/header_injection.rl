package protocol_attack

func matchHeaderInjection(data []byte) bool {
%% machine headerInjection;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    var ts, te, act int
    _, _, _ = ts, te, act

    %%{

        new_line = [\r\n];
        word = [^0-9a-z_];
        not_word = ^word;

        header = space | 'location' | 'refresh' | 'set-'? 'cookie' | 'x-'? 'forwarded-' ('for' | 'host' | 'server')
            | 'host' | 'via' | 'remote-ip' | 'remote-addr' | 'originating-ip' | 'content-' ('type' | 'length');

        main := |*
            new_line+ header space* ':' => {
                return true
            };

            any;
        *|;

        write init;
        write exec;
    }%%
        return false
}

