package main

import (
	"bytes"
	"fmt"
	"github.com/richardjennings/liquify"
	"github.com/richardjennings/liquify/expr"
	"github.com/richardjennings/liquify/parser"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var transpileCmd = &cobra.Command{
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		trans := args[0]
		fin := args[1]
		fout := args[2]

		switch trans {
		case "php":
			b, err := os.ReadFile(fin)
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
			cnf := parser.Config{
				Delims:  []string{"{{", "}}", "{%", "%}"},
				Grammar: parser.Grammer{},
			}
			l, err := liquify.Liquify(b, cnf)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			o, err := liquify.PHP{
				TagParsers: map[string]func(b *bytes.Buffer, t *parser.ASTTag, p liquify.PHP) error{
					"t": func(b *bytes.Buffer, t *parser.ASTTag, p liquify.PHP) error {
						v, err := expr.Parse(t.Args)
						if err != nil {
							panic(err)
						}
						b.Write([]byte(fmt.Sprintf(`<?php echo %s(%s);?>`, t.Name, p.Stmt(v))))
						return nil
					},
				},
			}.Transpile(l)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fh, err := os.OpenFile(fout, os.O_TRUNC|os.O_RDWR|os.O_CREATE, 0600)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			defer fh.Close()
			if _, err := fh.Write(o); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		default:
			log.Println("only php transpiler currently supported")
			os.Exit(1)
		}
	},
}

func main() {
	if err := transpileCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
