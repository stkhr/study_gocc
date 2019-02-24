package complexity

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func GetFuncNode(t *testing.T, code string) ast.Node {
	t.Helper()

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", code, 0)
	if err != nil {
		t.Fatal(err)
	}
	for _, decl := range file.Decls {
		if fd, ok := decl.(*ast.FuncDecl); ok {
			return fd
		}
	}
	t.Fatal("no function declear found")
	return nil
}

func TestCyclomaticComplexity(t * testing.T) {
	testcases := []struct {
		name		string // name of test case
		code		string // code of test
		complexity	int    // expected number
	}{
		// test case
		{
			name: "test1",
			code: `package main
func Count(node ast.Node) int{
	count := 1
	return count
}`,
			complexity: 1,
		},
		{
			name: "test2",
			code: `package main
func Double(n int) int {
	if n%2 == 0 {
		return 0
	}
	return n
}`,
			complexity: 2,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			// 抽象構文木を取得する
			a := GetFuncNode(t, testcase.code)

			// 循環複雑度を取得する
			c := Count(a)

			// 算出した循環複雑度が期待したものか調べる
			if c != testcase.complexity {
				t.Errorf("got=%d, want=%d", c, testcase.complexity)
			}
		})
	}
}