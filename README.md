# douSheng

## 抖声项目服务端

- 项目使用go语言开发,实现了抖声app的后端服务器

- 使用前先添加数据库sqlStatement/create.sql

- 在Const/theConst中修改参数  

- 使用gin开发,使用了gorm,ffmpeg,kitex(待定)等外部库,需要先进行go mod管理,再编译运行  

- 预计待redis补全,将信息先存到redis里,等待一定时间或者数量再发送到数据库

- 所有所需库已经补充在go.mod中

- kitex使用不正常可以查询[MyKitex](https://juejin.cn/post/7191552210696667196)的第一部分,它提供了kitex异常爆红的一些可能.  

> 使用go版本为19.5(不要使用大于等于1.20的版本)     
> 开发环境为linux/amd64   
> 使用数据库为 mysql Ver 15.1 Distrib 10.10.3-MariaDB, for Linux (x86_64) using readline 5.1

```shell
go build ./cmd/HttpService/main.go && ./main
```

```shell
go build ./cmd/feed/feedMain.go && ./feedMain
```

```shell
go build ./cmd/user/userMain.go && ./userMain
```

```shell
go build ./cmd/publish/publishMain.go && ./publishMain
```

```shell
go build ./cmd/favorite/favoriteMain.go && ./favoriteMain
```

```shell
go build ./cmd/comment/commentMain.go && ./commentMain
```

```shell
go build ./cmd/relation/relationMain.go && ./relationMain
```

```shell
go build ./cmd/message/messageMain.go && ./messageMain
```

### 一些注意

- 使用gorm预编译防止sql注入
- 使用ffmpeg "github.com/u2takey/ffmpeg-go",进行视频判断并且截图作为视频封面
- 使用jwt-go用账户和密码生成token.

### 项目分层功能说明
- cmd : 各个服务器及其handler
  - HttpService为http服务器
  - Class : 项目连接数据库使用的数据结构
  - rpc : 服务器之间连接的rpc服务启动
- Const : 项目启动前,需要配置的一些参数
- Controller : 项目的业务逻辑层
- public_cover : 存储视频封面或者用户封面的目录,预计未来将分离.
- public_videos : 存储视频的目录,预计未来将分离.
- Setting : 目前使用中的参数设置,预计未来合并入Const.
- Sql : 项目的持久化层,将数据读取,存储,修改入数据库.
- sqlStatement : 一些sql代码,包含mysql的初始化文件(DDL).

### 人员
yeyr2 : 全部(悲)

### 项目进度
> 当前 : 单体架构(完成) -> 初步完成微服务(完成)
> 预期 : 微服务  
> 当前目标 : 添加kitex使用gRPC连接业务逻辑层与持久化层(数据库)(完成) -> 添加redis缓存层

### [函数接口文档](funcJoggle.md)(未完成)

### [项目接口记录](接口文档记录.md)

#### 接口功能简易说明

- /douyin/feed/ - 视频流接口
  - 随机抽取视频,以传入的时间为最迟(传入的时间一般为now())
  - 返回视频列表和找到的最早的视频更新时间

- /douyin/user/register/ - 用户注册接口
  - 对用户账户和密码拼接后的token进行简单的加密
  - 将符合条件的用户添加入数据库

- /douyin/user/login/ - 用户登录接口
    - 使用用户账户和密码拼接后的token加密后的数据进行验证
    - 正确则返回部分用户信息

- /douyin/user/ - 用户信息
  - 使用用户账户和密码拼接后的token加密后的数据进行验证
  - 返回验证成功的用户信息

- /douyin/publish/action/ - 视频投稿
  - 登录用户选择视频上传。

- /douyin/publish/list/ - 发布列表
  - 登录用户的视频发布列表，直接列出用户所有投稿过的视频

- /douyin/favorite/action/ - 赞操作
  - 登录用户对视频的点赞和取消点赞操作。

- /douyin/favorite/list/ - 喜欢列表
  - 登录用户的所有点赞视频。

- /douyin/comment/action/ - 评论操作
  - 登录用户对视频进行评论。

- /douyin/comment/list/ - 视频评论列表
  - 查看视频的所有评论，按发布时间倒序。

- /douyin/relation/action/ - 关系操作
  - 登录用户对其他用户进行关注或取消关注。

- /douyin/relatioin/follow/list/ - 用户关注列表
  - 登录用户关注的所有用户列表。

- /douyin/relation/friend/list/ - 用户好友列表
  - 所有关注登录用户的粉丝列表。

- /douyin/message/chat/ - 聊天记录
  - 当前登录用户和其他指定用户的聊天消息记录

- /douyin/message/action/ - 消息操作
  - 登录用户对消息的相关操作，目前只支持消息发送
