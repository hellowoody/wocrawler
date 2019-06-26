package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte,error){
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Close = true
	req.Header.Set("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36")
	if err!=nil  {
		return nil,err
	}
	resp,err := client.Do(req)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil,fmt.Errorf("wrong statusCode, %d",resp.StatusCode)
	}
	//读取相应体并返回
	return ioutil.ReadAll(resp.Body)
}