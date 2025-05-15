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
	_ = faker.FakeData(&fakeUser)

	if fakeUser.ID == "" {
		if ints, err := faker.RandomInt(100000, 999999); err == nil && len(ints) > 0 {
			fakeUser.ID = fmt.Sprintf("%d", ints[0])
		} else {
			fakeUser.ID = "1234567"
		}
	}

	expectedResponse, _ := json.Marshal(fakeUser)

	mux.HandleFunc("/account/info", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
	})

	var want User
	_ = json.Unmarshal(expectedResponse, &want)

	got, err := client.Account().Read()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestAccountService_Read_invalidServer(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	mux.HandleFunc("/account/info", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.Account().Read()
	if err == nil || err.Error() != "account service error: Internal Server Error" {
		t.Errorf("Expected error 'account service error: Internal Server Error', got %v", err)
	}
}

func TestAccountService_Read_userNotFound(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	mux.HandleFunc("/account/info", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
	})

	_, err := client.Account().Read()
	if err == nil || err.Error() != "user not found in account information" {
		t.Errorf("Expected error 'user not found in account information', got %v", err)
	}
}
