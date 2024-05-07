package utho

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsoService_Create_happyPath(t *testing.T) {
	token := "token"
	var payload CreateISOParams
	_ = json.Unmarshal([]byte(dummyCreateIsoRequestJson), &payload)

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/iso/add", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateResponseJson)
	})

	got, err := client.ISO().CreateISO(payload)

	var want CreateResponse
	_ = json.Unmarshal([]byte(dummyCreateResponseJson), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestIsoService_Create_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.ISO().CreateISO(CreateISOParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestIsoService_ListAll_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	expectedResponse := dummyIsoRes

	mux.HandleFunc("/iso", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, dummyIsoReq)
	})

	var want []ISO
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.ISO().ListISOs()
	if len(got) != len(want) {
		t.Errorf("Was expecting %d iso to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestIsoService_ListAll_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	iso, err := client.ISO().ListISOs()
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

	got, _ := client.ISO().DeleteISO(isoId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestIsoService_Delete_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	delResponse, err := client.ISO().DeleteISO("someIsoId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
	}
}

const dummyIsoReq = `{
    "isos": [
        {
            "name": "qwdwqqs",
            "file": "SYNiH-197456.iso",
            "size": 6140.98,
            "added_at": "2024-05-05 14:33:36",
            "download": "100",
            "dc": "inmumbaizone2",
            "dclocation": {
                "dccc": "in",
                "location": "Mumbai"
            }
        }
    ]
}`

const dummyIsoRes = `[
        {
            "name": "qwdwqqs",
            "file": "SYNiH-197456.iso",
            "size": 6140.98,
            "added_at": "2024-05-05 14:33:36",
            "download": "100",
            "dc": "inmumbaizone2",
            "dclocation": {
                "dccc": "in",
                "location": "Mumbai"
            }
        }
    ]
`

const dummyCreateIsoRequestJson = `{
    "dcslug": "innoida",
    "url": "https://software.download.prss.microsoft.com/dbazure/Win10_22H2_English_x64v1.iso?t=d7dc55e3-3b50-4d99-a510-32723166ab49&P1=1715194710&P2=601&P3=2&P4=CHCbHgRCO7kWiaI%2blqxfj67KjzaJqo7V4FogqdZ9jikjPtP1QHJGENuQLTXC6FxE3wTPuxFguvHZcmJWGjHiIEyvPptOXi2GTANoggReg%2bABWyFJXQp%2fncY2SHMzz7%2beLEJ7gnTEoY9cu3LnFIr9YcFEwKityfZEJVPzlosk6UbH0sb44W4a54YDjFxyHZmHXvzs13Xq3y7SLoCG7xX9Os8jpcbHv1Q%2bPxLVAnZYBUZgFrqcyW6WzyAuqtGa%2fLLfFs64%2f2TsYDTp9xfHTmcIWIVofMPeO17I1csqq8X2DIHXbSURZNBAP%2b9G%2fAujttaT1LgYCfzJNT93ZLLgA4DumA%3d%3d",
    "name": "dqwd"
}`

const dummyCreateResponseJson = `{
	"id":"111",
    "status": "success",
    "message": "success"
}`

const dummyDeleteResponseJson = `{
    "status": "success",
    "message": "success"
}`
