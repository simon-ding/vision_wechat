package sdk500px

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	"vision_wechat/utils"
)

const (
	loginUrl  = "/user/v2/tologin"
	baseUrl   = "https://500px.com.cn/"
	userAgent = "Mozilla/5.0 (Linux; Android 10; Mi 10 Build/QKQ1.191117.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/81.0.4044.138 Mobile Safari/537.36"
)

func wrapperRequest(req *http.Request, token string) *http.Request {
	req.Header.Add("PF500MClient", "android")
	req.Header.Add("PF500MClientVersion", "404000")
	req.Header.Add("PF500MClientId", "a4e84d18f3414a788a234c76aee6067f")
	req.Header.Add("equipmentType", "Xiaomi Mi 10(Android 10)")
	req.Header.Add("netWorkState", "WIFI")
	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("access_token", token)
	req.Header.Add("accessToken", token)
	req.Header.Add("Host", "500px.com.cn")
	return req
}

type Client struct {
	client   *http.Client
	username string
	password string
	token    string
	UserId   string
}

func NewClient(username, password string) *Client {
	return &Client{
		client:   http.DefaultClient,
		username: username,
		password: password,
	}
}

type Response struct {
	Data    interface{}
	Message string
	Status  string
}

func (c *Client) Login() error {
	data := url.Values{}
	data.Add("userName", c.username)
	data.Add("password", c.password)
	data.Add("expires", "31536000")

	m := map[string]interface{}{}
	m["expireTime"] = time.Now().UnixNano()/1000/1000 + 300000
	m["phone"] = c.username
	res, err := json.Marshal(m)
	if err != nil {
		return errors.Wrap(err, "json marshal")
	}
	cipertext, err := DesEncrypt(res, key500px)
	if err != nil {
		return errors.Wrap(err, "des encrypt")
	}
	s, err := DesDecrypt(cipertext, key500px)
	fmt.Println(string(s), err)
	text := encodeBase64(cipertext)

	fmt.Println(text)
	//bb := decodeBase64("POSTHVWuTePIWX9K8H40kQldU/AN57bSR8CVxFLkjk/hNh2+hhgaaunvN5lAuqzT72GGLo2dOAG+uW0jLYrpAQ==")
	//d, _ := DesDecrypt(bb, key500px)
	//fmt.Println(string(d))
	data.Add("secretText", text)

	req, err := http.NewRequest(http.MethodPost, baseUrl+loginUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	req = wrapperRequest(req, "")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var user UserInfo
	err = json.Unmarshal(respData, &user)
	if err != nil {
		return errors.Wrap(err, "json")
	}
	if user.Status != "200" {
		return errors.New(user.Message)
	}

	c.token = user.UserAccountInfo.AccessToken
	c.UserId = user.UserAccountInfo.UserID
	fmt.Println(string(respData))
	log.Print("login sucess!")
	return nil
}

func (c *Client) TestLogin() error {
	if c.token == "" {
		return c.Login()
	}

	_, err := c.GetPage(1, 20)
	if err != nil {
		return c.Login()
	}
	return nil
}

func (c *Client) GetPage(page int, size int) (*IndexPage, error) {
	req, err := c.newRequest("GET", "/feedflow/index", nil)
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()

	query.Add("startTime", "")
	query.Add("page", strconv.Itoa(page))
	query.Add("size", strconv.Itoa(size))
	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var indexPage IndexPage

	err = jsoniter.Unmarshal(data, &indexPage)
	if err != nil {
		return nil, err
	}
	if indexPage.Status != "200" {
		msg := fmt.Sprintf("status: %s, %s", indexPage.Status, indexPage.Message)
		if strings.Contains(indexPage.Message, "login") {
			utils.NotifyServerChan(msg, "")
		}
		return nil, fmt.Errorf(msg)
	}
	return &indexPage, nil
}

func (c *Client) DoLike(id, uploadID string) error {
	req, err := c.newRequest("GET", "/community/doLike.do", nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("action", "1")
	q.Add("likedId", id)
	q.Add("byAction", uploadID)
	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	data, _ := ioutil.ReadAll(resp.Body)
	var res Response
	_ = jsoniter.Unmarshal(data, &res)
	if res.Status != "200" {
		msg := fmt.Sprintf("status: %s, %s", res.Status, res.Message)
		if strings.Contains(res.Message, "login") {
			utils.NotifyServerChan(msg, "")
		}
		return fmt.Errorf(msg)
	}
	return nil
}

func (c *Client) newRequest(method, refUrl string, body io.Reader) (*http.Request, error) {
	baseUrl, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}
	baseUrl, err = baseUrl.Parse(refUrl)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, baseUrl.String(), body)
	if err != nil {
		return nil, err
	}
	req = wrapperRequest(req, c.token)
	//req.Header.Add("Cookie", c.Cookie)
	//req.Header.Add("User-Agent", c.userAgent)
	return req, nil
}

func (c *Client) GetPhotoDetails(photoID string) (*PhotoDetail, error) {
	req, err := c.newRequest("GET", "community/photo-details/"+photoID, nil)
	if err != nil {
		logrus.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("type", "json")
	q.Add("imgsize", "p1,p2,p5,p6")
	req.URL.RawQuery = q.Encode()
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	var photoDetail PhotoDetail
	err = jsoniter.Unmarshal(data, &photoDetail)
	return &photoDetail, err
}

func (c *Client) DownloadPhoto(downLoadURL string) ([]byte, error) {
	resp, err := c.client.Get(downLoadURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	return data, err
}
