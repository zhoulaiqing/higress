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
    return slices.Contains(riskFunctions, name)
}

func matchPhpFunctions(data []byte) bool {
%% machine phpFunctions;
%% write data;
    cs, p, pe, eof := 0, 0, len(data), len(data)

    var ts, te, act int
    _, _, _ = ts, te, act

    pb := 0
    var funcName string

    %%{

        action mark {
            pb = p
        }

        action setFuncName {
            fmt.Printf("php func name: %s \n", string(data[pb:p]))
            funcName = string(data[pb:p])
        }

        iden = [a-z_];

        quotes = '"' | "'";
        comment1 = '/*' (any* -- '*/') '*/';
        comment2 = ('//'|'#') (any* -- '\n') '\n';
        comments = comment1 | comment2;

        fun = (any* -- iden) '('? quotes* iden+ >mark %setFuncName comments* ')'? quotes* space* '(' (any* -- ')') ')';

        main := |*
            fun => {
                if checkFunc(funcName) {
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
