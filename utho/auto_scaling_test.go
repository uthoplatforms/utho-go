package utho

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAutoScalingService_Create_happyPath(t *testing.T) {
	token := "token"

	var payload CreateAutoScalingParams
	_ = json.Unmarshal([]byte(dummyCreateAutoScaling), &payload)

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/autoscaling", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateResponseJson)
	})

	got, err := client.AutoScaling().Create(payload)

	var want CreateResponse
	_ = json.Unmarshal([]byte(dummyCreateResponseJson), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestAutoScalingService_Create_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.AutoScaling().Create(CreateAutoScalingParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestAutoScalingService_Read_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	autoscalingId := "11111"
	expectedResponse := dummyReadAutoScalingRes
	serverResponse := dummyAutoScalingServerRes

	mux.HandleFunc("/autoscaling/"+autoscalingId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want Groups
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.AutoScaling().Read(autoscalingId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestAutoScalingService_Read_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	apikey, err := client.AutoScaling().Read("someId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if apikey != nil {
		t.Errorf("Was not expecting any apikey to be returned, instead got %v", apikey)
	}
}

func TestAutoScalingService_List_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	expectedResponse := dummyListAutoScalingRes
	serverResponse := dummyAutoScalingServerRes

	mux.HandleFunc("/autoscaling", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want []Groups
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.AutoScaling().List()
	if len(got) != len(want) {
		t.Errorf("Was expecting %d autoscaling to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestAutoScalingService_List_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	autoscaling, err := client.AutoScaling().List()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if autoscaling != nil {
		t.Errorf("Was not expecting any autoscaling to be returned, instead got %v", autoscaling)
	}
}

func TestAutoScalingService_Delete_happyPath(t *testing.T) {
	token := "token"
	autoscalingId := "11111"
	autoscalingName := "Auto-scaling-Jz8hceLN.utho"

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/autoscaling/"+autoscalingId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyDeleteResponseJson)
	})

	want := DeleteResponse{Status: "success", Message: "success"}

	got, _ := client.AutoScaling().Delete(autoscalingId, autoscalingName)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestAutoScalingService_DeleteAutoScaling_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	delResponse, err := client.AutoScaling().Delete("someAutoScalingId", "autoscalingName")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
	}
}

// autoscaling Policy
func TestAutoScalingService_CreatePolicy_happyPath(t *testing.T) {
	token := "token"
	var payload CreateAutoScalingPolicyParams
	_ = json.Unmarshal([]byte(dummyCreateAutoScalingPolicy), &payload)

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/autoscaling/policy", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateResponseJson)
	})

	got, err := client.AutoScaling().CreatePolicy(payload)

	var want CreateResponse
	_ = json.Unmarshal([]byte(dummyCreateResponseJson), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestAutoScalingService_CreatePolicy_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.AutoScaling().CreatePolicy(CreateAutoScalingPolicyParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestAutoScalingService_ReadPolicy_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	autoscalingId := "11111"
	autoscalingPolicyId := "22222"
	expectedResponse := dummyReadAutoScalingPolicyRes
	serverResponse := dummyAutoScalingServerRes

	mux.HandleFunc("/autoscaling/"+autoscalingId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want Policy
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.AutoScaling().ReadPolicy(autoscalingId, autoscalingPolicyId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestAutoScalingService_ReadPolicy_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	apikey, err := client.AutoScaling().ReadPolicy("someId", "policyId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if apikey != nil {
		t.Errorf("Was not expecting any apikey to be returned, instead got %v", apikey)
	}
}

func TestAutoScalingService_ListPolicy_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	autoscalingId := "11111"
	expectedResponse := "[" + dummyReadAutoScalingPolicyRes + "]"
	serverResponse := dummyAutoScalingServerRes

	mux.HandleFunc("/autoscaling/"+autoscalingId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want []Policy
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.AutoScaling().ListPolicies(autoscalingId)
	if len(got) != len(want) {
		t.Errorf("Was expecting %d autoscaling to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestAutoScalingService_ListPolicy_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	autoscaling, err := client.AutoScaling().ListPolicies("id")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if autoscaling != nil {
		t.Errorf("Was not expecting any autoscaling to be returned, instead got %v", autoscaling)
	}
}

func TestAutoScalingService_DeletePolicy_happyPath(t *testing.T) {
	token := "token"
	autoScalingPolicyId := "22222"

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/autoscaling/policy/"+autoScalingPolicyId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyDeleteResponseJson)
	})

	want := DeleteResponse{Status: "success", Message: "success"}

	got, _ := client.AutoScaling().DeletePolicy(autoScalingPolicyId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestAutoScalingService_DeletePolicy_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	delResponse, err := client.AutoScaling().DeletePolicy("someAutoScalingId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
	}
}

// autoscaling Schedule
func TestAutoScalingService_CreateSchedule_happyPath(t *testing.T) {
	token := "token"
	var payload CreateAutoScalingScheduleParams
	_ = json.Unmarshal([]byte(dummyCreateAutoScalingSchedule), &payload)
	payload.AutoScalingId = "11111"

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/autoscaling/"+payload.AutoScalingId+"/schedulepolicy", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateResponseJson)
	})

	got, err := client.AutoScaling().CreateSchedule(payload)

	var want CreateResponse
	_ = json.Unmarshal([]byte(dummyCreateResponseJson), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestAutoScalingService_CreateSchedule_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.AutoScaling().CreateSchedule(CreateAutoScalingScheduleParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestAutoScalingService_ReadSchedule_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	autoscalingId := "11111"
	autoscalingScheduleId := "33333"
	expectedResponse := dummyReadAutoScalingScheduleRes
	serverResponse := dummyAutoScalingServerRes

	mux.HandleFunc("/autoscaling/"+autoscalingId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want Schedule
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.AutoScaling().ReadSchedule(autoscalingId, autoscalingScheduleId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestAutoScalingService_ReadSchedule_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	apikey, err := client.AutoScaling().ReadSchedule("someId", "scheduleId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if apikey != nil {
		t.Errorf("Was not expecting any apikey to be returned, instead got %v", apikey)
	}
}

func TestAutoScalingService_ListSchedule_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	autoscalingId := "11111"
	expectedResponse := "[" + dummyReadAutoScalingScheduleRes + "]"
	serverResponse := dummyAutoScalingServerRes

	mux.HandleFunc("/autoscaling/"+autoscalingId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want []Schedule
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.AutoScaling().ListSchedules(autoscalingId)
	if len(got) != len(want) {
		t.Errorf("Was expecting %d autoscaling to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestAutoScalingService_ListSchedule_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	autoscaling, err := client.AutoScaling().ListPolicies("id")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if autoscaling != nil {
		t.Errorf("Was not expecting any autoscaling to be returned, instead got %v", autoscaling)
	}
}

func TestAutoScalingService_DeleteSchedule_happyPath(t *testing.T) {
	token := "token"
	autoScalingeId := "11111"
	autoScalingScheduleId := "33333"

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/autoscaling/"+autoScalingeId+"/schedulepolicy/"+autoScalingScheduleId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyDeleteResponseJson)
	})

	want := DeleteResponse{Status: "success", Message: "success"}

	got, _ := client.AutoScaling().DeleteSchedule(autoScalingeId, autoScalingScheduleId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestAutoScalingService_DeleteSchedule_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	delResponse, err := client.AutoScaling().DeleteSchedule("someAutoScalingId", "id")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
	}
}

// autoscaling Loadbalancer
func TestAutoScalingService_CreateLoadbalancer_happyPath(t *testing.T) {
	token := "token"
	payload := CreateAutoScalingLoadbalancerParams{
		AutoScalingId:  "11111",
		LoadbalancerId: "44444",
	}

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/autoscaling/"+payload.AutoScalingId+"/loadbalancer/"+payload.LoadbalancerId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateResponseJson)
	})

	got, err := client.AutoScaling().CreateLoadbalancer(payload)

	var want CreateResponse
	_ = json.Unmarshal([]byte(dummyCreateResponseJson), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestAutoScalingService_CreateLoadbalancer_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.AutoScaling().CreateLoadbalancer(CreateAutoScalingLoadbalancerParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestAutoScalingService_ReadLoadbalancer_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	autoscalingId := "11111"
	autoscalingLoadbalancerId := "44444"
	expectedResponse := dummyReadAutoScalingLoadbalancerRes
	serverResponse := dummyAutoScalingServerRes

	mux.HandleFunc("/autoscaling/"+autoscalingId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want AutoScalingLoadbalancers
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.AutoScaling().ReadLoadbalancer(autoscalingId, autoscalingLoadbalancerId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestAutoScalingService_ReadLoadbalancer_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	apikey, err := client.AutoScaling().ReadLoadbalancer("someId", "loadbalancerId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if apikey != nil {
		t.Errorf("Was not expecting any apikey to be returned, instead got %v", apikey)
	}
}

func TestAutoScalingService_ListLoadbalancer_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	autoscalingId := "11111"
	expectedResponse := "[" + dummyReadAutoScalingLoadbalancerRes + "]"
	serverResponse := dummyAutoScalingServerRes

	mux.HandleFunc("/autoscaling/"+autoscalingId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want []AutoScalingLoadbalancers
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.AutoScaling().ListLoadbalancers(autoscalingId)
	if len(got) != len(want) {
		t.Errorf("Was expecting %d autoscaling to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestAutoScalingService_ListLoadbalancer_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	autoscaling, err := client.AutoScaling().ListPolicies("id")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if autoscaling != nil {
		t.Errorf("Was not expecting any autoscaling to be returned, instead got %v", autoscaling)
	}
}

func TestAutoScalingService_DeleteLoadbalancer_happyPath(t *testing.T) {
	token := "token"
	autoScalingeId := "11111"
	autoScalingLoadbalancerId := "44444"

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/autoscaling/"+autoScalingeId+"/loadbalancerpolicy/"+autoScalingLoadbalancerId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyDeleteResponseJson)
	})

	want := DeleteResponse{Status: "success", Message: "success"}

	got, _ := client.AutoScaling().DeleteLoadbalancer(autoScalingeId, autoScalingLoadbalancerId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestAutoScalingService_DeleteLoadbalancer_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	delResponse, err := client.AutoScaling().DeleteLoadbalancer("someAutoScalingId", "id")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
	}
}

// autoscaling SecurityGroup
func TestAutoScalingService_CreateSecurityGroup_happyPath(t *testing.T) {
	token := "token"
	payload := CreateAutoScalingSecurityGroupParams{
		AutoScalingId:              "11111",
		AutoScalingSecurityGroupId: "55555",
	}

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/autoscaling/"+payload.AutoScalingId+"/securitygroup/"+payload.AutoScalingSecurityGroupId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateResponseJson)
	})

	got, err := client.AutoScaling().CreateSecurityGroup(payload)

	var want CreateResponse
	_ = json.Unmarshal([]byte(dummyCreateResponseJson), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestAutoScalingService_CreateSecurityGroup_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.AutoScaling().CreateSecurityGroup(CreateAutoScalingSecurityGroupParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestAutoScalingService_ReadSecurityGroup_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	autoscalingId := "11111"
	autoscalingSecurityGroupId := "55555"
	expectedResponse := dummyReadAutoScalingSecurityGroupRes
	serverResponse := dummyAutoScalingServerRes

	mux.HandleFunc("/autoscaling/"+autoscalingId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want SecurityGroup
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.AutoScaling().ReadSecurityGroup(autoscalingId, autoscalingSecurityGroupId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestAutoScalingService_ReadSecurityGroup_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	apikey, err := client.AutoScaling().ReadSecurityGroup("someId", "SecurityGroupId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if apikey != nil {
		t.Errorf("Was not expecting any apikey to be returned, instead got %v", apikey)
	}
}

func TestAutoScalingService_ListSecurityGroup_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	autoscalingId := "11111"
	expectedResponse := "[" + dummyReadAutoScalingSecurityGroupRes + "]"
	serverResponse := dummyAutoScalingServerRes

	mux.HandleFunc("/autoscaling/"+autoscalingId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want []SecurityGroup
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.AutoScaling().ListSecurityGroups(autoscalingId)
	if len(got) != len(want) {
		t.Errorf("Was expecting %d autoscaling to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestAutoScalingService_ListSecurityGroup_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	autoscaling, err := client.AutoScaling().ListPolicies("id")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if autoscaling != nil {
		t.Errorf("Was not expecting any autoscaling to be returned, instead got %v", autoscaling)
	}
}

func TestAutoScalingService_DeleteSecurityGroup_happyPath(t *testing.T) {
	token := "token"
	autoScalingeId := "11111"
	autoScalingSecurityGroupId := "55555"

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/autoscaling/"+autoScalingeId+"/securitygroup/"+autoScalingSecurityGroupId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyDeleteResponseJson)
	})

	want := DeleteResponse{Status: "success", Message: "success"}

	got, _ := client.AutoScaling().DeleteSecurityGroup(autoScalingeId, autoScalingSecurityGroupId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestAutoScalingService_DeleteSecurityGroup_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	delResponse, err := client.AutoScaling().DeleteSecurityGroup("someAutoScalingId", "id")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
	}
}

// autoscaling Targetgroup
func TestAutoScalingService_CreateTargetgroup_happyPath(t *testing.T) {
	token := "token"
	payload := CreateAutoScalingTargetgroupParams{
		AutoScalingId:            "11111",
		AutoScalingTargetgroupId: "666666",
	}

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/autoscaling/"+payload.AutoScalingId+"/targetgroup/"+payload.AutoScalingTargetgroupId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateResponseJson)
	})

	got, err := client.AutoScaling().CreateTargetgroup(payload)

	var want CreateResponse
	_ = json.Unmarshal([]byte(dummyCreateResponseJson), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestAutoScalingService_CreateTargetgroup_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.AutoScaling().CreateTargetgroup(CreateAutoScalingTargetgroupParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestAutoScalingService_ReadTargetgroup_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	autoscalingId := "11111"
	autoscalingTargetgroupId := "666666"
	expectedResponse := dummyReadAutoScalingTargetgroupRes
	serverResponse := dummyAutoScalingServerRes

	mux.HandleFunc("/autoscaling/"+autoscalingId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want AutoScalingTargetGroup
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.AutoScaling().ReadTargetgroup(autoscalingId, autoscalingTargetgroupId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestAutoScalingService_ReadTargetgroup_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	apikey, err := client.AutoScaling().ReadTargetgroup("someId", "TargetgroupId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if apikey != nil {
		t.Errorf("Was not expecting any apikey to be returned, instead got %v", apikey)
	}
}

func TestAutoScalingService_ListTargetgroup_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	autoscalingId := "11111"
	expectedResponse := "[" + dummyReadAutoScalingTargetgroupRes + "]"
	serverResponse := dummyAutoScalingServerRes

	mux.HandleFunc("/autoscaling/"+autoscalingId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want []AutoScalingTargetGroup
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.AutoScaling().ListTargetgroups(autoscalingId)
	if len(got) != len(want) {
		t.Errorf("Was expecting %d autoscaling to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestAutoScalingService_ListTargetgroup_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	autoscaling, err := client.AutoScaling().ListPolicies("id")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if autoscaling != nil {
		t.Errorf("Was not expecting any autoscaling to be returned, instead got %v", autoscaling)
	}
}

func TestAutoScalingService_DeleteTargetgroup_happyPath(t *testing.T) {
	token := "token"
	autoScalingeId := "11111"
	autoScalingTargetgroupId := "666666"

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/autoscaling/"+autoScalingeId+"/targetgroup/"+autoScalingTargetgroupId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyDeleteResponseJson)
	})

	want := DeleteResponse{Status: "success", Message: "success"}

	got, _ := client.AutoScaling().DeleteTargetgroup(autoScalingeId, autoScalingTargetgroupId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestAutoScalingService_DeleteTargetgroup_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	delResponse, err := client.AutoScaling().DeleteTargetgroup("someAutoScalingId", "id")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
	}
}
