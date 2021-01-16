package sdk500px

import (
	"fmt"
	"testing"
)

func TestNewClient(t *testing.T) {

	c := NewClient("", "")
	err := c.Login()
	fmt.Println(err)
	page, err := c.GetPage(0, 30)
	fmt.Println(err)
	fmt.Println(*page)
	g, err := c.FetchGallery(c.UserId, 1, 20)
	fmt.Println(err)
	fmt.Println(g)
}
