package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/pkg/errors"
)

var tmplExt = ".tmpl"

func fillDir(in, out string, vals tmplValues) (err error) {
	in, err = filepath.Abs(in)
	if err != nil {
		return errors.Wrap(err, "Failed to find absolute path for "+in)
	}
	out, err = filepath.Abs(out)
	if err != nil {
		return errors.Wrap(err, "Failed to find absolute path for "+out)
	}
	return filepath.Walk(in, func(name string, finfo os.FileInfo, err error) error {
		if err != nil {
			return errors.Wrap(err, "Error walking input directory "+in)
		}
		if finfo.IsDir() {
			return nil
		}
		if filepath.Ext(name) != tmplExt {
			return nil
		}
		if err := fillTmpl(name, createOutputName(in, name, out), vals); err != nil {
			return errors.Wrap(err, "Error filling template")
		}
		return nil
	})
}

func fillTmpl(in, out string, vals tmplValues) error {
	t, err := template.ParseFiles(in)
	if err != nil {
		return errors.Wrap(err, "Error parsing template "+in)
	}
	outF, err := os.Create(out)
	if err != nil {
		log.Print("Mkdir", filepath.Dir(out))
		if err := os.MkdirAll(filepath.Dir(out), 0777); err != nil {
			return errors.Wrap(err, "Error creating directory structure "+filepath.Dir(out))
		}
		outF, err = os.Create(out)
		if err != nil {
			return errors.Wrap(err, "Failed to create output file")
		}
	}
	if err := t.Execute(outF, vals); err != nil {
		return errors.Wrap(err, "Failed to fill template")
	}
	return nil
}

func createOutputName(in, name, out string) string {
	outputName := filepath.Join(out, strings.TrimPrefix(name, in))
	return strings.TrimSuffix(outputName, filepath.Ext(outputName))
}
