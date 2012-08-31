package main

import (
	"go/parser"
	//"go/printer"
	"go/token"
	"go/ast"
	
	"io"
	"os"
	"os/exec"
	"strings"
	
	"path/filepath"
	
	"fmt"
)

var fs *token.FileSet

var imports = NewImportSet()

func main() {
	dirs := os.Args[1:]
	fs = token.NewFileSet()
	
	//dirs := parseInput(*dir)
	fmt.Println(dirs)
	
	for _, cd := range dirs {
		fmt.Println(cd)
		nodes, err := parser.ParseDir(fs, cd, sourceFilter, 0)
		
		if err != nil {
			continue
			//fmt.Printf("Error while parsing files in \"%s\".: %s\n", cd, err)
		}
		
		for _, node := range nodes {
			ast.Inspect(node, RootHandeler)
		}
	}
	
	fmt.Println(imports)
	imports.GetImports()

	//printer.Fprint(os.Stdout, fs, node)
}

func sourceFilter(fi os.FileInfo) bool {
	return strings.HasSuffix(fi.Name(), ".go")
}

func parseInput(dir string) []string {
	fmt.Println(dir)
	dir = filepath.Clean(dir)
	dirs, err := filepath.Glob(dir)
	if err != nil {
		return make([]string, 0)
	}
	return dirs
}

func RootHandeler(n ast.Node) bool {
    switch x := n.(type) {
    case *ast.ImportSpec:
		ast.Inspect(x, ImportHandeler)
    case *ast.FuncDecl:
		//ast.Inspect(x, ImportHandeler)
    }
    return true
}

func ImportHandeler(n ast.Node) bool {
	var s string
	switch x := n.(type) {
		case *ast.BasicLit:
			s = "[Import BasicLit] \t" + x.Value
			i := strings.Replace(x.Value, "\"", "", -1)
			imports.Put(i)
	}
	if s != "" {
		//fmt.Printf("%s:\t\t%s\n", fs.Position(n.Pos()), s)
	}
	return true
}

func (imp *ImportSet) GetImports() {
		l := append([]string{"get"}, imp.set...)
		cmd := exec.Command("go", l...)
		cmdOut, _ := cmd.StdoutPipe()
		
		fmt.Println("Running command:", cmd.Args)
		go io.Copy(os.Stdout, cmdOut)
		err := cmd.Run()
		
		if err != nil  {
			fmt.Println(err)
		}
}
