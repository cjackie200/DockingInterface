package DockingInterface

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const (
	checkFlag = true
)

const (
	urlSuspendAccount               = "https://api2288.s1288.net/suspend_player.aspx"
	urlDisableAccount               = "https://api2288.s1288.net/disable_player.aspx"
	urlActiveAccount                = "https://api2288.s1288.net/active_player.aspx"
	urlKickoutPlayer                = "https://api2288.s1288.net/kickout_player.aspx"
	urlGetBalance                   = "https://api2288.s1288.net/get_balance.aspx"
	urlDeposit                      = "https://api2288.s1288.net/deposit.aspx"
	urlWithdraw                     = "https://api2288.s1288.net/withdraw.aspx"
	urlCheckTransfer                = "https://api2288.s1288.net/check_transfer.aspx"
	urlGetTransfers                 = "https://api2288.s1288.net/get_transfer.aspx"
	urlGetTransfers2                = "https://api2288.s1288.net/get_transfer_2.aspx"
	urlGetCockfightOpenTicket       = "https://api2288.s1288.net/get_cockfight_open_ticket.aspx"
	urlGetCockfightOpenTicket2      = "https://api2288.s1288.net/get_cockfight_open_ticket_2.aspx"
	urlGetCockfightProcessedTicket  = "https://api2288.s1288.net/get_cockfight_processed_ticket.aspx"
	urlGetCockfightProcessedTicket2 = "https://api2288.s1288.net/get_cockfight_processed_ticket_2.aspx"
	urlLogintoGamePart1             = "https://api2288.s1288.net/get_session_id.aspx"
	urlLogintoGamePart2             = "http://www.1288128.net/api/auth_login.aspx"
	urlSetBetLimit                  = "https://api2288.s1288.net/set_bet_limit.aspx"
)

const (
	lenApiKey    = 50
	lenAgentCode = 10
	lenLoginId   = 20
	lenName      = 20
	lenRef       = 50
	lenOdds      = 2
	lenDatetime  = 16
	lenSessionId = 50
	lenLang      = 10
)

const (
	keyApiKey        = "api_key"
	keyAgentCode     = "agent_code"
	keyLoginId       = "login_id"
	keyName          = "name"
	keyAmount        = "amount"
	keyRef           = "ref_no"
	keyOdds          = "odds_type"
	keyStartDatetime = "start_datetime"
	keyEndDatetime   = "end_datetime"
	keyMarker        = "marker"
	keySessionId     = "session_id"
	keyLang          = "lang"
)

var (
	bad                = "bad "
	errorApiKey        = errors.New(bad + keyApiKey)
	errorAgentCode     = errors.New(bad + keyAgentCode)
	errorLoginId       = errors.New(bad + keyLoginId)
	errorName          = errors.New(bad + keyName)
	errorAmount        = errors.New(bad + keyAmount)
	errorRef           = errors.New(bad + keyRef)
	errorOdds          = errors.New(bad + keyOdds)
	errorStartDatetime = errors.New(bad + keyStartDatetime)
	errorEndDatetime   = errors.New(bad + keyEndDatetime)
	errorSessionId     = errors.New(bad + keySessionId)
	errorLang          = errors.New(bad + keyLang)
)

func SuspendAccount(apiKey string, agentCode string, loginId string) ([]byte, error) {
	v := makeApiKeyAgentCodeLoginId(apiKey, agentCode, loginId)
	return requst(urlSuspendAccount, v)
}

func DisableAccount(apiKey string, agentCode string, loginId string) ([]byte, error) {
	v := makeApiKeyAgentCodeLoginId(apiKey, agentCode, loginId)
	return requst(urlDisableAccount, v)
}

func ActiveAccount(apiKey string, agentCode string, loginId string) ([]byte, error) {
	v := makeApiKeyAgentCodeLoginId(apiKey, agentCode, loginId)
	return requst(urlActiveAccount, v)
}

func KickoutPlayer(apiKey string, agentCode string, loginId string) ([]byte, error) {
	v := makeApiKeyAgentCodeLoginId(apiKey, agentCode, loginId)
	return requst(urlKickoutPlayer, v)
}

func GetBalance(apiKey string, agentCode string, loginId string) ([]byte, error) {
	v := makeApiKeyAgentCodeLoginId(apiKey, agentCode, loginId)
	return requst(urlGetBalance, v)
}

func Deposit(apiKey string, agentCode string, loginId string, name string, amount float64, ref string, odds string) ([]byte, error) {
	v := makeApiKeyAgentCodeLoginIdAmountRef(apiKey, agentCode, loginId, amount, ref)
	v.Set(keyName, name)
	if "" != odds {
		v.Set(keyOdds, odds)
	}
	return requst(urlDeposit, v)
}

func Withdraw(apiKey string, agentCode string, loginId string, amount float64, ref string) ([]byte, error) {
	v := makeApiKeyAgentCodeLoginIdAmountRef(apiKey, agentCode, loginId, amount, ref)
	return requst(urlWithdraw, v)
}

func CheckTransfer(apiKey string, agentCode string, ref string) ([]byte, error) {
	v := makeApiKeyAgentCode(apiKey, agentCode)
	v.Set(keyRef, ref)
	return requst(urlCheckTransfer, v)
}

func GetTransfers(apiKey string, agentCode string, start string, end string, marker int) ([]byte, error) {
	v := makeApiKeyAgentStartEnd(apiKey, agentCode, start, end)
	v.Set(keyMarker, strconv.Itoa(marker))
	return requst(urlGetTransfers, v)
}

func GetTransfers2(apiKey string, agentCode string, start string, end string) ([]byte, error) {
	v := makeApiKeyAgentStartEnd(apiKey, agentCode, start, end)
	return requst(urlGetTransfers2, v)
}

func GetCockfightOpenTicket(apiKey string, agentCode string) ([]byte, error) {
	v := makeApiKeyAgentCode(apiKey, agentCode)
	return requst(urlGetCockfightOpenTicket, v)
}

func GetCockfightOpenTicket2(apiKey string, agentCode string) ([]byte, error) {
	v := makeApiKeyAgentCode(apiKey, agentCode)
	return requst(urlGetCockfightOpenTicket2, v)
}

func GetCockfightProcessedTicket(apiKey string, agentCode string, start string, end string, marker int) ([]byte, error) {
	v := makeApiKeyAgentStartEnd(apiKey, agentCode, start, end)
	v.Set(keyMarker, strconv.Itoa(marker))
	return requst(urlGetCockfightProcessedTicket, v)
}

func GetCockfightProcessedTicket2(apiKey string, agentCode string, start string, end string) ([]byte, error) {
	v := makeApiKeyAgentStartEnd(apiKey, agentCode, start, end)
	return requst(urlGetCockfightProcessedTicket2, v)
}

func LogintoGamePart1(apiKey string, agentCode string, loginId string, name string, odds string) ([]byte, error) {
	v := makeApiKeyAgentCodeLoginId(apiKey, agentCode, loginId)
	v.Set(keyName, name)
	if "" != odds {
		v.Set(keyOdds, odds)
	}
	return requst(urlLogintoGamePart1, v)
}

func LogintoGamePart2(apiKey string, sessionId string, lang string, loginId string) ([]byte, error) {
	v := makeApiKeySessionLangLogin(apiKey, sessionId, lang, loginId)
	return requst(urlLogintoGamePart2, v)
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

func makeApiKeyAgentCodeLoginIdAmountRef(apiKey string, agentCode string, loginId string, amount float64, ref string) url.Values {
	v := makeApiKeyAgentCodeLoginId(apiKey, agentCode, loginId)
	v.Set(keyAmount, strconv.FormatFloat(amount, 'f', 2, 64))
	v.Set(keyRef, ref)
	return v
}

func makeApiKeyAgentCodeLoginId(apiKey string, agentCode string, loginId string) url.Values {
	v := makeApiKeyAgentCode(apiKey, agentCode)
	v.Set(keyLoginId, loginId)
	return v
}

func makeApiKeyAgentStartEnd(apiKey string, agentCode string, start string, end string) url.Values {
	v := makeApiKeyAgentCode(apiKey, agentCode)
	v.Set(keyStartDatetime, start)
	v.Set(keyEndDatetime, end)
	return v
}

func makeApiKeyAgentCode(apiKey string, agentCode string) url.Values {
	v := url.Values{}
	v.Set(keyApiKey, apiKey)
	v.Set(keyAgentCode, agentCode)
	return v
}

func makeApiKeySessionLangLogin(apiKey string, sessionId string, lang string, loginId string) url.Values {
	v := url.Values{}
	v.Set(keyApiKey, apiKey)
	v.Set(keySessionId, sessionId)
	v.Set(keyLoginId, loginId)
	if "" != lang {
		v.Set(keyLang, lang)
	}
	return v
}

func httpRequst(url string, param url.Values) ([]byte, error) {
	c := newHttpClient()
	resp, err := c.PostForm(url, param)
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
		err = checkApiKey(value)
	case keyAgentCode:
		err = checkAgentCode(value)
	case keyLoginId:
		err = checkLoginId(value)
	case keyName:
		err = checkName(value)
	case keyRef:
		err = checkRef(value)
	case keyOdds:
		err = checkOdds(value)
	case keyStartDatetime:
		err = checkStartDatetime(value)
	case keyEndDatetime:
		err = checkEndDatetime(value)
	case keySessionId:
		err = checkSessionId(value)
	case keyLang:
		err = checkLang(value)
	default:
	}
	return
}

func checkApiKey(value []string) error {
	for _, apiKey := range value {
		if checkStrOut(apiKey, lenApiKey) {
			return errorApiKey
		}
	}
	return nil
}

func checkAgentCode(value []string) error {
	for _, agentCode := range value {
		if checkStrOut(agentCode, lenAgentCode) {
			return errorAgentCode
		}
	}
	return nil
}

func checkLoginId(value []string) error {
	for _, loginId := range value {
		if checkStrOut(loginId, lenLoginId) {
			return errorLoginId
		}
	}
	return nil
}

func checkName(value []string) error {
	for _, name := range value {
		if checkStrOut(name, lenName) {
			return errorName
		}
	}
	return nil
}

func checkRef(value []string) error {
	for _, ref := range value {
		if checkStrOut(ref, lenRef) {
			return errorRef
		}
	}
	return nil
}

func checkOdds(value []string) error {
	for _, odds := range value {
		if checkStrOut(odds, lenOdds) {
			return errorOdds
		}
	}
	return nil
}

func checkStartDatetime(value []string) error {
	for _, start := range value {
		if checkStrOut(start, lenDatetime) {
			return errorStartDatetime
		}
	}
	return nil
}

func checkEndDatetime(value []string) error {
	for _, end := range value {
		if checkStrOut(end, lenDatetime) {
			return errorEndDatetime
		}
	}
	return nil
}

func checkSessionId(value []string) error {
	for _, session := range value {
		if checkStrOut(session, lenSessionId) {
			return errorSessionId
		}
	}
	return nil
}

func checkLang(value []string) error {
	for _, lang := range value {
		if checkStrOut(lang, lenLang) {
			return errorLang
		}
	}
	return nil
}

func checkStrOut(str string, l int) bool {
	return len(str) > l
}
