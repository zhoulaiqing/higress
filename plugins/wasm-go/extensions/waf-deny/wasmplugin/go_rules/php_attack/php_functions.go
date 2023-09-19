
//line php_functions.rl:1
package php_attack

import (
    "fmt"
    "golang.org/x/exp/slices"
)

var riskFunctions = []string {
"array_diff_uassoc",
"array_diff_ukey",
"array_filter",
"array_intersect_uassoc",
"array_intersect_ukey",
"array_map",
"array_reduce",
"array_udiff",
"array_udiff_assoc",
"array_udiff_uassoc",
"array_uintersect",
"array_uintersect_assoc",
"array_uintersect_uassoc",
"assert",
"assert_options",
"base64_encode",
"bson_decode",
"bson_encode",
"bzopen",
"chr",
"convert_uuencode",
"create_function",
"curl_exec",
"curl_file_create",
"curl_init",
"debug_backtrace",
"error_reporting",
"escapeshellarg",
"escapeshellcmd",
"eval",
"exec",
"exif_imagetype",
"exif_read_data",
"exif_tagname",
"exif_thumbnail",
"file",
"file_exists",
"fileatime",
"filectime",
"filegroup",
"fileinode",
"filemtime",
"fileperms",
"finfo_open",
"fopen",
"fputs",
"ftp_connect",
"ftp_get",
"ftp_nb_get",
"ftp_nb_put",
"ftp_put",
"function_exists",
"fwrite",
"get_cfg_var",
"get_current_user",
"get_meta_tags",
"getcwd",
"getenv",
"getimagesize",
"getlastmod",
"getmygid",
"getmyinode",
"getmypid",
"getmyuid",
"glob",
"gzcompress",
"gzdeflate",
"gzencode",
"gzfile",
"gzopen",
"gzread",
"gzwrite",
"hash_file",
"hash_hmac_file",
"hash_update_file",
"header_register_callback",
"hex2bin",
"highlight_file",
"html_entity_decode",
"htmlentities",
"htmlspecialchars",
"htmlspecialchars_decode",
"image2wbmp",
"imagecreatefromgif",
"imagecreatefromjpeg",
"imagecreatefrompng",
"imagecreatefromwbmp",
"imagecreatefromxbm",
"imagecreatefromxpm",
"imagegd",
"imagegd2",
"imagegif",
"imagejpeg",
"imagepng",
"imagewbmp",
"imagexbm",
"ini_get",
"ini_get_all",
"ini_set",
"iptcembed",
"is_dir",
"is_executable",
"is_file",
"is_readable",
"is_writable",
"is_writeable",
"iterator_apply",
"json_decode",
"json_encode",
"mb_ereg",
"mb_ereg_match",
"mb_ereg_replace",
"mb_ereg_replace_callback",
"mb_eregi",
"mb_eregi_replace",
"mb_parse_str",
"md5_file",
"method_exists",
"mkdir",
"move_uploaded_file",
"mysql_query",
"ob_clean",
"ob_end_clean",
"ob_end_flush",
"ob_flush",
"ob_get_clean",
"ob_get_contents",
"ob_get_flush",
"ob_start",
"odbc_connect",
"odbc_exec",
"odbc_execute",
"odbc_result",
"odbc_result_all",
"opendir",
"parse_ini_file",
"parse_str",
"passthru",
"pg_connect",
"pg_execute",
"pg_prepare",
"pg_query",
"php_strip_whitespace",
"php_uname",
"phpinfo",
"phpversion",
"popen",
"posix_getegid",
"posix_geteuid",
"posix_getgid",
"posix_getlogin",
"posix_getpwnam",
"posix_kill",
"posix_mkfifo",
"posix_mknod",
"posix_ttyname",
"preg_match",
"preg_match_all",
"preg_replace",
"preg_replace_callback",
"preg_replace_callback_array",
"preg_split",
"print_r",
"proc_close",
"proc_get_status",
"proc_nice",
"proc_open",
"proc_terminate",
"putenv",
"rawurldecode",
"rawurlencode",
"read_exif_data",
"readdir",
"readfile",
"readgzfile",
"register_shutdown_function",
"register_tick_function",
"rename_function",
"runkit_constant_add",
"runkit_constant_redefine",
"runkit_function_add",
"runkit_function_copy",
"runkit_function_redefine",
"runkit_function_rename",
"runkit_method_add",
"runkit_method_copy",
"runkit_method_redefine",
"runkit_method_rename",
"session_set_save_handler",
"session_start",
"set_error_handler",
"set_exception_handler",
"set_include_path",
"set_magic_quotes_runtime",
"setdefaultstub",
"sha1_file",
"show_source",
"simplexml_load_file",
"simplexml_load_string",
"socket_connect",
"socket_create",
"spl_autoload_register",
"sqlite_array_query",
"sqlite_create_aggregate",
"sqlite_create_function",
"sqlite_exec",
"sqlite_open",
"sqlite_popen",
"sqlite_query",
"sqlite_single_query",
"sqlite_unbuffered_query",
"stream_context_create",
"stream_socket_client",
"stripcslashes",
"stripslashes",
"strrev",
"system",
"tmpfile",
"uasort",
"uksort",
"unpack",
"unserialize",
"urldecode",
"urlencode",
"usort",
"var_dump",
}

func checkFunc(name string) bool {
    fmt.Printf("Check php func name %s \n", name)
    return slices.Contains(riskFunctions, name)
}

func matchPhpFunctions(data []byte) bool {

//line php_functions.rl:244

//line php_functions.go:249
const phpFunctions_start int = 10
const phpFunctions_first_final int = 10
const phpFunctions_error int = -1

const phpFunctions_en_main int = 10


//line php_functions.rl:245
    cs, p, pe, eof := 0, 0, len(data), len(data)

    var ts, te, act int
    _, _, _ = ts, te, act

    pb := 0
    var funcName string

    
//line php_functions.go:267
	{
	cs = phpFunctions_start
	ts = 0
	te = 0
	act = 0
	}

//line php_functions.go:275
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 0:
		goto st_case_0
	case 1:
		goto st_case_1
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 12:
		goto st_case_12
	}
	goto st_out
tr0:
//line php_functions.rl:280
p = (te) - 1

	goto st10
tr11:
//line php_functions.rl:274
te = p+1
{
                if checkFunc(funcName) {
                    return true
                }
            }
	goto st10
tr18:
//line php_functions.rl:280
te = p+1

	goto st10
tr21:
//line php_functions.rl:280
te = p
p--

	goto st10
	st10:
//line NONE:1
ts = 0

		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line NONE:1
ts = p

//line php_functions.go:345
		switch data[p] {
		case 34:
			goto tr19
		case 95:
			goto tr20
		}
		switch {
		case data[p] > 40:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr20
			}
		case data[p] >= 39:
			goto tr19
		}
		goto tr18
tr19:
//line NONE:1
te = p+1

	goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
//line php_functions.go:371
		switch data[p] {
		case 34:
			goto st0
		case 39:
			goto st0
		case 95:
			goto tr2
		}
		if 97 <= data[p] && data[p] <= 122 {
			goto tr2
		}
		goto tr21
	st0:
		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
		switch data[p] {
		case 34:
			goto st0
		case 39:
			goto st0
		case 95:
			goto tr2
		}
		if 97 <= data[p] && data[p] <= 122 {
			goto tr2
		}
		goto tr0
tr2:
//line php_functions.rl:255

            pb = p
        
	goto st1
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
//line php_functions.go:412
		switch data[p] {
		case 32:
			goto tr3
		case 34:
			goto tr4
		case 35:
			goto tr5
		case 40:
			goto tr6
		case 47:
			goto tr7
		case 95:
			goto st1
		}
		switch {
		case data[p] < 39:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr3
			}
		case data[p] > 41:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto tr4
		}
		goto tr0
tr3:
//line php_functions.rl:259

            funcName = string(data[pb:p])
        
	goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line php_functions.go:451
		switch data[p] {
		case 32:
			goto st2
		case 40:
			goto st3
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st2
		}
		goto tr0
tr6:
//line php_functions.rl:259

            funcName = string(data[pb:p])
        
	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line php_functions.go:473
		if data[p] == 41 {
			goto tr11
		}
		goto st3
tr4:
//line php_functions.rl:259

            funcName = string(data[pb:p])
        
	goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line php_functions.go:489
		switch data[p] {
		case 32:
			goto st2
		case 34:
			goto st4
		case 39:
			goto st4
		case 40:
			goto st3
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st2
		}
		goto tr0
tr5:
//line php_functions.rl:259

            funcName = string(data[pb:p])
        
	goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line php_functions.go:515
		if data[p] == 10 {
			goto st6
		}
		goto st5
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch data[p] {
		case 32:
			goto st2
		case 34:
			goto st4
		case 35:
			goto st5
		case 40:
			goto st3
		case 47:
			goto st7
		}
		switch {
		case data[p] > 13:
			if 39 <= data[p] && data[p] <= 41 {
				goto st4
			}
		case data[p] >= 9:
			goto st2
		}
		goto tr0
tr7:
//line php_functions.rl:259

            funcName = string(data[pb:p])
        
	goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line php_functions.go:557
		switch data[p] {
		case 42:
			goto st8
		case 47:
			goto st5
		}
		goto tr0
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		if data[p] == 42 {
			goto st9
		}
		goto st8
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		switch data[p] {
		case 42:
			goto st9
		case 47:
			goto st6
		}
		goto st8
tr20:
//line NONE:1
te = p+1

//line php_functions.rl:255

            pb = p
        
	goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line php_functions.go:600
		switch data[p] {
		case 32:
			goto tr3
		case 34:
			goto tr4
		case 35:
			goto tr5
		case 40:
			goto tr6
		case 47:
			goto tr7
		case 95:
			goto st1
		}
		switch {
		case data[p] < 39:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr3
			}
		case data[p] > 41:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto tr4
		}
		goto tr21
	st_out:
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof0: cs = 0; goto _test_eof
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 11:
			goto tr21
		case 0:
			goto tr0
		case 1:
			goto tr0
		case 2:
			goto tr0
		case 3:
			goto tr0
		case 4:
			goto tr0
		case 5:
			goto tr0
		case 6:
			goto tr0
		case 7:
			goto tr0
		case 8:
			goto tr0
		case 9:
			goto tr0
		case 12:
			goto tr21
		}
	}

	}

//line php_functions.rl:285

        return false
}
