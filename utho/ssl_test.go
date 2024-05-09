package utho

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSslService_Create_happyPath(t *testing.T) {
	token := "token"
	payload := CreateSslParams{
		Name:             "example",
		Type:             "custom",
		CertificateKey:   "sdfjbfke.",
		PrivateKey:       "fds^wer4^r7!w.",
		CertificateChain: "wrfncjks",
	}

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/certificates", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateResponseJson)
	})

	got, err := client.Ssl().Create(payload)

	var want CreateResponse
	_ = json.Unmarshal([]byte(dummyCreateResponseJson), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestSslService_Create_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Ssl().Create(CreateSslParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestSslService_Read_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	sslId := "11111"
	expectedResponse := dummyReadSslRes
	serverResponse := dummySslServerRes

	mux.HandleFunc("/certificates", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want Certificates
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.Ssl().Read(sslId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestSslService_Read_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	apikey, err := client.Ssl().Read("someId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if apikey != nil {
		t.Errorf("Was not expecting any apikey to be returned, instead got %v", apikey)
	}
}

func TestSslService_List_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	expectedResponse := dummyListSslRes
	serverResponse := dummySslServerRes

	mux.HandleFunc("/certificates", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want []Certificates
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.Ssl().List()
	if len(got) != len(want) {
		t.Errorf("Was expecting %d ssl to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestSslService_List_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	ssl, err := client.Ssl().List()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if ssl != nil {
		t.Errorf("Was not expecting any ssl to be returned, instead got %v", ssl)
	}
}

func TestSslService_Delete_happyPath(t *testing.T) {
	token := "token"
	certId := "11111"

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/certificates/"+certId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyDeleteResponseJson)
	})

	want := DeleteResponse{Status: "success", Message: "success"}

	got, _ := client.Ssl().Delete(certId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestSslService_DeleteSsl_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	delResponse, err := client.Ssl().Delete("someSslId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
	}
}

const dummyReadSslRes = `{
	"id": "11111",
	"userid": "11111",
	"cloudid": "0",
	"name": "ssl-etahl5yf",
	"status": "Pending",
	"created_at": "2024-05-09 13:46:05",
	"ip": "103.209.111.111",
	"count": "0"
}
`

const dummySslServerRes = `{
    "certificates": [` + dummyReadSslRes + `]
}`

const dummyListSslRes = `[` + dummyReadSslRes + `]`
