package rce_attack

import (
    "fmt"
    "strings"
    "golang.org/x/exp/slices"
)

var commands = []string{
"fd", "xz", "ssh", "esh", "irb", "env", "scp", "7zx", "get", "bzz", "csh", "rcp", "rc", "ls", "svn", "gpg", "ksh",
"zsh", "sh", "7z", "lz4", "udp", "sed", "7zr", "lz", "pxz", "fg", "w3m", "df", "hup", "7za", "php", "pwd",
}

func checkCommand(name string) bool {

    name = strings.ReplaceAll(name, ";", "")

    fmt.Printf("possible command name %s \n", name)
    if slices.Contains(commands, name) {
        return true
    }
    if strings.HasPrefix(name, "gcc") {
        return true
    }

    return false
}

func matchShortCommandBase(data []byte) bool {
%% machine shortCommandBase;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    pb := 0
    var commandName string

    var ts, te, act int
    // _, _, _ = ts, te, act
    _ = act

    %%{

        action is_eof {
            true if p == eof - 1
        }

        action mark {
            pb = p
        }

        action setName {
            commandName = string(data[pb:p])
        }

        # eof
        EOF = zlen when is_eof;

        bracketed = '(' space* ')';
        escape = ';' | '{' | '|' | '&' | '\n' | '\r' | '$(' | '$((' | '<(' | '>(' | '`' | bracketed;

        cmdPrefix = '{' | space* '(' space* | '!' space* | '$';
        word = [0-9A-Za-z_];

        path = ('?' | '*' | '[' | ']' |'(' | ')' |'-' | '+' | word+ | '.' | '/' )? '/';

        quotes = '"' | "'"
        cmdLiteral = word | qutoes | '\\' | '^' ;

        command = escape space* cmdPrefix* space* path** cmdLiteral{2,} >mark %setName;

        main := |*
            command => {

                fmt.Printf("Full command: %s \n", string(data[ts:te]))
                if checkCommand(commandName) {
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
