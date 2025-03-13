# vvBot

> [!WARNING]

> 本项目仅供学习交流使用，请勿用作不良用途

一个基于 LagrangeGo-Template 的 qq 机器人，直接修改自 [LagrangeGo-Template](https://github.com/ExquisiteCore/LagrangeGo-Template) 项目

基于 [LagrangeGo](https://github.com/LagrangeDev/LagrangeGo) 的 Bot 模板参考自[MiraiGo-Template](https://github.com/Logiase/MiraiGo-Template)

## 基础配置

账号配置 [application.toml](./application.toml)

```toml
[bot]
# 账号 必填
account = 114514

# 密码 选填
password = "pwd"
```

> [!NOTE]
> 基本上只能扫码登录，密码登录成功率较低

## 功能实现

增加了基于 vv433 的 vv428 表情库，使用全 png 格式图片

自动检测群聊中`vv <参数>`格式的消息，如`vv 测试`，并对参数进行简单的匹配

若在文件目录中匹配不到相应表情包，则发送随机表情

若在文件目录中匹配到多个同优先级的表情包，则返回其中随机一个

文件目录以`filename.json`的方式存储在`logic`文件夹下

## 快速入门

### 二次开发

> [!NOTE]
> 若要重新设计 Bot 逻辑，建议直接使用模板 [LagrangeGo-Template](https://github.com/ExquisiteCore/LagrangeGo-Template)

使用`git clone`下载本项目

修改`/logic/custom_logic.go`中的事件监听逻辑，或是修改`/logic/custom_logic.go`中的搜索

## 欢迎对本项目进行贡献

## 引入的第三方 go module

- [LagrangeGo](https://github.com/LagrangeDev/LagrangeGo)

核心协议库

- [toml](https://github.com/BurntSushi/toml)

用于解析配置文件，同时可监听配 toml 置文件的修改

- [logrus](https://github.com/sirupsen/logrus)

功能丰富的 Logger

再次声明，本 Bot 基于模板 [LagrangeGo-Template](https://github.com/ExquisiteCore/LagrangeGo-Template) 进行开发，若要直接设计新的 Bot 逻辑，建议直接使用此模板进行开发
