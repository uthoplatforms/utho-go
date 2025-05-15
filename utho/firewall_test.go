package utho

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestFirewallService_Create_happyPath(t *testing.T) {
	token := "token"
	payload := CreateFirewallParams{}
	_ = faker.FakeData(&payload)

	client, mux, _, teardown := setup(token)
	defer teardown()

	var dummyResponse CreateFirewallResponse
	_ = faker.FakeData(&dummyResponse)

	serverResponse, _ := json.Marshal(dummyResponse)

	mux.HandleFunc("/firewall/create", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, string(serverResponse))
	})

	got, err := client.Firewall().Create(payload)

	assert.Nil(t, err)
	assert.Equal(t, dummyResponse, *got)
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

	firewallId := faker.UUIDDigit()
	var dummyFirewall Firewall
	_ = faker.FakeData(&dummyFirewall)

	serverResponse, _ := json.Marshal(Firewalls{
		Firewalls: []Firewall{dummyFirewall},
		Status:    "success",
	})

	mux.HandleFunc("/firewall/"+firewallId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, string(serverResponse))
	})

	got, _ := client.Firewall().Read(firewallId)
	if !reflect.DeepEqual(*got, dummyFirewall) {
		t.Errorf("Response = %v, want %v", *got, dummyFirewall)
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

	var dummyFirewallList []Firewall
	for i := 0; i < 3; i++ { // Generate a list of 3 dummy firewalls
		var firewall Firewall
		_ = faker.FakeData(&firewall)
		dummyFirewallList = append(dummyFirewallList, firewall)
	}

	serverResponse, _ := json.Marshal(Firewalls{
		Firewalls: dummyFirewallList,
		Status:    "success",
	})

	mux.HandleFunc("/firewall", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, string(serverResponse))
	})

	got, _ := client.Firewall().List()
	if len(got) != len(dummyFirewallList) {
		t.Errorf("Was expecting %d firewalls to be returned, instead got %d", len(dummyFirewallList), len(got))
	}

	if !reflect.DeepEqual(got, dummyFirewallList) {
		t.Errorf("Response = %v, want %v", got, dummyFirewallList)
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

	firewallId := faker.UUIDDigit()
	firewallRuleId := faker.UUIDDigit()
	var dummyFirewall Firewall
	_ = faker.FakeData(&dummyFirewall)

	serverResponse, _ := json.Marshal(Firewalls{
		Firewalls: []Firewall{dummyFirewall},
		Status:    "success",
	})

	mux.HandleFunc("/firewall/"+firewallId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, string(serverResponse))
	})

	var expectedRule FirewallRule
	for _, rule := range dummyFirewall.Rules {
		if rule.ID == firewallRuleId {
			expectedRule = rule
			break
		}
	}

	got, _ := client.Firewall().ReadFirewallRule(firewallId, firewallRuleId)
	if !reflect.DeepEqual(*got, expectedRule) {
		t.Errorf("Response = %v, want %v", *got, expectedRule)
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

	firewallId := faker.UUIDDigit()
	var dummyFirewall Firewall
	_ = faker.FakeData(&dummyFirewall)

	// Ensure the dummy firewall has rules for testing
	if len(dummyFirewall.Rules) == 0 {
		for i := 0; i < 3; i++ { // Add 3 dummy rules
			var rule FirewallRule
			_ = faker.FakeData(&rule)
			dummyFirewall.Rules = append(dummyFirewall.Rules, rule)
		}
	}

	serverResponse, _ := json.Marshal(Firewalls{
		Firewalls: []Firewall{dummyFirewall},
		Status:    "success",
	})

	mux.HandleFunc("/firewall/"+firewallId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, string(serverResponse))
	})

	got, _ := client.Firewall().ListFirewallRules(firewallId)
	if len(got) != len(dummyFirewall.Rules) {
		t.Errorf("Was expecting %d firewall rules to be returned, instead got %d", len(dummyFirewall.Rules), len(got))
	}

	if !reflect.DeepEqual(got, dummyFirewall.Rules) {
		t.Errorf("Response = %v, want %v", got, dummyFirewall.Rules)
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
