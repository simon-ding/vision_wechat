package wechat

import (
	"github.com/ahmdrz/goinsta/v2"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/cache"
	"github.com/silenceper/wechat/message"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"vision_wechat/db"
)

type Client struct {
	wc *wechat.Wechat
}

func NewClient(appID, appSecret, token, encodingAESKey string) *Client {
	config := wechat.Config{
		AppID:          appID,
		AppSecret:      appSecret,
		Token:          token,
		EncodingAESKey: encodingAESKey,
		Cache:          cache.NewMemory(),
	}
	wc := wechat.NewWechat(&config)
	return &Client{wc: wc}
}

func (c *Client) Handlerfunc() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		// 传入request和responseWriter
		server := c.wc.GetServer(request, writer)
		//设置接收消息的处理方法
		server.SetMessageHandler(messageHandler)

		//处理消息接收以及回复
		err := server.Serve()
		if err != nil {
			logrus.Error(err)
			return
		}
		//发送回复的消息
		server.Send()
	}
}

func messageHandler(msg message.MixMessage) *message.Reply {

	logrus.Info("message received: ", msg)
	switch msg.MsgType {
	//文本消息
	case message.MsgTypeText:
		if strings.HasPrefix(msg.Content, "500px") {
			cookie := strings.TrimLeft(msg.Content, "500px")
			cookie = strings.TrimSpace(cookie)
			userID := msg.FromUserName
			db.DefaultDB.Set500pxCookie(userID, cookie)
			return textReturn("成功设置500px cookie!")
		}
		return textReturn(msg.Content)

		//图片消息
	case message.MsgTypeImage:
		imgURL := msg.PicURL
		resp, err := http.Get(imgURL)
		if err != nil {
			logrus.Error("download pic error: ", err)
			return textReturn("上传失败 instagram 失败！")
		}
		insAccount := db.DefaultDB.GetInstagram(msg.FromUserName)
		insta := goinsta.New(insAccount.Username, insAccount.Password)
		err = insta.Login()
		if err != nil {
			logrus.Error("login to instagram fail: ", err)
			return textReturn("上传失败 instagram 失败！")
		}
		defer insta.Logout()
		_, err = insta.UploadPhoto(resp.Body, "11", 100, 0)
		if err != nil {
			logrus.Error("upload to instagram error, ", err)
			return textReturn("上传失败 instagram 失败！")
		}
		return textReturn("上传成功！")

		//语音消息
	case message.MsgTypeVoice:
		//do something

		//视频消息
	case message.MsgTypeVideo:
		//do something

		//小视频消息
	case message.MsgTypeShortVideo:
		//do something

		//地理位置消息
	case message.MsgTypeLocation:
		//do something

		//链接消息
	case message.MsgTypeLink:
		//do something

		//事件推送消息
	case message.MsgTypeEvent:

	}
	return nil
}

func textReturn(content string) *message.Reply {
	text := message.NewText(content)
	return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
}