package main

import (
	"flag"
	"strings"
	"github.com/antchfx/htmlquery"
	"fmt"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type Cmd struct {
	helpFlag bool
	versionFlag bool
	locationFlag bool
	reverseFlag bool
	option string
	args []string
}

type Response struct {
	Code int `json:"code"`
	Data struct {
		Area      string `json:"area"`
		AreaID    string `json:"area_id"`
		City      string `json:"city"`
		CityID    string `json:"city_id"`
		Country   string `json:"country"`
		CountryID string `json:"country_id"`
		County    string `json:"county"`
		CountyID  string `json:"county_id"`
		IP        string `json:"ip"`
		Isp       string `json:"isp"`
		IspID     string `json:"isp_id"`
		Region    string `json:"region"`
		RegionID  string `json:"region_id"`
	} `json:"data"`
}

func get(url string) string {
	client := &http.Client{Timeout: 5 * time.Second}
	req, _ := http.NewRequest("GET", url, nil)
    req.Header.Add("User-Agent","Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.79 Safari/537.36")
	resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    var buffer [512]byte
    result := bytes.NewBuffer(nil)
    for {
        n, err := resp.Body.Read(buffer[0:])
        result.Write(buffer[0:n])
        if err != nil && err == io.EOF {
            break
        } else if err != nil {
            panic(err)
        }
    }
    return result.String()
}

func getIPInfo(ip string) {
	data := get("http://ip.taobao.com/service/getIpInfo.php?ip="+ip)

	var resp Response
	err := json.Unmarshal([]byte(data), &resp)
	if err != nil {
		fmt.Println("Parse error: ", err)
	}
	fmt.Println(resp.Data.IP, ": ", resp.Data.Country, "-",resp.Data.Region,"-",resp.Data.City, "（", resp.Data.Isp, "）")
}

func reverseDomainParse(ip string) {
	url := "https://site.ip138.com/"+ip+"/"
	resp := get(url)
	doc, err := htmlquery.Parse(strings.NewReader(resp))
	if err != nil {
		panic(err)
	}
	fmt.Println(ip, "上的网站，绑定过的域名如下：\n")
	for idx, node := range htmlquery.Find(doc, "//*[@id='list']/li") {
		if idx > 1 {
			date := htmlquery.InnerText(htmlquery.FindOne(node, "/span"))
			domain := htmlquery.InnerText(htmlquery.FindOne(node, "/a"))

			var buf bytes.Buffer
			for i:=0; i < 60 - 25 - len(domain); i++ {
				buf.WriteString(" ")
			}
			fmt.Println(domain,buf.String(), date)
		}
	}
}

func printUsage() {
	fmt.Printf("Usage: ipinfo [-options] [args...]\n")
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage

	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.BoolVar(&cmd.locationFlag, "l", false, "print version and exit")
	flag.BoolVar(&cmd.reverseFlag, "r", false, "print version and exit")
	flag.Parse()
	args := flag.Args()
	cmd.args = args
	return cmd
}

func main() {
	cmd := parseCmd()
	var ip string

	if len(cmd.args) == 0 {
		ip = "myip"
	} else {
		ip = cmd.args[0]
	}
	
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if (cmd.helpFlag) {
		printUsage()
	} else if cmd.locationFlag {
		getIPInfo(ip)
	} else if cmd.reverseFlag {
		reverseDomainParse(ip)
	} else {
		getIPInfo(ip)
	}
}
