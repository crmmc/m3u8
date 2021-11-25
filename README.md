# M3U8

原项目地址: : https://github.com/oopsguy/m3u8/
M3U8 是一个使用了 Go 语言编写的迷你 M3U8 下载工具。你只需指定必要的 flag (`u`、`o`、`c`) 来运行, 工具就会自动帮你解析 M3U8 文件，并将 TS 片段下载下来合并成一个文件。
我只是汉化了一下帮助页面和报错信息,没有修改主要功能代码

## 功能

- 下载和解析 M3U8（仅限 VOD 类型）
- 下载 TS 失败重试
- 解析 Master playlist
- 解密 TS
- 合并 TS 片段

## 用法

### 源码方式

```bash
go run main.go -u=http://example.com/index.m3u8 -o=/data/example
```

### 二进制方式:

#### 编译:

windows,Linux 和 MAC OS 上大同小异

> 首先下载依赖
> go get "github.com/oopsguy/m3u8/dl"
> 然后就可以开始编译了,添加的参数是为了去掉 debug 信息,减小可执行文件体积
> go build -ldflags="-w -w"
> 找到源码目录下的 m3u8.exe \[Windows\] 或 m3u8\[Linux \& Mac os\] 直接运行即可
> 可以使用 upx --best m3u8.exe 来进一步压缩可执行文件的体积
> 我编译了一份 windows 下的 m3u8.exe 放在源码目录,已进行 upx 压缩,想用可以自取,不想要删掉即可

使用教程

```
./m3u8 -u=http://example.com/index.m3u8 -o=/data/example -c=16
```

参数说明：

```
-u=[M3U8 地址]
-o=[文件保存目录]
-c=[下载协程并发数，默认 4]
```

部分服务器可能限制请求频率，不想被封 IP 的话请合理调整下载并发数。

## 下载

[来自原作者的二进制文件,想用汉化版请自己编译](https://github.com/oopsguy/m3u8/releases)

## 截屏

![Demo](./screenshots/demo.gif)

## 参考资料

- [TS 科普 2 包头](https://blog.csdn.net/cabbage2008/article/details/49281729)
- [HTTP Live Streaming draft-pantos-http-live-streaming-23](https://tools.ietf.org/html/draft-pantos-http-live-streaming-23#section-4.3.4.2)
- [MPEG transport stream - Wikipedia](https://en.wikipedia.org/wiki/MPEG_transport_stream)

## License

[MIT License](./LICENSE)
