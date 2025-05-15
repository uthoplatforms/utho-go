package utho

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"github.com/go-faker/faker/v4"
)

func TestActionService_List_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	var fakeAction Action
	_ = faker.FakeData(&fakeAction)
	actionsResp := Actions{Actions: []Action{fakeAction}, Status: "success"}
	respBytes, _ := json.Marshal(actionsResp)

	mux.HandleFunc("/actions", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		w.Write(respBytes)
	})

	var want []Action
	_ = json.Unmarshal(respBytes, &struct {
		Actions []Action `json:"actions"`
		Status  string   `json:"status"`
	}{})
	want = []Action{fakeAction}

	got, _ := client.Action().List()
	if len(got) != len(want) {
		t.Errorf("Was expecting %d action to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestActionService_List_withFaker(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	// Generate fake actions
	var fakeActions []Action
	for i := 0; i < 2; i++ {
		var a Action
		_ = faker.FakeData(&a)
		fakeActions = append(fakeActions, a)
	}
	actionsResp := Actions{Actions: fakeActions, Status: "success"}
	respBytes, _ := json.Marshal(actionsResp)

	mux.HandleFunc("/actions", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		w.Write(respBytes)
	})

	got, err := client.Action().List()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != len(fakeActions) {
		t.Errorf("Was expecting %d action to be returned, instead got %d", len(fakeActions), len(got))
	}
}

func TestActionService_Read_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	// Use faker to generate a single action for the read endpoint
	var fakeAction Action
	_ = faker.FakeData(&fakeAction)
	respBytes, _ := json.Marshal(fakeAction)

	mux.HandleFunc("/actions/"+fakeAction.ID, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		w.Write(respBytes)
	})

	want := fakeAction

	got, err := client.Action().Read(fakeAction.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}
