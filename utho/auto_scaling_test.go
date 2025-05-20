package utho

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestAutoScalingService_Create_happyPath(t *testing.T) {
	token := "token"

	var payload CreateAutoScalingParams
	_ = faker.FakeData(&payload)

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp CreateAutoScalingResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/autoscaling", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	got, err := client.AutoScaling().Create(payload)

	var want CreateAutoScalingResponse
	_ = json.Unmarshal(respBytes, &want)

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

	var fakeGroup Groups
	_ = faker.FakeData(&fakeGroup)
	serverResp := struct {
		Groups []Groups `json:"groups"`
		Status string   `json:"status"`
	}{
		Groups: []Groups{fakeGroup},
		Status: "success",
	}
	serverResponse, _ := json.Marshal(serverResp)
	expectedResponse, _ := json.Marshal(fakeGroup)

	mux.HandleFunc("/autoscaling/"+fakeGroup.ID, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		w.Write(serverResponse)
	})

	var want Groups
	_ = json.Unmarshal(expectedResponse, &want)

	got, _ := client.AutoScaling().Read(fakeGroup.ID)
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

	var fakeGroups []Groups
	for i := 0; i < 2; i++ {
		var g Groups
		_ = faker.FakeData(&g)
		fakeGroups = append(fakeGroups, g)
	}
	serverResp := struct {
		Groups []Groups `json:"groups"`
		Status string   `json:"status"`
	}{
		Groups: fakeGroups,
		Status: "success",
	}
	serverResponse, _ := json.Marshal(serverResp)
	expectedResponse, _ := json.Marshal(fakeGroups)

	mux.HandleFunc("/autoscaling", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		w.Write(serverResponse)
	})

	var want []Groups
	_ = json.Unmarshal(expectedResponse, &want)

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

	var autoScalingeId, autoscalingName string
	_ = faker.FakeData(&autoScalingeId)
	_ = faker.FakeData(&autoscalingName)

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp DeleteResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	fakeResp.Message = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/autoscaling/"+autoScalingeId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	want := fakeResp

	got, _ := client.AutoScaling().Delete(autoScalingeId, autoscalingName)
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
	_ = faker.FakeData(&payload)

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp CreateResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/autoscaling/policy", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	got, err := client.AutoScaling().CreatePolicy(payload)

	var want CreateResponse
	_ = json.Unmarshal(respBytes, &want)

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

	var fakeGroup Groups
	_ = faker.FakeData(&fakeGroup)
	var fakePolicy Policy
	_ = faker.FakeData(&fakePolicy)
	fakeGroup.Policies = []Policy{fakePolicy}
	serverResp := struct {
		Groups []Groups `json:"groups"`
		Status string   `json:"status"`
	}{
		Groups: []Groups{fakeGroup},
		Status: "success",
	}
	serverResponse, _ := json.Marshal(serverResp)
	expectedResponse, _ := json.Marshal(fakePolicy)

	mux.HandleFunc("/autoscaling/"+fakeGroup.ID, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		w.Write(serverResponse)
	})

	var want Policy
	_ = json.Unmarshal(expectedResponse, &want)

	got, _ := client.AutoScaling().ReadPolicy(fakeGroup.ID, fakePolicy.ID)
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

	var fakeGroup Groups
	_ = faker.FakeData(&fakeGroup)
	var fakePolicies []Policy
	for i := 0; i < 2; i++ {
		var p Policy
		_ = faker.FakeData(&p)
		fakePolicies = append(fakePolicies, p)
	}
	fakeGroup.Policies = fakePolicies
	serverResp := struct {
		Groups []Groups `json:"groups"`
		Status string   `json:"status"`
	}{
		Groups: []Groups{fakeGroup},
		Status: "success",
	}
	serverResponse, _ := json.Marshal(serverResp)
	expectedResponse, _ := json.Marshal(fakePolicies)

	mux.HandleFunc("/autoscaling/"+fakeGroup.ID, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		w.Write(serverResponse)
	})

	var want []Policy
	_ = json.Unmarshal(expectedResponse, &want)

	got, _ := client.AutoScaling().ListPolicies(fakeGroup.ID)
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

	var autoScalingPolicyId string
	_ = faker.FakeData(&autoScalingPolicyId)

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp DeleteResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	fakeResp.Message = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/autoscaling/policy/"+autoScalingPolicyId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	want := fakeResp

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
	_ = faker.FakeData(&payload)
	payload.AutoScalingId = "11111"

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp CreateResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/autoscaling/"+payload.AutoScalingId+"/schedulepolicy", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	got, err := client.AutoScaling().CreateSchedule(payload)

	var want CreateResponse
	_ = json.Unmarshal(respBytes, &want)

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

	var fakeGroup Groups
	_ = faker.FakeData(&fakeGroup)
	var fakeSchedule Schedule
	_ = faker.FakeData(&fakeSchedule)
	fakeGroup.Schedules = []Schedule{fakeSchedule}
	serverResp := struct {
		Groups []Groups `json:"groups"`
		Status string   `json:"status"`
	}{
		Groups: []Groups{fakeGroup},
		Status: "success",
	}
	serverResponse, _ := json.Marshal(serverResp)
	expectedResponse, _ := json.Marshal(fakeSchedule)

	mux.HandleFunc("/autoscaling/"+fakeGroup.ID, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		w.Write(serverResponse)
	})

	var want Schedule
	_ = json.Unmarshal(expectedResponse, &want)

	got, _ := client.AutoScaling().ReadSchedule(fakeGroup.ID, fakeSchedule.ID)
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

	var fakeGroup Groups
	_ = faker.FakeData(&fakeGroup)
	var fakeSchedules []Schedule
	for i := 0; i < 2; i++ {
		var s Schedule
		_ = faker.FakeData(&s)
		fakeSchedules = append(fakeSchedules, s)
	}
	fakeGroup.Schedules = fakeSchedules
	serverResp := struct {
		Groups []Groups `json:"groups"`
		Status string   `json:"status"`
	}{
		Groups: []Groups{fakeGroup},
		Status: "success",
	}
	serverResponse, _ := json.Marshal(serverResp)
	expectedResponse, _ := json.Marshal(fakeSchedules)

	mux.HandleFunc("/autoscaling/"+fakeGroup.ID, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		w.Write(serverResponse)
	})

	var want []Schedule
	_ = json.Unmarshal(expectedResponse, &want)

	got, _ := client.AutoScaling().ListSchedules(fakeGroup.ID)
	if len(got) != len(want) {
		t.Errorf("Was expecting %d autoscaling to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestAutoScalingService_ListSchedule_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	autoscaling, err := client.AutoScaling().ListSchedules("id")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if autoscaling != nil {
		t.Errorf("Was not expecting any autoscaling to be returned, instead got %v", autoscaling)
	}
}

func TestAutoScalingService_DeleteSchedule_happyPath(t *testing.T) {
	token := "token"

	var autoScalingeId, autoScalingScheduleId string
	_ = faker.FakeData(&autoScalingeId)
	_ = faker.FakeData(&autoScalingScheduleId)

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp DeleteResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	fakeResp.Message = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/autoscaling/"+autoScalingeId+"/schedulepolicy/"+autoScalingScheduleId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	want := fakeResp

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

	var fakeResp CreateResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/autoscaling/"+payload.AutoScalingId+"/loadbalancer/"+payload.LoadbalancerId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	got, err := client.AutoScaling().CreateLoadbalancer(payload)

	var want CreateResponse
	_ = json.Unmarshal(respBytes, &want)

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

	var fakeGroup Groups
	_ = faker.FakeData(&fakeGroup)
	var fakeLoadbalancer AutoScalingLoadbalancers
	_ = faker.FakeData(&fakeLoadbalancer)
	fakeGroup.Loadbalancers = []AutoScalingLoadbalancers{fakeLoadbalancer}
	serverResp := struct {
		Groups []Groups `json:"groups"`
		Status string   `json:"status"`
	}{
		Groups: []Groups{fakeGroup},
		Status: "success",
	}
	serverResponse, _ := json.Marshal(serverResp)
	expectedResponse, _ := json.Marshal(fakeLoadbalancer)

	mux.HandleFunc("/autoscaling/"+fakeGroup.ID, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		w.Write(serverResponse)
	})

	var want AutoScalingLoadbalancers
	_ = json.Unmarshal(expectedResponse, &want)

	got, _ := client.AutoScaling().ReadLoadbalancer(fakeGroup.ID, fakeLoadbalancer.ID)
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

	var fakeGroup Groups
	_ = faker.FakeData(&fakeGroup)
	var fakeLoadbalancers []AutoScalingLoadbalancers
	for i := 0; i < 2; i++ {
		var lb AutoScalingLoadbalancers
		_ = faker.FakeData(&lb)
		fakeLoadbalancers = append(fakeLoadbalancers, lb)
	}
	fakeGroup.Loadbalancers = fakeLoadbalancers
	serverResp := struct {
		Groups []Groups `json:"groups"`
		Status string   `json:"status"`
	}{
		Groups: []Groups{fakeGroup},
		Status: "success",
	}
	serverResponse, _ := json.Marshal(serverResp)
	expectedResponse, _ := json.Marshal(fakeLoadbalancers)

	mux.HandleFunc("/autoscaling/"+fakeGroup.ID, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		w.Write(serverResponse)
	})

	var want []AutoScalingLoadbalancers
	_ = json.Unmarshal(expectedResponse, &want)

	got, _ := client.AutoScaling().ListLoadbalancers(fakeGroup.ID)
	if len(got) != len(want) {
		t.Errorf("Was expecting %d autoscaling to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestAutoScalingService_ListLoadbalancer_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	autoscaling, err := client.AutoScaling().ListLoadbalancers("id")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if autoscaling != nil {
		t.Errorf("Was not expecting any autoscaling to be returned, instead got %v", autoscaling)
	}
}

func TestAutoScalingService_DeleteLoadbalancer_happyPath(t *testing.T) {
	token := "token"

	var autoScalingeId, autoScalingLoadbalancerId string
	_ = faker.FakeData(&autoScalingeId)
	_ = faker.FakeData(&autoScalingLoadbalancerId)

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp DeleteResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	fakeResp.Message = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/autoscaling/"+autoScalingeId+"/loadbalancerpolicy/"+autoScalingLoadbalancerId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	want := fakeResp

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

	var fakeResp CreateResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/autoscaling/"+payload.AutoScalingId+"/securitygroup/"+payload.AutoScalingSecurityGroupId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	got, err := client.AutoScaling().CreateSecurityGroup(payload)

	var want CreateResponse
	_ = json.Unmarshal(respBytes, &want)

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

	autoScalingeId := "11111"
	securityGroupId := "55555"

	var fakeGroup Groups
	_ = faker.FakeData(&fakeGroup)
	var fakeSecurityGroup SecurityGroup
	fakeSecurityGroup.ID = securityGroupId
	fakeGroup.SecurityGroups = []SecurityGroup{fakeSecurityGroup}
	serverResp := struct {
		Groups []Groups `json:"groups"`
		Status string   `json:"status"`
	}{
		Groups: []Groups{fakeGroup},
		Status: "success",
	}
	serverResponse, _ := json.Marshal(serverResp)

	mux.HandleFunc("/autoscaling/"+autoScalingeId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		w.WriteHeader(http.StatusOK)
		w.Write(serverResponse)
	})

	want := fakeSecurityGroup

	got, err := client.AutoScaling().ReadSecurityGroup(autoScalingeId, securityGroupId)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
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

	autoScalingeId := "11111"

	var fakeGroup Groups
	_ = faker.FakeData(&fakeGroup)
	var fakeSecurityGroups []SecurityGroup
	for i := 0; i < 2; i++ {
		var sg SecurityGroup
		_ = faker.FakeData(&sg)
		fakeSecurityGroups = append(fakeSecurityGroups, sg)
	}
	fakeGroup.SecurityGroups = fakeSecurityGroups
	serverResp := struct {
		Groups []Groups `json:"groups"`
		Status string   `json:"status"`
	}{
		Groups: []Groups{fakeGroup},
		Status: "success",
	}
	serverResponse, _ := json.Marshal(serverResp)
	expectedResponse, _ := json.Marshal(fakeSecurityGroups)

	mux.HandleFunc("/autoscaling/"+autoScalingeId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		w.Write(serverResponse)
	})

	var want []SecurityGroup
	_ = json.Unmarshal(expectedResponse, &want)

	got, _ := client.AutoScaling().ListSecurityGroups(autoScalingeId)
	if len(got) != len(want) {
		t.Errorf("Was expecting %d autoscaling to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestAutoScalingService_ListSecurityGroup_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	autoscaling, err := client.AutoScaling().ListSecurityGroups("id")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if autoscaling != nil {
		t.Errorf("Was not expecting any autoscaling to be returned, instead got %v", autoscaling)
	}
}

func TestAutoScalingService_DeleteSecurityGroup_happyPath(t *testing.T) {
	token := "token"

	var autoScalingeId, autoScalingSecurityGroupId string
	_ = faker.FakeData(&autoScalingeId)
	_ = faker.FakeData(&autoScalingSecurityGroupId)

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp DeleteResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	fakeResp.Message = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/autoscaling/"+autoScalingeId+"/securitygroup/"+autoScalingSecurityGroupId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	want := fakeResp

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

	var fakeResp CreateResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/autoscaling/"+payload.AutoScalingId+"/targetgroup/"+payload.AutoScalingTargetgroupId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	got, err := client.AutoScaling().CreateTargetgroup(payload)

	var want CreateResponse
	_ = json.Unmarshal(respBytes, &want)

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

	autoScalingeId := "11111"
	autoscalingTargetgroupId := "666666"

	var fakeGroup Groups
	_ = faker.FakeData(&fakeGroup)

	var fakeTargetGroup AutoScalingTargetGroup
	_ = faker.FakeData(&fakeTargetGroup)
	fakeTargetGroup.ID = autoscalingTargetgroupId
	fakeGroup.TargetGroups = []AutoScalingTargetGroup{fakeTargetGroup}

	serverResp := struct {
		Groups []Groups `json:"groups"`
		Status string   `json:"status"`
	}{
		Groups: []Groups{fakeGroup},
		Status: "success",
	}
	serverResponse, _ := json.Marshal(serverResp)
	expectedResponse, _ := json.Marshal(fakeTargetGroup)

	mux.HandleFunc("/autoscaling/"+autoScalingeId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		w.Write(serverResponse)
	})

	var want AutoScalingTargetGroup
	_ = json.Unmarshal(expectedResponse, &want)

	got, err := client.AutoScaling().ReadTargetgroup(autoScalingeId, autoscalingTargetgroupId)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
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

	autoScalingeId := "11111"

	var fakeGroup Groups
	_ = faker.FakeData(&fakeGroup)
	var fakeTargetGroups []AutoScalingTargetGroup
	for i := 0; i < 2; i++ {
		var tg AutoScalingTargetGroup
		_ = faker.FakeData(&tg)
		fakeTargetGroups = append(fakeTargetGroups, tg)
	}
	fakeGroup.TargetGroups = fakeTargetGroups
	serverResp := struct {
		Groups []Groups `json:"groups"`
		Status string   `json:"status"`
	}{
		Groups: []Groups{fakeGroup},
		Status: "success",
	}
	serverResponse, _ := json.Marshal(serverResp)
	expectedResponse, _ := json.Marshal(fakeTargetGroups)

	mux.HandleFunc("/autoscaling/"+autoScalingeId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		w.Write(serverResponse)
	})

	var want []AutoScalingTargetGroup
	_ = json.Unmarshal(expectedResponse, &want)

	got, _ := client.AutoScaling().ListTargetgroups(autoScalingeId)
	if len(got) != len(want) {
		t.Errorf("Was expecting %d autoscaling to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestAutoScalingService_ListTargetgroup_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	autoscaling, err := client.AutoScaling().ListTargetgroups("id")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if autoscaling != nil {
		t.Errorf("Was not expecting any autoscaling to be returned, instead got %v", autoscaling)
	}
}

func TestAutoScalingService_DeleteTargetgroup_happyPath(t *testing.T) {
	token := "token"

	var autoScalingeId, autoScalingTargetgroupId string
	_ = faker.FakeData(&autoScalingeId)
	_ = faker.FakeData(&autoScalingTargetgroupId)

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp DeleteResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	fakeResp.Message = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/autoscaling/"+autoScalingeId+"/targetgroup/"+autoScalingTargetgroupId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	want := fakeResp

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

func TestAutoScalingService_Update_happyPath(t *testing.T) {
	token := "token"
	var autoScalingeId string
	_ = faker.FakeData(&autoScalingeId)

	var payload UpdateAutoScalingParams
	_ = faker.FakeData(&payload)
	payload.AutoScalingId = autoScalingeId

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp UpdateResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	fakeResp.Message = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/autoscaling/"+autoScalingeId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPut)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	want := fakeResp

	got, err := client.AutoScaling().Update(payload)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestAutoScalingService_Update_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.AutoScaling().Update(UpdateAutoScalingParams{AutoScalingId: "someId"})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestAutoScalingService_UpdatePolicy_happyPath(t *testing.T) {
	token := "token"
	var autoScalingeId, autoScalingPolicyId string
	_ = faker.FakeData(&autoScalingeId)
	_ = faker.FakeData(&autoScalingPolicyId)

	var payload UpdateAutoScalingPolicyParams
	_ = faker.FakeData(&payload)
	payload.AutoScalingPolicyId = autoScalingPolicyId

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp UpdateResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	fakeResp.Message = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/autoscaling/policy/"+autoScalingPolicyId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPut)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	want := fakeResp

	got, err := client.AutoScaling().UpdatePolicy(payload)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestAutoScalingService_UpdatePolicy_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.AutoScaling().UpdatePolicy(UpdateAutoScalingPolicyParams{AutoScalingPolicyId: "someId"})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestAutoScalingService_UpdateSchedule_happyPath(t *testing.T) {
	token := "token"
	var autoScalingeId, autoScalingScheduleId string
	_ = faker.FakeData(&autoScalingeId)
	_ = faker.FakeData(&autoScalingScheduleId)

	var payload UpdateAutoScalingScheduleParams
	_ = faker.FakeData(&payload)
	payload.AutoScalingeId = autoScalingeId
	payload.AutoScalingScheduleId = autoScalingScheduleId

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp UpdateResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	fakeResp.Message = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/autoscaling/"+autoScalingeId+"/schedulepolicy/"+autoScalingScheduleId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPut)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	want := fakeResp

	got, err := client.AutoScaling().UpdateSchedule(payload)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestAutoScalingService_UpdateSchedule_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.AutoScaling().UpdateSchedule(UpdateAutoScalingScheduleParams{AutoScalingeId: "someId", AutoScalingScheduleId: "someOtherId"})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}
