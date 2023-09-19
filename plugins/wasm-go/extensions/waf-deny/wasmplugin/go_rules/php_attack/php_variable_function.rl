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
        roundBracketContent = '(' (any+ -- ')') ')';
        mayEmptyRoundBracketContent = '(' (any* -- ')') ')';

        funcHead = varStart (iden | expr) misc* roundBracketContent;

        quotes = '"' | "'";
        quoteContent = quotes (alnum | '_')+ quotes;

        riskFun1 = roundBracketContent roundBracketContent;
        riskFun2 = roundBracketContent quoteContent roundBracketContent;
        riskFun3 = '[' digit+ ']' roundBracketContent;
        riskFun4 = '{' digit+ '}' roundBracketContent;

        temp = '(' | ')' | ',' | '.' | ';' | '/' | '\\';
        riskFun5 = '$' ^temp+ roundBracketContent;
        quoteContentExt = quotes (alnum | '_' | '\\')+ quotes;
        riskFun6 = quoteContentExt roundBracketContent;

        stringBracketed = '(' (any* -- ')') 'string' (any* -- ')') ')';
        temp2 = alnum | quotes | '.' | '{' | '}' | '[' | ']' | space;
        riskFun7 = stringBracketed temp2+ mayEmptyRoundBracketContent;

        riskFuncs = riskFun1 | riskFun2 | riskFun3 | riskFun4 | riskFun5 | riskFun6 | riskFun7;

        main := |*
            funcHead => {
                return true
            };

            riskFuncs => {
                return true
            };

            any;
        *|;

        write init;
        write exec;
    }%%
        return false
}