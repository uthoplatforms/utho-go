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

func TestEbsService_Create_happyPath(t *testing.T) {
	token := "token"
	payload := CreateEBSParams{}
	_ = faker.FakeData(&payload)

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/ebs", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateResponseJson)
	})

	got, err := client.Ebs().Create(payload)

	var want CreateResponse
	_ = json.Unmarshal([]byte(dummyCreateResponseJson), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestEbsService_Create_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Ebs().Create(CreateEBSParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestEbsService_Read_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	ebsId := faker.UUIDDigit()
	var dummyEbs Ebs
	_ = faker.FakeData(&dummyEbs)

	serverResponse, _ := json.Marshal(EBSs{
		Ebs:    []Ebs{dummyEbs},
		Status: "success",
	})

	mux.HandleFunc("/ebs/"+ebsId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, string(serverResponse))
	})

	got, _ := client.Ebs().Read(ebsId)
	if !reflect.DeepEqual(*got, dummyEbs) {
		t.Errorf("Response = %v, want %v", *got, dummyEbs)
	}
}

func TestEbsService_Read_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	apikey, err := client.Ebs().Read("someId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if apikey != nil {
		t.Errorf("Was not expecting any apikey to be returned, instead got %v", apikey)
	}
}

func TestEbsService_List_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	var dummyEbsList []Ebs
	for i := 0; i < 3; i++ {
		var ebs Ebs
		_ = faker.FakeData(&ebs)
		dummyEbsList = append(dummyEbsList, ebs)
	}

	serverResponse, _ := json.Marshal(EBSs{
		Ebs:    dummyEbsList,
		Status: "success",
	})

	mux.HandleFunc("/ebs", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, string(serverResponse))
	})

	got, _ := client.Ebs().List()
	if len(got) != len(dummyEbsList) {
		t.Errorf("Was expecting %d ebs to be returned, instead got %d", len(dummyEbsList), len(got))
	}

	if !reflect.DeepEqual(got, dummyEbsList) {
		t.Errorf("Response = %v, want %v", got, dummyEbsList)
	}
}

func TestEbsService_List_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	ebs, err := client.Ebs().List()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if ebs != nil {
		t.Errorf("Was not expecting any ebs to be returned, instead got %v", ebs)
	}
}

func TestEbsService_Delete_happyPath(t *testing.T) {
	token := "token"
	ebsId := "someEbsId"

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/ebs/"+ebsId+"/destroy", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyDeleteResponseJson)
	})

	want := DeleteResponse{Status: "success", Message: "success"}

	got, _ := client.Ebs().Delete(ebsId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestEbsService_Delete_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	delResponse, err := client.Ebs().Delete("someEbsId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
	}
}
