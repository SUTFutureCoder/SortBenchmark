##排序竞赛
TOPIC：如何在有限的内存排序超大号文件？

##使用方法
### 第一步：填充随机无序文件
```bash
cd SortBenchmark
export GOPATH=$PWD
cd src
go run Fill.go -size 要填充的文件大小(M) -file 要填充的文件地址
```

### 第二步：执行排序
待填坑