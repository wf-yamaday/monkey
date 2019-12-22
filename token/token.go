package token

// Type 字句の種類
type Type string

// Token 字句の定義
type Token struct {
	Type    Type
	Literal string
}

const (
	// ILLEGAL 無効な文字
	ILLEGAL = "ILLEGAL"
	// EOF 終端文字
	EOF     = "EOF"

	// IDENT 識別子・リテラル
	IDENT = "IDENT"
	// INT 整数トークン
	INT   = "INT"

	// ASSIGN 束縛演算子
	ASSIGN   = "="
	// PLUS 加算演算子
	PLUS     = "+"
	// MINUS 減算演算子
	MINUS    = "-"
	// BANG 否定演算子
	BANG     = "!"
	// ASTERISK 乗算演算子
	ASTERISK = "*"
	// SLASH 除算演算子
	SLASH    = "/"

	// LT 比較演算子 小なり
	LT = "<"
	// GT 比較演算子 大なり
	GT = ">"

	// EQ 等価演算子
	EQ     = "=="
	// NOTEQ 不等価演算子
	NOTEQ = "!="

	// COMMA デリミタ
	COMMA     = ","
	// SEMICOLON デリミタ
	SEMICOLON = ";"

	// LPAREN 左括弧
	LPAREN = "("
	// RPAREN 右括弧
	RPAREN = ")"
	// LBRACE 左中括弧
	LBRACE = "{"
	// RBRACE 右中括弧
	RBRACE = "}"

	// FUNCTION 関数キーワード
	FUNCTION = "FUNCTION"
	// LET 変数キーワード
	LET      = "LET"
	// TRUE 真
	TRUE     = "TRUE"
	// FALSE 偽
	FALSE    = "FALSE"
	// IF 条件
	IF       = "IF"
	// ELSE または
	ELSE     = "ELSE"
	// RETURN return
	RETURN   = "RETURN"
)

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent キーワードに登録されていないものはIDENTとする
func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
