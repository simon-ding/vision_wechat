package sdk500px

import (
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"strconv"
)

type Gallery struct {
	Message string
	Status  string
	Data    []struct {
		Rating      float64
		Title       string
		RatingMax   float64 `json:"ratingMax"`
		ID          string
		UploaderId  string `json:"uploaderId"`
		CreatedTime int64  `json:"createdTime"`
	}
}

func (c *Client) FetchGallery(userID string, page, size int) (*Gallery, error) {
	req, err := c.newRequest("GET", "community/v2/user/profile", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("resourceType", "0,2,4")
	q.Add("imgsize", "p1,p2,p3")
	q.Add("queriedUserId", userID)
	q.Add("page", strconv.Itoa(page))
	q.Add("size", strconv.Itoa(size))
	q.Add("type", "json")
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
	var gallery Gallery
	err = jsoniter.Unmarshal(data, &gallery)
	return &gallery, err
}

type UserInfo struct {
	UserAccountInfo struct {
		UserRoleIds struct {
			Signinvitephotographer       bool `json:"signinvitephotographer"`
			Creativecontractphotographer bool `json:"creativecontractphotographer"`
		} `json:"userRoleIds"`
		LoginType    string `json:"loginType"`
		IsBindSina   bool   `json:"isBindSina"`
		NickName     string `json:"nickName"`
		Firstlogin   int64  `json:"firstlogin"`
		WeixinID     string `json:"weixinId"`
		IsBindEmail  bool   `json:"isBindEmail"`
		IsBindWeixin bool   `json:"isBindWeixin"`
		Avatar       struct {
			A1      string `json:"a1"`
			BaseURL string `json:"baseUrl"`
		} `json:"avatar"`
		UserName    string `json:"userName"`
		UserID      string `json:"userId"`
		AccessToken string `json:"access_token"`
		IsBindPhone bool   `json:"isBindPhone"`
		Phone       string `json:"phone"`
		ID          string `json:"id"`
		SinaID      string `json:"sinaId"`
		Email       string `json:"email"`
		NoRecommend bool   `json:"noRecommend"`
	} `json:"userAccountInfo"`
	Message string `json:"message"`
	Status  string `json:"status"`
}
