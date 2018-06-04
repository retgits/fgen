package fgutil

import (
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"
)

func CreateFileFromTemplate(basePath string, fileName string, template string, data interface{}) error {
	filePath := path.Join(basePath, fileName)
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	RenderTemplate(f, template, data)
	f.Close()

	return nil
}

func CreateFileFromString(filePath string, data string) error {
	d1 := []byte(data)
	return ioutil.WriteFile(filePath, d1, 0644)
}

//RenderTemplate renders the specified template
func RenderTemplate(w io.Writer, text string, data interface{}) {
	t := template.New("top")
	t.Funcs(template.FuncMap{"trim": strings.TrimSpace, "capitalize": Capitalize})
	template.Must(t.Parse(text))
	if err := t.Execute(w, data); err != nil {
		panic(err)
	}
}

// Capitalize will return a copy of the provided string with the first letter capitalized
func Capitalize(s string) string {
	if s == "" {
		return s
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToTitle(r)) + s[n:]
}

// IsStringInList determines if the specified value is in the provided string list
func IsStringInList(str string, list []string) bool {
	for _, value := range list {
		if value == str {
			return true
		}
	}
	return false
}

// CaseInsensitiveContains will search in a string if a substring exists, reagardless of case.
// This isn't the most performant way, but this will be able to check if the string exists while ignoring any case sensitivity.
func CaseInsensitiveContains(s, substr string) bool {
	s, substr = strings.ToUpper(s), strings.ToUpper(substr)
	return strings.Contains(s, substr)
}

// LeftPad pads a string with a padStr for a certain number of times
func LeftPad(s string, padStr string, pLen int) string {
	return strings.Repeat(padStr, pLen) + s
}

// RightPadToLen pads a string to a certain length
func RightPadToLen(s string, padStr string, overallLen int) string {
	padCountInt := 1 + ((overallLen - len(padStr)) / len(padStr))
	retStr := s + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}
