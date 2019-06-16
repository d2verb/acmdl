package main

import (
	"fmt"
	"strings"

	acmdl "github.com/d2verb/acmdl/lib"
	"github.com/urfave/cli"
)

func ExecCLI(c *cli.Context) error {
	words := c.String("words")
	year := c.Int("year")
	start := c.Int("start")

	req := acmdl.NewRequest()
	for _, word := range strings.Split(words, ",") {
		req.Q.AddWord(word)
	}
	req.Q.Year = year
	req.Q.Start = start - 1

	results, err := req.Send()
	if err != nil {
		fmt.Errorf("%s", err)
		return err
	}

	Pprint(results, start)
	return nil
}
