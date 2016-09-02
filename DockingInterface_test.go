package DockingInterface

import (
	"fmt"
	"testing"
)

func TestA(*testing.T) {
	a := make([]byte, 60)
	apiKey := string(a[0:50])
	agentCode := string(a[0:10])
	loginId := string(a[0:20])
	name := string(a[0:20])
	ref := string(a[0:50])
	odds := string(a[0:2])
	start := string(a[0:16])
	end := string(a[0:16])
	sessionId := string(a[0:50])
	lang := string(a[0:10])
	res, err := SuspendAccount(apiKey, agentCode, loginId)
	print("SuspendAccount", res, err)
	res, err = DisableAccount(apiKey, agentCode, loginId)
	print("DisableAccount", res, err)
	res, err = ActiveAccount(apiKey, agentCode, loginId)
	print("ActiveAccount", res, err)
	res, err = KickoutPlayer(apiKey, agentCode, loginId)
	print("KickoutPlayer", res, err)
	res, err = GetBalance(apiKey, agentCode, loginId)
	print("GetBalance", res, err)
	res, err = Deposit(apiKey, agentCode, loginId, name, 123124.123123, ref, odds)
	print("Deposit", res, err)
	res, err = Withdraw(apiKey, agentCode, loginId, 123124.123123, ref)
	print("Withdraw", res, err)
	res, err = Withdraw(apiKey, agentCode, loginId, 123124.123123, ref)
	print("Withdraw", res, err)
	res, err = CheckTransfer(apiKey, agentCode, ref)
	print("CheckTransfer", res, err)
	res, err = GetTransfers(apiKey, agentCode, start, end, 12312)
	print("GetTransfers", res, err)
	res, err = GetTransfers2(apiKey, agentCode, start, end)
	print("GetTransfers2", res, err)
	res, err = GetCockfightOpenTicket(apiKey, agentCode)
	print("GetCockfightOpenTicket", res, err)
	res, err = GetCockfightOpenTicket2(apiKey, agentCode)
	print("GetCockfightOpenTicket2", res, err)
	res, err = GetCockfightProcessedTicket(apiKey, agentCode, start, end, 123213)
	print("GetCockfightProcessedTicket", res, err)
	res, err = GetCockfightProcessedTicket2(apiKey, agentCode, start, end)
	print("GetCockfightProcessedTicket2", res, err)
	res, err = LogintoGamePart1(apiKey, agentCode, loginId, name, odds)
	print("LogintoGamePart1", res, err)
	res, err = LogintoGamePart2(apiKey, sessionId, lang, loginId)
	print("LogintoGamePart2", res, err)
}

func print(f string, res []byte, err error) {
	if nil == err {
		fmt.Printf("%v res = %s\n", f, res)
	} else {
		fmt.Printf("%v err = %s\n", f, err.Error())
	}
}
