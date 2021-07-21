package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/akamensky/argparse"
	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	tw "github.com/olekukonko/tablewriter"
	cf "github.com/redcode-labs/Coldfire"
)

func findRedirects(url string) {
	if !strings.Contains(url, "https://") && !strings.Contains(url, "http:") {
		url = "https://" + url
	}
	var tableData [][]string
	nextUrl := url
	for i := 0; i < 50; i++ {
		redir := &http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}}
		res, err := redir.Get(nextUrl)
		if err != nil {
			cf.PrintError(err.Error())
		}
		nextUrl = res.Header.Get("Location")
		sc := cf.IntToStr(res.StatusCode)
		urlColor := cf.Red
		u := nextUrl
		if sc == "200" {
			urlColor = cf.Green
			u = url
		}
		id := cf.IntToStr(i)
		tableData = append(tableData, []string{id, urlColor(u), sc})
		if sc == "200" {
			break
		}
	}
	table := tw.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "URL", "STATUS CODE"})
	table.SetAutoWrapText(false)
	table.SetCenterSeparator("*")
	table.SetAlignment(tw.ALIGN_CENTER)
	table.SetRowSeparator("-")
	for v := range tableData {
		table.Append(tableData[v])
	}
	if len(tableData) != 0 {
		fmt.Println("")
		cf.PrintInfo("URL => " + url)
		table.Render()
		fmt.Println("")
	}
}

func printBanner() {
	banner := figure.NewFigure("UnChain", "", true)
	color.Set(color.FgMagenta)
	fmt.Println("")
	banner.Print()
	color.Unset()
	fmt.Println("")
	fmt.Println("\tCreated by: redcodelabs.io " + cf.Red("<*>"))
	fmt.Println("")
}

func main() {
	printBanner()
	parser := argparse.NewParser("unchain", "")
	var URLS = parser.String("u", "url", &argparse.Options{Required: true, Help: "File containing urls or a single url"})
	err := parser.Parse(os.Args)
	cf.ExitOnError(err)
	var urls []string
	if cf.Exists(*URLS) {
		urls = cf.FileToSlice(*URLS)
	} else {
		urls = []string{*URLS}
	}
	for _, u := range urls {
		findRedirects(u)
	}
}
