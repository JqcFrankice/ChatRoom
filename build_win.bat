@echo off
SET GOOS=windows
if not exist bin (
   md bin
)
cd bin
if not exist windows (
   md windows
)
cd windows
@echo on
go build ../../server/main/main.go

main.exe