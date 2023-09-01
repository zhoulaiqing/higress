package rule_tasks

import (
	ahocorasick "github.com/petar-dambovaliev/aho-corasick"
	"github.com/wasilibs/go-re2"
)

const (
	Ptn951110 = `(?i:JET Database Engine|Access Database Engine|\[Microsoft\]\[ODBC Microsoft Access Driver\])`
	Ptn951120 = `(?i:ORA-[0-9][0-9][0-9][0-9]|java\.sql\.SQLException|Oracle error|Oracle.*Driver|Warning.*oci_.*|Warning.*ora_.*)`
	Ptn951130 = `(?i:DB2 SQL error:|\[IBM\]\[CLI Driver\]\[DB2/6000\]|CLI Driver.*DB2|DB2 SQL error|db2_\w+\()`
	Ptn951140 = `(?i:\[DM_QUERY_E_SYNTAX\]|has occurred in the vicinity of:)`
	Ptn951150 = `(?i)Dynamic SQL Error`
	Ptn951160 = `(?i)Exception (?:condition )?\d+\. Transaction rollback\.`
	Ptn951170 = `(?i)org\.hsqldb\.jdbc`
	Ptn951180 = `(?i:An illegal character has been found in the statement|com\.informix\.jdbc|Exception.*Informix)`
	Ptn951190 = `(?i:Warning.*ingres_|Ingres SQLSTATE|Ingres\W.*Driver)`
	Ptn951200 = `(?i:<b>Warning</b>: ibase_|Unexpected end of command in statement)`
	Ptn951210 = `(?i:SQL error.*POS[0-9]+.*|Warning.*maxdb.*)`
	Ptn951220 = `(?i)(?:System\.Data\.OleDb\.OleDbException|\[Microsoft\]\[ODBC SQL Server Driver\]|\[Macromedia\]\[SQLServer JDBC Driver\]|\[SqlException|System\.Data\.SqlClient\.SqlException|Unclosed quotation mark after the character string|'80040e14'|mssql_query\(\)|Microsoft OLE DB Provider for ODBC Drivers|Microsoft OLE DB Provider for SQL Server|Incorrect syntax near|Sintaxis incorrecta cerca de|Syntax error in string in query expression|Procedure or function .* expects parameter|Unclosed quotation mark before the character string|Syntax error .* in query expression|Data type mismatch in criteria expression\.|ADODB\.Field \(0x800A0BCD\)|the used select statements have different number of columns|OLE DB.*SQL Server|Warning.*mssql_.*|Driver.*SQL[ _-]*Server|SQL Server.*Driver|SQL Server.*[0-9a-fA-F]{8}|Exception.*\WSystem\.Data\.SqlClient\.)`
	Ptn951230 = `(?i)(?:supplied argument is not a valid |SQL syntax.*)MySQL|Column count doesn't match(?: value count at row)?|mysql_fetch_array\(\)|on MySQL result index|You have an error in your SQL syntax(?:;| near)|MyS(?:QL server version for the right syntax to use|qlClient\.)|\[MySQL\]\[ODBC|(?:Table '[^']+' doesn't exis|valid MySQL resul)t|Warning.{1,10}mysql_(?:[\(-\)_a-z]{1,26})?|ERROR [0-9]{4} \([0-9a-z]{5}\):`
	Ptn951240 = `(?i)P(?:ostgreSQL(?: query failed:|.{1,20}ERROR)|G::[a-z]*Error)|pg_(?:query|exec)\(\) \[:|Warning.{1,20}\bpg_.*|valid PostgreSQL result|Npgsql\.|Supplied argument is not a valid PostgreSQL .*? resource|Unable to connect to PostgreSQL server`
	Ptn951250 = `(?i)(?:Warning.*sqlite_.*|Warning.*SQLite3::|SQLite/JDBCDriver|SQLite\.Exception|System\.Data\.SQLite\.SQLiteException)`
	Ptn951260 = `(?i)(?:Sybase message:|Warning.{2,20}sybase|Sybase.*Server message.*)`
)

var Re951110, Re951120, Re951130, Re951140, Re951150, Re951160, Re951170, Re951180, Re951190, Re951200, Re951210,
	Re951220, Re951230, Re951240, Re951250, Re951260 *re2.Regexp

var SQL_ERRORS = []string{
	"MySqlClient.",
	"Server message",
	"SQL error",
	"Oracle error",
	"JET Database Engine",
	"Procedure or function",
	"SQLite.Exception",
	"[IBM][CLI Driver][DB2/6000]",
	"the used select statements have different number of columns",
	"org.postgresql.util.PSQLException",
	"Access Database Engine",
	"Incorrect syntax near",
	"Syntax error in string in query expression",
	"SQLiteException",
	"' doesn't exist",
	"CLI Driver",
	"on MySQL result index",
	"sybase",
	"com.informix.jdbc",
	"[MySQL][ODBC",
	"Error",
	"has occurred in the vicinity of:",
	"Sintaxis incorrecta cerca de",
	"MySQL server version for the right syntax to use",
	"com.mysql.jdbc.exceptions",
	"You have an error in your SQL syntax near",
	"You have an error in your SQL syntax;",
	"An illegal character has been found in the statement",
	"pg_query() [:",
	"supplied argument is not a valid MySQL",
	"mssql_query()",
	"mysql_fetch_array()",
	"Exception",
	"java.sql.SQLException",
	"Column count doesn't match value count at row",
	"Sybase message",
	"SQL Server",
	"PostgreSQL query failed:",
	"Dynamic SQL Error",
	"System.Data.SQLite.SQLiteException",
	"SQLite/JDBCDriver",
	"Unclosed quotation mark before the character string",
	"System.Data.SqlClient.",
	"Unclosed quotation mark after the character string",
	"System.Data.OleDb.OleDbException",
	"[DM_QUERY_E_SYNTAX]",
	"[SqlException",
	"Unexpected end of command in statement",
	"valid PostgreSQL result",
	"pg_exec() [:",
	"SQL Server",
	"[SQLITE_ERROR]",
	"Microsoft OLE DB Provider for ODBC Drivers",
	"PostgreSQL",
	"org.hsqldb.jdbc",
	"ADODB.Field (0x800A0BCD)",
	"SQL syntax",
	"Exception",
	"System.Data.SqlClient.SqlException",
	"Data type mismatch in criteria expression.",
	"Driver",
	"DB2 SQL error",
	"Sybase message:",
	"ORA-",
	"[Microsoft][ODBC SQL Server Driver]",
	"'80040e14'",
	"Microsoft OLE DB Provider for SQL Server",
	" in query expression",
	"Npgsql.",
	"valid MySQL result",
	"supplied argument is not a valid PostgreSQL result",
	"db2_",
	"Ingres SQLSTATE",
	"Column count doesn't match",
	"Warning",
	"[Microsoft][ODBC Microsoft Access Driver]",
	"[Macromedia][SQLServer JDBC Driver]",
	"<b>Warning</b>: ibase_",
	"Roadhouse.Cms.",
	"DB2 SQL error:",
	"SQLSTATE[",
}
var Rule951100Matcher ahocorasick.AhoCorasick

func init() {
	Rule951100Matcher = AHO_CORASICK_BUILDER.Build(SQL_ERRORS)

	Re951110, _ = re2.Compile(Ptn951110)
	Re951120, _ = re2.Compile(Ptn951120)
	Re951130, _ = re2.Compile(Ptn951130)
	Re951140, _ = re2.Compile(Ptn951140)
	Re951150, _ = re2.Compile(Ptn951150)
	Re951160, _ = re2.Compile(Ptn951160)
	Re951170, _ = re2.Compile(Ptn951170)
	Re951180, _ = re2.Compile(Ptn951180)
	Re951190, _ = re2.Compile(Ptn951190)
	Re951200, _ = re2.Compile(Ptn951200)
	Re951210, _ = re2.Compile(Ptn951210)
	Re951220, _ = re2.Compile(Ptn951220)
	Re951230, _ = re2.Compile(Ptn951230)
	Re951240, _ = re2.Compile(Ptn951240)
	Re951250, _ = re2.Compile(Ptn951250)
	Re951260, _ = re2.Compile(Ptn951260)
}
