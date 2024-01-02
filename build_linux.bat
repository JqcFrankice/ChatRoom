@echo off
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
SET GODEBUG=asyncpreemptoff=1
if not exist bin (
   md bin
)
cd bin
if not exist linux (
   md linux
)
cd linux
@echo on
go build ../../server/main/main.go
cd ../..
pause