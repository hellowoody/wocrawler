package model

import (
	"sync"
	"log"
)

type WoChan struct{
	ReqChan chan *Request
	ItemChan chan *Request
	TorrentChan chan *Request
	ImgChan chan *Request
	ExitChan chan bool
}
var once sync.Once
var instance *WoChan

func GetWoChanInstance() *WoChan{
	once.Do(func(){
		instance = &WoChan{
			ReqChan : make(chan *Request,100),
			ItemChan : make(chan *Request,100),
			TorrentChan : make(chan *Request,100),
			ImgChan : make(chan *Request,100),
			ExitChan : make(chan bool,3),
		}
	})
	return instance
}

func CloseChan(){
	woChan := GetWoChanInstance()
	for i:=0;i<3;i++{
		_,ok := <-woChan.ExitChan
		if i== 0 {
			if !ok {
				log.Println("Exit ReqChan==>?????????????")
				break
			}
			close(woChan.ReqChan)
		}else if i==1 {
			if !ok {
				log.Println("Exit ItemChan&ImgChan==>?????????????")
				break
			}
			close(woChan.ItemChan)
			close(woChan.ImgChan)
		}else if i==2 {
			if !ok {
				log.Println("Exit TorrentChan==>?????????????")
				break
			}
			close(woChan.TorrentChan)
		}
	}
}