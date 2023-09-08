package rule_941

func matchXSS(data []byte) bool {
%% machine xss;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof
    %%{

        action setMatched {
            return true
        }

        word_bound = zlen | (space | punct)*;
        word_ele = [_0-9A-Za-z];
        identifier = [_A-Za-z][_0-9A-Za-z]*;

        # 941110
        script_tag = any* '<script' any* '>' @setMatched;

        # 941130
        x_tag = word_bound ('xlink:href' | 'xhtml' | 'xmlns' | 'data:text/html' | 'formaction' | '@import' | ';base64')  %/setMatched word_bound @setMatched;
        pattern_tag = word_bound 'pattern' word_bound word_ele* '=' word_ele* %/setMatched word_bound @setMatched;
        entity_tag = word_bound /!ENTITY/i space+ ('%' space+)? identifier+ space+ (/SYSTEM/i | /PUBLIC/i) %/setMatched word_bound @setMatched;

        # 941140
        # [^:=]
        temp1 = any - (':' | '=');
        url_js_tag = lower+ '=' (temp1+ ':' any+ ';')* temp1+ ':url(javascript' %setMatched;

        # 941160
        # [^0-9<>A-Z_a-z]
        temp2 = [^0-9<>A-Z_a-z];
        temp3 = any - (space | '"' | '\'' | '<' | '>');
        temp4 = [^0-9>A-Z_a-z];

        tag_prefix_like = temp2* ( temp3 ':' )? temp2* ^word_ele*;

        s_like = 's' ^word_ele*;
        c_like = 'c' ^word_ele*;
        r_like = 'r' ^word_ele*;
        i_like = 'i' ^word_ele*;
        t_like = 't' ^word_ele*;
        y_like = 'y' ^word_ele*;
        l_like = 'l' ^word_ele*;
        e_like = 'e' ^word_ele*;
        f_like = 'f' ^word_ele*;
        o_like = 'o' ^word_ele*;
        m_like = 'm' ^word_ele*;
        a_like = 'a' ^word_ele*;
        v_like = 'v' ^word_ele*;
        n_like = 'n' ^word_ele*;
        b_like = 'b' ^word_ele*;
        j_like = 'j' ^word_ele*;
        d_like = 'd' ^word_ele*;
        g_like = 'g' ^word_ele*;
        q_like = 'q' ^word_ele*;
        u_like = 'u' ^word_ele*;
        p_like = 'p' ^word_ele*;

        script_like = s_like c_like r_like i_like 't';
        style_like = s_like t_like y_like l_like 'e';
        svg_like = s_like v_like 'g';
        set_like = s_like e_like t_like;
        form_like = f_like o_like r_like 'm';
        marquee_like = m_like a_like r_like q_like u_like e_like 'e';
        meta_like = m_like e_like t_like a_like;
        link_like = l_like i_like n_like 'k';
        object_like = o_like b_like j_like e_like c_like 't';
        embed_like = e_like m_like b_like e_like 'd';
        applet_like = a_like p_like p_like l_like e_like 't';
        audio_like = a_like u_like d_like i_like 'o';
        animate_like = a_like n_like i_like m_like a_like t_like 'e';
        param_like = p_like a_like r_like a_like 'm';
        iframe_like = i_like f_like r_like a_like m_like 'e';
        base_like = b_like a_like s_like 'e';
        body_like = b_like o_like d_like 'y';
        bindings_like = b_like i_like n_like d_like i_like n_like g_like 's';
        image_like = i_like m_like a_like? g_like e_like?;
        video_like = v_like i_like d_like e_like 'o';

        tag_name_like = script_like | style_like | svg_like | set_like | form_like | marquee_like | meta_like | link_like
            | object_like | embed_like | applet_like | audio_like | animate_like | param_like | iframe_like | base_like
            | body_like | bindings_like | image_like | video_like;

        html_tag = '<' tag_prefix_like tag_name_like   ;

        quotes = '"' | '\'';
        attribute_start = ('<' word_ele any* (space | '/')) | (quotes (any* (space | '/'))?);
        attribute_name = 'background' | 'formaction' | 'lowsrc' | 'on' | 'abort' | 'activate' | 'adapteradded' | 'addtrack'
            | 'afterprint' | 'afterscriptexecute' | 'afterupdate' | 'altering' | 'animationend' | 'animationiteration'
            | 'animationstart' | 'antennastatechange' | 'appcommand' | 'audioend' | 'audioprocess' | 'audiostart'
            | 'beforeactivate' | 'beforedeactivate' | 'beforescriptexecute' | 'beforecopy' | 'beforecut' | 'beforeeditfocus'
            | 'beforepaste' | 'beforeprint' | 'beforeunload' | 'beforeupdate' | 'begin' | 'beginevent' | 'blocked' | 'blur'
            | 'bounce' | 'boundary' | 'broadcast' | 'busy' | 'cached' | 'callschanged' | 'canplay' | 'canplaythrough'
            | 'cardstatechange' | 'cellchange' | 'cfstatechange' | 'change' | 'chargingchange' | 'chargingtimechange'
            | 'checking' | 'click' | 'close' | 'command' | 'commandupdate' | 'complete' | 'compositionend' | 'compositionstart'
            | 'compositionupdate' | 'connected' | 'connecting' | 'contextmenu' | 'controlselect' | 'copy' | 'cuechange'
            | 'cut' | 'dataavailable' | 'datachange' | 'dataerror' | 'datasetchanged' | 'datasetcomplete' | 'dblclick'
            | 'deactivate' | 'deliveryerror' | 'deliverysuccess' | 'devicefound' | 'devicelight' | 'devicemotion'
            | 'deviceorientation' | 'deviceproximity' | 'dialing' | 'disabled' | 'dischargingtimechange' | 'disconnected'
            | 'disconnecting' | 'domactivate' | 'domattrmodified' | 'domcharacterdatamodified' | 'domsubtreemodified'
            | 'domfocusin' | 'domfocusout' | 'dommousescroll' | 'domnodeinserted' | 'domnodeinsertedintodocument'
            | 'domnoderemoved' | 'domnoderemovedfromdocument' | 'downloading' | 'dragdrop' | 'dragend' | 'dragenter'
            | 'dragexit' | 'draggesture' | 'dragleave' | 'dragover' | 'dragstart' | 'drop' | 'durationchange' | 'emptied'
            | 'enabled' | 'end' | 'ended' | 'endevent' | 'enter' | 'error' | 'errorupdate' | 'exit' | 'failed' | 'finish'
            | 'filterchange' | 'focus' | 'focusin' | 'focusout' | 'formchange' | 'forminput' | 'gamepadaxismove'
            | 'gamepadbuttondown' | 'gamepadbuttonup' | 'gamepadconnected' | 'gamepaddisconnected' | 'get' | 'hashchange'
            | 'headphoneschange' | 'held' | 'help' | 'holding' | 'icccardlockerror' | 'iccinfochange' | 'incoming' | 'input'
            | 'invalid' | 'keydown' | 'keypress' | 'keyup' | 'levelchange' | 'load' | 'loadeddata' | 'loadedmetadata'
            | 'loadend' | 'loadstart' | 'losecapture' | 'ly' |


        main := script_tag | x_tag | pattern_tag | entity_tag | url_js_tag;
        write init;
        write exec;
    }%%
        return false
}
