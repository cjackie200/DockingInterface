package DockingInterface

import (
	"fmt"
	"testing"
)

func TestA(*testing.T) {
	a := make([]byte, 60)
	test_apiKey := string(a[0:50])
	test_agentCode := string(a[0:10])
	test_loginId := string(a[0:20])
	res, err := SuspendAccount(test_apiKey, test_agentCode, test_loginId)
	print(res, err)
	res, err = DisableAccount(test_apiKey, test_agentCode, test_loginId)
	print(res, err)
}

func print(res []byte, err error) {
	if nil == err {
		fmt.Printf("res = %s\n", res)
	} else {
		fmt.Printf("err = %s\n", err.Error())
	}
}
