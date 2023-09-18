package protocol_attack

func matchBodyProcessorBypass(data []byte) bool {
%% machine bodyProcessorBypass;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    %%{

        action setMatched {
            return true
        }

        ctypePrefix = 'application/' | 'text/';
        keyword = space | ',' | ';';

        main := ^keyword+ keyword any* ctypePrefix @setMatched ;

        write init;
        write exec;
    }%%
        return false
}

