package lexer

import "go_interpreter/token"

// supports ASCII characters
// TODO: support Unicode
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// this gives us the next character and advance our position in the input string
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) { // check if reached the end of input
		l.ch = 0 // ASCII Code for the "NUL" character, it means "end of file"
	} else {
		l.ch = l.input[l.readPosition] // sets l.ch to the next character
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// look at the current character under examination (l.ch) and return a token depending on which character it is.
// before returning the token we advance our pointers into the input so when we call NextToken() again the l.ch field is already updated.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	default:
		if isLetter(l.ch) { // check if the current character is a letter
			tok.Literal = l.readIdentifier()          // calls readChar() repeatedly and advance our readPosition and position fields pas the last character of the current identifier.
			tok.Type = token.LookupIdent(tok.Literal) // check if it is a identifier or a keyword.
			return tok                                // returns tok, so we don't need to call the readChar() after the switch statement
		} else if isDigit(l.ch) { // check if the current character is a number
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// it reads in an identifier and advances our lexer’s positions until it encounters a non-letter-character
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// check if the given argument is a letter.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// skips whitespace and newline characters.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// it reads in a number and advances our lexer’s positions until it encounters a non-number-character.
// doesn't support floats, numbers in hex notation, or Octal notation yet!
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// check if the given argument is a number.
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// peeks ahead in the input and not move around in it.
// similiar to readChar(), except that it doesn't increment l.position and l.readPosition.
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// TODO:
// 1.add suport for character escaping, such as string literals like "hello \"world\"", "hello\n world" and "hello\t\t\tworld"
// 2. make readString report an error instead of simply returning when it reaches the end of the input
func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}
