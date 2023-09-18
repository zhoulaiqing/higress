package php_attack

func matchPhpVariableFunction(data []byte) bool {
%% machine phpVariableFunction;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)

    var ts, te, act int
    _, _, _ = ts, te, act

    %%{

        varStart = '$'+;
        iden = [a-zA-Z_\x7f-\xff][a-zA-Z0-9_\x7f-\xff]*;

        body = '{' (any+ -- '}') '}';
        attr = '[' (any+ -- ']') ']';

        expr = space* body;

        comment1 = '/*' (any* -- '*/') '*/';
        comment2 = ('//'|'#') (any* -- '\n') '\n';
        comments = comment1 | comment2;

        misc = space | attr | body | comments;

        funcParam = '(' (any+ -- ')') ')';

        funcHead = varStart (iden | expr) misc* funcParam;

        main := |*
            funcHead => {
                return true
            };

            any;
        *|;

        write init;
        write exec;
    }%%
        return false
}