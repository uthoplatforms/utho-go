package utho

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestApiKeyService_Create_happyPath(t *testing.T) {
	token := "token"

	// Use faker for payload and response
	var payload CreateApiKeyParams
	_ = faker.FakeData(&payload)

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp CreateApiKeyResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/api/generate", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes) // <-- Fix: actually write the response!
	})

	got, err := client.ApiKey().Create(payload)

	var want CreateApiKeyResponse
	_ = json.Unmarshal(respBytes, &want)

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

func TestApiKeyService_List_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	// Generate fake API keys for both server and expected response
	var fakeApiKeys []ApiKey
	for i := 0; i < 2; i++ {
		var k ApiKey
		_ = faker.FakeData(&k)
		fakeApiKeys = append(fakeApiKeys, k)
	}

	expectedResponse := ApiKeys{Status: "success", API: fakeApiKeys}
	respBytes, _ := json.Marshal(expectedResponse)

	mux.HandleFunc("/api", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		w.WriteHeader(http.StatusOK)
		w.Write(respBytes) // Write the response
	})

	got, err := client.ApiKey().List()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != len(fakeApiKeys) {
		t.Errorf("Was expecting %d apikey to be returned, instead got %d", len(fakeApiKeys), len(got))
	}

	if !reflect.DeepEqual(got, fakeApiKeys) {
		t.Errorf("Response = %v, want %v", got, fakeApiKeys)
	}
}

func TestApiKeyService_List_invalidServer(t *testing.T) {
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

	expectedResponse := DeleteResponse{Status: "success", Message: "success"}
	respBytes, _ := json.Marshal(expectedResponse)

	mux.HandleFunc("/api/"+apiKeyId+"/delete", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.WriteHeader(http.StatusOK)
		w.Write(respBytes) // Write the response
	})

	got, err := client.ApiKey().Delete(apiKeyId)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(*got, expectedResponse) {
		t.Errorf("Response = %v, want %v", *got, expectedResponse)
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
