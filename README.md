# EnvTmpl

EnvTmpl is a simple tool to fill a go templates with variables from the environment.

## Installing

Go get

```shell
go get github.com/xxacc/envtmpl
```

## Usage

### Templating

EnvTmpl uses [go templates](https://golang.org/pkg/text/template/).
Environment variables are exposed under the `.Env` value, e.g. `{{.Env.VARIABLE}}`.  
EnvTmpl will scan the given directory for files ending in .tmpl and create their outputs with the
.tmpl extenstion stripped (sample.txt.tmpl -> sample.txt). It will ignore all files that don't
have the .tmpl extension. See [testdata](testdata) for examples

### Options

```shell
envtmpl [options] template_directory
-e, --env=".env": Path to the .env file to use
-o, --out="gen": Directory to put output files
-p, --prefix="ENVTMPL": Will assume all variables to will be found at <prefix>_<name>
```

- `-e, --env`: EnvTmpl will add any variables found a `.env` file to its environment
  before executing the templates. It will search for a file called `.env` in the
  current directory by default. Variables in the `.env` file should be specified
  as `VARIABLE=value` with the prefix included in the name.
- `-o, --out`: EnvTmpl will put generated files into a directory called `gen` by default.
- `-p, --prefix`: EnvTmpl will only use variables with the prefix `ENVTMPL` by default.
  Template values should not include the prefix, but variables defined in the `.env` file must.  
  If the prefix is set to "", EnvTmpl will expose all you environment variables to the template.

## Todo

- Add tests!
- Improve template file handling. Currently will only work if given a single directory.
  Allow combination of single files and directories.
- Remove requirement for variables in the `.env.` file to include the prefix.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/xxacc/envtmpl/tags).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
