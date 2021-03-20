package main

import (
    "fmt"
    "github.com/bitly/go-simplejson"
    "github.com/gocolly/colly"
    "gopkg.in/gomail.v2"
    "os"
    "time"
)

var(
    gmailPwd string
)

func main() {
    gmailPwd = os.Getenv("gmail_pwd")
    fmt.Println("gmail_pwd", gmailPwd)

    // Instantiate default collector
    c := colly.NewCollector()
    c.OnResponse(func(response *colly.Response) {
        newJson, err := simplejson.NewJson(response.Body)
        if err != nil{
            fmt.Println("response.Body json err", err)
            return
        }
        array := newJson.Get("rows").MustArray()
        for _, a := range array {
            stock,_ := a.(map[string]interface{})
            if stock["id"].(string) == "163406" {
                cell := stock["cell"].(map[string]interface{})
                fmt.Println(cell["fund_nm"])
                fmt.Println(cell["discount_rt"])
            }
        }
    })
    url := fmt.Sprintf("https://www.jisilu.cn/data/lof/stock_lof_list/?___jsl=LST___t=%d&rp=25&page=1", time.Now().UnixNano())
    err := c.Visit(url)
    if err != nil {
        fmt.Println(err)
    }

    sendEmail()
}

func sendEmail() {
    m := gomail.NewMessage()
    m.SetHeader("From", "znyalor@gmail.com")
    m.SetHeader("To", "657345933@qq.com")
    m.SetHeader("Subject", "Hello!")
    m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")

    d := gomail.NewDialer("smtp.gmail.com", 465, "znyalor@gmail.com", gmailPwd)
    if err := d.DialAndSend(m); err != nil {
        fmt.Println("sendEmail err", err)
        panic(err)
    }
}
