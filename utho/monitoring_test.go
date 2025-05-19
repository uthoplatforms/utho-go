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

func TestMonitoringService_CreateAlert_happyPath(t *testing.T) {
	token := "token"
	payload := CreateAlertParams{}
	err := faker.FakeData(&payload)
	if err != nil {
		t.Fatalf("Failed to fake data for CreateAlertParams: %v", err)
	}

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/alert", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		// We're no longer relying on dummy JSON constants, so we can mock a basic successful response
		fmt.Fprint(w, `{"status":"success","message":"Alert created successfully"}`)
	})

	got, err := client.Monitoring().CreateAlert(payload)

	var want BasicResponse
	// Adjust the expected response to match the mocked successful response
	_ = json.Unmarshal([]byte(`{"status":"success","message":"Alert created successfully"}`), &want)

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

	alertId := faker.UUIDDigit()

	// Create a faker instance for the expected Alert object
	var want Alert
	err := faker.FakeData(&want)
	if err != nil {
		t.Fatalf("Failed to fake data for Alert: %v", err)
	}
	want.ID = alertId // Ensure the ID matches the requested one

	// Marshal the faked 'want' object into JSON for the server response
	serverResponse, err := json.Marshal(map[string][]Alert{"alerts": {want}})
	if err != nil {
		t.Fatalf("Failed to marshal faked alert: %v", err)
	}

	mux.HandleFunc("/alert", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, string(serverResponse))
	})

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

	var alert1 Alert
	err := faker.FakeData(&alert1)
	if err != nil {
		t.Fatalf("Failed to fake data for Alert: %v", err)
	}

	var alert2 Alert
	err = faker.FakeData(&alert2)
	if err != nil {
		t.Fatalf("Failed to fake data for Alert: %v", err)
	}

	want := []Alert{alert1, alert2} // Create a slice of faked alerts

	// Marshal the slice of faked alerts into JSON for the server response
	serverResponse, err := json.Marshal(map[string][]Alert{"alerts": want})
	if err != nil {
		t.Fatalf("Failed to marshal faked alerts: %v", err)
	}

	mux.HandleFunc("/alert", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, string(serverResponse))
	})

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
// 	alertId := faker.UUIDDigit()
// 	client, mux, _, teardown := setup(token)
// 	defer teardown()

// 	mux.HandleFunc("/alert/"+alertId, func(w http.ResponseWriter, req *http.Request) {
// 		testHttpMethod(t, req, "DELETE")
// 		testHeader(t, req, "Authorization", "Bearer "+token)
// 		fmt.Fprint(w, `{"status":"success","message":"success"}`)
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

func TestMonitoringService_CreateContact_happyPath(t *testing.T) {
	token := "token"
	payload := CreateContactParams{}
	err := faker.FakeData(&payload)
	if err != nil {
		t.Fatalf("Failed to fake data for CreateContactParams: %v", err)
	}

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/alert/contact/add", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, `{"status":"success","message":"Contact created successfully"}`)
	})

	got, err := client.Monitoring().CreateContact(payload)

	var want CreateResponse
	_ = json.Unmarshal([]byte(`{"status":"success","message":"Contact created successfully"}`), &want)

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

	contactId := faker.UUIDDigit()

	// Create a faker instance for the expected Contact object
	var want Contact
	err := faker.FakeData(&want)
	if err != nil {
		t.Fatalf("Failed to fake data for Contact: %v", err)
	}
	want.ID = contactId // Ensure the ID matches the requested one

	// Marshal the faked 'want' object into JSON for the server response
	serverResponse, err := json.Marshal(map[string][]Contact{"contacts": {want}})
	if err != nil {
		t.Fatalf("Failed to marshal faked contact: %v", err)
	}

	mux.HandleFunc("/alert/contact/list", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, string(serverResponse))
	})

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

	var contact1 Contact
	err := faker.FakeData(&contact1)
	if err != nil {
		t.Fatalf("Failed to fake data for Contact: %v", err)
	}

	var contact2 Contact
	err = faker.FakeData(&contact2)
	if err != nil {
		t.Fatalf("Failed to fake data for Contact: %v", err)
	}

	want := []Contact{contact1, contact2} // Create a slice of faked contacts

	// Marshal the slice of faked contacts into JSON for the server response
	serverResponse, err := json.Marshal(map[string][]Contact{"contacts": want})
	if err != nil {
		t.Fatalf("Failed to marshal faked contacts: %v", err)
	}

	mux.HandleFunc("/alert/contact/list", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, string(serverResponse))
	})

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
	contactId := faker.UUIDDigit()

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/alert/contact/"+contactId+"/delete", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, `{"status":"success","message":"success"}`)
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
