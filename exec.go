package main

import (
	"fmt"
	"strings"
	"time"

	acmdl "github.com/d2verb/acmdl/lib"
	"github.com/fatih/color"
	"github.com/urfave/cli"
)

func idleRoutine(ch chan int) {
	c := color.New(color.FgCyan).Add(color.Bold)
	roller := []string{"[-]", "[\\]", "[/]"}
	i := 0
	for {
		select {
		case _ = <-ch:
			ch <- 1
			break
		default:
			c.Printf("%s Wait a minute...\r", roller[i])
			time.Sleep(time.Second / 2)
			i = (i + 1) % len(roller)
		}
	}
	fmt.Printf("\r\n")
}

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

	ch := make(chan int)
	go idleRoutine(ch)

	results, err := req.Send()
	ch <- 1
	_ = <-ch

	if err != nil {
		fmt.Errorf("%s", err)
		return err
	}

	Pprint(results, start)
	return nil
}
