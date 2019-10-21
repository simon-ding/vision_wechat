package main

import (
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/cache"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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

	logrus.Println(config)

}
