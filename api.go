package fufik

import (
	"github.com/fobus89/fufik/internal/ast"
	"github.com/fobus89/fufik/internal/parser"
	"github.com/fobus89/fufik/internal/token"
)

type (
	Parser       = parser.Parser
	Expr         = ast.Expr
	Token        = token.Token
	TokenType    = token.TokenType
	BindingPower = parser.BindingPower
)

const (
	Lowest BindingPower = iota
	Comma
	Assigment
	Logical
	Relational
	Additive
	Muptiplicative
	Unary
	Call
	Member
	Primary
	Highest
)

const (
	token_start TokenType = iota

	// Special
	EOF     // end of file
	ILLEGAL // unknown token

	// Literals
	literal_start
	STRING_LITERAL // "string"
	STRING_FORMAT  // "{expr}"
	NUMBER_LITERAL // int | float
	FLOAT_LITERAL  // 0.0-9.0
	INT_LITERAL    // 0-9
	literal_end

	// Identifiers
	IDENT // identifier

	// Symbols (operators and punctuation)
	symbol_start

	// Arithmetic operators
	arithmetic_start
	PLUS        // +
	MINUS       // -
	STAR        // *
	STAR_STAR   // **
	SLASH       // /
	PERCENT     // %
	PLUS_PLUS   // ++
	MINUS_MINUS // --
	arithmetic_end

	// Bitwise operators
	bitwise_start
	AMP   // &
	PIPE  // |
	CARET // ^
	TILDE // ~
	LT_LT // <<
	GT_GT // >>
	bitwise_end

	// Comparison operators
	comparison_start
	EQ_EQ   // ==
	BANG_EQ // !=
	GT      // >
	LT      // <
	GT_EQ   // >=
	LT_EQ   // <=
	comparison_end

	// Assignment operators
	assignment_start
	EQ         // =
	PLUS_EQ    // +=
	MINUS_EQ   // -=
	STAR_EQ    // *=
	SLASH_EQ   // /=
	PERCENT_EQ // %=
	AMP_EQ     // &=
	PIPE_EQ    // |=
	LT_LT_EQ   // <<=
	GT_GT_EQ   // >>=
	assignment_end

	// Logical operators
	logical_start
	BANG      // !
	AMP_AMP   // && | and
	PIPE_PIPE // || | or
	logical_end

	// Other operators
	operator_start
	EQ_GT       // =>
	MINUS_GT    // ->
	LT_MINUS    // <-
	DOT_DOT     // ..
	DOT_DOT_EQ  // ..=
	DOT_DOT_DOT // ...
	operator_end

	// Punctuation
	punctuation_start
	QUESTION    // ?
	COLON       // :
	COLON_COLON // ::
	SEMICOLON   // ;
	COMMA       // ,
	DOT         // .
	AT          // @
	HASH        // #
	DOLLAR      // $
	LPARENT     // (
	RPARENT     // )
	LBRACE      // {
	RBRACE      // }
	LBRACKET    // [
	RBRACKET    // ]
	punctuation_end

	symbol_end

	// Keywords
	keyword_start
	AS        // as
	ASYNC     // async
	AWAIT     // await
	BETWEEN   // between
	BREAK     // break
	CATCH     // catch
	CONST     // const
	DEFER     // defer
	DO        // do
	ELSE      // else
	ENUM      // enum
	EXPORT    // export
	FALSE     // false
	FINALLY   // finally
	FN        // fn
	FOR       // for
	FROM      // from
	IF        // if
	IMPL      // impl
	IMPORT    // import
	IN        // in
	INTERFACE // interface
	IS        // is
	LET       // let
	LOOP      // loop
	MATCH     // match
	MUT       // mut
	NIL       // nil
	PUB       // pub
	REF       // ref
	RETURN    // return
	SELF      // self
	SPAWN     // spawn
	// STRUCT    // struct
	SUPER  // super
	THROW  // throw
	TRUE   // true
	TRY    // try
	TYPE   // type
	TYPEOF // typeof
	USE    // use
	WHERE  // where
	WHILE  // while
	YIELD  // yield

	// SQL keywords
	ALL      // all
	ALTER    // alter
	ASC      // asc
	BY       // by
	CREATE   // create
	CROSS    // cross
	DESC     // desc
	DISTINCT // distinct
	DROP     // drop
	EXISTS   // exists
	FULL     // full
	GROUP    // group
	HAVING   // having
	INNER    // inner
	INTO     // into
	JOIN     // join
	LEFT     // left
	LIKE     // like
	LIMIT    // limit
	OFFSET   // offset
	ON       // on
	ORDER    // order
	OUTER    // outer
	RIGHT    // right
	SELECT   // select
	SET      // set
	TABLE    // table
	UNION    // union
	VALUES   // values
	keyword_end

	// Built-in functions
	builtin_start
	PRINT   // print
	PRINTLN // println
	builtin_end

	// Compile-time / macros
	compiletime_start
	MACROS         // macro fn
	EXPR           // expr
	CONSTEXPR      // $constexpr
	COMPTIME_IDENT // $"{expr}string"
	compiletime_end

	// Types
	type_start
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Float32
	Float64
	String
	Char
	Byte
	Rune
	Array
	Slice
	Map
	Tuple
	Struct
	Enum
	Interface
	Any
	Void
	Never
	type_end

	token_end
)
