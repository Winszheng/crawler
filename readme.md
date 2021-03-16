本爬虫爬取了珍爱网所有城市第一页用户的信息。信息存储在docker运行的elasticsearch。

开发流程：单任务版->并发版。

### Run

用docker跑elasticsearch，版本号根据自己的docker进行设置：

```
docker run -d -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:7.11.1
```

然后直接运行项目的main.go

### 结果

搜索结果如下，实际上，爬取的数据种类比前端显示出来的更多，但是前端不是本项目的重点(所以就随便整整得了吧)。

![image-20210315184458452](img/image-20210315184458452.png)