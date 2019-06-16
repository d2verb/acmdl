package acmdl

import (
	"bytes"
	"net/url"
	"strconv"
	"strings"
)

type Query struct {
	Words []string
	Year  int
	Start int
}

type WordType string

func NewQuery() *Query {
	return &Query{Year: -1, Start: 0}
}

func (q *Query) AddWord(word string) {
	q.Words = append(q.Words, word)
}

func (q *Query) Encode() string {
	v := url.Values{}
	v.Set("query", q.encodeWords())
	v.Set("dte", q.encodeYear())
	v.Set("start", strconv.Itoa(q.Start))
	return v.Encode()
}

func (q *Query) encodeWords() string {
	var out bytes.Buffer

	out.WriteString("(")

	for i, word := range q.Words {
		if i != 0 {
			out.WriteString(" ")
		}

		// "+security" should be converted to "%2Bsecurity" before url encoding
		// So after url encoding "%2Bsecurity" will be "%252Bsecurity"
		out.WriteString(strings.Replace(word, "+", "%2B", -1))
	}

	out.WriteString(")")

	return out.String()
}

func (q *Query) encodeYear() string {
	if q.Year != -1 {
		return strconv.Itoa(q.Year)
	} else {
		return ""
	}
}
