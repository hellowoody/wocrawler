package kits

import (
	"os"
	"io"
	"bufio"
	"log"
	"strings"
	"bytes"
)

func Mkdir(filepath string) error{
	_,err := os.Stat(filepath)
	if os.IsNotExist(err) {
		err=os.MkdirAll(filepath,os.ModePerm)
	}
	return err
}

func CreateFile(filepath string) error {
	_,err := os.Stat(filepath)
	if os.IsNotExist(err) {
		_,err=os.Create(filepath)
	}
	return err
}

func DownloadFile(filepath string,contents []byte) {
	file, err := os.Create(filepath)
	if err != nil {
		log.Printf("open file err=%v\n", err)
		return 
	}
	//及时关闭file句柄
	defer file.Close()
	io.Copy(file,bytes.NewReader(contents))
}

func WriteFile(filepath string,content string){
	file, err := os.OpenFile(filepath, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Println("filekit===>> open file err=", err)
		return 
	}
	//及时关闭file句柄
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString(content+"\n")
	writer.Flush()
}

func LoadFile(filename string) (m map[string]interface{} , err error){
	file,err := os.Open(filename)
	if err != nil {
		log.Println("open file error=>",err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var lineNo int = 0
	m = make(map[string]interface{},0)
	for {
		line,err := reader.ReadString('\n')
		if err == io.EOF {
			lineParse(&lineNo,&line,&m)
			break
		}
		if err != nil {
			log.Println(err)
			break
		}
		lineParse(&lineNo,&line,&m)
	}
	return
}

func lineParse(lineNo *int,line *string,m *map[string]interface{}){
	*lineNo++
	content := strings.TrimSpace(*line)
	if len(content) ==0 || content[0] == '\n' || content[0] == ';' || content[0] == '#'{
		return
	}
	contentSlice := strings.SplitN(content,"=",2)
	if len(contentSlice) == 0 {
		log.Printf("invalid config ,line: %d ",lineNo)
		return 
	}
	key := strings.TrimSpace(contentSlice[0])
	if len(key) == 0 {
		log.Printf("invalid config, line :%d ",lineNo)
		return
	}
	value := strings.TrimSpace(contentSlice[1])
	(*m)[key] = value
	log.Printf("config key : %s ,value : %s",key,value)
	return
}