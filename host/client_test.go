package host

import (
	"fmt"
	"github.com/weihxa/zabbix-go/common"
	"testing"
)

func TestClient_GetHosts(t *testing.T) {

	session, err := common.NewClient("http://127.0.0.1/api_jsonrpc.php", common.PlaceToken("0768e85be0f6de1a93fe9ee8ed26fdbd7800ddceba3cceb454bd3080788d20e0"))
	if err != nil {
		panic(err)
	}
	client, _ := NewClient(session)
	hosts, err := client.GetHosts(HostGetParams{
		GroupIDs:     []string{"1", "4"},
		SelectGroups: common.SelectCount,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(hosts)
}
