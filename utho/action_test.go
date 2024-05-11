package utho

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestActionService_List_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	expectedResponse := dummyListActionRes
	serverResponse := dummyActionServerRes

	mux.HandleFunc("/actions", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want []Action
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.Action().List()
	if len(got) != len(want) {
		t.Errorf("Was expecting %d action to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

const dummyReadActionRes = `{
	"userid": "11111",
	"id": "124214",
	"action": "start",
	"resource_type": "cloud",
	"resource_id": "1277721",
	"started_at": "2024-05-11 07:00:28",
	"completed_at": "0000-00-00 00:00:00",
	"process": "95",
	"status": "Support"
}`

const dummyActionServerRes = `{
    "actions": [` + dummyReadActionRes + `]
}`

const dummyListActionRes = `[` + dummyReadActionRes + `]`
