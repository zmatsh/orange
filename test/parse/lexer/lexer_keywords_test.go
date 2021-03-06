package test

import (
	"github.com/orange-lang/orange/parse/lexer/token"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
)

var _ = Describe("Keywords", func() {
	DescribeTable("should get lexed when the keyword is", expectToken,
		Entry("int", "int", token.Int),
		Entry("int8", "int8", token.Int8),
		Entry("int16", "int16", token.Int16),
		Entry("int32", "int32", token.Int32),
		Entry("int64", "int64", token.Int64),
		Entry("uint", "uint", token.UInt),
		Entry("uint8", "uint8", token.UInt8),
		Entry("uint16", "uint16", token.UInt16),
		Entry("uint32", "uint32", token.UInt32),
		Entry("uint64", "uint64", token.UInt64),
		Entry("var", "var", token.Var),
		Entry("enum", "enum", token.Enum),
		Entry("class", "class", token.Class),
		Entry("public", "public", token.Public),
		Entry("protected", "protected", token.Protected),
		Entry("private", "private", token.Private),
		Entry("if", "if", token.If),
		Entry("elif", "elif", token.Elif),
		Entry("else", "else", token.Else),
		Entry("for", "for", token.For),
		Entry("while", "while", token.While),
		Entry("do", "do", token.Do),
		Entry("break", "break", token.Break),
		Entry("continue", "continue", token.Continue),
		Entry("def", "def", token.Def),
		Entry("extern", "extern", token.Extern),
		Entry("interface", "interface", token.Interface),
		Entry("package", "package", token.Package),
		Entry("import", "import", token.Import),
		Entry("new", "new", token.New),
		Entry("delete", "delete", token.Delete),
		Entry("get", "get", token.Get),
		Entry("set", "set", token.Set),
		Entry("virtual", "virtual", token.Virtual),
		Entry("final", "final", token.Final),
		Entry("where", "where", token.Where),
		Entry("data", "data", token.Data),
		Entry("extend", "extend", token.Extend),
		Entry("const", "const", token.Const),
		Entry("try", "try", token.Try),
		Entry("catch", "catch", token.Catch),
		Entry("finally", "finally", token.Finally),
		Entry("throw", "throw", token.Throw),
		Entry("of", "of", token.Of),
		Entry("property", "property", token.Property),
		Entry("this", "this", token.This),
	)
})
