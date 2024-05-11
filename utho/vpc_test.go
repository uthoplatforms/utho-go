package utho

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVpcService_Create_happyPath(t *testing.T) {
	token := "token"
	payload := CreateVpcParams{
		Dcslug:  "innoida",
		Name:    "testq",
		Planid:  "1008",
		Network: "10.200.210.1",
		Size:    "24",
	}

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/vpc/create", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateResponseJson)
	})

	got, err := client.Vpc().Create(payload)

	var want CreateResponse
	_ = json.Unmarshal([]byte(dummyCreateResponseJson), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestVpcService_Create_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Vpc().Create(CreateVpcParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestVpcService_Read_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	vpcId := "qwsdrf1-1bfa-46ef-8b94-f69f3qwszcf"
	expectedResponse := dummyReadVpcRes
	serverResponse := dummyVpcServerRes

	mux.HandleFunc("/vpc", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want Vpc
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.Vpc().Read(vpcId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestVpcService_Read_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	apikey, err := client.Vpc().Read("someId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if apikey != nil {
		t.Errorf("Was not expecting any apikey to be returned, instead got %v", apikey)
	}
}

func TestVpcService_List_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	expectedResponse := dummyListVpcRes
	serverResponse := dummyVpcServerRes

	mux.HandleFunc("/vpc", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want []Vpc
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.Vpc().List()
	if len(got) != len(want) {
		t.Errorf("Was expecting %d stacks to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestVpcService_List_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	stacks, err := client.Vpc().List()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if stacks != nil {
		t.Errorf("Was not expecting any stacks to be returned, instead got %v", stacks)
	}
}

func TestVpcService_Delete_happyPath(t *testing.T) {
	token := "token"
	vpcId := "qwsdrf1-1bfa-46ef-8b94-f69f3qwszcf"

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/vpc/"+vpcId+"/destroy", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyDeleteResponseJson)
	})

	want := DeleteResponse{Status: "success", Message: "success"}

	got, _ := client.Vpc().Delete(vpcId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestVpcService_Delete_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	delResponse, err := client.Vpc().Delete("someVpcId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
	}
}

const dummyReadVpcRes = `{
	"id": "qwsdrf1-1bfa-46ef-8b94-f69f3qwszcf",
	"total": 254,
	"available": 248,
	"network": "10.210.100.0",
	"name": "test",
	"size": "24",
	"dcslug": "innoida",
	"dclocation": {
		"dccc": "in",
		"location": "Delhi (Noida)"
	},
	"is_default": "0"
}`

const dummyVpcServerRes = `{
    "vpc": [` + dummyReadVpcRes + `]
}`

const dummyListVpcRes = `[` + dummyReadVpcRes + `]`
