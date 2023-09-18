package php_attack

func matchPhpInjection(data []byte) bool {
%% machine phpInjection;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    var ts, te, act int
    _, _, _ = ts, te, act

    %%{

        dangerTags = '<?' !'xml';
        otherPhpTags = '[' ('/' | '\\') 'php' ']';

        io = 'std' ('in' |'out'|'err') | ('in'|'out') 'put' | 'fd' | 'memory' | 'temp'| 'filter';
        phpIO = 'php://' io;

        ssh2 = 'ssh2' ('.' ('shell' | 'sftp' | 'scp' | 'exec' | 'tunnel'))?;
        wrapper = 'bzip2' | 'expect' | 'glob' | 'ogg' | 'phar' |'rar' | ssh2 | 'zip' | 'zlib';
        phpWrapper = wrapper '://';

        objectInjection = ('o' | 'c') ':' digit+ ':"' (any+ -- '"') '":' digit+ ':{' (any* -- '}') '}';

        main := |*

            dangerTags | otherPhpTags => {
                return true
            };

            phpIO => {
                return true
            };

            phpWrapper => {
                return true
            };

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