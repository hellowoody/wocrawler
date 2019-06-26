package main

import (
	"log"
	// "io"
	// "os"
	"wocrawler/engine"
	"wocrawler/parser"
	"wocrawler/kits"
	"wocrawler/model"
)

var (
	url string 
	woChan = model.GetWoChanInstance()
	cacheKit = kits.GetInstance()
)

func init(){
	url = cacheKit.Config["baseUrl"].(string)+cacheKit.Config["url"].(string)
}

func main(){
	// logFile , _ := os.Create("wocrawler.log")
	// io.MultiWriter(logFile,os.Stdout)
	log.Println("wocrawler start...")
	log.Println("******************************")
	go kits.ShowBar()
	go engine.Run(model.Request{
		Url:url,
		Content:"none",
		ParserFunc:parser.ParseLevelOne,
	})
	go model.CloseChan()
	engine.Crawl()
	log.Println("******************************")
	log.Println("wocrawler end.")
	log.Println("Bye bye.")
}

// func showBar(){
// 	str :=""
// 	for {
// 		if !woChan.QuitFlag {
// 			str += "#"
// 			fmt.Fprintf(os.Stdout, "正在下载图片和种子文件 [%s]\r",str)
// 			time.Sleep(time.Second * 1)
// 		}
// 	}
// }