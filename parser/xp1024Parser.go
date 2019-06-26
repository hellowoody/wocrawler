package parser

import (
	"strings"
	"log"
	"regexp"
	// "fmt"
	"github.com/PuerkitoBio/goquery"
	"wocrawler/model"
	"wocrawler/kits"
)

func ParseLevelOne(contents []byte,param interface{}) {
	woChan := model.GetWoChanInstance()
	reqChan := woChan.ReqChan
	exitReq := woChan.ExitChan
	doc,err := goquery.NewDocumentFromReader(strings.NewReader(string(contents)))
	if err != nil {
		log.Println("goquery parseleveone error : ",err)
	}
	doc.Find("div.main-wrap > div > div.t > table > tbody:nth-child(2) > tr.tr3 > td:nth-child(2) ").Each(func(i int,s *goquery.Selection){
		r,_ := regexp.Compile("^\\[+[\u4e00-\u9fa5]") //正则匹配 开头是[ 并且至少一个中文
		ifnext := r.MatchString(s.Children().First().Text())
		// log.Println(res,s.Children().First().Text())
		if ifnext {
			item := s.Children().Eq(1)
			url,ok := item.Find("a").Attr("href")
			if ok {
				cacheKit := kits.GetInstance()
				url = cacheKit.Config["baseUrl"].(string)+url
				// fmt.Println(item.Text(),url)
			}
			req := model.Request{
				Url:url,
				Content:item.Text(),
				ParserFunc:ParseLevelTwo,
			}
			reqChan <- &req
		}
	})
	exitReq <- true  //reqChan 可以close了
}

func ParseLevelTwo(contents []byte,param interface{}){
	woChan := model.GetWoChanInstance()
	itemChan := woChan.ItemChan
	imgChan := woChan.ImgChan
	title := param.(model.Request).Content
	doc,err := goquery.NewDocumentFromReader(strings.NewReader(string(contents)))
	if err != nil {
		log.Println("goquery ParseLevelTwo error : ",err)
	}
	downloadUrl := doc.Find("#read_tpc > a")
	url,ok := downloadUrl.Last().Attr("href")
	if ok {
		req := model.Request{
			Url:url,
			Content:title,
			ParserFunc:ParseLevelThree,
		}
		itemChan <- &req
		// fmt.Println(downloadUrl.Text())
	}
	downloadImgUrl := doc.Find("#read_tpc img")
	imgUrl,ok := downloadImgUrl.First().Attr("src")
	if ok {
		req := model.Request{
			Url:imgUrl,
			Content:title,
			ParserFunc:DownloadImg,
		}
		imgChan <- &req
	}
}

func ParseLevelThree(contents []byte,param interface{}) {
	woChan := model.GetWoChanInstance()
	torrentChan := woChan.TorrentChan
	title := param.(model.Request).Content
	baseUrl := param.(model.Request).Url
	doc,err := goquery.NewDocumentFromReader(strings.NewReader(string(contents)))
	if err != nil {
		log.Println("goquery ParseLevelTwo error : ",err)
	}
	downloadUrl := doc.Find("a.uk-button ")
	resInt := strings.Index(baseUrl,"/torrent")
	if resInt > 0 {
		rs := []rune(baseUrl)
		baseUrl = string(rs[:resInt])
		url,ok := downloadUrl.Last().Attr("href")
		if ok {
			req := model.Request{
				Url:baseUrl+url,
				Content:title,
				ParserFunc:DownloadFile,
			}
			torrentChan <- &req
			// fmt.Println(title)
		}
	}
}

func DownloadFile(contents []byte,param interface{}) {
	title := param.(model.Request).Content
	title = strings.TrimSpace(title)
	title = strings.Replace(title," ","",-1)
	title = strings.Replace(title,"/","",-1)
	cacheKit := kits.GetInstance()
	downloadPath := cacheKit.Config["downloadPath"].(string)
	filePath := cacheKit.Config["filePath"].(string)
	err := kits.Mkdir(downloadPath+filePath)
	if err != nil {
		log.Println("Parser===>> Mkdir torrent err:",err)
	}
	kits.DownloadFile(downloadPath+filePath+"/"+title+".torrent",contents)
}

func DownloadImg(contents []byte,param interface{}){
	title := strings.TrimSpace(param.(model.Request).Content)
	title = strings.Replace(title," ","",-1)
	title = strings.Replace(title,"/","",-1)
	baseUrl := param.(model.Request).Url
	cacheKit := kits.GetInstance()
	downloadPath := cacheKit.Config["downloadPath"].(string)
	filePath := cacheKit.Config["imgPath"].(string)
	err := kits.Mkdir(downloadPath+filePath)
	if err != nil {
		log.Println("Parser===>> Mkdir img err:",err)
	}
	resInt := strings.LastIndex(baseUrl,".")
	if resInt > 0 {
		rs := []rune(baseUrl)
		suffix := string(rs[resInt:])
		// fmt.Println(downloadPath+filePath+"/"+title+suffix)
		kits.DownloadFile(downloadPath+filePath+"/"+title+suffix,contents)
	}
}