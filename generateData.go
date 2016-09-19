package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/doc"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type Data struct {
	Data []Mapping
}

type Mapping struct {
	FFmpeg string             `json:"ffmpeg"`
	GoRep  []GoRepresentation `json:"goRep"`
}

type GoRepresentation struct {
	Name string `json:"name"`
	Doc  string `json:"doc"`
}

var cRepMatcher *regexp.Regexp

func init() {
	cRepMatcher = regexp.MustCompile("C-(?:Function|Attribute|Method|Member|Field): ([a-zA-Z0-9_.:->]+)")
}

func extractCRepresentation(str string) (string, bool) {
	matches := cRepMatcher.FindStringSubmatch(str)
	if len(matches) < 2 {
		return "", false
	}
	return matches[1], true
}

func handleFunction(f *doc.Func, namePrefix string, mping map[string]Mapping) map[string]Mapping {
	cRep, ok := extractCRepresentation(f.Doc)
	if !ok {
		return mping
	}

	m, ok := mping[cRep]
	if !ok {
		m = Mapping{
			FFmpeg: cRep,
			GoRep:  make([]GoRepresentation, 0, 1),
		}
	}

	goName := namePrefix + f.Name
	var buf bytes.Buffer
	doc.ToHTML(&buf, f.Doc, make(map[string]string))
	m.GoRep = append(m.GoRep, GoRepresentation{Name: goName, Doc: buf.String()})
	mping[cRep] = m

	return mping
}

func parseDir(path string, namePrefix string, mping map[string]Mapping) (map[string]Mapping, error) {
	fset := token.NewFileSet()
	packages, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	for pkgName, tree := range packages {
		pkg := doc.New(tree, pkgName, 0)
		for _, f := range pkg.Funcs {
			mping = handleFunction(f, namePrefix+pkgName+".", mping)
		}
		for _, t := range pkg.Types {
			for _, f := range t.Methods {
				mping = handleFunction(f, namePrefix+pkgName+"."+t.Name+"::", mping)
			}
		}
	}

	return mping, nil
}

func parseDirRecursive(path string, namePrefix string, mping map[string]Mapping) (map[string]Mapping, error) {
	var err error
	mping, err = parseDir(path, namePrefix, mping)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if !f.IsDir() || f.Name()[0] == '.' {
			continue
		}
		pref := ""
		i := strings.LastIndex(path, string(os.PathSeparator))
		if i == -1 {
			pref = path
		} else {
			pref = path[i+1:]
		}
		mping, err = parseDirRecursive(path+string(os.PathSeparator)+f.Name(), pref+"/"+namePrefix, mping)
		if err != nil {
			return nil, err
		}
	}
	return mping, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run generateData.go <path>")
		os.Exit(1)
		return
	}
	path := os.Args[1]

	mping, err := parseDirRecursive(strings.TrimRight(path, string(os.PathSeparator)), "", make(map[string]Mapping))
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
		return
	}

	var out []Mapping

	for _, val := range mping {
		out = append(out, val)
	}

	outTxt, err := json.Marshal(out)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
		return
	}

	fmt.Print("DATA = ")
	fmt.Print(string(outTxt))
	fmt.Println(";")
}
