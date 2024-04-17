package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/k0kubun/pp"
	"github.com/tidwall/gjson"
)

func main() {
	url := "https://dashscope-result-bj.oss-cn-beijing.aliyuncs.com/5fc5c860/2023-12-27/72e0abf8-1a04-4663-aa8f-36f6a93dbe63_output_1703666324177.txt.gz?Expires=1703925524&OSSAccessKeyId=LTAI5tQZd8AEcZX6KZV4G8qL&Signature=KMIuwFs4IPiFBXIdy0xaLx9vE34%3D"

	filePath := "file.txt" // 替换为要保存文件的路径和文件名

	// 下载压缩文件
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("下载压缩文件失败：%v", err)
	}
	defer resp.Body.Close()

	// 创建本地文件
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("创建本地文件失败：%v", err)
	}
	defer file.Close()

	// 解压缩文件
	reader, err := gzip.NewReader(resp.Body)
	if err != nil {
		log.Fatalf("解压缩文件失败：%v", err)
	}
	defer reader.Close()

	// 读取解压缩后的内容
	content, err := io.ReadAll(reader)
	if err != nil {
		log.Fatalf("读取解压缩后的内容失败：%v", err)
	}
	resjson := gjson.Parse(string(content))
	pp.Println(resjson)
	pp.Println(resjson.Get("output.embedding"))
	// 写入文件
	_, err = file.Write(content)
	if err != nil {
		log.Fatalf("写入文件失败：%v", err)
	}

	fmt.Println("文件已保存到本地。")

}
