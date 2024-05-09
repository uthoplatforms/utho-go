package utho

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMonitoringService_CreateAlert_happyPath(t *testing.T) {
	token := "token"
	payload := CreateAlertParams{
		Compare:  "below",
		Contacts: "27",
		For:      "5m",
		Name:     "wqe",
		RefIds:   "1277623",
		RefType:  "cloud",
		Status:   "Active",
		Type:     "cpu",
		Value:    "23",
	}

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/alert", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateBasicResponseJson)
	})

	got, err := client.Monitoring().client.Monitoring().CreateAlert(payload)

	var want BasicResponse
	_ = json.Unmarshal([]byte(dummyCreateBasicResponseJson), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestMonitoringService_CreateAlert_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Monitoring().CreateAlert(CreateAlertParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestMonitoringService_ReadAlert_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	alertId := "11111"
	expectedResponse := dummyReadAlertRes
	serverResponse := dummyReadAlertServerRes

	mux.HandleFunc("/alert", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want Alert
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.Monitoring().ReadAlert(alertId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestMonitoringService_ReadAlert_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	apikey, err := client.Monitoring().ReadAlert("someId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if apikey != nil {
		t.Errorf("Was not expecting any apikey to be returned, instead got %v", apikey)
	}
}

func TestMonitoringService_ListAlert_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	expectedResponse := dummyReadAlertRes
	serverResponse := "[" + dummyReadAlertServerRes + "]"

	mux.HandleFunc("/alert", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want []Alert
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.Monitoring().ListAlerts()
	if len(got) != len(want) {
		t.Errorf("Was expecting %d alert to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestMonitoringService_ListAlert_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	alert, err := client.Monitoring().ListAlerts()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if alert != nil {
		t.Errorf("Was not expecting any alert to be returned, instead got %v", alert)
	}
}

// func TestMonitoringService_DeleteAlert_happyPath(t *testing.T) {
// 	token := "token"
// 	alertId := "someAlertId"

// 	client, mux, _, teardown := setup(token)
// 	defer teardown()

// 	mux.HandleFunc("/alert/"+alertId, func(w http.ResponseWriter, req *http.Request) {
// 		testHttpMethod(t, req, "DELETE")
// 		testHeader(t, req, "Authorization", "Bearer "+token)
// 		fmt.Fprint(w, dummyDeleteResponseJson)
// 	})

// 	want := DeleteResponse{Status: "success", Message: "success"}

// 	got, _ := client.Monitoring().DeleteAlert(alertId)
// 	if !reflect.DeepEqual(*got, want) {
// 		t.Errorf("Response = %v, want %v", *got, want)
// 	}
// }

// func TestMonitoringService_DeleteAlert_invalidServer(t *testing.T) {
// 	client, _ := NewClient("token")

// 	delResponse, err := client.Monitoring().Delete("someAlertId")
// 	if err == nil {
// 		t.Errorf("Expected error to be returned")
// 	}
// 	if delResponse != nil {
// 		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
// 	}
// }

// Contact
func TestMonitoringService_CreateContact_happyPath(t *testing.T) {
	token := "token"
	payload := CreateContactParams{
		Email:        "23@dwq.cw",
		Mobilenumber: "123456",
		Name:         "23",
		Status:       "1",
	}

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/alert/contact/add", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateResponseJson)
	})

	got, err := client.Monitoring().CreateContact(payload)

	var want CreateResponse
	_ = json.Unmarshal([]byte(dummyCreateResponseJson), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestMonitoringService_CreateContact_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Monitoring().CreateContact(CreateContactParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestMonitoringService_ReadContact_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	contactId := "11111"
	expectedResponse := dummyReadContactRes
	serverResponse := dummyReadContactServerRes

	mux.HandleFunc("/alert/contact/list", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want Contact
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.Monitoring().ReadContact(contactId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestMonitoringService_ReadContact_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	apikey, err := client.Monitoring().ReadContact("someId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if apikey != nil {
		t.Errorf("Was not expecting any apikey to be returned, instead got %v", apikey)
	}
}

func TestMonitoringService_ListContact_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	expectedResponse := dummyReadContactRes
	serverResponse := "[" + dummyReadContactServerRes + "]"

	mux.HandleFunc("/alert/contact/list", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want []Contact
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.Monitoring().ListContacts()
	if len(got) != len(want) {
		t.Errorf("Was expecting %d contact to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestMonitoringService_ListContact_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	contact, err := client.Monitoring().ListContacts()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if contact != nil {
		t.Errorf("Was not expecting any contact to be returned, instead got %v", contact)
	}
}

func TestMonitoringService_DeleteContact_happyPath(t *testing.T) {
	token := "token"
	contactId := "someContactId"

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/alert/contact/"+contactId+"/delete", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyDeleteResponseJson)
	})

	want := DeleteResponse{Status: "success", Message: "success"}

	got, _ := client.Monitoring().DeleteContact(contactId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestMonitoringService_DeleteContact_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	delResponse, err := client.Monitoring().DeleteContact("someContactId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
	}
}

const dummyReadAlertServerRes = `{
    "alerts": [
        {
            "id": "11111",
            "type": "cpu",
            "name": "wqew",
            "ref_ids": "12344",
            "ref_type": "cloud",
            "compare": "below",
            "value": "231",
            "for": "5m",
            "contacts": "24",
            "status": "1"
        }
    ]
}`

const dummyReadAlertRes = `{
	"id": "11111",
	"type": "cpu",
	"name": "wqew",
	"ref_ids": "12344",
	"ref_type": "cloud",
	"compare": "below",
	"value": "231",
	"for": "5m",
	"contacts": "24",
	"status": "1"
}`

const dummyReadContactServerRes = `{
    "contacts": [
        {
            "id": "11111",
            "name": "tested",
            "email": "test22@test.com",
            "slack": "",
            "mobilenumber": "11111111"
        }
    ]
}`

const dummyReadContactRes = `{
	"id": "11111",
	"name": "tested",
	"email": "test22@test.com",
	"slack": "",
	"mobilenumber": "11111111"
}`
