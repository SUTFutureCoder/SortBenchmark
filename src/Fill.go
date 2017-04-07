/**
 * 用于在磁盘中写入指定大小随机数
 */
package main

import (
	"flag"
	"logs"
	"os"
	"math"
	"math/rand"
	"strconv"
)

var (
	size   = flag.Int("size", 0, "size of file to fill")
	file   = flag.String("file", "/home/sort_raw_file", "file to fill")
	buffer = 1    //用于分片，避免内存先爆炸
)

func main()  {
	//编译命令行命令
	flag.Parse()

	if *size == 0  || *file == ""{
		logs.Fatal("file fill size can not be zero")
		os.Exit(0)
	}

	fd, _ := os.OpenFile(*file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	defer fd.Close()

	//当前数据大小
	currentSize := 0

	logs.Trace("START GENERATE")
	for currentSize < (*size * 1024 * 1024){
		data          := ""
		tmpDataLength := 0
		for {
			tmpData := strconv.FormatInt(rand.Int63n(math.MaxInt64), 10) + " "
			data    += tmpData
			tmpDataLength += len([]byte(tmpData))
			if tmpDataLength >= (buffer * 1024 * 1024){
				break;
			}
		}
		fd.Write([]byte(data))
		currentSize += (buffer * 1024 * 1024)
	}
	logs.Trace("done")
}
