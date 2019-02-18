package main

import (
	"log"

	"github.com/ogier/pflag"

	"github.com/joho/godotenv"
)

var (
	env    string
	out    string
	prefix string
)

func init() {
	pflag.StringVarP(&env, "env", "e", ".env", "Path to the .env file to use")
	pflag.StringVarP(&out, "out", "o", "gen", "Directory to put output files")
	pflag.StringVarP(&prefix, "prefix", "p", "", "Will assume all variables to will be found at <prefix>_<name>")
}

func main() {
	pflag.Parse()
	err := godotenv.Load(env)
	if err != nil {
		log.Print("Not using .env file: ", err)
	}
	tmplVals := loadTmplValues(prefix)
	for _, a := range pflag.Args() {
		if err := fillDir(a, out, tmplVals); err != nil {
			log.Fatal(err)
		}
	}
}
