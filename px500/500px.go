package px500

import (
	"bytes"
	"github.com/TheForgotten69/goinsta/v2"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
	"vision_wechat/db"
	"vision_wechat/sdk500px"
	"vision_wechat/utils"
)

var m = make(map[string]*sdk500px.Client)

func Heart500px() {
	logrus.Info("500px 点赞开始运行")
	accounts := db.DefaultDB.GetAll500px()
	for _, account := range accounts {
		client, ok := m[account.UserID]
		if !ok {
			client = sdk500px.NewClient(account.Username, account.Password)
			m[account.UserID] = client
		}
		err := client.TestLogin()
		if err != nil {
			logrus.Errorf("login error %s", err)
			continue
		}
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
			time.Sleep(10 * time.Second)
		}
	}
}

func Upload2Instagram(duration time.Duration) func() {
	return func() {
		logrus.Errorf("instagram photo syncing begin...")
		accounts := db.DefaultDB.GetAll500px()
		for i, account := range accounts {
			logrus.Infof("%d: 500px account: %v", i, account)

			insAccount := db.DefaultDB.GetInstagram(account.UserID)
			if insAccount.Username == "" {
				logrus.Infof("user %s has no ins account", account.UserID)
				continue
			}
			logrus.Infof("%d: instagram account: %v", i, insAccount)

			client, ok := m[account.UserID]
			if !ok {
				client = sdk500px.NewClient(account.Username, account.Password)
				m[account.UserID] = client
			}
			err := client.TestLogin()
			if err != nil {
				logrus.Errorf("login error %s", err)
				continue
			}
			g, err := client.FetchGallery(client.UserId, 1, 20)
			if err != nil {
				logrus.Error(err)
				continue
			}

			insta, err := goinsta.Import(account.UserID)
			if err != nil {
				insta = goinsta.New(insAccount.Username, insAccount.Password)
				err = insta.Login()
				if err != nil {
					logrus.Error("login to instagram fail: ", err)
					continue
				}
				logrus.Infof("login instagram account %s success", insAccount.Username)
				err := insta.Export(account.UserID)
				if err != nil {
					logrus.Errorf("export ins: %v", err)
					continue
				}
			}

			for _, p := range g.Data {
				t := time.Unix(p.CreatedTime/1000, 0)
				if time.Now().Sub(t) > duration {
					continue
				}
				logrus.Infof("begin uploading image %s to instagram", p.Title)

				detail, err := client.GetPhotoDetails(p.ID)
				if err != nil {
					logrus.Errorf("get photo %s error: %v", p.Title, err)
					continue
				}
				data, err := client.DownloadPhoto(detail.DownLoadURL)
				if err != nil {
					logrus.Errorf("%v", err)
					continue
				}
				_, err = insta.UploadPhoto(bytes.NewReader(data), "11", 100, 0)
				if err != nil {
					logrus.Errorf("upload to instagram error: %v", err)
					continue
				}
				logrus.Info("upload instagram done")
			}
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

var replyWordsCN = []string{"谢谢！", "谢谢啦！", "谢谢老师！", "谢谢🙏"}

func ReplyComments() {
	logrus.Info("回复评论程序开始运行")
	accounts := db.DefaultDB.GetAll500px()
	for _, account := range accounts {
		client, ok := m[account.UserID]
		if !ok {
			client = sdk500px.NewClient(account.Username, account.Password)
			m[account.UserID] = client
		}
		err := client.TestLogin()
		if err != nil {
			logrus.Errorf("login error %s", err)
			continue
		}

		userId := client.UserId
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
					n := rand.Intn(len(replyWordsCN)) //随机选择
					replyMsg = replyWordsCN[n]
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
				time.Sleep(10 * time.Second)
			}
		}
	}
}
