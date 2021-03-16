package main

import(
    "fmt"
    "os"
    "net/http"
    "strings"
    "github.com/akamensky/argparse"
	"github.com/olekukonko/tablewriter"
    "github.com/fatih/color"
    "github.com/common-nighthawk/go-figure"
    . "github.com/redcode-labs/Coldfire"
)

func find_redirects(url string){
    if !strings.Contains(url, "https://") && !strings.Contains(url, "http:"){
        url = "https://"+url
    }
    table_data := [][]string{}
    next_url := url
    for i := 0; i < 50; i++{
        redir := &http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
            return http.ErrUseLastResponse }}
        res, err := redir.Get(next_url)
        if err != nil{
            PrintError(err.Error())
        }
        next_url = res.Header.Get("Location")
        sc := IntToStr(res.StatusCode)
        url_color := Red
        u := next_url
        if sc == "200"{
            url_color = Green
            u = url
        }
        id := IntToStr(i)
        table_data = append(table_data, []string{id, url_color(u), sc})
        if sc == "200"{
            break
        }
    }
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "URL", "STATUS CODE"})
	table.SetAutoWrapText(false)
	table.SetCenterSeparator("*")
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetRowSeparator("-")
	for v := range table_data {
		table.Append(table_data[v])
	}
    if len(table_data) != 0{
        fmt.Println("")
        PrintInfo("URL => "+url)
	    table.Render()
        fmt.Println("")
    }
}

func print_banner(){
    banner := figure.NewFigure("UnChain", "", true)
    color.Set(color.FgMagenta)
	fmt.Println("")
    banner.Print()
    color.Unset()
	fmt.Println("")
    fmt.Println(F("\tCreated by: redcodelabs.io "+Red("<*>")))
	fmt.Println("")
}

func main(){
    print_banner()
    parser := argparse.NewParser("unchain", "")
    var URLS *string = parser.String("u", "url", &argparse.Options{Required: true, Help: "File containing urls or a single url"})
    err := parser.Parse(os.Args)
    ExitOnError(err)
    urls := []string{}
    if Exists(*URLS){
        urls = FileToSlice(*URLS)
    } else{
        urls = []string{*URLS}
    }
    for _, u := range urls{
        find_redirects(u)
    }
}