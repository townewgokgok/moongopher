package moongopher

import (
	"fmt"
	"go/ast"
	"log"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

type generator struct {
	inspector *inspector.Inspector
	indent    string
}

func newGenerator(pass *analysis.Pass) *generator {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	return &generator{
		inspector: inspect,
	}
}

func (g *generator) lf() {
	log.Println()
}

func (g *generator) printf(format string, args ...interface{}) {
	log.Printf(g.indent+format, args...)
}

func (g *generator) print(args ...interface{}) {
	log.Print(g.indent + fmt.Sprint(args...))
}

func (g *generator) enter() {
	g.indent += "\t"
}

func (g *generator) leave() {
	if len(g.indent) == 0 {
		return
	}
	g.indent = g.indent[1:]
}

func (g *generator) processSpec(s ast.Spec) {
	g.printf("%#v", s)
}

func (g *generator) processGenDecl(d *ast.GenDecl) {
	g.lf()
	g.processCommentGroup(d.Doc)
	g.printf("%s (", d.Tok)
	g.enter()
	for _, s := range d.Specs {
		g.processSpec(s)
	}
	g.leave()
	g.print(")")
}

func (g *generator) processBlockStmt(d *ast.BlockStmt) {
	for _, l := range d.List {
		g.printf("%#v", l)
	}
}

func (g *generator) processFuncDecl(d *ast.FuncDecl) {
	g.lf()
	g.processCommentGroup(d.Doc)
	g.printf("func %#v %s %#v {", d.Recv, d.Name.Name, d.Type)
	g.enter()
	g.processBlockStmt(d.Body)
	g.leave()
	g.print("}")
}

func (g *generator) processDecl(d ast.Decl) {
	switch d := d.(type) {
	case *ast.GenDecl:
		g.processGenDecl(d)
	case *ast.FuncDecl:
		g.processFuncDecl(d)
	default:
		g.lf()
		g.printf("? %#v", d)
	}
}

func (g *generator) processCommentGroup(c *ast.CommentGroup) {
	if c == nil {
		return
	}
	for _, l := range c.List {
		g.print(l.Text)
	}
}

func (g *generator) processFile(f *ast.File) {
	g.printf("// file: %s", f.Name.Name)
	for _, i := range f.Imports {
		g.processCommentGroup(i.Comment)
		g.processCommentGroup(i.Doc)
		g.printf("// import %s", i.Path.Value)
	}
	for _, d := range f.Decls {
		g.processDecl(d)
	}
}

func (g *generator) process() {
	filter := []ast.Node{
		(*ast.File)(nil),
	}
	g.inspector.Preorder(filter, func(n ast.Node) {
		g.processFile(n.(*ast.File))
	})
}
