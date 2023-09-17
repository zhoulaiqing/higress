package generic_attack

import (
    "strings"
    "golang.org/x/exp/slices"
)

var fnsForProcess = []string{
	"access", "appendfile", "argv", "availability",
	"caveats", "chmod", "chown", "close", "copyfile", "cp", "createreadstream", "createwritestream",
	"exec", "execfile", "exists",
	"fchmod", "fchown", "fdata", "fdatasync", "fstat", "fsync", "futimes",
	"inodes",
	"lchmod", "link", "lstat", "lutimes",
	"mkdir", "mkdtemp",
	"open", "opendir",
	"read", "readdir", "readfile", "readlink", "readv", "rename", "rm",
	"spawn", "spawnfile", "stat", "symlink",
	"truncate",
	"unlink", "unwatchfile", "utimes",
	"watchfile", "write", "writefile", "writev",
}

var keywordsForProcess = []string{
	"binding", "constructor", "env", "global", "main", "mainmodule", "process", "require",
}

var fnsForConsole = []string{
	"debug", "error", "info", "trace", "warn",
}

var fnsForRequire = []string{
	"resolve", "main", "extensions", "cache",
}

func checkPattern1(value string) bool {
	snippets := strings.Split(value, ".")

	switch snippets[0] {
	case "this":
		return snippets[1] == "constructor"
	case "string":
	    return snippets[1] == "fromcharcode"
	case "module":
		return snippets[1] == "exports"
	case "console":
		return slices.Contains(fnsForConsole, snippets[1])
	case "require":
		return slices.Contains(fnsForRequire, snippets[1])
	case "process":
		return slices.Contains(fnsForProcess, snippets[1]) || slices.Contains(keywordsForProcess, snippets[1])
	}

	return false
}

func checkPattern2(value string) bool {
	idx := strings.IndexByte(value, '[')
	key := value[:idx]

	// ["xxx"]
	attr := value[idx+2 : len(value)-2]

	switch key {
	case "console":
		return slices.Contains(fnsForConsole, attr)
	case "require":
		return slices.Contains(fnsForRequire, attr)
	case "process":
		return slices.Contains(fnsForProcess, attr) || slices.Contains(keywordsForProcess, attr)
	}

	return false
}



func matchInsecureUnserialize(data []byte) bool {
%% machine insecure_unserialize;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)

    var ts, te, act int
    _ = act

    %%{

        action is_eof {
            true if p == eof - 1
        }

        # eof
        EOF = zlen when is_eof;

        not_word = [^0-9a-zA-Z];

        libs = '_' ('$$nd_func$$_' | '_js_function');
        snippet1 = ('eval' | 'newfunction') '(';
        snippet2 = 'function' '(' ')'  '{';
        snippet3 = '(' not_word 'child_process' not_word ')';
        snippet4 = ('process' | 'binding' | 'constructor' | 'env' | 'global' | 'main' ('module')? | 'require') '[';

        ptn1 = alpha+ ('.' alpha+ )+ ;
        quotes = '"' | "'" | '`' ;
        ptn2 = alpha+ '[' quotes alpha+ quotes (']' | EOF);

        main := |*
            libs | snippet1 | snippet2 | snippet3 | snippet4 => {
                return true
            };

            ptn1 => {
                if checkPattern1(string(data[ts:te])) {
                    return true
                }
            };

            ptn2 => {
                if checkPattern2(string(data[ts:te])) {
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