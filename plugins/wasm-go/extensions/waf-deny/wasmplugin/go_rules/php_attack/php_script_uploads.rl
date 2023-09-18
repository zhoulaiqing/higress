package php_attack

func matchPhpScriptUploads(data []byte) bool {
%% machine phpScriptUploads;
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

        ext = 'php' digit* | 'phtml' | 'phar' | 'phps' | 'pht' | 'phpt';

        main := |*

            '.' ext '.'* EOF => {
                return true
            };

            any;
        *|;

        write init;
        write exec;
    }%%
        return false
}