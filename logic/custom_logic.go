package logic

import (
	"encoding/json"
	"strings"
	"os"
	"math/rand"
	"time"
	
	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/client/event"
	"github.com/LagrangeDev/LagrangeGo/message"
)

// RegisterCustomLogic 注册所有自定义逻辑
func RegisterCustomLogic() {
	// 注册私聊消息处理逻辑
	Manager.RegisterPrivateMessageHandler(func(client *client.QQClient, event *message.PrivateMessage) {
		// client.SendPrivateMessage(event.Sender.Uin, []message.IMessageElement{message.NewText("Hello World!")})
	})

	// 注册群消息处理逻辑
	Manager.RegisterGroupMessageHandler(func(client *client.QQClient, event *message.GroupMessage) {
		rand.Seed(time.Now().UnixNano())

		const jsonFilePath = "./logic/filename.json"
		// 提前加载文件目录
		jsonData, err := os.ReadFile(jsonFilePath)
		if err != nil {
			reply := []message.IMessageElement{
				message.NewText("读取文件目录出错"),
        	}
			client.SendGroupMessage(event.GroupUin, reply)
		}

		var filenames []string
		if err := json.Unmarshal(jsonData, &filenames); err != nil {
			reply := []message.IMessageElement{
				message.NewText("解析文件目录出错"),
        	}
			client.SendGroupMessage(event.GroupUin, reply)
		}

		// 添加1~2秒的随机延时, 求生欲
		randomDelay := 1000 + rand.Intn(1001)
		time.Sleep(time.Duration(randomDelay) * time.Millisecond)

		// 提取文本消息
		var msgText string
    	for _, elem := range event.Elements {
        	if textElem, ok := elem.(*message.TextElement); ok {
            	msgText += textElem.Content
        	}

			if strings.HasPrefix(msgText, "vv ") {
				args := strings.TrimPrefix(msgText, "vv ")

				// 关键词模糊搜索逻辑
				filepath := searchImageByKeyword(args, filenames)

				// 读取本地图片
				imageData, _ := os.ReadFile(filepath)

				imgElement := message.NewImage(imageData)
				
				reply := []message.IMessageElement{
					imgElement,
        		}
        		client.SendGroupMessage(event.GroupUin, reply)
    		}

			if strings.HasPrefix(msgText, "求生欲"){
				client.SendPrivateMessage(event.GroupUin, []message.IMessageElement{message.NewText("我不是 bot !")})
			}
    	}
	})

	Manager.RegisterNewFriendRequestHandler(func(client *client.QQClient, event *event.NewFriendRequest) {
		//event.SourceUid
		//logrus.Println("UID" + event.SourceUid)
		//client.SetFriendRequest(true, event.SourceUid)
	})
}


