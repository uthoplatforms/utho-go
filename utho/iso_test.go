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

func TestIsoService_Create_happyPath(t *testing.T) {
	token := "token"
	var payload CreateISOParams
	_ = faker.FakeData(&payload)

	client, mux, _, teardown := setup(token)
	defer teardown()

	var dummyResponse CreateResponse
	_ = faker.FakeData(&dummyResponse)
	dummyResponse.Status = "success"

	serverResponse, _ := json.Marshal(dummyResponse)

	mux.HandleFunc("/iso/add", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(serverResponse))
	})

	got, err := client.ISO().Create(payload)

	assert.Nil(t, err)
	assert.Equal(t, dummyResponse, *got)
}

func TestIsoService_Create_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.ISO().Create(CreateISOParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestIsoService_ListAll_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	var dummyIsoList []ISO
	for i := 0; i < 3; i++ { // Generate a list of 3 dummy ISOs
		var iso ISO
		_ = faker.FakeData(&iso)
		dummyIsoList = append(dummyIsoList, iso)
	}

	serverResponse, _ := json.Marshal(ISOs{
		ISOs:    dummyIsoList,
		Status:  "success",
		Message: "success",
	})

	mux.HandleFunc("/iso", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, string(serverResponse))
	})

	got, _ := client.ISO().List()
	if len(got) != len(dummyIsoList) {
		t.Errorf("Was expecting %d ISOs to be returned, instead got %d", len(dummyIsoList), len(got))
	}

	if !reflect.DeepEqual(got, dummyIsoList) {
		t.Errorf("Response = %v, want %v", got, dummyIsoList)
	}
}

func TestIsoService_ListAll_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	iso, err := client.ISO().List()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if iso != nil {
		t.Errorf("Was not expecting any iso to be returned, instead got %v", iso)
	}
}

func TestIsoService_Delete_happyPath(t *testing.T) {
	token := "token"
	isoId := "someIsoId"

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/iso/"+isoId+"/delete", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyDeleteResponseJson)
	})

	want := DeleteResponse{Status: "success", Message: "success"}

	got, _ := client.ISO().Delete(isoId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestIsoService_Delete_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	delResponse, err := client.ISO().Delete("someIsoId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
	}
}
