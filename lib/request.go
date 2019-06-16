package acmdl

import (
	"net/http"
	"net/url"
)

type Request struct {
	Q *Query
}

func NewRequest() *Request {
	q := NewQuery()
	return &Request{Q: q}
}

func (r *Request) Send() ([]Result, error) {
	u := url.URL{
		Scheme:   "https",
		Host:     "dl.acm.org",
		Path:     "results.cfm",
		RawQuery: r.Q.Encode(),
	}
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	results, err := ParseResponse(resp)
	if err != nil {
		return nil, err
	}

	return results, nil
}
