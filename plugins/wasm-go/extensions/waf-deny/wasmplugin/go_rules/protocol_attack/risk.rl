package protocol_attack

func matchRisk(data []byte) bool {
%% machine protocol_attack_risk;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)

    var ts, te, act int
    _, _, _ = ts, te, act

    %%{

        action is_eof {
            true if p == eof - 1
        }

        # eof
        EOF = zlen when is_eof;

        method = ('get' | 'post' | 'head' | 'options' | 'connect' | 'put' | 'delete' | 'trace' | 'track'
         | 'patch' | 'propfind' | 'propatch' | 'mkcol' | 'copy' | 'move' | 'lock' | 'unlock');

        purl = (^space)+;

        http = 'http/' digit;

        tag = 'html' | 'meta';
        word = [_0-9a-z];
        not_word = ^word;

        new_line = [\r\n];
        key_word = ('content-type' | 'content-length' | 'set-cookie' | 'location') ':';

        main := |*
            method space+ purl space+ http => {
                return true
            };

            '<' tag (not_word | EOF) => {
                return true
            };

            new_line not_word* key_word space* word => {
                return true
            };

            any;
        *|;

        write init;
        write exec;
    }%%
        return false
}
