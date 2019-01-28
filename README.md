# goblog
基于Golang的个人简易博客系统

## goblog介绍

> goblog基于go语言开发的一个简约版个人博客系统,基于Golang语言编写,后端基于了Beego的web框架,目前具备博文系统最基础的功能模块.基本上是一个拿来即用的个人博文平台,只需要部署一个mysql数据存储服务,即可作为个人博文的发布平台使用.

## goblog界面截图

**后台管理**

- ![](https://ws2.sinaimg.cn/large/006tNc79gy1fzmd65xx1xj31ne0u0hdu.jpg)
- ![](https://ws2.sinaimg.cn/large/006tNc79gy1fzmd7o3vkzj31mz0u0tji.jpg)
- ![](https://ws3.sinaimg.cn/large/006tNc79gy1fzmdatx3ejj31n30u0n5v.jpg)
- ![](https://ws2.sinaimg.cn/large/006tNc79gy1fzmdbgwjnbj31ne0u048m.jpg)

**PC前端展示**

- ![](https://ws1.sinaimg.cn/large/006tNc79gy1fzmd8ka8mbj31mq0u0n6r.jpg)
- ![](https://ws4.sinaimg.cn/large/006tNc79gy1fzmdd7bvluj31mw0u0h1z.jpg)
- ![](https://ws1.sinaimg.cn/large/006tNc79gy1fzmdeyausgj31mx0u0qbe.jpg)
- ![](https://ws4.sinaimg.cn/large/006tNc79gy1fzmdeet12bj31nb0u0dmf.jpg)

**手机前端展示**
- ![](https://ws2.sinaimg.cn/large/006tNc79gy1fzmdinfw1mj30u01hcq8s.jpg)
- ![](https://ws2.sinaimg.cn/large/006tNc79gy1fzmdigavjlj30u01hcgrm.jpg)
- ![](https://ws3.sinaimg.cn/large/006tNc79gy1fzmdko5i9qj30u01hcgop.jpg)

## goblog技术组件

-	基于go语言,
-	集成于beego的web框架

  > https://beego.me/
-	数据持久mysql
-	博文撰写组件
	> 支持 [富文本编辑](https://summernote.org/)
	> 支持 [Markdown编辑](http://pandao.github.io/editor.md/)
	> 目前两款编译器中设计到图片上传资源均存储在当前服务器中,暂时没有使用第三方云存储服务
-	站内全文检索riot

  > https://github.com/go-ego/riot

## goblog安装部署

>	**好了,现在让我们来手动搭建一个goblog吧！**

###	安装

#### 获取goblog源码
1. 我们先通过github拉取goblog源码
2. github地址: https://github.com/liangyd1024/goblog
3. 这里拉取git库需要安装git,git的安装步骤这里就不在重复(网上一大把呦)
```linux
	git clone https://github.com/liangyd1024/goblog.git
```

#### 安装go运行环境
**各个操作系统安装go的步骤大同小异，这里我们以Linux来做示例**

1.	获取go对应的版本安装包,这里我们到go官网获取最新版本的安装文件

 ```linux
	wget https://dl.google.com/go/go1.11.5.linux-amd64.tar.gz
 ```
-	通过 https://golang.org/dl/ 我们可查看到go的所有版本资源

2.	解压下载包

```linux
	tar -xvf go1.11.5.linux-amd64.tar.gz
```
3.	配置go的环境变量
```linux
	cd ~
	vi ~/.bash_profile
	export GOROOT=$HOME/App/go
	export GOPATH=$HOME/Project
	export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```
-	编辑文件并保存
-	这里GOROOT是指go安装包解压后所在的目录
-	GOPATH为后续应用源码所在工作目录
```linux
	source ~/.bash_profile
	go version
```
-	刷新配置文件后验证go是否安装正常,这里会输出`go version go1.11.5 linux/amd64`


#### 安装mysql
**goblog采用mysql作为数据存储服务,so 我们需要在我们对应的服务器上安装上一个mysql服务实例并启动他提供服务(对于mysql的安装本篇幅就不做多描述,网上已经有许多实例)**
-	在安装好mysql实例后我们只需要再实例上手动建立一个mysql的Shcema,Schema的名字为为goblog,所有的表结构在goblog启动时会自动创建生成(下文有特别介绍)


#### 编译并部署goblog
**通过git获取goblog源码后通过`go build`命令进行编译**
```linux
	cd $HOME/Project/src/goblog
	go build ./
```
-	源码拉取请在\$GOPATH后先建立src目录(因为golang中对\$GOPATH目录约定有3个子目录:src、pkg、bin),src下存放所有项目的源代码
-	编译期间由于国内网络受限一些相关的lib会下载失败: go: golang.org/x/net@v0.0.0-20181220203305-927f97764cc3: unrecognized import path "golang.org/x/net" (https fetch: Get https://golang.org/x/net?go-get=1: dial tcp 216.239.37.1:443: i/o timeout)
-	我们需要单独设置下go下载包的代理
-	继续编辑配置文件`~/.bash_profile`
```linux
	export GOPROXY=https://goproxy.io
```
-	配置好后重新执行`go build ./`命令 **’chua‘** 的一下依赖包就下好了
-	执行完构建命令后我们可以在当前目录下找到刚刚构建好后的包`goblog`

#### 运行goblog
**上面通过`go build`命令构建完成后,接下来我们就可以把goblog运行起来了**
> 在运行前我们需要先说明下几个目录的作用
> 1. conf/-------项目配置文件
>   - beego.conf------goblog中一些web配置
>   - db.conf------数据库配置
>   - log.cong------日志配置
> 2. src/------go源码文件
> 3. static/------静态资源(js/css/img...)
> 4. views/------页面模板文件(html)
> 5. main.go------主程序入口文件
> 6. go.mod go.sum------go的模块lib依赖配置文件
> 7. README.MD------项目介绍文件

- 在conf下我们可以将对应要链接的数据库的地址进行配置,编辑db.conf

- 直接运行刚才我们构建好的应用文件

- ```linux
   ./goblog
   ```

- 或者通过`go run main.go`命令运行goblog

- 运行后可以通过控台查看到启动日志
![](https://ws1.sinaimg.cn/large/006tNc79gy1fzmd1byxnfj31no0u0gqn.jpg)
- goblog启动后会建立默认的账号和密码：admin/goblog,我们通过访问127.0.0.1:9090/admin可以登录到后台进行博文的发布,发布完成后可以通过127.0.0.1:9090查阅已发布的博文信息
- 登录界面
![](https://ws2.sinaimg.cn/large/006tNc79gy1fzmd65xx1xj31ne0u0hdu.jpg)
- 博文管理
- ![](https://ws2.sinaimg.cn/large/006tNc79gy1fzmd7o3vkzj31mz0u0tji.jpg)
- 博文浏览(前端展示)
- ![](https://ws1.sinaimg.cn/large/006tNc79gy1fzmd8ka8mbj31mq0u0n6r.jpg)



**这里有个地方要特别强调下,因为goblog是直接采用beego提供的orm框架来做DB操作,所以在conf/db.conf中有个配置项`mysqlForce = false`要特别指出,设置为true时每次启动应用时都会将对应的表结构数据清除,所以此配置项只需在首次应用使用时指定true即可(自动创建表模型),后续在生产环境或者开发环境下不需要重新格式化数据情况时请慎重开启.**


