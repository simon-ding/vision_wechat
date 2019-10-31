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
		Rating     float64
		Title      string
		RatingMax  float64 `json:"ratingMax"`
		ID         string
		UploaderId string `json:"uploaderId"`
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
