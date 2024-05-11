package utho

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqsService_Create_happyPath(t *testing.T) {
	token := "token"
	payload := CreateSqsParams{
		Name:   "example",
		Dcslug: "innoida",
		Planid: "10195",
	}

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/sqs", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateResponseJson)
	})

	got, err := client.Sqs().Create(payload)

	var want CreateResponse
	_ = json.Unmarshal([]byte(dummyCreateResponseJson), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestSqsService_Create_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Sqs().Create(CreateSqsParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestSqsService_Read_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	sqsId := "11111"
	expectedResponse := dummyReadSqsRes
	serverResponse := dummySqsServerRes

	mux.HandleFunc("/sqs/"+sqsId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want Sqs
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.Sqs().Read(sqsId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestSqsService_Read_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	apikey, err := client.Sqs().Read("someId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if apikey != nil {
		t.Errorf("Was not expecting any apikey to be returned, instead got %v", apikey)
	}
}

func TestSqsService_List_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	expectedResponse := dummyListSqsRes
	serverResponse := dummySqsServerRes

	mux.HandleFunc("/sqs", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want []Sqs
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.Sqs().List()
	if len(got) != len(want) {
		t.Errorf("Was expecting %d sqs to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestSqsService_List_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	sqs, err := client.Sqs().List()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if sqs != nil {
		t.Errorf("Was not expecting any sqs to be returned, instead got %v", sqs)
	}
}

func TestSqsService_Delete_happyPath(t *testing.T) {
	token := "token"
	sqsId := "someSqsId"
	sqsname := "someSqsname"

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/sqs/"+sqsId+"/destroy", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyDeleteResponseJson)
	})

	want := DeleteResponse{Status: "success", Message: "success"}

	got, _ := client.Sqs().Delete(sqsId, sqsname)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestSqsService_Delete_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	delResponse, err := client.Sqs().Delete("someSqsId", "sqsname")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
	}
}

const dummyReadSqsRes = `{
	"id": "729002",
	"userid": "197456",
	"cloudid": "0",
	"name": "sqs-etahl5yf",
	"status": "Pending",
	"created_at": "2024-05-09 13:46:05",
	"ip": "103.209.146.150",
	"count": "0"
}
`

const dummySqsServerRes = `{
    "sqs": [` + dummyReadSqsRes + `]
}`

const dummyListSqsRes = `[` + dummyReadSqsRes + `]`
