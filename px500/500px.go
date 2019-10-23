package px500

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"vision_wechat/db"
)

var client = http.DefaultClient

const url500px = "https://500px.me"

type mainPage struct {
	Data []Data `json:data`
}

type Data struct {
	ID         string `json:id`
	UploaderID string `json:uploaderId`
}

func Heart500px() {
	accounts := db.DefaultDB.GetAll500px()
	for _, account := range accounts {
		logrus.Info("do heart for account ", account.UserID)
		indexPage(1, account.Cookie)
	}
}

func indexPage(page int, cookie string) error {
	indexURL := url500px + "/feedflow/index"
	req, err := http.NewRequest("GET", indexURL, nil)
	if err != nil {
		return err
	}
	query := req.URL.Query()

	query.Add("startTime", "")
	query.Add("page", strconv.Itoa(page))
	query.Add("size", "21")
	req.URL.RawQuery = query.Encode()

	req.Header.Add("Cookie", cookie)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var mainData mainPage

	err = jsoniter.Unmarshal(data, &mainData)
	if err != nil {
		return err
	}
	logrus.Infof("%+v\n", mainData)

	for _, pic := range mainData.Data {
		err = doLike(pic.ID, pic.UploaderID, cookie)
		if err != nil {
			logrus.Error(err)
		}
	}

	return nil
}

func doLike(id, uploadID, cookie string) error {
	likeURL := url500px + "/community/doLike.do"
	req, err := http.NewRequest("GET", likeURL, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Cookie", cookie)

	q := req.URL.Query()
	q.Add("action", "1")
	q.Add("likedId", id)
	q.Add("byAction", uploadID)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	data, _ := ioutil.ReadAll(resp.Body)
	logrus.Info(string(data))
	return nil
}
