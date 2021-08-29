package providergen

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"go/types"
	"golang.org/x/tools/go/packages"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

//go:embed templates/*.tmpl
var templateFS embed.FS

var tmpls = template.Must(template.ParseFS(templateFS, "templates/*"))

func GenerateProvider(cmd *cobra.Command, args []string) {
	// Load the go package
	pkgPath, typeName := parseSourceType(args[0])
	pkg, err := loadPackage(pkgPath)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	// check if type exists
	obj := pkg.Types.Scope().Lookup(typeName)
	if _, ok := obj.(*types.TypeName); !ok {
		log.Fatalf("%s is not a named type", typeName)
		return
	}

	// check if type is a struct
	t, ok := obj.Type().Underlying().(*types.Struct)
	if !ok {
		log.Fatalf("%s is not a struct", typeName)
		return
	}

	// get the package name for the resource
	pkgParts := strings.Split(pkg.String(), "/")

	// parse the resource
	var resource *Resource
	resource, err = ResourceFromStruct(typeName, pkgParts[len(pkgParts)-1], t)
	if err != nil {
		log.Fatal(err.Error())
	}



	// render the provider template
	err = render(resource)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

func loadPackage(path string) (*packages.Package, error) {
	cfg := &packages.Config{Mode: packages.NeedTypes | packages.NeedImports, Env:  append(os.Environ(), "CGO_ENABLED=0")}
	pkgs, err := packages.Load(cfg, path)
	if err != nil {
		return nil, err
	}

	if packages.PrintErrors(pkgs) > 0 {
		os.Exit(1)
	}

	return pkgs[0], nil
}

//parseSourceType returns the package path and type name
func parseSourceType(sourceType string) (string, string) {
	idx := strings.LastIndexByte(sourceType, '.')
	if idx == -1 {
		return "", ""
	}

	return sourceType[0:idx], sourceType[idx+1:]
}

func render(data *Resource) error {
	tmplBuf := bytes.Buffer{}
	fmtBuf := bytes.Buffer{}
	importBuf := bytes.Buffer{}
	errBuf := bytes.Buffer{}
	filename := fmt.Sprintf("zz_%s_provider.go", data.Name)

	fmtC := exec.Command("gofmt")
	fmtC.Stdin = &tmplBuf
	fmtC.Stdout = &fmtBuf
	fmtC.Stderr = &errBuf

	importC := exec.Command("goimports")
	importC.Stdin = &fmtBuf
	importC.Stdout = &importBuf
	importC.Stderr = &errBuf

	if err := tmpls.ExecuteTemplate(&tmplBuf, "root.tmpl", data); err != nil {
		return err
	}

	if err := fmtC.Run(); err != nil {
		return errors.New(string(errBuf.Bytes()))
	}

	if err := importC.Run(); err != nil {
		return errors.New(string(errBuf.Bytes()))
	}

	dest, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() { _ = dest.Close() }()

	_, err = io.Copy(dest, &importBuf)
	if err != nil {
		return err
	}

	println(filename)

	return nil
}
