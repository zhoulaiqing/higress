package protocol_attack

func matchLdapInjection(data []byte) bool {
%% machine ldapInjection;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof


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

        main := rpart lpart %/setMatched any* @setMatched;


        write init;
        write exec;
    }%%
        return false
}

