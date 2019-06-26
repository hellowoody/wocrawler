package kits

import (
	"time"
	"fmt"
	"os"
	"strconv"
)

func ShowBar(){
	str :=""
	cacheKit := GetInstance()
	for {
		quitFlag,_ := strconv.ParseBool(cacheKit.Config["quitFlag"].(string))
		if !quitFlag{
			str += "#"
			fmt.Fprintf(os.Stdout, "正在下载图片和种子文件 [%s]\r",str)
			time.Sleep(time.Second * 1)
		}
	}
}