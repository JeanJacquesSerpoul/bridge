$flag= "-X main.version=$(git describe --always)"
$filename="deskserver.go"
Set-Location -Path server
Remove-Item *.html
Remove-Item *.outclear
$env:GOOS="linux"
$env:GOARCH="arm"
go build -o deskserver_arm -ldflags="$flag" $filename 
# Put in /usr/local/bin and chmod +x
$env:GOOS="linux"
$env:GOARCH="amd64"
go build -o deskserver_linux -ldflags="$flag" $filename 
$env:GOOS="windows"
$env:GOARCH="amd64"
go build -o deskserver_windows.exe -ldflags="$flag" $filename 
Copy-Item  deskserver_windows.exe "C:\applications\bridge_server\deskserver_windows.exe"
Set-Location -Path ..
Set-Location -Path distribution
Remove-Item *.tmp
Remove-Item *.pbn
Remove-Item *.out
Remove-Item *.html
Remove-Item *.test
Set-Location -Path ..