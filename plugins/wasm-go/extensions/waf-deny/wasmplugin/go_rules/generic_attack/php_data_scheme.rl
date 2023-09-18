package generic_attack

func matchPhpDataScheme(data []byte) bool {
%% machine phpDataScheme;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    var ts, te, act int
    _, _, _ = ts, te, act

    %%{
        action setMatched {
            return true
        }

        simple = [^!-\"\(-\),/:-\?\[-\]\{\}]+ | '*';

        contentTypeFormat = '*' | simple '/' simple;

        main := 'data:' contentTypeFormat %/setMatched any @setMatched;

        write init;
        write exec;
    }%%
        return false
}

