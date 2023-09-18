package php_attack

import (
    "fmt"
    "golang.org/x/exp/slices"
)

var wrappers = []string{
"bzip2", "expect", "glob", "ogg", "phar", "rar", "zip", "zlib",
"ssh2", "ssh2.shell", "ssh2.sftp", "ssh2.scp", "ssh2.exec", "ssh2.tunnel",
}

func checkWrapper(value string) bool {
    v := value[:len(value)-3]
    fmt.Printf("May be wrapper: %s \n", v)
    return slices.Contains(wrappers, v)
}

func matchPhpWrapper(data []byte) bool {
%% machine phpWrapper;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    var ts, te, act int
    _ = act

    %%{

        wrapper = alpha+ ('.' alpha+)?;
        phpWrapper = wrapper '://';

        main := |*
            phpWrapper => {
                if checkWrapper(string(data[ts:te])) {
                    return true
                }
            };

            any;
        *|;

        write init;
        write exec;
    }%%
        return false
}