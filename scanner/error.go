package scanner

import "github.com/goccy/go-yaml/token"

type InvalidTokenError struct {
	Token *token.Token
}

func (e *InvalidTokenError) Error() string { _ = "STUB: not implemented"; return "" }

func ErrInvalidToken(tk *token.Token) *InvalidTokenError { _ = "STUB: not implemented"; return nil }
