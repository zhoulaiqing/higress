package protocol_attack

func matchLdapInjection(data []byte) bool {
%% machine ldapInjection;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof


    var ts, te, act int
    _, _, _ = ts, te, act

    startWithRightPart := false

    %%{

        action setMatched {
            return true
        }

        kwd = ':' | '(' | ')' | '&' | '|' | '!' | '<' | '>' | '~';

        simple = ('>' | '<' | '~')? '=' space*;
        logic = '&' | '|' | '!' space*;
        attr = ^(kwd | ',' | '=')+;

        rpart = ^kwd* ')' (space* ')')?;

        lpart =  logic? '(' logic? (attr simple ^kwd*)?;

        main := |*
            lpart => {
                if startWithRightPart {
                    return true
                } else {
                    return false
                }
            };

            rpart => {
                if !startWithRightPart {
                    startWithRightPart = true
                }
            };

            any;

        *|;


        write init;
        write exec;
    }%%
        return false
}

