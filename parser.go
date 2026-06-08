package fhtml

import (
	"bytes"
	"io"
	"strings"

	"golang.org/x/net/html"
)

var _ io.Reader = (*Parser)(nil)

type Parser struct {
	reader io.Reader

	variableSource Valuable

	buffer    bytes.Buffer
	tokenizer html.Tokenizer

	attributes      []html.Attribute
	attributesByKey map[string]*html.Attribute
}

func New(reader io.Reader) *Parser {
	return &Parser{
		reader:    reader,
		buffer:    bytes.Buffer{},
		tokenizer: *html.NewTokenizer(reader),
	}
}

func (p *Parser) SetVariablesMap(variables map[string]any) {
	p.variableSource = ValuableMap(variables)
}

func (p *Parser) SetVariablesContext(context ValuableContext) {
	p.variableSource = ValuableContext(context)
}

func (p *Parser) Read(b []byte) (n int, err error) {
	for {
		switch p.tokenizer.Next() {
		case html.StartTagToken, html.SelfClosingTagToken:
			token := p.tokenizer.Token()
			_ = token
		case html.TextToken:

		default:

		}
	}
}

func (p *Parser) processAttribute(attr html.Attribute) {
	var key string = attr.Key
	var values []string
	var flags []string

	if strings.HasPrefix(key, ":") {
		key = strings.TrimPrefix(key, ":")

		// it's a computed attribute; need to process the value
	} else {
		values = append(values, attr.Val)
	}

	key, flags = parseAttributeKeyFlags(key)

	if _, ok := p.attributesByKey[key]; !ok {
		p.attributes = append(p.attributes, attr)
		p.attributesByKey[key] = &p.attributes[len(p.attributes)-1]
	}

	p.processAttributeKeyFlags(key, flags, values)
}

func (p *Parser) processAttributeKeyFlags(key string, flags []string, values []string) {
	if _, ok := p.attributesByKey[key]; !ok {
		panic("failed assertion: attribute should've been processed before")
	}

	if flags[0] == "replace" {
		sep := " "

		if flags[0] == "comma" {
			sep = ","
		} else if flags[0] == "semicolon" {
			sep = ";"
		}

		if flags[1] == "space" {
			sep += " "
		}

		p.attributesByKey[key].Val = strings.Join(values, sep)
		return
	}

	{ // append
		sep := " "

		if flags[0] == "comma" {
			sep = ","
		} else if flags[0] == "semicolon" {
			sep = ";"
		}

		if flags[1] == "space" {
			sep += " "
		}

		p.attributesByKey[key].Val = strings.Join(values, sep)
	}
}

func parseAttributeKeyFlags(key string) (string, []string) {
	flags := strings.Split(key, ":")
	key = flags[0]
	flags = flags[1:]

	return key, flags
}
