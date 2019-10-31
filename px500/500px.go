package px500

import (
	"github.com/sirupsen/logrus"
	"vision_wechat/db"
	"vision_wechat/sdk500px"
)

func Heart500px() {
	logrus.Info("500px scheduler begins")
	accounts := db.DefaultDB.GetAll500px()
	for _, account := range accounts {
		logrus.Info("do heart for account ", account.UserID)
		client := sdk500px.NewClientUseCookie(account.Cookie)
		page, err := client.GetPage(1, 30)
		if err != nil {
			logrus.Error(err)
			continue
		}
		for _, item := range page.Data {
			err := client.DoLike(item.ID, item.UploaderID)
			if err != nil {
				logrus.Error(err)
			}
		}
	}
}
