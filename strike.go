package main

import(
    "github.com/codegangsta/cli"
    "os"
    "net/http"
    "net/url"
    "io/ioutil"
    "log"
    "fmt"
    "encoding/json"
    "strings"
    "strconv"
)

type ApiResponse struct {
    Results uint
    Statuscode uint
    Responsetime float64
    Torrents []Torrent
}

type Torrent struct {
    Torrent_title string
    Torrent_hash string
    Torrent_category string
    Sub_category string
    Seeds int
    Leeches int
    File_count int
    Size int
    Upload_date string
    Uploader_username string
    File_info struct {
        File_names []string
        File_lengths []int
    }
    Magnet_uri string
}

func main() {
    app := cli.NewApp()
    app.Name = "strikecli"
    app.Usage = "Use getstrike search from your terminal"
    app.Version = "0.0.1"
    app.Flags = []cli.Flag {
        cli.StringFlag {
            Name: "category, c",
            Value: "",
            Usage: "filter by category",
        },
        cli.StringFlag {
            Name: "format, f",
            Usage:
            `set the output formatting, by matching the following sequences:
             %t: the torrent title
             %m: the torrent magnet link
             %h: the torrent hash
             %S: the number of seeders
             %l: the number of leechers
             %s: the torrent size
             %d: the torrent upload date
             %c: the torrent category
             %sc: the torrent sub category
             %f: the torrent file count
             %u: the uploader username`,
            Value: "%t",
        },
        cli.IntFlag {
            Name: "limit, l",
            Usage: "set the maximum output torrents",
            Value: 100,
        },
    }
    app.Commands = []cli.Command {
        {
            Name: "search",
            Usage: "Search for the given keywords",
            Action: search,
            Flags: app.Flags,
        },
        {
            Name: "info",
            Usage: "Get informations on the given hash",
            Action: info,
            Flags: app.Flags,
        },
    }
    app.Run(os.Args)
}

func search(c *cli.Context) {
    query := url.QueryEscape(strings.Join(c.Args()," "))
    uri := "https://getstrike.net/api/v2/torrents/search/?phrase=" + query
    if c.String("category") != "" {
        uri = uri + "&category=" + url.QueryEscape(c.String("category"))
    }
    data := request(uri)
    var obj ApiResponse
    err := json.Unmarshal(data,&obj)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Your search for \"%s\" returned %d results:\n",strings.Join(c.Args()," "),obj.Results)
    for i,item := range obj.Torrents {
        if(i >= c.Int("limit")) { break }
        println(formatString(item,c.String("format")))
    }
}

func info(c *cli.Context) {
    uri := "https://getstrike.net/api/v2/torrents/info/?hashes=" + strings.Join(c.Args(),",")
    data := request(uri)
    var obj ApiResponse
    err := json.Unmarshal(data,&obj)
    if err != nil {
        log.Fatal(err)
    }
    for _,item := range obj.Torrents {
        println(formatString(item,c.String("format")))
    }
}

func request(url string) []byte {
    res,err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    data,err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    if err != nil {
        log.Fatal(err)
    }
    return data
}

func formatString(torrent Torrent,format string) string {
    str := format
    str = strings.Replace(str,"%t",torrent.Torrent_title,-1)
    str = strings.Replace(str,"%m",torrent.Magnet_uri,-1)
    str = strings.Replace(str,"%h",torrent.Torrent_hash,-1)
    str = strings.Replace(str,"%S",strconv.Itoa(torrent.Seeds),-1)
    str = strings.Replace(str,"%l",strconv.Itoa(torrent.Leeches),-1)
    str = strings.Replace(str,"%s",strconv.Itoa(torrent.Size),-1)
    str = strings.Replace(str,"%d",torrent.Upload_date,-1)
    str = strings.Replace(str,"%c",torrent.Torrent_category,-1)
    str = strings.Replace(str,"%sc",torrent.Sub_category,-1)
    str = strings.Replace(str,"%f",strconv.Itoa(torrent.File_count),-1)
    str = strings.Replace(str,"%u",torrent.Uploader_username,-1)
    return str
}
