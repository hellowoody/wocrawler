# wocrawler
golang crawler

### 注意 可能需要手动导入golang.org/net包
    $cd $GOPATH/src/golang.org/x/
    $git clone https://github.com/golang/net.git net
    $go install net  //这句可以不执行;如果执行,可以执行go install,执行后之后没有提示，就说明安装好了

## 编译
go build -o bin/wocrawler/run wocrawler/
## 执行
./bin/wocrawler/run
## 打包win64
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/wocrawler/run.exe wocrawler/
## 打包win32
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o bin/wocrawler/run32.exe wocrawler/
## 打包linux 需要手动赋权 chmod -R 777 文件夹
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o  bin/wocrawler/startup wocrawler/