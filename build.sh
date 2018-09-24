SERVER_NAME="api-server"
# 打包
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/darwin/$SERVER_NAME -v -ldflags "-s -w" main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/linux/$SERVER_NAME -v -ldflags "-s -w" main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/windows/$SERVER_NAME.exe -v -ldflags "-s -w" main.go


# 压缩
upx -9 bin/darwin/$SERVER_NAME
upx -9 bin/linux/$SERVER_NAME
upx -9 bin/windows/$SERVER_NAME.exe

# 配置文件 也许你需要 对生产环境新增一个配置文件 config.prod.yml
cp config.yml bin/darwin/config.yml
cp config.yml bin/linux/config.yml
cp config.yml bin/windows/config.yml

# 静态资源
cp -R public bin/darwin/
cp -R public bin/linux/
cp -R public bin/windows/

# unix管理文件
cp manage.sh bin/darwin/manage.sh
cp manage.sh bin/linux/manage.sh
