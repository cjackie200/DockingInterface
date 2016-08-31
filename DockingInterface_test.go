package DockingInterface

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

const (
	UrlSuspendAccount = "https://api2288.s1288.net/suspend_player.aspx"
)

func NewHttpClient() *http.Client {
	return &http.Client{}
}

func TestA(*testing.T) {
	c := NewHttpClient()
	v := url.Values{}
	v.Set("api_key", "11111111111111111111111111111111111111111111111111")
	v.Set("agent_code", "1111111111")
	v.Set("login_id", "1111111111")

	resp, err := c.PostForm(UrlSuspendAccount, v)
	if nil != err {
		fmt.Printf("err = %v\n", err.Error())
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("respone = %v\n", string(data))
	}
	resp.Body.Close()
}

func TestB(*testing.T) {
	v := url.Values{}
	v.Set("api_key", "11111111111111111111111111111111111111111111111111")
	v.Set("agent_code", "1111111111")
	v.Set("login_id", "1111111111")
	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{}
	req, _ := http.NewRequest("POST", UrlSuspendAccount, body)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8") //这个一定要加，不加form的值post不过去，被坑了两小时
	// fmt.Printf("%+v\n", req)                                                         //看下发送的结构

	resp, err := client.Do(req) //发送
	data, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close() //一定要关闭resp.Body
	fmt.Println(string(data), err)
}

var c = NewHttpClient()
var v = NewValue()

func NewValue() url.Values {
	v := url.Values{}
	v.Set("api_key", "11111111111111111111111111111111111111111111111111")
	v.Set("agent_code", "1111111111")
	v.Set("login_id", "1111111111")
	return v
}
func BenchmarkA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resp, _ := c.PostForm(UrlSuspendAccount, v)
		resp.Body.Close()
	}
}

func BenchmarkB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
		req, _ := http.NewRequest("POST", UrlSuspendAccount, body)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8") //这个一定要加，不加form的值post不过去，被坑了两小时
		resp, _ := c.Do(req)                                                              //发送
		if resp != nil {
			resp.Body.Close()
		}
	}
}
