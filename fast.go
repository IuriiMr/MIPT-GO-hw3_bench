package main

import (
	"bufio"
	json "encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

type User struct {
	Browsers []string
	Name     string
	Email    string
}

// вам надо написать более быструю оптимальную этой функции
func FastSearch(out io.Writer) {

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	browsers := make(map[string]bool, 120)
	i := -1

	fmt.Fprintln(out, "found users:")
	scanner := bufio.NewScanner(file)
	buf := make([]byte, 1024)
	scanner.Buffer(buf, bufio.MaxScanTokenSize)

	data := &User{}

	for scanner.Scan() {
		i++
		line := scanner.Text()
		//err := json.Unmarshal([]byte(line), &data)
		err = data.UnmarshalJSON([]byte(line))
		if err != nil {
			panic(err)
		}

		isAndroid := false
		isMsie := false
		for _, browser := range data.Browsers {

			if strings.Contains(browser, "Android") {
				isAndroid = true
				browsers[browser] = true
			} else if strings.Contains(browser, "MSIE") {
				isMsie = true
				browsers[browser] = true
			}
		}

		if !(isAndroid && isMsie) {
			continue
		}

		email := strings.Replace(data.Email, "@", " [at] ", 1)
		fmt.Fprintf(out, "[%d] %s <%s>\n", i, data.Name, email)

	}

	fmt.Fprintln(out) // blank newline
	fmt.Fprintln(out, "Total unique browsers", len(browsers))
}

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson3486653aDecodeCourseraGoWsHw3Bench(in *jlexer.Lexer, out *User) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "email":
			out.Email = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "browsers":
			if in.IsNull() {
				in.Skip()
				out.Browsers = nil
			} else {
				in.Delim('[')
				if out.Browsers == nil {
					if !in.IsDelim(']') {
						out.Browsers = make([]string, 0, 4)
					} else {
						out.Browsers = []string{}
					}
				} else {
					out.Browsers = (out.Browsers)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Browsers = append(out.Browsers, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3486653aDecodeCourseraGoWsHw3Bench(&r, v)
	return r.Error()
}
