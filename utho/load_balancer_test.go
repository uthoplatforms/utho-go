package utho

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadbalancerService_Create_happyPath(t *testing.T) {
	token := "token"
	payload := CreateLoadbalancerParams{
		Dcslug: "innoida",
		Name:   "example",
		Type:   "application",
	}

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/loadbalancer", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateLoadbalancerResponseJson)
	})

	got, err := client.Loadbalancers().client.Loadbalancers().Create(payload)

	var want CreateLoadbalancerResponse
	_ = json.Unmarshal([]byte(dummyCreateLoadbalancerResponseJson), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestLoadbalancerService_Create_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Loadbalancers().Create(CreateLoadbalancerParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestLoadbalancerService_Read_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	loadbalancerId := "11111"
	expectedResponse := dummyReadLoadbalancerRes
	serverResponse := dummyReadLoadbalancerServerRes

	mux.HandleFunc("/loadbalancer/"+loadbalancerId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want Loadbalancer
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.Loadbalancers().Read(loadbalancerId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestLoadbalancerService_Read_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	apikey, err := client.Loadbalancers().Read("someId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if apikey != nil {
		t.Errorf("Was not expecting any apikey to be returned, instead got %v", apikey)
	}
}

func TestLoadbalancerService_List_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	expectedResponse := dummyReadLoadbalancerRes
	serverResponse := "[" + dummyReadLoadbalancerServerRes + "]"

	mux.HandleFunc("/loadbalancer", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want []Loadbalancer
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.Loadbalancers().List()
	if len(got) != len(want) {
		t.Errorf("Was expecting %d loadbalancer to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestLoadbalancerService_List_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	loadbalancer, err := client.Loadbalancers().List()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if loadbalancer != nil {
		t.Errorf("Was not expecting any loadbalancer to be returned, instead got %v", loadbalancer)
	}
}

func TestLoadbalancerService_Delete_happyPath(t *testing.T) {
	token := "token"
	loadbalancerId := "someLoadbalancerId"

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/loadbalancer/"+loadbalancerId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyDeleteResponseJson)
	})

	want := DeleteResponse{Status: "success", Message: "success"}

	got, _ := client.Loadbalancers().Delete(loadbalancerId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestLoadbalancerService_Delete_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	delResponse, err := client.Loadbalancers().Delete("someLoadbalancerId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
	}
}

// loadbalancer ACL
func TestLoadbalancerService_CreateACL_happyPath(t *testing.T) {
	token := "token"
	payload := CreateLoadbalancerACLParams{
		LoadbalancerId: "1231",
		Name:           "example",
		ConditionType:  "http_user_agent",
		Value:          "{'backend_id':'12324','type':'http_user_agent','data':['value1','value2']}",
	}

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/loadbalancer/"+payload.LoadbalancerId+"/acl", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateResponseJson)
	})

	got, err := client.Loadbalancers().client.Loadbalancers().CreateACL(payload)

	var want CreateResponse
	_ = json.Unmarshal([]byte(dummyCreateResponseJson), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestLoadbalancerService_CreateACL_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Loadbalancers().Create(CreateLoadbalancerParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestLoadbalancerService_ReadACL_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	loadbalancerId := "11111"
	loadbalancerACLId := "22222"
	expectedResponse := dummyReadLoadbalancAclRes
	serverResponse := dummyReadLoadbalancerServerRes

	mux.HandleFunc("/loadbalancer/"+loadbalancerId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want ACLs
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.Loadbalancers().ReadACL(loadbalancerId, loadbalancerACLId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestLoadbalancerService_ReadACL_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	apikey, err := client.Loadbalancers().Read("someId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if apikey != nil {
		t.Errorf("Was not expecting any apikey to be returned, instead got %v", apikey)
	}
}

func TestLoadbalancerService_ListACL_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	loadbalancerId := "11111"
	expectedResponse := "[" + dummyReadLoadbalancAclRes + "]"
	serverResponse := dummyReadLoadbalancerServerRes

	mux.HandleFunc("/loadbalancer/"+loadbalancerId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want []ACLs
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.Loadbalancers().ListACLs(loadbalancerId)
	if len(got) != len(want) {
		t.Errorf("Was expecting %d loadbalancer to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestLoadbalancerService_ListACL_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	loadbalancer, err := client.Loadbalancers().List()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if loadbalancer != nil {
		t.Errorf("Was not expecting any loadbalancer to be returned, instead got %v", loadbalancer)
	}
}

func TestLoadbalancerService_DeleteACL_happyPath(t *testing.T) {
	token := "token"
	loadbalancerId := "someLoadbalancerId"
	loadbalancerACLId := "22222"

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/loadbalancer/"+loadbalancerId+"/acl/"+loadbalancerACLId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyDeleteResponseJson)
	})

	want := DeleteResponse{Status: "success", Message: "success"}

	got, _ := client.Loadbalancers().DeleteACL(loadbalancerId, loadbalancerACLId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestLoadbalancerService_DeleteACL_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	delResponse, err := client.Loadbalancers().Delete("someLoadbalancerId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
	}
}

// loadbalancer Frontend
func TestLoadbalancerService_CreateFrontend_happyPath(t *testing.T) {
	token := "token"
	payload := CreateLoadbalancerFrontendParams{
		LoadbalancerId: "1231",
		Name:           "example",
		Proto:          "http",
		Port:           "80",
		Algorithm:      "roundrobin",
		Redirecthttps:  "1",
		Cookie:         "1",
	}

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/loadbalancer/"+payload.LoadbalancerId+"/frontend", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateResponseJson)
	})

	got, err := client.Loadbalancers().client.Loadbalancers().CreateFrontend(payload)

	var want CreateResponse
	_ = json.Unmarshal([]byte(dummyCreateResponseJson), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestLoadbalancerService_CreateFrontend_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Loadbalancers().Create(CreateLoadbalancerParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestLoadbalancerService_ReadFrontend_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	loadbalancerId := "11111"
	loadbalancerFrontendId := "22222"
	expectedResponse := dummyReadLoadbalancFrontendRes
	serverResponse := dummyReadLoadbalancerServerRes

	mux.HandleFunc("/loadbalancer/"+loadbalancerId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want Frontends
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.Loadbalancers().ReadFrontend(loadbalancerId, loadbalancerFrontendId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestLoadbalancerService_ReadFrontend_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	apikey, err := client.Loadbalancers().ReadFrontend("someId", "id")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if apikey != nil {
		t.Errorf("Was not expecting any apikey to be returned, instead got %v", apikey)
	}
}

func TestLoadbalancerService_ListFrontend_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	loadbalancerId := "11111"
	expectedResponse := "[" + dummyReadLoadbalancFrontendRes + "]"
	serverResponse := dummyReadLoadbalancerServerRes

	mux.HandleFunc("/loadbalancer/"+loadbalancerId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want []Frontends
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.Loadbalancers().ListFrontends(loadbalancerId)
	if len(got) != len(want) {
		t.Errorf("Was expecting %d loadbalancer to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestLoadbalancerService_ListFrontend_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	loadbalancer, err := client.Loadbalancers().List()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if loadbalancer != nil {
		t.Errorf("Was not expecting any loadbalancer to be returned, instead got %v", loadbalancer)
	}
}

func TestLoadbalancerService_DeleteFrontend_happyPath(t *testing.T) {
	token := "token"
	loadbalancerId := "someLoadbalancerId"
	loadbalancerFrontendId := "22222"

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/loadbalancer/"+loadbalancerId+"/frontend/"+loadbalancerFrontendId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyDeleteResponseJson)
	})

	want := DeleteResponse{Status: "success", Message: "success"}

	got, _ := client.Loadbalancers().DeleteFrontend(loadbalancerId, loadbalancerFrontendId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestLoadbalancerService_DeleteFrontend_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	delResponse, err := client.Loadbalancers().Delete("someLoadbalancerId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
	}
}

// loadbalancer Backend
func TestLoadbalancerService_CreateBackend_happyPath(t *testing.T) {
	token := "token"
	payload := CreateLoadbalancerBackendParams{
		LoadbalancerId: "1231",
		BackendPort:    "43",
		Cloudid:        "1277662",
		FrontendID:     "169",
	}

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/loadbalancer/"+payload.LoadbalancerId+"/backend", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateResponseJson)
	})

	got, err := client.Loadbalancers().client.Loadbalancers().CreateBackend(payload)

	var want CreateResponse
	_ = json.Unmarshal([]byte(dummyCreateResponseJson), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestLoadbalancerService_CreateBackend_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Loadbalancers().Create(CreateLoadbalancerParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestLoadbalancerService_ReadBackend_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	loadbalancerId := "11111"
	loadbalancerBackendId := "22222"
	expectedResponse := dummyReadLoadbalancBackendRes
	serverResponse := dummyReadLoadbalancerServerRes

	mux.HandleFunc("/loadbalancer/"+loadbalancerId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want Backends
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.Loadbalancers().ReadBackend(loadbalancerId, loadbalancerBackendId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestLoadbalancerService_ReadBackend_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	apikey, err := client.Loadbalancers().ReadBackend("someId", "id")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if apikey != nil {
		t.Errorf("Was not expecting any apikey to be returned, instead got %v", apikey)
	}
}

func TestLoadbalancerService_ListBackend_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	loadbalancerId := "11111"
	expectedResponse := "[" + dummyReadLoadbalancBackendRes + "]"
	serverResponse := dummyReadLoadbalancerServerRes

	mux.HandleFunc("/loadbalancer/"+loadbalancerId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want []Backends
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.Loadbalancers().ListBackends(loadbalancerId)
	if len(got) != len(want) {
		t.Errorf("Was expecting %d loadbalancer to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestLoadbalancerService_ListBackend_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	loadbalancer, err := client.Loadbalancers().List()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if loadbalancer != nil {
		t.Errorf("Was not expecting any loadbalancer to be returned, instead got %v", loadbalancer)
	}
}

func TestLoadbalancerService_DeleteBackend_happyPath(t *testing.T) {
	token := "token"
	loadbalancerId := "someLoadbalancerId"
	loadbalancerBackendId := "22222"

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/loadbalancer/"+loadbalancerId+"/backend/"+loadbalancerBackendId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyDeleteResponseJson)
	})

	want := DeleteResponse{Status: "success", Message: "success"}

	got, _ := client.Loadbalancers().DeleteBackend(loadbalancerId, loadbalancerBackendId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestLoadbalancerService_DeleteBackend_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	delResponse, err := client.Loadbalancers().Delete("someLoadbalancerId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
	}
}

// loadbalancer Route
func TestLoadbalancerService_CreateRoute_happyPath(t *testing.T) {
	token := "token"
	payload := CreateLoadbalancerRouteParams{
		LoadbalancerId: "1231",
		FrontendID:     "169",
		ACLID:          "1223",
		RouteCondition: "true",
		TargetGroups:   "231,243,3234,4234",
	}

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/loadbalancer/"+payload.LoadbalancerId+"/route", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateResponseJson)
	})

	got, err := client.Loadbalancers().client.Loadbalancers().CreateRoute(payload)

	var want CreateResponse
	_ = json.Unmarshal([]byte(dummyCreateResponseJson), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestLoadbalancerService_CreateRoute_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Loadbalancers().Create(CreateLoadbalancerParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestLoadbalancerService_ReadRoute_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	loadbalancerId := "11111"
	loadbalancerRouteId := "22222"
	expectedResponse := dummyReadLoadbalancRouteRes
	serverResponse := dummyReadLoadbalancerServerRes

	mux.HandleFunc("/loadbalancer/"+loadbalancerId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want Routes
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.Loadbalancers().ReadRoute(loadbalancerId, loadbalancerRouteId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestLoadbalancerService_ReadRoute_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	apikey, err := client.Loadbalancers().ReadRoute("someId", "id")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if apikey != nil {
		t.Errorf("Was not expecting any apikey to be returned, instead got %v", apikey)
	}
}

func TestLoadbalancerService_ListRoute_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	loadbalancerId := "11111"
	expectedResponse := "[" + dummyReadLoadbalancRouteRes + "]"
	serverResponse := dummyReadLoadbalancerServerRes

	mux.HandleFunc("/loadbalancer/"+loadbalancerId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want []Routes
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.Loadbalancers().ListRoutes(loadbalancerId)
	if len(got) != len(want) {
		t.Errorf("Was expecting %d loadbalancer to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestLoadbalancerService_ListRoute_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	loadbalancer, err := client.Loadbalancers().List()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if loadbalancer != nil {
		t.Errorf("Was not expecting any loadbalancer to be returned, instead got %v", loadbalancer)
	}
}

func TestLoadbalancerService_DeleteRoute_happyPath(t *testing.T) {
	token := "token"
	loadbalancerId := "someLoadbalancerId"
	loadbalancerRouteId := "22222"

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/loadbalancer/"+loadbalancerId+"/route/"+loadbalancerRouteId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyDeleteResponseJson)
	})

	want := DeleteResponse{Status: "success", Message: "success"}

	got, _ := client.Loadbalancers().DeleteRoute(loadbalancerId, loadbalancerRouteId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestLoadbalancerService_DeleteRoute_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	delResponse, err := client.Loadbalancers().Delete("someLoadbalancerId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
	}
}
