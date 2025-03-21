# vvBot

> [!WARNING]
> 本项目仅供学习交流使用，请勿用作不良用途

一个基于 LagrangeGo-Template 的 qq 机器人，直接修改自 [LagrangeGo-Template](https://github.com/ExquisiteCore/LagrangeGo-Template) 项目

基于 [LagrangeGo](https://github.com/LagrangeDev/LagrangeGo) 的 Bot 模板参考自[MiraiGo-Template](https://github.com/Logiase/MiraiGo-Template)

## 基础配置

配置 [application.toml](./application.toml)

```toml
[bot]
# 账号 必填
account = 114514

# 密码 选填
password = "pwd"

[ai]
# 是否开启 AI 搜索
isAISearch = true
# 使用的平台 url
url = "https://api.siliconflow.cn/v1/chat/completions"
# 使用的模型
model = "Qwen/Qwen2.5-7B-Instruct"
# 使用的平台api key
APIKey = ""
```

> [!NOTE]
> 基本上只能扫码登录，密码登录成功率较低

## 功能实现

### vv 表情

增加了基于 vv433 的 vv428 表情库，使用全 png 格式图片

自动检测群聊中`vv <参数>`格式的消息，如`vv 测试`，并对参数进行简单的匹配

默认使用 AI 进行关键词匹配，如果关闭或者匹配失败则转入普通匹配

普通匹配时，若在文件目录中匹配不到相应表情包，则发送随机表情

若在文件目录中匹配到多个同优先级的表情包，则返回其中随机一个

文件目录以`filename.json`的方式存储在`logic`文件夹下

### vvv 聊天

自动检测群聊中`vvv <参数>`格式的消息，如`vvv 测试`，并对参数进行~~神秘~~简单的回复

## 快速入门

### 快速部署

下载 release 中的压缩包，在您的 linux 环境中合适的位置解压，填写好 toml 文件后，运行可执行文件即可

若想长期部署，建议使用`nohup`命令后台运行，可以使用`nohup ./vvBot > app.log 2>&1 &`命令，将日志重定向到 app.log 文件

或者使用`screen`命令

### 二次开发

> [!NOTE]
> 若要重新设计 Bot 逻辑，建议直接使用模板 [LagrangeGo-Template](https://github.com/ExquisiteCore/LagrangeGo-Template)

使用`git clone`下载本项目

修改`/logic/custom_logic.go`中的事件监听逻辑，或是修改`/logic/custom_logic.go`中的搜索

## 欢迎对本项目进行贡献

待做

- 实现更加清晰的事件监听逻辑
- ~~实现模糊语义搜索~~
- 增加 vv428 武器库
- 增加防大手 👊🐔 逻辑

## 引入的第三方 go module

- [LagrangeGo](https://github.com/LagrangeDev/LagrangeGo)

核心协议库

- [toml](https://github.com/BurntSushi/toml)

用于解析配置文件，同时可监听配 toml 置文件的修改

- [logrus](https://github.com/sirupsen/logrus)

功能丰富的 Logger

再次声明，本 Bot 基于模板 [LagrangeGo-Template](https://github.com/ExquisiteCore/LagrangeGo-Template) 进行开发，若要直接设计新的 Bot 逻辑，建议直接使用此模板进行开发
