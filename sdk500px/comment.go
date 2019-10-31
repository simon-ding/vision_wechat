package sdk500px

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"
)

type Comments struct {
	CommentCount int `json:"commentCount"`
	Message      string
	Status       string
	Comments     []*Comment
}

type Comment struct {
	ChildComments []Comment `json:"childComments"`
	CountLike     int       `json:"countLike"`
	CreateDate    int       `json:"createDate"`
	FanyiFlag     int       `json:"fanyiFlag"`
	ID            string
	IP            string
	Like          bool
	Message       string //回复消息
	MessageFanyi  string `json:"messageFanyi"`
	ParentId      string `json:"parentId"`
	PlatformType  int    `json:"platformType"`
	ResourceId    string `json:"resourceId"`
	Sort          int
	State         int
	Type          int
	UserId        string   `json:"userId"`
	UserInfo      userInfo `json:"userInfo"`
}
type userInfo struct {
	Avatar struct {
		A1      string
		BaseUrl string `json:"baseUrl"`
	}
	ID       string
	NickName string `json:"nickName"`
	UserName string `json:"userName"`
}

func (c *Client) FetchComments(photoID string, page, size int) (*Comments, error) {
	req, err := c.newRequest("GET", "community/comment/list", nil)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	q := req.URL.Query()
	q.Add("type", "json")
	q.Add("page", strconv.Itoa(page))
	q.Add("size", strconv.Itoa(size))
	q.Add("resId", photoID)
	req.URL.RawQuery = q.Encode()
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var com Comments
	err = jsoniter.Unmarshal(data, &com)
	return &com, err
}

func (c *Client) ReplyComment(comment *Comment, msg string) error {
	form := url.Values{}
	form.Add("message", msg)
	form.Add("resId", comment.ResourceId)
	form.Add("toUserId", comment.UserId)
	form.Add("parentId", comment.ID)
	return c.doComment(form)
}
func (c *Client) Comment(message string, resID string) error {
	form := url.Values{}
	form.Add("message", message)
	form.Add("resId", resID)
	return c.doComment(form)
}

func (c Client) doComment(form url.Values) error {
	if req, err := c.newRequest("POST", "/community/comment/add", strings.NewReader(form.Encode())); err != nil {
		return err
	} else {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		if resp, err := c.client.Do(req); err != nil {
			return err
		} else {
			defer resp.Body.Close()
			data, _ := ioutil.ReadAll(resp.Body)
			var ret Response
			err := jsoniter.Unmarshal(data, &ret)
			if err != nil {
				return err
			}
			if ret.Status == "200" {
				return nil
			}
			return fmt.Errorf("%s", ret.Message)
		}
	}

}
