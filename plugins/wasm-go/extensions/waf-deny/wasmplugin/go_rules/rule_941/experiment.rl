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
    "param", "iframe", "frame", "base", "body", "bindings", "image", "img", "video", "importimplementation",
}

var riskAttrs = []string{
	"background", "formaction", "lowsrc", "ping", "src", "style",
}

func isRiskAttr(attr string) bool {
    if slices.Contains(riskAttrs, attr) {
        return true
    }

    if strings.HasPrefix(attr, "on") {
        return true
    }

    return false
}

func matchExp(data []byte) bool {
%% machine xss;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof
    pbt, pba := 0, 0
    maybeTag := false
    var tagNameBuilder strings.Builder
    var tagName string

    maybeAttr := false
    var attrName string

    %%{

        action markTag {
            if maybeTag {
                pbt = p
            }
        }

        action startTag {
            maybeTag = true
            tagNameBuilder.Reset()
        }

        action endTag {
            if !maybeTag {
                maybeTag = false
                tagNameBuilder.Reset()
            }
        }

        action resetTag {
            maybeTag = true
            tagNameBuilder.Reset()
        }

        action addTagNameChar {
            if maybeTag {
                tagNameBuilder.WriteString(string(data[pbt:p]))
                tagName = tagNameBuilder.String()
                fmt.Printf("tag name: %s \n", tagName)

                if slices.Contains(riskTags, tagName) {
                    fmt.Println("Matched")
                    return true
                }
            }
        }

        action checkMatchTag {
            if maybeTag {
                tagName = tagNameBuilder.String()
                fmt.Printf("tag name: %s \n", tagName)

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
            pba = p
        }

        action endAttr {
            if maybeAttr {
                attrName = string(data[pba:p])
                fmt.Printf("Attr name: %d %d %s \n", pba, p, attrName)
                maybeAttr = false
            }
        }

        action checkMatchAttr {
            if len(attrName) > 0 && isRiskAttr(attrName) {
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
        temp3 = [^\s\v\"'<>] ;
        temp4 = [^0-9>A-Z_a-z];
        tag_prefix = temp2* ( temp3* ':' %resetTag )? temp2*;
        alpha_like = ^(word_ele|':')* alpha >markTag;
        left_angle = '<' ;
        right_angle = '>' ;
        html_tag = any* left_angle %startTag tag_prefix (alpha_like %addTagNameChar)+ %/checkMatchTag right_angle >endTag ;

        quotes = '"' | '\'' ;
        attr_start = (left_angle word_ele any* (space | '/')) | (quotes (any* (space | '/'))?);
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
