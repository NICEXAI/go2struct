package util

import (
	"bytes"
	"go/format"
	"go/parser"
	"go/token"
	"strings"
	"unicode"
)

// UnderscoreToUpperCamelCase word change from underscore to upperCamelCase
func UnderscoreToUpperCamelCase(n string) string {
	var cList []string

	for _, cell := range strings.Split(n, "_") {
		cList = append(cList, strings.Title(cell))
	}

	return strings.Join(cList, "")
}

// UpperCamelCaseToUnderscore word change from upperCamelCase to underscore
func UpperCamelCaseToUnderscore(n string) string {
	var (
		cList   []rune
		isUpper bool
	)

	for i, cell := range n {
		if unicode.IsUpper(cell) {
			if i != 0 && !isUpper {
				cList = append(cList, '_')
			}
			cList = append(cList, unicode.ToLower(cell))
			isUpper = true
		} else {
			cList = append(cList, cell)
			isUpper = false
		}
	}

	return string(cList)
}

// FormatGoStruct format the go struct text
func FormatGoStruct(txt string) (string, error) {
	buffer := bytes.NewBufferString("")
	choreTxt := "package main\n\n"
	newTxt := choreTxt + string(txt)

	f, err := parser.ParseFile(token.NewFileSet(), "", []byte(newTxt), 0)
	if err != nil {
		return "", err
	}

	if err = format.Node(buffer, token.NewFileSet(), f); err != nil {
		return "", err
	}

	structTxt := strings.Replace(buffer.String(), choreTxt, "", 1)
	return structTxt, nil
}
