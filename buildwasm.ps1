$Env:GOOS = 'js'
$Env:GOARCH = 'wasm'
go build -ldflags='-s -w' -trimpath -o ichimoudajin.wasm .
Remove-Item Env:GOOS
Remove-Item Env:GOARCH
