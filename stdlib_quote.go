// Copied and trimmed down from https://github.com/golang/go/blob/e3769299cd3484e018e0e2a6e1b95c2b18ce4f41/src/strconv/quote.go
// We want to use the standard library's private "quoteWith" function rather than write our own so that we get robust unicode support.
// Every private function called by quoteWith was copied.
// There are 2 modifications to simplify the code:
// 1. The unicode.IsPrint function was substituted for the custom implementation of IsPrint
// 2. All code paths reachable only when ASCIIonly or grphicOnly are set to true were removed.

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package yaml

const (
	lowerhex = "0123456789abcdef"
)

func quoteWith(s string, quote byte) string { _ = "STUB: not implemented"; return "" }

func appendQuotedWith(buf []byte, s string, quote byte) []byte {
	_ = "STUB: not implemented"
	// Often called with big strings, so preallocate. If there's quoting,
	// this is conservative but still helps a lot.
	return nil
}

func appendEscapedRune(buf []byte, r rune, quote byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

// goccy/go-yaml patch on top of the standard library's appendEscapedRune function.
//
// We use this to implement the YAML single-quoted string, where the only escape sequence is '', which represents a single quote.
// The below snippet from the standard library is for escaping e.g. \ with \\, which is not what we want for the single-quoted string.
//
// if r == rune(quote) || r == '\\' { // always backslashed
// 	buf = append(buf, '\\')
// 	buf = append(buf, byte(r))
// 	return buf
// }
