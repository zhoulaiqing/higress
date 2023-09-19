package php_attack

func matchPhpObjectInjection(data []byte) bool {
%% machine phpObjectInjection;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    var ts, te, act int
    _, _, _ = ts, te, act

    %%{
        objectInjection = ('o' | 'c') ':' digit+ ':"' (any+ -- '"') '":' digit+ ':{' (any* -- '}') '}';

        main := |*

            objectInjection => {
                return true
            };

            any;
        *|;

        write init;
        write exec;
    }%%
        return false
}