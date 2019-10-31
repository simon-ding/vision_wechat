package px500

import (
	"github.com/sirupsen/logrus"
	"time"
	"vision_wechat/db"
	"vision_wechat/sdk500px"
	"vision_wechat/utils"
)

func Px500Scheduler() {
	Heart500px()
	ReplyComments()
}

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
		for _, photo := range page.Data {
			err := client.DoLike(photo.ID, photo.UploaderID)
			if err != nil {
				logrus.Error(err)
				continue
			}
			logrus.Infof("成功点赞了 %s 的作品 %s", photo.UploaderInfo.NickName, photo.Title)
			time.Sleep(5 * time.Second)
		}
	}
}

var flutteredWordsCN = []string{
	"👍👍👍",
	"好作品",
	"精彩拍摄",
	"美👍👍",
	"美拍",
}
var flutteredWordsEN = []string{}

func ReplyComments() {
	accounts := db.DefaultDB.GetAll500px()
	for _, account := range accounts {
		client := sdk500px.NewClientUseCookie(account.Cookie)
		userId := client.OwnerID()
		gallery, err := client.FetchGallery(userId, 1, 20)
		if err != nil {
			logrus.Error(err)
			continue
		}
		if gallery.Status != "200" {
			logrus.Error(gallery.Message)
			continue
		}
		for _, g := range gallery.Data {
			comments, err := client.FetchComments(g.ID, 1, 20)
			if err != nil {
				logrus.Error(err)
				continue
			}
			if comments.Status != "200" {
				logrus.Error(comments.Message)
				continue
			}
			for _, c := range comments.Comments {
				if c.ChildComments != nil && len(c.ChildComments) > 0 {
					//已有回复
					continue
				}
				var replyMsg string
				if utils.IsChineseChar(c.Message) || utils.IsChineseChar(c.UserInfo.NickName) {
					//中文评论
					replyMsg = "谢谢啦！"
				} else {
					//英文评论
					replyMsg = "thanks!"
				}
				err := client.ReplyComment(c, replyMsg)
				if err != nil {
					logrus.Error(err)
				} else {
					logrus.Infof("成功回复 %s 的评论！", c.UserInfo.NickName)
				}
				//避免回复太快
				time.Sleep(5 * time.Second)
			}
		}
	}
}
