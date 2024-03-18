set GOROOT=D:\dev\sdk\go1.22\go1.22.1
set GOPATH=D:\dev\sdk\gopath
set CGO_ENABLED=0
set GOARCH=amd64
set GOOS=windows
D:\dev\sdk\go1.22\go1.22.1\bin\go.exe build -o D:\dev\project\chrelyonly-frp\build\frp-luobo-windows.exe github.com/fatedier/frp/cmd/frpc
set GOARCH=amd64
set GOOS=linux
D:\dev\sdk\go1.22\go1.22.1\bin\go.exe build -o D:\dev\project\chrelyonly-frp\build\frp-luobo-linux-amd64 github.com/fatedier/frp/cmd/frpc
set GOARCH=amd64
set GOOS=darwin
D:\dev\sdk\go1.22\go1.22.1\bin\go.exe build -o D:\dev\project\chrelyonly-frp\build\frp-luobo-darwin-amd64 github.com/fatedier/frp/cmd/frpc