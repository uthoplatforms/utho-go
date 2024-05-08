package utho

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiKeyService_Create_happyPath(t *testing.T) {
	token := "token"

	var payload CreateApiKeyParams
	_ = json.Unmarshal([]byte(dummyCreateApiKeyRequestJson), &payload)

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/api/generate", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateApiKeyResponseJson)
	})

	got, err := client.ApiKey().Create(payload)

	var want CreateApiKeyResponse
	_ = json.Unmarshal([]byte(dummyCreateApiKeyResponseJson), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestApiKeyService_Create_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.ApiKey().Create(CreateApiKeyParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestApiKeyService_ListAll_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	expectedResponse := dummyListApiKeyRes
	serverResponse := dummyListApiKeyServerRes

	mux.HandleFunc("/api", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want []ApiKey
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.ApiKey().List()
	if len(got) != len(want) {
		t.Errorf("Was expecting %d apikey to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestApiKeyService_ListAll_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	apikey, err := client.ApiKey().List()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if apikey != nil {
		t.Errorf("Was not expecting any apikey to be returned, instead got %v", apikey)
	}
}

func TestApiKeyService_Delete_happyPath(t *testing.T) {
	token := "token"
	apiKeyId := "someApiKeyId"

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/api/"+apiKeyId+"/delete", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyDeleteResponseJson)
	})

	want := DeleteResponse{Status: "success", Message: "success"}

	got, _ := client.ApiKey().Delete(apiKeyId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestApiKeyService_Delete_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	delResponse, err := client.ApiKey().Delete("someApiKeyId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
	}
}

const dummyCreateApiKeyRequestJson = `{
    "name": "example-name",
	"write": "on"
}`

const dummyCreateApiKeyResponseJson = `{
    "status": "success",
    "apikey": "api_key_value",
    "message": "New API has been created"
}`

const dummyListApiKeyRes = `[
	{
		"id": "10000",
		"name": "name",
		"write": "1",
		"created_at": "2024-04-22 01:12:36"
	},
	{
		"id": "10001",
		"name": "nam2",
		"write": "0",
		"created_at": "2024-04-22 01:16:51"
	}
]`

const dummyListApiKeyServerRes = `{
	"status": "success",
	"api": [
		{
			"id": "10000",
			"name": "name",
			"write": "1",
			"created_at": "2024-04-22 01:12:36"
		},
		{
			"id": "10001",
			"name": "nam2",
			"write": "0",
			"created_at": "2024-04-22 01:16:51"
		}
	]
}`
