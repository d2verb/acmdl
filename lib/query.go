package acmdl

import (
	"bytes"
	"net/url"
	"strconv"
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
		out.WriteString(word)
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
