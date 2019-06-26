# wocrawler
golang crawler

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