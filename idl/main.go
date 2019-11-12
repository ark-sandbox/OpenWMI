package main

import (
	"os"

	"github.com/alecthomas/kong"
	"github.com/alecthomas/repr"

	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
)

type Interf struct {
	Header *InterfHeader `"[" @@`
	Body   *InterfBody   `"{" @@ "}"`
}

type InterfHeader struct {
	Attributes []*InterfAttribute `@@ ("," @@)* "]"`
	InterfName string             `"interface" @Ident`
}

type InterfAttribute struct {
	InterfUUID     *AttribUUID      `@@`
	Version        *AttribVersion   `|@@`
	Endpoint       *AttribEndpoint  `|@@`
	Exception      *AttribException `|@@`
	Local          string           `|"local"`
	DefaultPointer *AttribPointer   `|@@`
}

type AttribUUID struct {
	Value string `"uuid""(" @UUID ")"`
}

type AttribVersion struct {
	Value string `"version""(" @Number @("." Number)? ")"`
}

type AttribEndpoint struct {
	Value []string `"endpoint""(" @String ("," @String)*")"`
}

type AttribException struct {
	Value []string `"exceptions""(" @Ident ("," @Ident)*")"`
}

type AttribPointer struct {
	Value string `"pointer_default""(" @("ref"|"unique"|"ptr") ")"`
}

type InterfBody struct {
	Imports     []*ImportStatment `(@@)*`
	InterfComps []*InterfComp     `{@@}`
}

type ImportStatment struct {
	ImportList []string `"import" @String ("," @String)* ";"`
}

type InterfComp struct {
	Export *ExportDeclarator `@@ ";"`
	Oper   *OpDeclarator     `|@@ ";"`
}

type ExportDeclarator struct {
	Type  *TypeDeclarator  `@@`
	Const *ConstDeclarator `|@@`
	//Tagged TaggedDeclarator `|@@`
}

type TypeDeclarator struct {
	AttributeList *TypeAttributeList `"typedef" (@@)?`
	Spec          *TypeSpec          `@@`
	Declrs        []*Declarator      `(@@) ("," @@)*`
}

type TypeAttributeList struct {
	AttributeList []*TypeAttribute `"[" (@@) ("," @@)* "]"`
}

type TypeAttribute struct {
	TransmitAs       *SimpleTypeSpec `"transmit_as" "("@@")"`
	Handle           string          `|@"Handle"`
	UsageAttribute   string          `|@("string" | "context_handle")`
	UniontSwitchAttr *SwitchTypeSpec `| "switch_type""(" @@ ")"`
	PtrAttr          string          `| @("ref" | "unique" | "ptr")`
}

type TypeSpec struct {
	Simple      *SimpleTypeSpec      `@@`
	Constructed *ConstructedTypeSpec `|@@`
}

/*(80) <predefined_type_spec> ::= error_status_t
	| <international_character_type>
(81) <international_character_type> ::= ISO_LATIN_1
	| ISO_MULTI_LINGUAL
	| ISO_UCS
*/

type SimpleTypeSpec struct {
	BaseType       *BaseTypeSpec `@@`
	PredefinedType string        `|@("error_status_t"|"ISO_LATIN_1"|"ISO_MULTI_LINGUAL"|"ISO_UCS")`
	Identifier     string        `| @Ident`
}

/*<base_type_spec> ::= <floating_pt_type>
| <integer_type>
| <char_type>
| <boolean_type>
| <byte_type>
| <void_type>
| <handle_type>
*/
type BaseTypeSpec struct {
	FloatT  string       `@("float"|"double")`
	IntT    *IntegerType `|@@`
	ChatT   string       `|@(("unsigned")?"char")`
	BoolT   string       `|@("boolean")`
	ByteT   string       `|@("byte")`
	VoidT   string       `|@("void")`
	HandleT string       `|@("handle_t")`
}

type IntegerType struct {
	PrimIntT    *PrimitiveIntegerType `@@`
	HyperT      string                `|@("hyper" "unsigned"? "int"?)`
	UnsignHyerT string                `|@("unsigned" "hyper" "int"?)`
}

type PrimitiveIntegerType struct {
	SignedIntT   *SignedIntType   `@@`
	UnsignedIntT *UnsignedIntType `|@@`
}

type SignedIntType struct {
	IntSizeT *IntegerSizeType `@@ ("int")?`
}

type UnsignedIntType struct {
	UIntSizeT *IntegerSizeType `("unsigned" @@ ("int")?)|(@@ "unsigned" ("int")?)`
}

type IntegerSizeType struct {
	Size string `@("long"|"short"|"small"|"int")`
}

/*(38) <constructed_type_spec> ::= <struct_type>
| <union_type>
| <enumeration_type>
| <tagged_declarator>
| <pipe_type>
*/
type ConstructedTypeSpec struct {
	StructT *StructType `@@`
	UnionT  *UnionType  `|@@`
	EnumT   *EnumType   `|@@`
	/* SKIPPED
	TaggedDeclT TaggedDeclarator `|@@`
	PipeT TypeSpec `|"pipe" @@`
	*/
}

type StructType struct {
	Members []*StructMember `"struct" "{"(@@)+"}"`
}

type StructMember struct {
	Attributes []*StructFieldAttrb `("["(@@) ("," @@)*"]")?`
	Type       *TypeSpec           `@@`
	Declrs     []*Declarator       `(@@) ("," @@)* ";"`
}

/*(67) <field_attribute> ::= first_is ( <attr_var_list> )
	| last_is ( <attr_var_list> )
	| length_is ( <attr_var_list> )
	| min_is ( <attr_var_list> )
	| max_is ( <attr_var_list> )
	| size_is ( <attr_var_list> )
	| <usage_attribute>
	| <union_instance_switch_attr>
	| ignore
	| <ptr_attr>
(68) <attr_var_list> ::= <attr_var> [ , <attr_var> ] ...
(69) <attr_var> ::= [ [ * ] <Identifier> ]
*/
type StructFieldAttrb struct {
	FirstIs                 *AttribVarList `"first_is" @@`
	LastIs                  *AttribVarList `| "last_is" @@`
	LengthIs                *AttribVarList `| "length_is" @@`
	MinIs                   *AttribVarList `|"min_is" @@`
	MaxIs                   *AttribVarList `|"max_is" @@`
	SizeIs                  *AttribVarList `|"size_is" @@`
	Range                   *AttribRange   `|"range" @@`
	UsageAttribute          string         `|@("string" | "context_handle")`
	UnionInstanceSwitchAttr *AttribVar     `| "switch_is" "(" @@ ")"`
	Ignore                  string         `| "ignore"`
	PtrAttr                 string         `| ("ref" | "unique" | "ptr")`
}

type AttribVarList struct {
	Attributes []*AttribVar `"(" ((@@) ("," @@)*)? ")"`
}

type AttribVar struct {
	Variable string `@( ("*")? Ident)`
}

type AttribRange struct {
	Start string `"(" @Number ","`
	End   string `@Number ")"`
}

type UnionType struct {
	Members []*UnionCase `"union" "{"(@@)+"}"`
}

/*(53) <union_case> ::= <union_case_label> [ <union_case_label> ] ...
	<union_arm>
	| <default_case>
(54) <union_case_label> ::= case <const_exp> :
(55) <default_case> ::= default : <union_arm>
(55.2) <union_arm> ::= [ <field_declarator> ] ;
*/

//TODO: Create a typecast to StructMember to UnionMember...
type UnionCase struct {
	CaseLabel *UnionCaseLabel `@@`
	Field     *StructMember   `@@`
}

//FIXME: Does current ConstExpression handle enumerators etc.
type UnionCaseLabel struct {
	NormalCase  []*ConstExpression `"[""case" "(" @@ ")""]"`
	DefaultCase string             `"[" "default" "]"`
}

type EnumType struct {
	Identifiers []string `(@Ident) ("," @Ident)*`
}

/*(49) <switch_type_spec> ::= <primitive_integer_type>
| <char_type>
| <boolean_type>
| <Identifier>
*/
type SwitchTypeSpec struct {
	IntT       *PrimitiveIntegerType `@@`
	ChatT      string                `|@(("unsigned")?"char")`
	BoolT      string                `|@("boolean")`
	Identifier string                `|@Ident`
}

type ConstDeclarator struct {
	ConstType       *ConstantType   `"const" @@`
	ConstIdentifier string          `@Ident "=" `
	ConstExpression ConstExpression `@@`
}

type ConstantType struct {
	IntT   *IntegerType `@@`
	OtherT string       `|@("char"|"boolean"|"void"|"void *")`
}

/*
//Skipping didn't see any use of it.
type TaggedDeclarator struct {
}
*/
/* SKIPPED: all the grammar except primitive constants */
type ConstExpression struct {
	Int        string `@Int`
	Number     string `|@Number`
	Identifier string `|@Ident`
	String     string `|@String`
	Char       string `|@Char`
	NTF        string `|@("NULL"|"TRUE"|"FALSE")`
}

/*(71) <op_declarator> ::= [ <operation_attributes> ]
<simple_type_spec> <Identifier> <param_declarators>
*/

type OpDeclarator struct {
	OpAttrs []*OpAttribute    `("["(@@) ("," @@)*"]")?`
	Type    *SimpleTypeSpec   `@@`
	OpName  string            `@Ident`
	Params  *ParamDeclarators `@@`
}

/*<operation_attribute> ::= idempotent
| broadcast
| maybe
| reflect_deletions
| <usage_attribute>
| <ptr_attr>
*/
type OpAttribute struct {
	OtherAttribute string `@("idempotent"|"broadcast"|"maybe"|"reflect_deletions")`
	UsageAttribute string `|@("string" | "context_handle")`
	PtrAttr        string `| @("ref" | "unique" | "ptr")`
}

type ParamDeclarators struct {
	Params  []*ParamDeclarator `"("(@@) ("," @@)*")"`
	NoParam string             `| "(" @"void" ")"`
}

/*
(75) <param_declarator> ::= <param_attributes> <type_spec> <declarator>
(76) <param_attributes> ::= <[> <param_attribute>
	[ , <param_attribute> ] ... <]>
(77) <param_attribute> ::= <directional_attribute>
	| <field_attribute>
(78) <directional_attribute> ::= in
	| out
(79) <function_declarator> ::= <direct_declarator> <param_declarators>
*/
type ParamDeclarator struct {
	ParamAttrs []*ParamAttribute `"["(@@) ("," @@)*"]"`
	ParamType  *TypeSpec         `@@`
	ParamDeclr *Declarator       `@@`
}

type ParamAttribute struct {
	DirectionAttr string            `@("in"|"out")`
	FieldAttr     *StructFieldAttrb `|@@`
}

type Declarator struct {
	PointerOpt string            `@("*")*`
	DirectDecl *DirectDeclarator `@@`
}

/*
(24) <direct_declarator> ::= <Identifier>
	| ( <declarator> )
	| <array_declarator>
	| <function_declarator>
*/
type DirectDeclarator struct {
	ArrayDeclr   *ArrayDeclarator    `@@`
	FuncDeclr    *FunctionDeclarator `|@@`
	OfDeclarator *Declarator         `|"("@@")"`
	Identifier   string              `|@Ident`
}

/* (59) <array_declarator> ::= <direct_declarator> <array_bounds_declarator>
(61) <array_bounds_declarator> ::= <[> [ <array_bound> ] <]>
	| <[> <array_bounds_pair> <]>
(62) <array_bounds_pair> ::= <array_bound> .. <array_bound>
(63) <array_bound> ::= *
	| <integer_const_exp>
	| <Identifier>
*/
type ArrayDeclarator struct {
	Identifier string                `@Ident`
	Bounds     *ArrayBoundDeclarator `@@`
}

type ArrayBoundDeclarator struct {
	BoundPair *ArrayBoundPair ` "[" @@ "]"`
	Bound     *ArrayBound     `| "[" @@ "]"`
	NoBound   string          `| "[" "]"`
}

type ArrayBound struct {
	Value string `@("*"|Number|Ident)`
}

type ArrayBoundPair struct {
	Start *ArrayBound `@@ ".""."`
	End   *ArrayBound `@@`
}

/* (79) <function_declarator> ::= <direct_declarator> <param_declarators> */
type FunctionDeclarator struct {
	//Declr *DirectDeclarator `@@`
	Identifier string             `@Ident`
	Params     []*ParamDeclarator `"("(@@) ("," @@)*")?"`
}

var (
	IDLLexer = lexer.Must(lexer.Regexp(
		`(?m)` + //Multiline Regex Flag
			`(\s+)` + //Not capturing whitespaces.
			`|(?s)(?P<MultiLineComment>/\*.*?\*/)` +
			`|(?P<Comment>^//.*?$)` +
			`|(?P<UUID>[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12})` +
			`|\b(?P<Keyword>boolean|byte|case|char|const|default|double|
                enum|false|float|handle_t|hyper|import|int|interface|long|NULL|pipe|short|small|struct|switch|TRUE|typedef|union|unsigned|void|string)\b` +
			`|(?P<Ident>[a-zA-Z][a-zA-Z_\d]*)` +
			`|(?P<Number>\d+)` +
			`|(?P<Int>[\-+]\d+)` +
			`|(?P<Char>'(?:\\.|[^'])')` +
			`|(?P<String>"(?:\\.|[^"])*")` +
			`|(?P<Float>\d+(?:\.\d+)?)` +
			`|(?P<Punct>[][{}\]\]=(),.*;#\-<>\|])`,
	))

	parser = participle.MustBuild(&Interf{},
		participle.Lexer(IDLLexer),
		participle.Elide("Comment", "MultiLineComment"),
		participle.UseLookahead(2),
	)
	cli struct {
		Files []string `arg:"" type:"existingfile" required:"" help:"IDL defination files to parse."`
	}
)

func main() {
	ctx := kong.Parse(&cli)
	for _, file := range cli.Files {
		idl := &Interf{}
		r, err := os.Open(file)
		ctx.FatalIfErrorf(err)
		err = parser.Parse(r, idl)
		r.Close()
		repr.Println(idl)
		ctx.FatalIfErrorf(err)
	}
}
