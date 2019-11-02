package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"unicode"
)

func IsChineseChar(str string) bool {
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			return true
		}
	}
	return false
}

func NotifyServerChan(msg, desp string) error {
	token := viper.GetString("serverChan.token")
	urlChan := fmt.Sprintf("https://sc.ftqq.com/%s.send", token)
	req, err := http.NewRequest("GET", urlChan, nil)
	if err != nil {
		logrus.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("text", msg)
	q.Add("desp", desp)
	req.URL.RawQuery = q.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
