package parse

import (
	"errors"

	"github.com/orange-lang/orange/ast"
	"github.com/orange-lang/orange/parse/lexer"
	"github.com/orange-lang/orange/parse/lexer/token"
)

func isStatementToken(t token.Token) bool {
	return t == token.Var || t == token.Package || t == token.Import ||
		t == token.If || t == token.Alias
}

func (p parser) parseStatement() (ast.Statement, error) {
	if ok, _ := p.peekFrom(isStatementToken); !ok {
		return nil, errors.New("Expected statement")
	}

	switch lexeme, _ := p.stream.Peek(); lexeme.Token {
	case token.Var:
		return p.parseVarDecl()
	case token.Package:
		return p.parsePackageDecl()
	case token.Import:
		return p.parseImportDecl()
	case token.If:
		return p.parseIf()
	case token.Alias:
		return p.parseAlias()
	}

	return nil, errors.New("Unexpected lexeme")
}

func (p parser) parseAlias() (*ast.AliasDecl, error) {
	if _, err := p.expect(token.Alias); err != nil {
		return nil, err
	}

	nameLexeme, err := p.expect(token.Identifier)
	if err != nil {
		return nil, errors.New("Expected identifier")
	}

	if _, err := p.expect(token.Assign); err != nil {
		return nil, err
	}

	targetType, err := p.parseType()
	if err != nil {
		return nil, err
	}

	return &ast.AliasDecl{Name: nameLexeme.Value, Type: targetType}, nil
}

func (p parser) parseIf() (*ast.IfStmt, error) {
	if _, err := p.expect(token.If); err != nil {
		return nil, err
	}

	mainCondition, err := p.parseCondition()
	if err != nil {
		return nil, err
	}

	return mainCondition, err
}

// Parses the condition part of an if statement, and then
// elif or else.
func (p parser) parseCondition() (*ast.IfStmt, error) {
	var condition ast.Expression
	var ifPart *ast.BlockStmt
	var elsePart ast.Node
	var err error

	if _, err := p.expect(token.OpenParen); err != nil {
		return nil, errors.New("Expected open parenthesis")
	}

	condition, err = p.parseExpr()
	if err != nil {
		return nil, err
	}

	if _, err := p.expect(token.CloseParen); err != nil {
		return nil, errors.New("Expected close parenthesis")
	}

	ifPart, err = p.parseBlock()
	if err != nil {
		return nil, err
	}

	if ok, _ := p.allow(token.Elif); err == nil && ok {
		elsePart, err = p.parseCondition()
	} else if ok, _ := p.allow(token.Else); err == nil && ok {
		elsePart, err = p.parseBlock()
	}

	if err != nil {
		return nil, err
	}

	return &ast.IfStmt{
		Condition: condition,
		Body:      ifPart,
		Else:      elsePart,
	}, nil
}

func (p parser) parseImportDecl() (*ast.ImportDecl, error) {
	var fullPackageName string

	if _, err := p.expect(token.Import); err != nil {
		return nil, err
	}

	for true {
		name, err := p.expect(token.Identifier)
		if err != nil {
			return nil, err
		}

		fullPackageName += name.Value

		if ok, _ := p.allow(token.Dot); !ok {
			break
		}

		fullPackageName += "."
	}

	return &ast.ImportDecl{Name: fullPackageName}, nil
}

func (p parser) parsePackageDecl() (*ast.PackageDecl, error) {
	var fullPackageName string

	if _, err := p.expect(token.Package); err != nil {
		return nil, err
	}

	for true {
		name, err := p.expect(token.Identifier)
		if err != nil {
			return nil, err
		}

		fullPackageName += name.Value

		if ok, _ := p.allow(token.Dot); !ok {
			break
		}

		fullPackageName += "."
	}

	return &ast.PackageDecl{Name: fullPackageName}, nil
}

func (p parser) parseVarDecl() (*ast.VarDecl, error) {
	var idLexeme lexer.Lexeme
	var nodeType ast.Type
	var nodeValue ast.Expression
	var err error

	if _, err := p.expect(token.Var); err != nil {
		return nil, err
	}

	if idLexeme, err = p.expect(token.Identifier); err != nil {
		return nil, err
	}

	if nodeType, err = p.tryParseColonType(); err != nil {
		return nil, err
	}

	if nodeValue, err = p.tryParseEqualValue(); err != nil {
		return nil, err
	}

	return &ast.VarDecl{
		Name:  idLexeme.Value,
		Type:  nodeType,
		Value: nodeValue,
	}, nil
}

func (p parser) tryParseColonType() (ast.Type, error) {
	if ok, err := p.allow(token.Colon); err != nil {
		return nil, err
	} else if ok {
		return p.parseType()
	}

	return nil, nil
}

func (p parser) tryParseEqualValue() (ast.Expression, error) {
	if ok, err := p.allow(token.Assign); err != nil {
		return nil, err
	} else if ok {
		return p.parseExpr()
	}

	return nil, nil
}
