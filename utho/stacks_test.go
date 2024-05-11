package utho

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStacksService_Create_happyPath(t *testing.T) {
	token := "token"
	payload := CreateStacksParams{
		Title:       "stackname",
		Description: "desc",
		Images:      "ubuntu-18.10-x86_64",
		IsPublic:    "yes",
		Script:      "echo 'hello'",
	}

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/stacks", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateResponseJson)
	})

	got, err := client.Stacks().Create(payload)

	var want CreateResponse
	_ = json.Unmarshal([]byte(dummyCreateResponseJson), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestStacksService_Create_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Stacks().Create(CreateStacksParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestStacksService_Read_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	stacksId := "11111"
	expectedResponse := dummyReadStacksRes
	serverResponse := dummyStacksServerRes

	mux.HandleFunc("/stacks", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want Stack
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.Stacks().Read(stacksId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestStacksService_Read_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	apikey, err := client.Stacks().Read("someId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if apikey != nil {
		t.Errorf("Was not expecting any apikey to be returned, instead got %v", apikey)
	}
}

func TestStacksService_List_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	expectedResponse := dummyListStacksRes
	serverResponse := dummyStacksServerRes

	mux.HandleFunc("/stacks", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want []Stack
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.Stacks().List()
	if len(got) != len(want) {
		t.Errorf("Was expecting %d stacks to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestStacksService_List_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	stacks, err := client.Stacks().List()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if stacks != nil {
		t.Errorf("Was not expecting any stacks to be returned, instead got %v", stacks)
	}
}

func TestStacksService_Update_happyPath(t *testing.T) {
	token := "token"
	payload := UpdateStacksParams{
		Title:       "stackname updated",
		Description: "desc",
		Images:      "ubuntu-18.10-x86_64",
		IsPublic:    "yes",
		Script:      "echo 'hello'",
	}

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/stacks/"+payload.StackId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPut)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyUpdateResponseJson)
	})

	got, err := client.Stacks().Update(payload)

	var want UpdateResponse
	_ = json.Unmarshal([]byte(dummyUpdateResponseJson), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestStacksService_Update_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Stacks().Update(UpdateStacksParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestStacksService_Delete_happyPath(t *testing.T) {
	token := "token"
	stackId := "someStacksId"

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/stacks/"+stackId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyDeleteResponseJson)
	})

	want := DeleteResponse{Status: "success", Message: "success"}

	got, _ := client.Stacks().Delete(stackId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestStacksService_Delete_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	delResponse, err := client.Stacks().Delete("someStacksId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
	}
}

const dummyReadStacksRes = `{
	"id": "11111",
	"is_owner": "0",
	"title": "qwer",
	"description": "fdsa",
	"distro": [
		"almalinux-9.2-x86_64"
	],
	"logo_url": "",
	"is_public": "1",
	"is_marketplace": "0",
	"status": "0",
	"script": "liufvged",
	"fields": []
}
`

const dummyStacksServerRes = `{
    "stacks": [` + dummyReadStacksRes + `]
}`

const dummyListStacksRes = `[` + dummyReadStacksRes + `]`
