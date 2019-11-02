package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	db2 "vision_wechat/db"
	"vision_wechat/px500"
	"vision_wechat/wechat"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true, FullTimestamp: true})
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
	Scheduler struct {
		Px500 string
	}
}

func main() {
	var config *Config
	if err := viper.Unmarshal(&config); err != nil {
		logrus.Panic(err)
	}
	db2.InitConnection()
	db2.DefaultDB.Migrate()

	scheduledTasks(config)

	wc := wechat.NewClient(config.Wechat.AppId, config.Wechat.Secret, config.Wechat.Token, config.Wechat.EncodingAESKey)
	logrus.Println(config)

	http.HandleFunc("/", wc.Handlerfunc())
	serverPort := viper.GetString("server.port")
	err := http.ListenAndServe(":"+serverPort, nil)
	if err != nil {
		fmt.Printf("start server error , err=%v", err)
	}

}

func scheduledTasks(config *Config) {
	cr := cron.New()
	cr.AddFunc("@every "+config.Scheduler.Px500, px500.Heart500px)
	cr.AddFunc("CRON_TZ=Asia/Shanghai 30 23 * * * *", px500.ReplyComments)

	cr.Start()
}
