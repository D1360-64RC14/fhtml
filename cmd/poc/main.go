package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

var variables = map[string]string{
	"variable": "has a value!",
	"greeting": "Hello World!",
}

var h = `
<html>
    <head>
        <title>Test Templating</title>
    </head>

    <body :param="variable">
        <h1>{{greeting}}</h1>
    </body>
</html>
`

func main() {
	t := html.NewTokenizer(strings.NewReader(h))

	for ttype := t.Next(); ttype != html.ErrorToken; ttype = t.Next() {
		switch ttype {
		case html.StartTagToken, html.SelfClosingTagToken:
			tk := t.Token()

			newAttrs := make([]html.Attribute, 0, len(tk.Attr))
			for i := range tk.Attr {
				if strings.HasPrefix(tk.Attr[i].Key, ":") {
					if v, ok := variables[tk.Attr[i].Val]; ok {
						newAttrs = append(newAttrs, html.Attribute{"", strings.TrimPrefix(tk.Attr[i].Key, ":"), v})
					}
				} else {
					newAttrs = append(newAttrs, tk.Attr[i])
				}
			}

			tk.Attr = newAttrs

			fmt.Printf("%s", tk)
		case html.TextToken:
			textBytes := t.Raw()
			textBytes = regexp.MustCompile("{{\\w+}}").ReplaceAllFunc(textBytes, func(b []byte) []byte {
				b = bytes.TrimPrefix(b, []byte("{{"))
				b = bytes.TrimSuffix(b, []byte("}}"))
				b = bytes.Trim(b, " ")

				if v, ok := variables[string(b)]; ok {
					return []byte(v)
				}

				return []byte{}
			})

			fmt.Printf("%s", textBytes)
		default:
			fmt.Printf("%s", t.Raw())
		}
	}
}
