package rce_attack

import (
    "fmt"
    "golang.org/x/exp/slices"
)

var commands = []string{
"fd", "xz", "ssh", "esh", "irb", "env", "scp", "7zx", "get", "bzz", "csh", "rcp", "rc", "ls", "svn", "gpg", "ksh",
"zsh", "sh", "7z", "lz4", "udp", "sed", "7zr", "lz", "pxz", "fg", "w3m", "df", "hup", "7za", "php", "pwd",
}

func checkCommand(name string) bool {
    return false
}

func matchShortCommandBase(data []byte) bool {
%% machine shortCommandBase;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    var ts, te, act int
    _ = act

    %%{

        bracketed = '(' space* ')';
        escape = ';' | '{' | '|' | '&' | '\n' | '\r' | '$(' | '$((' | '<(' | '>(' | '`' | bracketed;

        cmdPrefix = '{' | space* '(' space* | '!' space* | '$';

        path =

        command = escape space* cmdPrefix*

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
