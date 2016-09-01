package DockingInterface

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	checkFlag = true
)

const (
	UrlSuspendAccount = "https://api2288.s1288.net/suspend_player.aspx"
	UrlDisableAccount = "https://api2288.s1288.net/disable_player.aspx"
)

const (
	lenApiKey    = 50
	lenAgentCode = 10
	lenLoginId   = 20
)

const (
	keyApiKey    = "api_key"
	keyAgentCode = "agent_code"
	keyLoginId   = "login_id"
)

var (
	bad            = "bad "
	errorApiKey    = errors.New(bad + keyApiKey)
	errorAgentCode = errors.New(bad + keyAgentCode)
	errorLoginId   = errors.New(bad + keyLoginId)
)

func SuspendAccount(apiKey string, agentCode string, loginId string) ([]byte, error) {
	v := url.Values{}
	v.Set(keyApiKey, apiKey)
	v.Set(keyAgentCode, agentCode)
	v.Set(keyLoginId, loginId)
	return requst(UrlSuspendAccount, v)
}

func DisableAccount(apiKey string, agentCode string, loginId string) ([]byte, error) {
	v := url.Values{}
	v.Set(keyApiKey, apiKey)
	v.Set(keyAgentCode, agentCode)
	v.Set(keyLoginId, loginId)
	return requst(UrlDisableAccount, v)
}

func newHttpClient() *http.Client {
	return &http.Client{}
}

func requst(url string, v url.Values) ([]byte, error) {
	if checkFlag {
		err := checkParameters(v)
		if nil == err {
			return httpRequst(url, v)
		} else {
			return nil, err
		}
	} else {
		return httpRequst(url, v)
	}
}

func httpRequst(url string, param url.Values) ([]byte, error) {
	c := newHttpClient()
	resp, err := c.PostForm(UrlSuspendAccount, param)
	if nil != err {
		return nil, err
	} else {
		data, err2 := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		return data, err2
	}
}

func checkParameters(v url.Values) (err error) {
	for key, value := range v {
		err = checkUrlValues(key, value)
		if nil != err {
			return
		}
	}
	return
}

func checkUrlValues(key string, value []string) (err error) {
	switch key {
	case keyApiKey:
		err = checkApiKeyAll(value)
	case keyAgentCode:
		err = checkAgentCodeAll(value)
	case keyLoginId:
		err = checkLoginIdAll(value)
	default:
	}
	return
}

func checkApiKeyAll(value []string) error {
	for _, apiKey := range value {
		if checkApiKey(apiKey) {
			return errorApiKey
		}
	}
	return nil
}

func checkApiKey(apiKey string) bool {
	return lenApiKey != len(apiKey)
}

func checkAgentCodeAll(value []string) error {
	for _, agentCode := range value {
		if checkAgentCode(agentCode) {
			return errorAgentCode
		}
	}
	return nil
}

func checkAgentCode(agentCode string) bool {
	return lenAgentCode != len(agentCode)
}

func checkLoginIdAll(value []string) error {
	for _, loginId := range value {
		if checkLoginId(loginId) {
			return errorAgentCode
		}
	}
	return nil
}

func checkLoginId(loginId string) bool {
	return lenLoginId != len(loginId)
}
