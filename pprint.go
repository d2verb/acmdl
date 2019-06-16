package main

import (
	"fmt"

	acmdl "github.com/d2verb/acmdl/lib"
	"github.com/fatih/color"
)

func Pprint(results []acmdl.Result, numBase int) {
	numColor := color.New(color.FgCyan).Add(color.Bold)
	titleColor := color.New(color.FgGreen).Add(color.Bold)
	keywordsColor := color.New(color.FgYellow).Add(color.Bold)
	urlColor := color.New(color.FgMagenta).Add(color.Bold)

	for i, result := range results {
		numColor.Printf("%-2d ", i+numBase)
		titleColor.Printf("%s\n", result.Title)
		if result.Keywords == "" {
			keywordsColor.Printf("Keywords:\n")
		} else {
			keywordsColor.Printf("%s\n", result.Keywords)
		}
		urlColor.Printf("URL: %s\n", result.Url)
		fmt.Printf("%s\n\n", result.Abstract)
	}
}
