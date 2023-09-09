package rule_941

import (
    "fmt"
    "strings"
    "golang.org/x/exp/slices"
)

type Machine struct {

}

type XssChecker struct {

}

var riskTags = []string {
    "script", "style", "svg", "set", "form", "marquee", "meta", "link", "object", "embed", "applet", "audio", "animate",
    "param", "iframe", "frame", "base", "body", "bindings", "image", "img", "video",
}



func matchExp(data []byte) bool {
%% machine xss;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof
    pb := 0
    maybeTag := false
    var tagNameBuilder strings.Builder
    var tagName string

    maybeAttr := false
    var attrName string

    %%{

        action markTag {
            if maybeTag {
                pb = p
            }
        }

        action startTag {
            maybeTag = true
            tagNameBuilder.Reset()
        }

        action addTagNameChar {
            if maybeTag {
                tagNameBuilder.WriteString(string(data[pb:p]))
            }
        }

        action checkMatchTag {
            if maybeTag {
                tagName = tagNameBuilder.String()
                fmt.Println(tagName)

                if slices.Contains(riskTags, tagName) {
                    fmt.Println("Matched")
                    return true
                }

                tagNameBuilder.Reset()
                maybeTag = false
            }
        }

        action startAttr {
            maybeAttr = true
            attrName = ""
            pb = p
        }

        action endAttr {
            if maybeAttr {
                attrName = string(data[pb:p])
                fmt.Printf("Attr name: %d %d %s \n", pb, p, attrName)
                maybeAttr = false
            }
        }

        action checkMatchAttr {
            if len(attrName) > 0 {
                return true
            }
        }

        action setMatched {
            return true
        }

        word_bound = zlen | (space | punct)*;
        word_ele = [_0-9A-Za-z];
        identifier = [_A-Za-z][_0-9A-Za-z]*;

        temp2 = [^0-9<>A-Z_a-z];
        temp3 = any - (space | '"' | '\'' | '<' | '>');
        temp4 = [^0-9>A-Z_a-z];
        tag_prefix = temp2* ( temp3 ':' )? temp2*;
        alpha_like = ^word_ele* alpha >markTag;
        html_tag = any* '<' %startTag tag_prefix (alpha_like %addTagNameChar)+ %/checkMatchTag ^word_ele @checkMatchTag ;

        quotes = '"' | '\'';
        attr_start = ('<' word_ele any* (space | '/')) | (quotes (any* (space | '/'))?);
        # attr_name = 'background' | 'formaction' | 'lowsrc' | 'on' alpha{3,30} | 'ping' | 'src' | 'style' ;
        attr_name = alpha{3,30} ;
        # try with space
        attr_sep = space;

        html_attr = any* attr_start %startAttr attr_name %endAttr attr_sep* '=' $checkMatchAttr;

        main := html_tag | html_attr;
        write init;
        write exec;
    }%%
        return false
}
