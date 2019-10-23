package main

import (
	"fmt"
	"github.com/ahmdrz/goinsta/v2"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/cache"
	"github.com/silenceper/wechat/message"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	db2 "vision_wechat/db"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/wechat")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Panic(err)
	}

}

type Config struct {
	Wechat struct {
		Token          string
		AppId          string
		Secret         string
		EncodingAESKey string
	}

	Server struct {
		Port    string
		DataDir string
	}
}

var wc *wechat.Wechat
var db *db2.DB

func main() {
	var config *Config
	if err := viper.Unmarshal(&config); err != nil {
		logrus.Panic(err)
	}
	db = db2.NewConnection()
	defer db.Close()
	db.Migrate()

	memCache := cache.NewMemory()
	wcConfig := &wechat.Config{
		AppID:          config.Wechat.AppId,
		AppSecret:      config.Wechat.Secret,
		Token:          config.Wechat.Token,
		EncodingAESKey: config.Wechat.EncodingAESKey,
		Cache:          memCache,
	}
	wc = wechat.NewWechat(wcConfig)
	logrus.Println(config)

	http.HandleFunc("/", handler)
	serverPort := viper.GetString("server.port")
	err := http.ListenAndServe(":"+serverPort, nil)
	if err != nil {
		fmt.Printf("start server error , err=%v", err)
	}

}

func handler(writer http.ResponseWriter, request *http.Request) {
	// 传入request和responseWriter
	server := wc.GetServer(request, writer)
	//设置接收消息的处理方法
	server.SetMessageHandler(messageHandler)

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	server.Send()
}

func messageHandler(msg message.MixMessage) *message.Reply {

	logrus.Info("message received: ", msg)
	switch msg.MsgType {
	//文本消息
	case message.MsgTypeText:
		return textReturn(msg.Content)

		//图片消息
	case message.MsgTypeImage:
		imgURL := msg.PicURL
		resp, err := http.Get(imgURL)
		if err != nil {
			logrus.Error("download pic error: ", err)
			return textReturn("上传失败 instagram 失败！")
		}
		insAccount := db.GetInstagram(msg.FromUserName)
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
