package engine

import (
	"log"
	"sync"
	"wocrawler/model"
	"wocrawler/fetcher"
	"wocrawler/kits"
)

func Run(seed model.Request){
	var req model.Request = seed
	body,err := fetcher.Fetch(req.Url)
	if err!= nil {
		// log.Printf("engine===>>fetch error,url:%s,err:%v",req.Url,err)
		return
	}
	req.ParserFunc(body,req)
}

func CrawlRunSync(wg *sync.WaitGroup,req model.Request){
	Run(req)
	wg.Done()
}

func addReqChan(wg *sync.WaitGroup){
	woChan := model.GetWoChanInstance()
	for{
		req,ok := <-woChan.ReqChan
		if !ok {
			log.Println("ReqChan finish !!!!!!!!")
			break
		}
		// log.Println(req.Content)
		wg.Add(1)
		go CrawlRunSync(wg,*req)
	}
}

func addItemChan(wg *sync.WaitGroup){
	cacheKit := kits.GetInstance()
	woChan := model.GetWoChanInstance()
	downloadPath := cacheKit.Config["downloadPath"].(string)
	kits.Mkdir(downloadPath)
	filename := cacheKit.Config["filename"].(string)
	for{
		req,ok := <-woChan.ItemChan
		if !ok {
			log.Println("ItemChan finish !!!!!!!!")
			break
		}
		wg.Add(1)
		go CrawlRunSync(wg,*req)
		content := req.Content+";"+req.Url
		kits.WriteFile(downloadPath+filename,content)
	}
}

func addTorrentAndImgChan(wg *sync.WaitGroup){
	woChan := model.GetWoChanInstance()
	for{
		req,ok := <-woChan.TorrentChan
		reqImg,okImg := <-woChan.ImgChan
		if !ok && !okImg {
			log.Println("TorrentChan&ImgChan finish !!!!!!!!")
			break
		}
		if ok {
			wg.Add(1)
			go CrawlRunSync(wg,*req)
		}
		if okImg {
			wg.Add(1)
			go CrawlRunSync(wg,*reqImg)
		}
	}
}

func Crawl(){
	woChan := model.GetWoChanInstance()
	wg := sync.WaitGroup{}
	addReqChan(&wg)
	wg.Wait()
	woChan.ExitChan <- true  //ItemChan 可以close了
	addItemChan(&wg)
	wg.Wait()
	woChan.ExitChan <- true  //TorrentChan 可以close了
	addTorrentAndImgChan(&wg)
	wg.Wait()
	cacheKit := kits.GetInstance()
	cacheKit.Config["quitFlag"] = true
}
