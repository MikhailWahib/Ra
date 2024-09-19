package main

import (
	"Ra/evaluator"
	"Ra/lexer"
	"Ra/object"
	"Ra/parser"
	"Ra/repl"
	"fmt"
	"io"
	"os"
	"os/user"
)

func startRepl() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Ra 𓋹 programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Woops! We ran into some issues!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

func executeFile(args []string) {
	path := args[1]

	// check if the file end with .ra
	if path[len(path)-3:] != ".ra" {
		panic("File must end with .ra")
	}

	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	env := object.NewEnvironment()

	l := lexer.New(string(file))
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(os.Stdout, p.Errors())
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		io.WriteString(os.Stdout, evaluated.Inspect())
		io.WriteString(os.Stdout, "\n")
	}
}

func main() {
	args := os.Args

	if len(args) < 2 {
		startRepl()
	}

	if len(args) > 2 {
		panic("Wrong number of arguments!")
	}

	executeFile(args)
}
