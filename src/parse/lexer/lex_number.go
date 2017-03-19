package lexer

import (
	"errors"
	"fmt"
	"parse/lexer/token"
	"strconv"
	"strings"
	"unicode"
)

func lexNumber(s RuneStream) (Lexeme, error) {
	base := 10
	l := Lexeme{Token: token.IntVal}
	shouldParseSuffix := true

	// Get the prefix so we can get the base of the number
	if p, err := lexPrefix(s); err != nil {
		return l, err
	} else if p != NoPrefix {
		base = p.GetBase()
		l.Token = token.UIntVal
	}

	// Get the "body" of the number.
	for isHexDigit(s.Peek()) {
		val := s.Peek()

		// f and d are suffixes, but we may have peeked them
		// here thinking they were numbers.
		if base != 16 && isFloatingPointSuffix(val) {
			break
		}

		l.Value += string(s.Next())

		consumeSeparators(s)

		// Next sequence is .[character]
		if next := s.Lookahead(2); next[0] == '.' && !isHexDigit(next[1]) {
			shouldParseSuffix = false
			break
		} else if next[0] == '.' && l.Token != token.DoubleVal {
			// We've come across the first .; let's consume
			// and change to being a Double
			l.Value += string(s.Next())
			l.Token = token.DoubleVal
		}
	}

	if !validStringForBase(l.Value, base) {
		return l, fmt.Errorf("Invalid number %v for base %v", l.Value, base)
	}

	if shouldParseSuffix {
		if tokFromSuffix, err := lexNumberSuffix(s); err != nil {
			return l, err
		} else if tokFromSuffix != token.EOF {
			l.Token = tokFromSuffix
		}

		if base != 10 && (l.Token == token.FloatVal || l.Token == token.Double) {
			return l, errors.New("Number of non-decimal base cannot be floating-point")
		}

		if base != 10 && l.Token.SignedValue() {
			return l, errors.New("Number of non-decimal base cannot be signed")
		}
	}

	if l.Token != token.FloatVal && l.Token != token.DoubleVal &&
		strings.Contains(l.Value, ".") {
		return l, errors.New("Floating-point value cannot have integral suffix")
	}

	if base != 10 {
		i, err := strconv.ParseInt(l.Value, base, 64)
		if err != nil {
			return l, errors.New("Number out of range")
		}

		l.Value = fmt.Sprintf("%v", i)
	}

	return l, nil
}

func lexPrefix(s RuneStream) (prefix, error) {
	if lookahead := s.Lookahead(2); lookahead[0] != '0' || lookahead[1] == '.' {
		return NoPrefix, nil
	} else if prefix := makePrefix(lookahead[1]); prefix != NoPrefix {
		s.Get(2)
		return prefix, nil
	} else if !validSuffixStarter(lookahead[1]) {
		// If the character that we're treating as a prefix would be valid
		// for a suffix, then we can ignore it. Otherwise, it's an
		// invalid suffix.
		return NoPrefix, fmt.Errorf("Invalid numeric prefix %v", lookahead[0])
	}

	return NoPrefix, nil
}

func lexNumberSuffix(s RuneStream) (token.Token, error) {
	suffix := ""

	for unicode.IsLetter(s.Peek()) || unicode.IsDigit(s.Peek()) {
		suffix += string(s.Next())
	}

	suffixTable := map[string]token.Token{
		"f":   token.FloatVal,
		"d":   token.DoubleVal,
		"u":   token.UIntVal,
		"u8":  token.UInt8Val,
		"u16": token.UInt16Val,
		"u32": token.UInt32Val,
		"u64": token.UInt64Val,
		"i":   token.IntVal,
		"i8":  token.Int8Val,
		"i16": token.Int16Val,
		"i32": token.Int32Val,
		"i64": token.Int64Val,
	}

	tok, ok := suffixTable[suffix]
	if !ok && suffix != "" {
		return token.EOF, fmt.Errorf("Invalid suffix %v", suffix)
	}

	return tok, nil
}

func consumeSeparators(s RuneStream) {
	for s.Peek() == '_' {
		s.Next()
	}
}
