# logCollector

## 介绍
背景：日志文件是几乎任何系统必备的一项，当某个集群机器比较少时，查看日志不是难事，但是机器比较多时，需要有一套日志系统来方便管理。

目标：把集群机器上的日志实时收集，统一储存到中心系统，对这些系统建立索引，通过搜索即可找到对应日志，通过提供界面有好的web界面，通过web即可完成日志搜索。
日志准实时搜集，延时控制在分钟级别。

## 系统架构
![avatar](https://github.com/LinlinDu/myMarkdownPhotos/blob/master/photos/logcollect2.jpg)
## 安装
```
go get github.com/LinlinDu/logCollector
```
## 如何使用
整个项目分两个独立app:logAgent和logCollector。他们之间通过kafka传递数据。</br>
集群中安装kafka，选择一台机器用于管理日志，安装elasticsearch，kibana。 </br>
在conf文件夹中填写相应配置，在集群中需要收集日志的机器中使用logAgent，在管理日志的机器中使用logCollector。

## 运行案例
首先会等待日志文件在你设定的位置出现，然后开始逐条搜集日志
在dubug模式下输出如下
```
2019/07/01 13:09:53 Waiting for logs/20190701.log to appear...
2019/07/01 13:09:53.698 [I]  Log Agent start running...
2019/07/01 13:15:05.402 [D]  topic: [log] pid: [6], offset: [50]
2019/07/01 13:15:05.409 [D]  topic: [log] pid: [3], offset: [50]
2019/07/01 13:16:47.192 [D]  topic: [log] pid: [7], offset: [49]
2019/07/01 13:17:14.159 [D]  topic: [log] pid: [1], offset: [59]
2019/07/01 13:17:31.252 [D]  topic: [log] pid: [3], offset: [51]
2019/07/01 13:17:33.021 [D]  topic: [log] pid: [7], offset: [50]
```
访问http://127.0.0.1:5601 ，打开kibana，可看到搜集到的日志
![avatar](https://github.com/LinlinDu/myMarkdownPhotos/blob/master/photos/logcollect.jpg)
## 使用
+ kafka
+ zookeeper
+ tail
+ elasticsearch
+ kibana
