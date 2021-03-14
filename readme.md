本爬虫爬取了珍爱网所有城市第一页用户的信息。信息存储在docker运行的elasticsearch。

### Run

用docker跑elasticsearch，版本号根据自己的docker进行设置：

```
docker run -d -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:7.11.1
```

然后直接运行项目的main.go