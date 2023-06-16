#### 相关文档
* https://github.com/akavel/rsrc/releases
* https://github.com/lxn/walk

#### 打包流程

* rscr
```text
go get github.com/akavel/rsrc
rsrc -manifest test.manifest -o rsrc.syso
```
* build

```text
go build
```

```text
要摆脱cmd窗口，请运行
go build -ldflags="-H windowsgui"
```

* run
```text
file-tool.exe
```
