package sdk500px

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	Cookie    string
	client    *http.Client
	userAgent string
	baseUrl   string
}

func NewClientUseCookie(cookie string) *Client {
	return &Client{
		Cookie:    cookie,
		client:    http.DefaultClient,
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.70 Safari/537.36",
		baseUrl:   "https://500px.me",
	}
}

type Page struct {
	Data    []photoData
	Message string
	Status  string
}
type uploaderInfo struct {
	NickName string `json:"nickName"`
}

type photoData struct {
	ID           string `json:"id"`
	UploaderID   string `json:"uploaderId"`
	Title        string
	UploaderInfo uploaderInfo `json:"uploaderInfo"`
}

type Response struct {
	Data    interface{}
	Message string
	Status  string
}

func (c *Client) OwnerID() string {
	header := http.Header{}
	header.Add("Cookie", c.Cookie)
	req := http.Request{Header: header}
	userId, err := req.Cookie("userId")
	if err != nil {
		logrus.Error(err)
	}
	return userId.Value
}

func (c *Client) GetPage(page int, size int) (*Page, error) {
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
	var page1 Page

	err = jsoniter.Unmarshal(data, &page1)
	if err != nil {
		return nil, err
	}
	if page1.Status != "200" {
		return nil, fmt.Errorf("%s", page1.Message)
	}
	return &page1, nil
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
		return fmt.Errorf("%s", res.Message)
	}
	return nil
}

func (c *Client) newRequest(method, refUrl string, body io.Reader) (*http.Request, error) {
	baseUrl, err := url.Parse(c.baseUrl)
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
	req.Header.Add("Cookie", c.Cookie)
	req.Header.Add("User-Agent", c.userAgent)
	return req, nil
}
