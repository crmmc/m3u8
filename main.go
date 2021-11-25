package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/oopsguy/m3u8/dl"
)

var (
	url      string
	output   string
	chanSize int
)

func init() {
	flag.StringVar(&url, "u", "", "缺少m3u8地址")
	flag.IntVar(&chanSize, "c", 4, "缺少下载协程数")
	flag.StringVar(&output, "o", "", "缺少输出文件的目录")
}

func help() {
	fmt.Println("=== 一个支持多线程下载的M3U8下载器 ===")
	fmt.Println("原项目地址: https://github.com/oopsguy/m3u8/")
	fmt.Println("范例:")
	fmt.Println("./m3u8 -u=http://example.com/index.m3u8 -o=/data/example")
	fmt.Println("-u=[M3U8 地址]\n-o=[文件保存目录]\n-c=[下载协程并发数，默认 4]")

}

func main() {
	flag.Parse()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[出错]", r)
			os.Exit(-1)
		}
	}()
	if url == "" {
		help()
		panicParameter("u")
	}
	if output == "" {
		help()
		panicParameter("o")
	}
	if chanSize <= 0 {
		help()
		panic("最大协程数设置不可小于0")
	}
	downloader, err := dl.NewTask(output, url)
	if err != nil {
		panic(err)
	}
	if err := downloader.Start(chanSize); err != nil {
		panic(err)
	}
	fmt.Println("Done!")
}

func panicParameter(name string) {
	panic("缺少必要选项'-" + name + "=[]'")
}
