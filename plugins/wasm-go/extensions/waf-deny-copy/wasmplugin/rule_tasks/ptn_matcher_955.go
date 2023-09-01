package rule_tasks

import (
	ahocorasick "github.com/petar-dambovaliev/aho-corasick"
	"github.com/wasilibs/go-re2"
)

const (
	Ptn955110 = `(<title>r57 Shell Version [0-9.]+</title>|<title>r57 shell</title>)`
	Ptn955120 = `^<html><head><meta http-equiv='Content-Type' content='text/html; charset=Windows-1251'><title>.*? - WSO [0-9.]+</title>`
	Ptn955130 = `B4TM4N SH3LL</title>.*<meta name='author' content='k4mpr3t'/>`
	Ptn955140 = `<title>Mini Shell</title>.*Developed By LameHacker`
	Ptn955150 = `<title>\.:: .* ~ Ashiyane V [0-9.]+ ::\.</title>`
	Ptn955160 = `<title>Symlink_Sa [0-9.]+</title>`
	Ptn955170 = `<title>CasuS [0-9.]+ by MafiABoY</title>`
	Ptn955180 = `^<html>\r\n<head>\r\n<title>GRP WebShell [0-9.]+ `
	Ptn955190 = `<small>NGHshell [0-9.]+ by Cr4sh</body></html>\n$`
	Ptn955200 = `<title>SimAttacker - (?:Version|Vrsion) : [0-9.]+ - `
	Ptn955210 = `^<!DOCTYPE html>\n<html>\n<!-- By Artyum .*<title>Web Shell</title>`
	Ptn955220 = `<title>lama's'hell v. [0-9.]+</title>`
	Ptn955230 = `^ *<html>\n[ ]+<head>\n[ ]+<title>lostDC - `
	Ptn955240 = `^<title>PHP Web Shell</title>\r\n<html>\r\n<body>\r\n    <!-- Replaces command with Base64-encoded Data -->`
	Ptn955250 = "^<html>\n<head>\n<div align=\"left\"><font size=\"1\">Input command :</font></div>\n<form name=\"cmd\" method=\"POST\" enctype=\"multipart/form-data\">"
	Ptn955260 = `^<html>\n<head>\n<title>Ru24PostWebShell - `
	Ptn955270 = `<title>s72 Shell v[0-9.]+ Codinf by Cr@zy_King</title>`
	Ptn955280 = "^<html>\\r\\n<head>\\r\\n<meta http-equiv=\\\"Content-Type\\\" content=\\\"text/html; charset=gb2312\\\">\\r\\n<title>PhpSpy Ver [0-9]+</title>"
	Ptn955290 = `^ <html>\n\n<head>\n\n<title>g00nshell v[0-9.]+ `
	Ptn955310 = `^<html>\n      <head>\n             <title>azrail [0-9.]+ by C-W-M</title>`
	Ptn955320 = `>SmEvK_PaThAn Shell v[0-9]+ coded by <a href=`
	Ptn955330 = `^<html>\n<title>.*? ~ Shell I</title>\n<head>\n<style>`
	Ptn955340 = `^ <html><head><title>:: b374k m1n1 [0-9.]+ ::</title>`
)

var Re955110, Re955120, Re955130, Re955140, Re955150, Re955160, Re955170, Re955180, Re955190, Re955200, Re955210,
	Re955220, Re955230, Re955240, Re955250, Re955260, Re955270, Re955280, Re955290, Re955310, Re955320, Re955330,
	Re955340 *re2.Regexp

var WEB_SHELLS_PHP = []string{
	"<title>=[ 1n73ct10n privat shell ]=</title>",
	">Ajax/PHP Command Shell<",
	".:: :[ AK-74 Security Team Web-shell ]: ::.",
	"~ ALFA TEaM Shell -",
	"<title>--==[[ Andela Yuwono Priv8 Shell ]]==--</title>",
	"<title>Ani-Shell | India</title>",
	"<input type='submit' value='file' /></form>AnonymousFox",
	"- Antichat Shell</title>",
	"Ayyildiz Tim  | AYT",
	"<link rel='SHORTCUT ICON' href='data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAyJpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADw/eHBhY2tldCBiZWdpbj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDUuMy1jMDExIDY2LjE0NTY2MSwgMjAxMi8wMi8wNi0xNDo1NjoyNyAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvIiB4bWxuczp4bXBNTT0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL21tLyIgeG1sbnM6c3RSZWY9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9zVHlwZS9SZXNvdXJjZVJlZiMiIHhtcDpDcmVhdG9yVG9vbD0iQWRvYmUgUGhvdG9zaG9wIENTNiAoV2luZG93cykiIHhtcE1NOkluc3RhbmNlSUQ9InhtcC5paWQ6MkRFNDY2MDQ4MDgyMTFFM0FDRDdBN0MzOTAxNzZFQUYiIHhtcE1NOkRvY3VtZW50SUQ9InhtcC5kaWQ6MkRFNDY2MDU4MDgyMTFFM0FDRDdBN0MzOTAxNzZFQUYiPiA8eG1wTU06RGVyaXZlZEZyb20gc3RSZWY6aW5zdGFuY2VJRD0ieG1wLmlpZDoyREU0NjYwMjgwODIxMUUzQUNEN0E3QzM5MDE3NkVBRiIgc3RSZWY6ZG9jdW1lbnRJRD0ieG1wLmRpZDoyREU0NjYwMzgwODIxMUUzQUNEN0E3QzM5MDE3NkVBRiIvPiA8L3JkZjpEZXNjcmlwdGlvbj4gPC9yZGY6UkRGPiA8L3g6eG1wbWV0YT4gPD94cGFja2V0IGVuZD0iciI/Pu6UWJYAAAKySURBVHjafFNdSJNhFH7e/fhDkrm2i03QphsxhYSgMIUgIeiiK6/SCAKTKNlFoEtBRfEvXYhM+0GQMtMUL7qSgqS0QCNKTDS6cJWGi6n577Zv3/e+b+934ZgxPfDBd3jP85xznnOOzufz4SCr7R7knKOg4eaVd9WPBgsZY/3NZcWJ0TGaaKeuZzgz2ueMgFF+p6WnL0OAjzMK+f8k+wg4xXxN91D5ns8ok8CRH5S2GogS8HBKk1xud+uBBIwpm5zyRvW/+sHAJuM8nsrMIElHi0/aHAmFl/OI2WRyOevrK/YwJFoD0ecFkfWthpDNRH1Cct4ZOzRaglX/DsY+TcNqTUd2phEjo1OiWg5KKUhJTbua6XTT7SKvSlLpGWB6DUjuWQeW/m4iJIWho8DvBT+2tgOwpZsxM/tm/sn9Trsar2OMq6rOV3X19wncJUNSEsnKSsWifx0BKYTgdhDxiENBfjZCuxJejX0W4frZiAZNZUVxVKYfmcyuKTI15ZxKw4IA74aCCIiMeqZDptWIuV8+hAkXOlFo9eaLNyrvOfdp4Gp/FjKlpMSbLMlY2dhCaCcEnUJgt5sF4QqkkIKsDAtGXn9QSThlMmFCg8gUmELpkXg99FoNwgEJ2jBBWpoBP/8sC7AMi/EY/EvLUBQJCpOMT921hDG5JkIglPd8/7EIFpShCQMnrAYsrW0gLERUwTNfv2FyaloddWmvu25NxTzvaG6MELRVXK/SgL8fHZ9AjsMCKUzFqBhSjQZAkrC6viqyy+ILdxU775bH3APVblW3j3POzuc4bGIHNPgyM4dAcFdtslT07OWcvhRVJIvVtg0/9nhJrGMqqWzpFb1eFYuiVfdbACcGOlvzYx0cOewaVStyuiY5U3JFVbahhx3eQ48plr3obDtHqSxTRZ6K9f5PgAEAm/hvADIkGOQAAAAASUVORK5CYII='>",
	"<title>BloodSecurity Hackers Shell</title>",
	"<font color='red' size='6px' face='Fredericka the Great'> Bypass Attack Shell </font>",
	"title='.::[c0derz shell]::.'>",
	"<font face=Webdings size=6><b>!</b></font>",
	"<title>Con7ext Shell V.2</title>",
	"<font face=\"Wingdings 3\" size=\"5\">y</font><b>Crystal shell v.",
	"~ CWShell ~</font></a>",
	"&dir&pic=o.b height= width=>",
	"<b>[ Defacing Tool Pro v",
	"<title>Dive Shell - Emperor Hacking Team</title>",
	"<script>document.getElementById(\"cmd\").focus();</script>",
	"color=DeepSkyBlue   size=6>    <p align=\"center\" class=\"style4\">FaTaLSheLL v",
	"<title>G-Security Webshell</title>",
	"<title>h4ntu shell [powered by tsoi]</title>",
	"<H1><center>-=[+] IDBTEAM SHELLS",
	"<title>IndoXploit</title>",
	"<KAdot Universal Shell>     |",
	">LIFKA SHELL</span></big></big></big></a>",
	"<title>Loader'z WEB shell</title>",
	"b>--[ x2300 Locus7Shell v.",
	"<title>Lolipop.php - Edited By KingDefacer -",
	"<link rel=\"icon\" href=\"//0x5a455553.github.io/MARIJUANA/icon.png\" />",
	"<title> Matamu Mat </title>",
	"<b>MyShell</b> &copy;2001 Digitart Producciones</a>",
	"<h1>.:NCC:. Shell v",
	"<font size=3>PHPShell by Macker - Version",
	"PHPShell by MAX666, Private Exploit, For Server Hacking",
	"<form action=\"\" METHOD=\"GET\" >Execute Shell Command (safe mode is off): <input type=\"text\" name=\"c\"><input type=\"submit\" value=\"Go\"></form>",
	"<p align=\"center\"><font face=\"Verdana\" size=\"2\">Rootshell v",
	"<font color=lime>./rusuh</font>",
	"<font color=\"navy\"><strong><center><h1>Watch Your system Shany was here.</h1></center><center><h1>Linux Shells</h1></center><hr><hr>",
	"<!-- Simple PHP backdoor by DK",
	"<title>SimShell - Simorgh Security MGZ</title>",
	"<title>:: AventGrup ::.. - Sincap",
	"<title>Small Shell - Edited By KingDefacer</title>",
	"<title>small web shell by zaco",
	"<title>SoldiersofAllah Private Shell |",
	"<title>Sosyete Safe Mode Bypass Shell -",
	"&nbsp;&nbsp;STNC&nbsp;WebShell&nbsp;",
	"<font face=\"Wingdings 3\" size=\"5\">y</font><b>StresBypass<span lang=\"en-us\">v",
	"<title>SyRiAn Sh3ll ~",
	"<head><title>Wardom | Ne Mutlu T",
	"<hr>to browse go to http://?d=[directory here]",
	"<font color=\"red\">USTADCAGE_48</font> <font color=\"dodgerblue\">FILE MANAGER</font>",
	"<title>WebRoot Hack Tools</title>",
	"<div align=\"center\"><span class=\"style6\">By BLaSTER</span><br />",
	"<title>-:[GreenwooD]:- WinX Shell</title>",
	"<sup><a href=\"<title>Yourman.sh Mini Shell</title>",
	"</div><center><br />Zerion Mini Shell <font color=",
	"<title>0byt3m1n1-V2</title>",
	"<title>ZEROSHELL | ZEROSTORE</title>",
	"<input type=submit name=find value='find writeable'>",
}
var Rule955100Matcher ahocorasick.AhoCorasick

func init() {
	Rule955100Matcher = AHO_CORASICK_BUILDER.Build(WEB_SHELLS_PHP)

	Re955110, _ = re2.Compile(Ptn955110)
	Re955120, _ = re2.Compile(Ptn955120)
	Re955130, _ = re2.Compile(Ptn955130)
	Re955140, _ = re2.Compile(Ptn955140)
	Re955150, _ = re2.Compile(Ptn955150)
	Re955160, _ = re2.Compile(Ptn955160)
	Re955170, _ = re2.Compile(Ptn955170)
	Re955180, _ = re2.Compile(Ptn955180)
	Re955190, _ = re2.Compile(Ptn955190)
	Re955200, _ = re2.Compile(Ptn955200)
	Re955210, _ = re2.Compile(Ptn955210)
	Re955220, _ = re2.Compile(Ptn955220)
	Re955230, _ = re2.Compile(Ptn955230)
	Re955240, _ = re2.Compile(Ptn955240)
	Re955250, _ = re2.Compile(Ptn955250)
	Re955260, _ = re2.Compile(Ptn955260)
	Re955270, _ = re2.Compile(Ptn955270)
	Re955280, _ = re2.Compile(Ptn955280)
	Re955290, _ = re2.Compile(Ptn955290)
	Re955310, _ = re2.Compile(Ptn955310)
	Re955320, _ = re2.Compile(Ptn955320)
	Re955330, _ = re2.Compile(Ptn955330)
	Re955340, _ = re2.Compile(Ptn955340)
}
