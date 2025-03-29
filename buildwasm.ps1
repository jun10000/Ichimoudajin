$Env:GOOS = 'js'
$Env:GOARCH = 'wasm'
go build -o ichimoudajin.wasm .
Remove-Item Env:GOOS
Remove-Item Env:GOARCH
