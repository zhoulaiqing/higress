package rule_tasks

import (
	ahocorasick "github.com/petar-dambovaliev/aho-corasick"
	"github.com/wasilibs/go-re2"
)

const (
	PTN_944100   = `java\.lang\.(?:runtime|processbuilder)`
	PTN_944110_1 = `(?:runtime|processbuilder)`
	PTN_944110_2 = `(?:unmarshaller|base64data|java\.)`
	PTN_944120_1 = `(?:clonetransformer|forclosure|instantiatefactory|instantiatetransformer|invokertransformer|prototypeclonefactory|prototypeserializationfactory|whileclosure|getproperty|filewriter|xmldecoder)`
	PTN_944120_2 = `(?:runtime|processbuilder)`
	PTN_944140   = `.*\.(?:jsp|jspx)\.*$`
	PTN_944150   = `(?i)(?:\$|&dollar;?)(?:\{|&l(?:brace|cub);?)(?:[^\}]{0,15}(?:\$|&dollar;?)(?:\{|&l(?:brace|cub);?)|jndi|ctx)`
)

var Re944100, Re9441101, Re9441102, Re9441201, Re9441202, Re944140, Re944150 *re2.Regexp

var JAVA_CLASSES = []string{
	"com.opensymphony.xwork2",
	"com.sun.org.apache",
	"freemarker.core",
	"freemarker.template",
	"freemarker.ext.rhino",
	"java.io.BufferedInputStream",
	"java.io.BufferedReader",
	"java.io.ByteArrayInputStream",
	"java.io.ByteArrayOutputStream",
	"java.io.CharArrayReader",
	"java.io.DataInputStream",
	"java.io.File",
	"java.io.FileOutputStream",
	"java.io.FilePermission",
	"java.io.FileWriter",
	"java.io.FilterInputStream",
	"java.io.FilterOutputStream",
	"java.io.FilterReader",
	"java.io.InputStream",
	"java.io.InputStreamReader",
	"java.io.LineNumberReader",
	"java.io.ObjectOutputStream",
	"java.io.OutputStream",
	"java.io.PipedOutputStream",
	"java.io.PipedReader",
	"java.io.PrintStream",
	"java.io.PushbackInputStream",
	"java.io.Reader",
	"java.io.StringReader",
	"java.lang.Class",
	"java.lang.Integer",
	"java.lang.Number",
	"java.lang.Object",
	"java.lang.Process",
	"java.lang.ProcessBuilder",
	"java.lang.reflect",
	"java.lang.Runtime",
	"java.lang.String",
	"java.lang.StringBuilder",
	"java.lang.System",
	"javassist",
	"javax.script.ScriptEngineManager",
	"org.apache.commons",
	"org.apache.struts",
	"org.apache.struts2",
	"org.omg.CORBA",
	"java.beans.XMLDecode",
	"sun.reflect",
}
var Rule944130Matcher ahocorasick.AhoCorasick

func init() {
	Re944100, _ = re2.Compile(PTN_944100)
	Re9441101, _ = re2.Compile(PTN_944110_1)
	Re9441102, _ = re2.Compile(PTN_944110_2)
	Re9441201, _ = re2.Compile(PTN_944120_1)
	Re9441202, _ = re2.Compile(PTN_944120_2)
	Re944140, _ = re2.Compile(PTN_944140)
	Re944150, _ = re2.Compile(PTN_944150)

	Rule944130Matcher = AHO_CORASICK_BUILDER.Build(JAVA_CLASSES)
}
