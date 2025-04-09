package utho

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKubernetesService_Create_happyPath(t *testing.T) {
	token := "token"

	createNodepoolsParams := CreateNodepoolsParams{
		Label:    "pool",
		Size:     "10045",
		PoolType: "static",
		Count:    "2",
	}
	payload := CreateKubernetesParams{
		Dcslug:         "innoida",
		ClusterLabel:   "My_kubernetes",
		ClusterVersion: "1.27.0",
		Vpc:            "f1dd58f1-1bfa-46ef-8b94-f69f312c0245",
		SecurityGroups: "23432613,23432615",
		NetworkType:    "both",
		Firewall:       "23433480",
		Cpumodel:       "intel",
		Nodepools:      []CreateNodepoolsParams{createNodepoolsParams},
	}

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/kubernetes/deploy", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateResponseJson)
	})

	ctx := context.Background()
	got, err := client.Kubernetes().Create(ctx, payload)

	var want CreateResponse
	_ = json.Unmarshal([]byte(dummyCreateResponseJson), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestKubernetesService_Create_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	ctx := context.Background()
	_, err := client.Kubernetes().Create(ctx, CreateKubernetesParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestKubernetesService_Read_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	clusterId := "11111"
	expectedResponse := dummyReadKubernetesRes
	serverResponse := dummyKubernetesServerRes

	mux.HandleFunc("/kubernetes/"+clusterId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want KubernetesCluster
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	ctx := context.Background()
	got, _ := client.Kubernetes().Read(ctx, clusterId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestKubernetesService_Read_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	ctx := context.Background()
	apikey, err := client.Kubernetes().Read(ctx, "someId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if apikey != nil {
		t.Errorf("Was not expecting any apikey to be returned, instead got %v", apikey)
	}
}

func TestKubernetesService_List_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	expectedResponse := dummyListKubernetesRes
	serverResponse := dummyKubernetesServerRes

	mux.HandleFunc("/kubernetes", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want []K8s
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	ctx := context.Background()
	got, _ := client.Kubernetes().List(ctx)
	if len(got) != len(want) {
		t.Errorf("Was expecting %d stacks to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestKubernetesService_List_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	ctx := context.Background()
	stacks, err := client.Kubernetes().List(ctx)
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if stacks != nil {
		t.Errorf("Was not expecting any stacks to be returned, instead got %v", stacks)
	}
}

func TestKubernetesService_Delete_happyPath(t *testing.T) {
	token := "token"
	payload := DeleteKubernetesParams{
		ClusterId: "11111",
		Confirm:   "I am aware this action will delete data and cluster permanently",
	}

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/kubernetes/"+payload.ClusterId+"/destroy", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyDeleteResponseJson)
	})

	want := DeleteResponse{Status: "success", Message: "success"}

	ctx := context.Background()
	got, _ := client.Kubernetes().Delete(ctx, payload)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestKubernetesService_Delete_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	ctx := context.Background()
	delResponse, err := client.Kubernetes().Delete(ctx, DeleteKubernetesParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
	}
}

// K8s Loadbalancer test
func TestKubernetesServices_CreateLoadbalancer_happyPath(t *testing.T) {
	token := "token"

	payload := CreateKubernetesLoadbalancerParams{
		KubernetesId:   "11111",
		LoadbalancerId: "22222",
	}

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/kubernetes/"+payload.KubernetesId+"/loadbalancer/"+payload.LoadbalancerId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateBasicResponseJson)
	})

	ctx := context.Background()
	got, err := client.Kubernetes().CreateLoadbalancer(ctx, payload)

	var want CreateResponse
	_ = json.Unmarshal([]byte(dummyCreateBasicResponseJson), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestKubernetesServices_CreateLoadbalancer_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	ctx := context.Background()
	_, err := client.Kubernetes().CreateLoadbalancer(ctx, CreateKubernetesLoadbalancerParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestKubernetesServices_ReadLoadbalancer_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	kubernetesId := "11111"
	loadbalancerID := "22222"
	expectedResponse := dummyReadKubernetesLoadbalancerRes
	serverResponse := dummyKubernetesServerRes

	mux.HandleFunc("/kubernetes/"+kubernetesId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want K8sLoadbalancers
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	ctx := context.Background()
	got, _ := client.Kubernetes().ReadLoadbalancer(ctx, kubernetesId, loadbalancerID)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestKubernetesServices_ReadLoadbalancer_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	ctx := context.Background()
	apikey, err := client.Kubernetes().ReadLoadbalancer(ctx, "11111", "122134")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if apikey != nil {
		t.Errorf("Was not expecting any apikey to be returned, instead got %v", apikey)
	}
}

func TestKubernetesServices_ListLoadbalancer_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	kubernetesId := "11111"
	expectedResponse := dummyListKubernetesLoadbalancerRes
	serverResponse := dummyKubernetesServerRes

	mux.HandleFunc("/kubernetes/"+kubernetesId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want []K8sLoadbalancers
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	ctx := context.Background()
	got, _ := client.Kubernetes().ListLoadbalancers(ctx, kubernetesId)
	if len(got) != len(want) {
		t.Errorf("Was expecting %d kubernetes loadbalancer to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestKubernetesServices_ListLoadbalancer_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	ctx := context.Background()
	loadbalancer, err := client.Kubernetes().ListLoadbalancers(ctx, "11111")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if loadbalancer != nil {
		t.Errorf("Was not expecting any kubernetes loadbalancer to be returned, instead got %v", loadbalancer)
	}
}

func TestKubernetesServices_DeleteLoadbalancer_happyPath(t *testing.T) {
	token := "token"
	kubernetesId := "11111"
	loadbalancerId := "22222"

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/kubernetes/"+kubernetesId+"/loadbalancerpolicy/"+loadbalancerId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyDeleteResponseJson)
	})

	want := DeleteResponse{Status: "success", Message: "success"}

	ctx := context.Background()
	got, _ := client.Kubernetes().DeleteLoadbalancer(ctx, kubernetesId, loadbalancerId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestKubernetesServices_DeleteLoadbalancer_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	ctx := context.Background()
	delResponse, err := client.Kubernetes().DeleteLoadbalancer(ctx, "someLoadbalancerName", "123543")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
	}
}

// K8s SecurityGroup test
func TestKubernetesServices_CreateSecurityGroup_happyPath(t *testing.T) {
	token := "token"

	payload := CreateKubernetesSecurityGroupParams{
		KubernetesId:              "11111",
		KubernetesSecurityGroupId: "44444",
	}

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/kubernetes/"+payload.KubernetesId+"/securitygroup/"+payload.KubernetesSecurityGroupId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateBasicResponseJson)
	})

	ctx := context.Background()
	got, err := client.Kubernetes().CreateSecurityGroup(ctx, payload)

	var want CreateResponse
	_ = json.Unmarshal([]byte(dummyCreateBasicResponseJson), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestKubernetesServices_CreateSecurityGroup_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	ctx := context.Background()
	_, err := client.Kubernetes().CreateSecurityGroup(ctx, CreateKubernetesSecurityGroupParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestKubernetesServices_ReadSecurityGroup_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	kubernetesId := "11111"
	securityGroupId := "44444"
	expectedResponse := dummyReadKubernetesSecurityGroupRes
	serverResponse := dummyKubernetesServerRes

	mux.HandleFunc("/kubernetes/"+kubernetesId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want K8sSecurityGroups
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	ctx := context.Background()
	got, _ := client.Kubernetes().ReadSecurityGroup(ctx, kubernetesId, securityGroupId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestKubernetesServices_ReadSecurityGroup_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	ctx := context.Background()
	apikey, err := client.Kubernetes().ReadSecurityGroup(ctx, "11111", "122134")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if apikey != nil {
		t.Errorf("Was not expecting any apikey to be returned, instead got %v", apikey)
	}
}

func TestKubernetesServices_ListSecurityGroup_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	kubernetesId := "11111"
	expectedResponse := dummyListKubernetesSecurityGroupRes
	serverResponse := dummyKubernetesServerRes

	mux.HandleFunc("/kubernetes/"+kubernetesId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want []K8sSecurityGroups
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	ctx := context.Background()
	got, _ := client.Kubernetes().ListSecurityGroups(ctx, kubernetesId)
	if len(got) != len(want) {
		t.Errorf("Was expecting %d kubernetes securitygroup to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestKubernetesServices_ListSecurityGroup_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	ctx := context.Background()
	securitygroup, err := client.Kubernetes().ListSecurityGroups(ctx, "11111")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if securitygroup != nil {
		t.Errorf("Was not expecting any kubernetes securitygroup to be returned, instead got %v", securitygroup)
	}
}

func TestKubernetesServices_DeleteSecurityGroup_happyPath(t *testing.T) {
	token := "token"
	kubernetesId := "11111"
	securityGroupId := "44444"

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/kubernetes/"+kubernetesId+"/securitygroup/"+securityGroupId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyDeleteResponseJson)
	})

	want := DeleteResponse{Status: "success", Message: "success"}

	ctx := context.Background()
	got, _ := client.Kubernetes().DeleteSecurityGroup(ctx, kubernetesId, securityGroupId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestKubernetesServices_DeleteSecurityGroup_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	ctx := context.Background()
	delResponse, err := client.Kubernetes().DeleteSecurityGroup(ctx, "someSecurityGroupName", "123543")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
	}
}

// K8s Targetgroup test
func TestKubernetesServices_CreateTargetgroup_happyPath(t *testing.T) {
	token := "token"

	payload := CreateKubernetesTargetgroupParams{
		KubernetesId:            "11111",
		KubernetesTargetgroupId: "33333",
	}

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/kubernetes/"+payload.KubernetesId+"/targetgroup/"+payload.KubernetesTargetgroupId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateBasicResponseJson)
	})

	ctx := context.Background()
	got, err := client.Kubernetes().CreateTargetgroup(ctx, payload)

	var want CreateResponse
	_ = json.Unmarshal([]byte(dummyCreateBasicResponseJson), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestKubernetesServices_CreateTargetgroup_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	ctx := context.Background()
	_, err := client.Kubernetes().CreateTargetgroup(ctx, CreateKubernetesTargetgroupParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestKubernetesServices_ReadTargetgroup_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	kubernetesId := "11111"
	targetgroupId := "33333"
	expectedResponse := dummyReadKubernetesTargetgroupRes
	serverResponse := dummyKubernetesServerRes

	mux.HandleFunc("/kubernetes/"+kubernetesId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want K8sTargetGroups
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	ctx := context.Background()
	got, _ := client.Kubernetes().ReadTargetgroup(ctx, kubernetesId, targetgroupId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestKubernetesServices_ReadTargetgroup_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	ctx := context.Background()
	apikey, err := client.Kubernetes().ReadTargetgroup(ctx, "11111", "122134")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if apikey != nil {
		t.Errorf("Was not expecting any apikey to be returned, instead got %v", apikey)
	}
}

func TestKubernetesServices_ListTargetgroup_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	kubernetesId := "11111"
	expectedResponse := dummyListKubernetesTargetgroupRes
	serverResponse := dummyKubernetesServerRes

	mux.HandleFunc("/kubernetes/"+kubernetesId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want []K8sTargetGroups
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	ctx := context.Background()
	got, _ := client.Kubernetes().ListTargetgroups(ctx, kubernetesId)
	if len(got) != len(want) {
		t.Errorf("Was expecting %d kubernetes targetgroup to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestKubernetesServices_ListTargetgroup_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	ctx := context.Background()
	targetgroup, err := client.Kubernetes().ListTargetgroups(ctx, "11111")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if targetgroup != nil {
		t.Errorf("Was not expecting any kubernetes targetgroup to be returned, instead got %v", targetgroup)
	}
}

func TestKubernetesServices_DeleteTargetgroup_happyPath(t *testing.T) {
	token := "token"
	kubernetesId := "11111"
	targetgroupId := "33333"

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/kubernetes/"+kubernetesId+"/targetgroup/"+targetgroupId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyDeleteResponseJson)
	})

	want := DeleteResponse{Status: "success", Message: "success"}

	ctx := context.Background()
	got, _ := client.Kubernetes().DeleteTargetgroup(ctx, kubernetesId, targetgroupId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestKubernetesServices_DeleteTargetgroup_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	ctx := context.Background()
	delResponse, err := client.Kubernetes().DeleteTargetgroup(ctx, "someTargetgroupName", "123543")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
	}
}

func TestKubernetesService_PowerOn_happyPath(t *testing.T) {
	token := "token"
	kubernetesId := "someId"

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/kubernetes/"+kubernetesId+"/start", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateBasicResponseJson)
	})

	ctx := context.Background()
	got, err := client.Kubernetes().PowerOn(ctx, kubernetesId)

	var want BasicResponse
	_ = json.Unmarshal([]byte(dummyCreateBasicResponseJson), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestKubernetesService_PowerOn_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	ctx := context.Background()
	_, err := client.Kubernetes().PowerOn(ctx, "kubernetesId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestKubernetesService_PowerOff_happyPath(t *testing.T) {
	token := "token"
	kubernetesId := "someId"

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/kubernetes/"+kubernetesId+"/stop", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateBasicResponseJson)
	})

	ctx := context.Background()
	got, err := client.Kubernetes().PowerOff(ctx, kubernetesId)

	var want BasicResponse
	_ = json.Unmarshal([]byte(dummyCreateBasicResponseJson), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestKubernetesService_PowerOff_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	ctx := context.Background()
	_, err := client.Kubernetes().PowerOff(ctx, "kubernetesId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}
