package utho

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/go-faker/faker/v4"
)

func TestAccountService_Read_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	var fakeUser User
	err := faker.FakeData(&fakeUser)
	if err != nil {
		fmt.Println(err)
	}

	expectedResponse, _ := json.Marshal(struct {
		User   User   `json:"user"`
		Status string `json:"status"`
	}{
		User:   fakeUser,
		Status: "success",
	})

	mux.HandleFunc("/account/info", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		w.WriteHeader(http.StatusOK)
		w.Write(expectedResponse)
	})

	want := fakeUser

	got, err := client.Account().Read()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestAccountService_Read_userNotFound(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	mux.HandleFunc("/account/info", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"success","user":null}`))
	})

	_, err := client.Account().Read()
	if err == nil || err.Error() != "user not found in account information" {
		t.Errorf("Expected error 'user not found in account information', got %v", err)
	}
}
