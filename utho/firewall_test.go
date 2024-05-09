package utho

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirewallService_Create_happyPath(t *testing.T) {
	token := "token"
	payload := CreateFirewallParams{Name: "example"}

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/firewall/create", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateResponseJson)
	})

	got, err := client.Firewall().Create(payload)

	var want CreateResponse
	_ = json.Unmarshal([]byte(dummyCreateResponseJson), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestFirewallService_Create_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Firewall().Create(CreateFirewallParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestFirewallService_Read_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	firewallId := "11111"
	expectedResponse := dummyReadFirewallRes
	serverResponse := dummyReadFirewallServerRes

	mux.HandleFunc("/firewall/"+firewallId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want Firewall
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.Firewall().Read(firewallId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestFirewallService_Read_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	apikey, err := client.Firewall().Read("someId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if apikey != nil {
		t.Errorf("Was not expecting any apikey to be returned, instead got %v", apikey)
	}
}

func TestFirewallService_List_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	expectedResponse := dummyReadFirewallRes
	serverResponse := dummyListFirewallServerRes

	mux.HandleFunc("/firewall", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want []Firewall
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.Firewall().List()
	if len(got) != len(want) {
		t.Errorf("Was expecting %d firewall to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestFirewallService_List_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	firewall, err := client.Firewall().List()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if firewall != nil {
		t.Errorf("Was not expecting any firewall to be returned, instead got %v", firewall)
	}
}

func TestFirewallService_Delete_happyPath(t *testing.T) {
	token := "token"
	firewallId := "someFirewallId"

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/firewall/"+firewallId+"/destroy", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyDeleteResponseJson)
	})

	want := DeleteResponse{Status: "success", Message: "success"}

	got, _ := client.Firewall().Delete(firewallId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestFirewallService_DeleteFirewall_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	delResponse, err := client.Firewall().Delete("someFirewallId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
	}
}

// Firewall Rules test
func TestFirewallService_CreateFirewallRule_happyPath(t *testing.T) {
	token := "token"

	payload := CreateFirewallRuleParams{
		FirewallId: "111111",
		Type:       "incoming",
		Protocol:   "tcp",
		Port:       "5060",
		Addresses:  "1.1.1.1",
	}

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/firewall/"+payload.FirewallId+"/rule/add", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateBasicResponseJson)
	})

	got, err := client.Firewall().CreateFirewallRule(payload)

	var want CreateResponse
	_ = json.Unmarshal([]byte(dummyCreateBasicResponseJson), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestFirewallService_CreateFirewallRule_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Firewall().CreateFirewallRule(CreateFirewallRuleParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestFirewallService_ReadFirewallRule_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	firewallId := "23432613"
	firewallRuleID := "42344749"
	expectedResponse := dummyReadFirewallRuleRes
	serverResponse := dummyReadFirewallServerRes

	mux.HandleFunc("/firewall/"+firewallId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want FirewallRule
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.Firewall().ReadFirewallRule(firewallId, firewallRuleID)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestFirewallService_ReadFirewallRule_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	apikey, err := client.Firewall().ReadFirewallRule("11111", "122134")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if apikey != nil {
		t.Errorf("Was not expecting any apikey to be returned, instead got %v", apikey)
	}
}

func TestFirewallService_ListFirewallRule_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	firewallId := "11111"
	expectedResponse := dummyReadDomainRes
	serverResponse := dummyListDomainServerRes

	mux.HandleFunc("/firewall/"+firewallId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want []FirewallRule
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.Firewall().ListFirewallRules(firewallId)
	if len(got) != len(want) {
		t.Errorf("Was expecting %d firewallrule to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestFirewallService_ListFirewallRule_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	firewallrule, err := client.Firewall().ListFirewallRules("11111")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if firewallrule != nil {
		t.Errorf("Was not expecting any firewallrule to be returned, instead got %v", firewallrule)
	}
}

func TestFirewallService_DeleteFirewallRule_happyPath(t *testing.T) {
	token := "token"
	firewallId := "someFirewallId"
	firewallRuleId := "53211"

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/firewall/"+firewallId+"/rule/"+firewallRuleId+"/delete", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyDeleteResponseJson)
	})

	want := DeleteResponse{Status: "success", Message: "success"}

	got, _ := client.Firewall().DeleteFirewallRule(firewallId, firewallRuleId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestFirewallService_DeleteFirewallRule_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	delResponse, err := client.Firewall().DeleteFirewallRule("someFirewallRuleName", "123543")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
	}
}

const dummyReadFirewallRes = `{
	"id": "23432613",
	"name": "",
	"created_at": "2024-04-24 21:11:32",
	"rulecount": "5",
	"serverscount": "0",
	"servers": [],
	"rules": [
		{
			"id": "42344749",
			"firewallid": "23432613",
			"type": "incoming",
			"service": "SSH",
			"protocol": "TCP",
			"port": "22",
			"addresses": "0",
			"note": null
		},
		{
			"id": "42344750",
			"firewallid": "23432613",
			"type": "outgoing",
			"service": "PING",
			"protocol": "ICMP",
			"port": "ICMP",
			"addresses": "0",
			"note": null
		}
	],
	"rule": {
		"id": "42344997",
		"firewallid": "23432613",
		"type": "incoming",
		"service": "CUSTOM",
		"protocol": "TCP",
		"port": "23",
		"addresses": "0",
		"note": null
	},
	"scaling_groups": []
}
`

const dummyReadFirewallServerRes = `{
    "firewalls": [
        {
            "id": "23432613",
            "name": "",
            "created_at": "2024-04-24 21:11:32",
            "rulecount": "5",
            "serverscount": "0",
            "servers": [],
            "rules": [
                {
                    "id": "42344749",
                    "firewallid": "23432613",
                    "type": "incoming",
                    "service": "SSH",
                    "protocol": "TCP",
                    "port": "22",
                    "addresses": "0",
                    "note": null
                },
                {
                    "id": "42344750",
                    "firewallid": "23432613",
                    "type": "outgoing",
                    "service": "PING",
                    "protocol": "ICMP",
                    "port": "ICMP",
                    "addresses": "0",
                    "note": null
                }
            ],
            "rule": {
                "id": "42344997",
                "firewallid": "23432613",
                "type": "incoming",
                "service": "CUSTOM",
                "protocol": "TCP",
                "port": "23",
                "addresses": "0",
                "note": null
            },
            "scaling_groups": []
        },
        {
            "id": "23432614",
            "name": "testq",
            "created_at": "2024-04-24 21:22:44",
            "rulecount": "4",
            "serverscount": "50",
            "servers": [
                {
                    "id": null,
                    "firewallid": null,
                    "cloudid": "1277094",
                    "ip": null,
                    "name": "cloudserver-VBl82tPl.mhc",
                    "country": "India",
                    "cc": "in",
                    "city": "Bangalore"
                }
            ],
            "server": {
                "id": null,
                "firewallid": null,
                "cloudid": "1277220",
                "ip": "103.146.242.55",
                "name": "cloudserver-VBl82tPl.mhc",
                "country": "India",
                "cc": "in",
                "city": "Bangalore"
            },
            "rules": [
                {
                    "id": "42344753",
                    "firewallid": "23432614",
                    "type": "incoming",
                    "service": "SSH",
                    "protocol": "TCP",
                    "port": "22",
                    "addresses": "0",
                    "note": null
                }
            ],
            "rule": {
                "id": "42344756",
                "firewallid": "23432614",
                "type": "outgoing",
                "service": "ALL UDP",
                "protocol": "UDP",
                "port": "ALL",
                "addresses": "0",
                "note": null
            },
            "scaling_groups": [
                {
                    "id": "23492452",
                    "name": "Auto-scaling-Ve6zUYjs.uthoq"
                },
                {
                    "id": "23492453",
                    "name": "qwedc"
                }
            ]
        }
    ]
}
`

const dummyListFirewallServerRes = `[
	{
		"firewall": "examqweple.com",
		"nspoint": "NO",
		"created_at": "2024-05-03 21:29:23",
		"firewallrule_count": "2",
		"records": [
			{
				"id": "25244",
				"hostname": "example22.com.examqweple.com",
				"type": "A",
				"value": "1.1.1.1",
				"ttl": "65444",
				"priority": "10"
			},
			{
				"id": "25245",
				"hostname": "example.examqweple.com",
				"type": "A",
				"value": "1.1.12.1",
				"ttl": "65444",
				"priority": "10"
			}
		],
		"record": {
			"id": "25245",
			"hostname": "example.examqweple.com",
			"type": "A",
			"value": "1.1.12.1",
			"ttl": "65444",
			"priority": "10"
		}
	},{
		"firewall": "example2.com",
		"nspoint": "NO",
		"created_at": "2024-05-03 21:29:23",
		"firewallrule_count": "2",
		"records": [
			{
				"id": "25244",
				"hostname": "example.example2.com",
				"type": "A",
				"value": "1.1.1.1",
				"ttl": "65444",
				"priority": "10"
			},
			{
				"id": "25245",
				"hostname": "example.example2.com",
				"type": "A",
				"value": "1.1.12.1",
				"ttl": "65444",
				"priority": "10"
			}
		],
		"record": {
			"id": "25245",
			"hostname": "example22.com.examqweple.com",
			"type": "A",
			"value": "1.1.12.1",
			"ttl": "65444",
			"priority": "10"
		}
	}
]`

const dummyReadFirewallRuleRes = `{
	"id": "42344749",
	"firewallid": "23432613",
	"type": "incoming",
	"service": "SSH",
	"protocol": "TCP",
	"port": "22",
	"addresses": "0",
	"note": null
}
`
