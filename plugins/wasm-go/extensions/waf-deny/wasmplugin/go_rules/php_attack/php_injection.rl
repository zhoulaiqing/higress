package php_attack

func matchPhpInjection(data []byte) bool {
%% machine phpInjection;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    %%{

        action setMatched {
            return true
        }

        

        write init;
        write exec;
    }%%
        return false
}