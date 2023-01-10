package common

import (
	"fmt"
	"testing"
)

func TestNewClient(t *testing.T) {

	session, err := NewClient("http://127.0.0.1/api_jsonrpc.php", Login("Admin", "zabbix"))
	if err != nil {
		panic(err)
	}
	fmt.Println(session)
}
func TestNewClient2(t *testing.T) {

	session, err := NewClient("http://127.0.0.1/api_jsonrpc.php", PlaceToken("0768e85be0f6de1a93fe9ee8ed26fdbd7800ddceba3cceb454bd3080788d20e0"))
	if err != nil {
		panic(err)
	}
	fmt.Println(session)
}
