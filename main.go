package main

import (
	"fmt"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/cache"
	"github.com/silenceper/wechat/message"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
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
	Instagram struct {
		Username string
		Password string
	}
	Wechat struct {
		Token          string
		AppId          string
		Secret         string
		EncodingAESKey string
	}

	Server struct {
		Port string
	}
}

func main() {
	var config *Config
	if err := viper.Unmarshal(&config); err != nil {
		logrus.Panic(err)
	}
	memCache := cache.NewMemory()

	wcConfig := &wechat.Config{
		AppID:          config.Wechat.AppId,
		AppSecret:      config.Wechat.Secret,
		Token:          config.Wechat.Token,
		EncodingAESKey: config.Wechat.EncodingAESKey,
		Cache:          memCache,
	}
	wc := wechat.NewWechat(wcConfig)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// 传入request和responseWriter
		server := wc.GetServer(request, writer)
		//设置接收消息的处理方法
		server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {

			//回复消息：演示回复用户发送的消息
			text := message.NewText(msg.Content)
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
		})

		//处理消息接收以及回复
		err := server.Serve()
		if err != nil {
			fmt.Println(err)
			return
		}
		//发送回复的消息
		server.Send()

	})
	serverPort := viper.GetString("server.port")
	err := http.ListenAndServe(":"+serverPort, nil)
	if err != nil {
		fmt.Printf("start server error , err=%v", err)
	}

}
