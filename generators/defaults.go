// Copyright 2015 The go-vector Authors
// This file is part of the go-vector library.
//
// The go-vector library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-vector library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-vector library. If not, see <http://www.gnu.org/licenses/>.

//go:generate go run defaults.go default.json defs.go

package main //build !none

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func fatal(str string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, str, v...)
	os.Exit(1)
}

type setting struct {
	Value   int64  `json:"v"`
	Comment string `json:"d"`
}

func main() {
	if len(os.Args) < 3 {
		fatal("usage %s <input> <output>\n", os.Args[0])
	}

	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fatal("error reading file %v\n", err)
	}

	m := make(map[string]setting)
	json.Unmarshal(content, &m)

	filepath := filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "vector", "go-vector", "params", os.Args[2])
	output, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm /*0777*/)
	if err != nil {
		fatal("error opening file for writing %v\n", err)
	}

	output.WriteString(`// DO NOT EDIT!!!
// AUTOGENERATED FROM generators/defaults.go

package params

import "math/big"

var (
`)

	for name, setting := range m {
		output.WriteString(fmt.Sprintf("%s=big.NewInt(%d) // %s\n", strings.Title(name), setting.Value, setting.Comment))
	}

	output.WriteString(")\n")
	output.Close()

	cmd := exec.Command("gofmt", "-w", filepath)
	if err := cmd.Run(); err != nil {
		fatal("gofmt failed: %v\n", err)
	}
}
