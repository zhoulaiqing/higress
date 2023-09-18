package generic_attack

func matchNodejsDos(data []byte) bool {
%% machine nodejsDos;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    var ts, te, act int
    _, _, _ = ts, te, act

    %%{

        sgn = '+' | '-';

        emptyStr = '"'{2} | "'"{2} | '`'{2};
        notEmptyStr = '"' (any -- '"')+ '"' | "'" (any -- "'")+ "'" | '`' (any -- '`')+ '`';
        riskBody = '{' any* '}' | '[' any* ']';

        falseCond = 'false' | 'null' | 'undefined' | 'nan' | sgn? '0' | emptyStr;

        notZero = sgn? ('infinity' | [1-9] alpha*);
        newSent = 'new' space+ [a-z][0-9a-z]*;
        trueCond = 'true' | 'this' | notZero | newSent | 'window' | 'string' | 'boolean' | 'function' | 'object' | 'array'
            | riskBody;

        whileStart = 'while' space* '(' (space | '(')*;

        trueStmt = '!'+ falseCond | '!!'* trueCond;

        main := |*

            whileStart trueStmt any* ')' => {
                return true
            };

            any;
        *|;

        write init;
        write exec;
    }%%
        return false
}