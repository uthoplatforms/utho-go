package utho

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTargetGroupService_Create_happyPath(t *testing.T) {
	token := "token"
	payload := CreateTargetGroupParams{
		Name:                "d12d",
		Protocol:            "HTTP",
		Port:                "1",
		HealthCheckPath:     "1",
		HealthCheckProtocol: "HTTP",
		HealthCheckInterval: "1",
		HealthCheckTimeout:  "1",
		HealthyThreshold:    "1",
		UnhealthyThreshold:  "1",
	}

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/targetgroup", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateResponseJson)
	})

	got, err := client.TargetGroup().Create(payload)

	var want CreateResponse
	_ = json.Unmarshal([]byte(dummyCreateResponseJson), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestTargetGroupService_Create_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.TargetGroup().Create(CreateTargetGroupParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestTargetGroupService_Read_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	targetgroupId := "11111"
	expectedResponse := dummyReadTargetGroupRes
	serverResponse := dummyTargetGroupServerRes

	mux.HandleFunc("/targetgroup", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want TargetGroup
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.TargetGroup().Read(targetgroupId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestTargetGroupService_Read_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	apikey, err := client.TargetGroup().Read("someId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if apikey != nil {
		t.Errorf("Was not expecting any apikey to be returned, instead got %v", apikey)
	}
}

func TestTargetGroupService_List_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	expectedResponse := dummyListTargetGroupRes
	serverResponse := dummyTargetGroupServerRes

	mux.HandleFunc("/targetgroup", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want []TargetGroup
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.TargetGroup().List()
	if len(got) != len(want) {
		t.Errorf("Was expecting %d targetgroup to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestTargetGroupService_List_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	targetgroup, err := client.TargetGroup().List()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if targetgroup != nil {
		t.Errorf("Was not expecting any targetgroup to be returned, instead got %v", targetgroup)
	}
}

func TestTargetGroupService_Update_happyPath(t *testing.T) {
	token := "token"
	payload := UpdateTargetGroupParams{
		Name:                "d12d updated",
		Protocol:            "HTTP",
		Port:                "1",
		HealthCheckPath:     "1",
		HealthCheckProtocol: "HTTP",
		HealthCheckInterval: "1",
		HealthCheckTimeout:  "1",
		HealthyThreshold:    "1",
		UnhealthyThreshold:  "1",
	}

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/targetgroup/"+payload.TargetGroupId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPut)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyUpdateResponseJson)
	})

	got, err := client.TargetGroup().Update(payload)

	var want UpdateResponse
	_ = json.Unmarshal([]byte(dummyUpdateResponseJson), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestTargetGroupService_Update_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.TargetGroup().Update(UpdateTargetGroupParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestTargetGroupService_Delete_happyPath(t *testing.T) {
	token := "token"
	targetGroupId := "someTargetGroupId"
	targetGroupName := "someTargetGroupId"

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/targetgroup/"+targetGroupId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyDeleteResponseJson)
	})

	want := DeleteResponse{Status: "success", Message: "success"}

	got, _ := client.TargetGroup().Delete(targetGroupId, targetGroupName)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestTargetGroupService_DeleteTargetGroup_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	delResponse, err := client.TargetGroup().Delete("someTargetGroupId", "target Group Name")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
	}
}

func TestTargetGroupService_CreateTarget_happyPath(t *testing.T) {
	token := "token"
	payload := CreateTargetGroupTargetParams{
		TargetGroupId:   "11111",
		BackendProtocol: "HTTP",
		BackendPort:     "123",
		IP:              "11.11.11.11",
		Cloudid:         "33333",
	}

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/targetgroup/"+payload.TargetGroupId+"/target", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateResponseJson)
	})

	got, err := client.TargetGroup().CreateTarget(payload)

	var want CreateResponse
	_ = json.Unmarshal([]byte(dummyCreateResponseJson), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestTargetGroupService_CreateTarget_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.TargetGroup().CreateTarget(CreateTargetGroupTargetParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestTargetGroupService_ReadTarget_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	targetgroupId := "11111"
	targetgroupTargetId := "22222"
	expectedResponse := dummyReadTargetGroupTargetRes
	serverResponse := dummyTargetGroupServerRes

	mux.HandleFunc("/targetgroup", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want Target
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.TargetGroup().ReadTarget(targetgroupId, targetgroupTargetId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestTargetGroupService_ReadTarget_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	apikey, err := client.TargetGroup().ReadTarget("someId", "id")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if apikey != nil {
		t.Errorf("Was not expecting any apikey to be returned, instead got %v", apikey)
	}
}

func TestTargetGroupService_ListTarget_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	targetgroupId := "11111"
	expectedResponse := dummyListTargetGroupTargetRes
	serverResponse := dummyTargetGroupServerRes

	mux.HandleFunc("/targetgroup", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want []Target
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.TargetGroup().ListTarget(targetgroupId)
	if len(got) != len(want) {
		t.Errorf("Was expecting %d targetgroup to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestTargetGroupService_ListTarget_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	targetgroup, err := client.TargetGroup().ListTarget("id")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if targetgroup != nil {
		t.Errorf("Was not expecting any targetgroup to be returned, instead got %v", targetgroup)
	}
}

func TestTargetGroupService_DeleteTarget_happyPath(t *testing.T) {
	token := "token"
	targetGroupId := "11111"
	targetId := "22222"

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/targetgroup/"+targetGroupId+"/target/"+targetId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyDeleteResponseJson)
	})

	want := DeleteResponse{Status: "success", Message: "success"}

	got, _ := client.TargetGroup().DeleteTarget(targetGroupId, targetId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestTargetGroupService_DeleteTarget_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	delResponse, err := client.TargetGroup().DeleteTarget("someTargetGroupId", "target Group Name")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
	}
}

const dummyReadTargetGroupRes = `{
	"id": "11111",
	"name": "d12d4",
	"port": "1",
	"protocol": "HTTP",
	"health_check_path": "1",
	"health_check_interval": "1",
	"health_check_protocol": "HTTP",
	"health_check_timeout": "1",
	"healthy_threshold": "1",
	"unhealthy_threshold": "1",
	"created_at": "2024-04-29 23:09:02",
	"updated_at": "2024-05-06 15:20:55",
	"targets": [
		{
			"lbid": "0",
			"ip": "103.111.111.111",
			"cloudid": "1277220",
			"status": "",
			"scaling_groupid": "0",
			"kubernetes_clusterid": "0",
			"backend_port": "123",
			"backend_protocol": "HTTP",
			"targetgroup_id": "11111",
			"frontend_id": "0",
			"id": "22222"
		}
	]
}`

const dummyTargetGroupServerRes = `{
    "targetgroups": [` + dummyReadTargetGroupRes + `]
}`

const dummyListTargetGroupRes = `[` + dummyReadTargetGroupRes + `]`

const dummyReadTargetGroupTargetRes = `{
	"lbid": "0",
	"ip": "103.111.111.111",
	"cloudid": "1277220",
	"status": "",
	"scaling_groupid": "0",
	"kubernetes_clusterid": "0",
	"backend_port": "123",
	"backend_protocol": "HTTP",
	"targetgroup_id": "11111",
	"frontend_id": "0",
	"id": "22222"
}`

const dummyListTargetGroupTargetRes = `[` + dummyReadTargetGroupTargetRes + `]`
