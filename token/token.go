package token 


type TokenType string 

type Token struct {
    Type TokenType 
    Literal string 
}

const ( 
    ILLEGAL = "ILLEGAL" 
    EOF = "EOF"


    // Identifiers + literals 
    IDENT = "IDENT" // add, foo, x, y 
    INT = "INT" // 3221

    // Operators 
    ASSIGN = "=" 
    PLUS = "+" 

    //Delimeters 
    COMMA = "," 
    SEMICOLON = ";" 

    LPAREN = "(" 
    RPAREN = ")" 
    LBRACE = "{" 
    RBRACE = "}" 

    // keywords 
    FUNCTION = "FUNCTION" 
    LET = "LET" 

)



 var keywords = map[string]TokenType{ 
     "fn": FUNCTION,
     "let": LET, 
 } 


 func LookupIndent(ident string) TokenType { 
     if tok, ok := keywords[ident]; ok { 
	 return tok 
     } 
     return IDENT 
 } 

