package utho

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestLoadbalancerService_Create_happyPath(t *testing.T) {
	token := "token"
	payload := CreateLoadblancerParams{}
	_ = faker.FakeData(&payload)

	client, mux, _, teardown := setup(token)
	defer teardown()

	var dummyResponse CreateLoadbalancerResponse
	_ = faker.FakeData(&dummyResponse)
	dummyResponse.Status = "success"

	serverResponse, _ := json.Marshal(dummyResponse)

	mux.HandleFunc("/loadbalancer/add", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(serverResponse))
	})

	got, err := client.Loadbalancers().Create(payload)

	// Ensure the error is nil and the response matches the expected dummy response
	assert.Nil(t, err, "Expected no error, but got one")
	assert.NotNil(t, got, "Expected a valid response, but got nil")
	assert.Equal(t, dummyResponse, *got, "Response does not match the expected dummy response")
}

func TestLoadbalancerService_Create_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Loadbalancers().Create(CreateLoadblancerParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestLoadbalancerService_Read_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	loadbalancerId := faker.UUIDDigit()
	var dummyLoadbalancer Loadbalancer
	_ = faker.FakeData(&dummyLoadbalancer)

	serverResponse, _ := json.Marshal(Loadbalancers{
		Loadbalancers: []Loadbalancer{dummyLoadbalancer},
		Status:        "success",
	})

	mux.HandleFunc("/loadbalancer/"+loadbalancerId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, string(serverResponse))
	})

	got, _ := client.Loadbalancers().Read(loadbalancerId)
	assert.Equal(t, dummyLoadbalancer, *got)
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

	var dummyLoadbalancers []Loadbalancer
	for i := 0; i < 3; i++ {
		var loadbalancer Loadbalancer
		_ = faker.FakeData(&loadbalancer)
		dummyLoadbalancers = append(dummyLoadbalancers, loadbalancer)
	}

	serverResponse, _ := json.Marshal(Loadbalancers{
		Loadbalancers: dummyLoadbalancers,
		Status:        "success",
	})

	mux.HandleFunc("/loadbalancer", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, string(serverResponse))
	})

	got, _ := client.Loadbalancers().List()
	assert.Equal(t, dummyLoadbalancers, got)
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

	mux.HandleFunc("/loadbalancer/"+loadbalancerId+"/destroy", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyDeleteResponseJson)
	})

	want := DeleteResponse{Status: "success", Message: "success"}

	got, _ := client.Loadbalancers().Delete(loadbalancerId)
	assert.Equal(t, want, *got)
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
	payload := CreateLoadbalancerACLParams{}
	_ = faker.FakeData(&payload)

	client, mux, _, teardown := setup(token)
	defer teardown()

	var dummyResponse CreateResponse
	_ = faker.FakeData(&dummyResponse)
	dummyResponse.Status = "success"

	serverResponse, _ := json.Marshal(dummyResponse)

	mux.HandleFunc("/loadbalancer/"+payload.LoadbalancerId+"/acl", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, string(serverResponse))
	})

	got, err := client.Loadbalancers().CreateACL(payload)

	assert.Nil(t, err)
	assert.Equal(t, dummyResponse, *got)
}

func TestLoadbalancerService_CreateACL_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Loadbalancers().Create(CreateLoadblancerParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestLoadbalancerService_ReadACL_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	loadbalancerId := faker.UUIDDigit()
	loadbalancerACLId := faker.UUIDDigit()
	var dummyACL ACLs
	_ = faker.FakeData(&dummyACL)
	dummyACL.ID = loadbalancerACLId

	serverResponse, _ := json.Marshal(Loadbalancers{
		Loadbalancers: []Loadbalancer{
			{
				Acls: []ACLs{dummyACL},
			},
		},
		Status: "success",
	})

	mux.HandleFunc("/loadbalancer/"+loadbalancerId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, string(serverResponse))
	})

	got, _ := client.Loadbalancers().ReadACL(loadbalancerId, loadbalancerACLId)
	assert.Equal(t, dummyACL, *got)
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

	loadbalancerId := faker.UUIDDigit()
	var dummyACLs []ACLs
	for i := 0; i < 3; i++ {
		var acl ACLs
		_ = faker.FakeData(&acl)
		dummyACLs = append(dummyACLs, acl)
	}

	serverResponse, _ := json.Marshal(Loadbalancers{
		Loadbalancers: []Loadbalancer{
			{
				Acls: dummyACLs,
			},
		},
		Status: "success",
	})

	mux.HandleFunc("/loadbalancer/"+loadbalancerId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, string(serverResponse))
	})

	got, _ := client.Loadbalancers().ListACLs(loadbalancerId)
	assert.Equal(t, dummyACLs, got)
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
	assert.Equal(t, want, *got)
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
	payload := CreateLoadbalancerFrontendParams{}
	_ = faker.FakeData(&payload)

	client, mux, _, teardown := setup(token)
	defer teardown()

	var dummyResponse CreateResponse
	_ = faker.FakeData(&dummyResponse)
	dummyResponse.Status = "success"

	serverResponse, _ := json.Marshal(dummyResponse)

	mux.HandleFunc("/loadbalancer/"+payload.LoadbalancerId+"/frontend", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, string(serverResponse))
	})

	got, err := client.Loadbalancers().CreateFrontend(payload)

	assert.Nil(t, err)
	assert.Equal(t, dummyResponse, *got)
}

func TestLoadbalancerService_CreateFrontend_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Loadbalancers().Create(CreateLoadblancerParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestLoadbalancerService_ReadFrontend_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	loadbalancerId := faker.UUIDDigit()
	frontendId := faker.UUIDDigit()
	var dummyFrontend Frontends
	_ = faker.FakeData(&dummyFrontend)
	dummyFrontend.ID = frontendId

	serverResponse, _ := json.Marshal(Loadbalancers{
		Loadbalancers: []Loadbalancer{
			{
				Frontends: []Frontends{dummyFrontend},
			},
		},
		Status: "success",
	})

	mux.HandleFunc("/loadbalancer/"+loadbalancerId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, string(serverResponse))
	})

	got, _ := client.Loadbalancers().ReadFrontend(loadbalancerId, frontendId)
	assert.Equal(t, dummyFrontend, *got)
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

	loadbalancerId := faker.UUIDDigit()
	var dummyFrontends []Frontends
	for i := 0; i < 3; i++ {
		var frontend Frontends
		_ = faker.FakeData(&frontend)
		dummyFrontends = append(dummyFrontends, frontend)
	}

	serverResponse, _ := json.Marshal(Loadbalancers{
		Loadbalancers: []Loadbalancer{
			{
				Frontends: dummyFrontends,
			},
		},
		Status: "success",
	})

	mux.HandleFunc("/loadbalancer/"+loadbalancerId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, string(serverResponse))
	})

	got, _ := client.Loadbalancers().ListFrontends(loadbalancerId)
	assert.Equal(t, dummyFrontends, got)
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
	assert.Equal(t, want, *got)
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
	payload := CreateLoadbalancerBackendParams{}
	_ = faker.FakeData(&payload)

	client, mux, _, teardown := setup(token)
	defer teardown()

	var dummyResponse CreateResponse
	_ = faker.FakeData(&dummyResponse)
	dummyResponse.Status = "success"

	serverResponse, _ := json.Marshal(dummyResponse)

	mux.HandleFunc("/loadbalancer/"+payload.LoadbalancerId+"/backend", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, string(serverResponse))
	})

	got, err := client.Loadbalancers().CreateBackend(payload)

	assert.Nil(t, err)
	assert.Equal(t, dummyResponse, *got)
}

func TestLoadbalancerService_CreateBackend_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Loadbalancers().Create(CreateLoadblancerParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestLoadbalancerService_ReadBackend_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	loadbalancerId := faker.UUIDDigit()
	backendId := faker.UUIDDigit()
	var dummyBackend Backends
	_ = faker.FakeData(&dummyBackend)
	dummyBackend.ID = backendId

	serverResponse, _ := json.Marshal(Loadbalancers{
		Loadbalancers: []Loadbalancer{
			{
				Backends: []Backends{dummyBackend},
			},
		},
		Status: "success",
	})

	mux.HandleFunc("/loadbalancer/"+loadbalancerId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, string(serverResponse))
	})

	got, _ := client.Loadbalancers().ReadBackend(loadbalancerId, backendId)
	assert.Equal(t, dummyBackend, *got)
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

	loadbalancerId := faker.UUIDDigit()
	var dummyBackends []Backends
	for i := 0; i < 3; i++ {
		var backend Backends
		_ = faker.FakeData(&backend)
		dummyBackends = append(dummyBackends, backend)
	}

	serverResponse, _ := json.Marshal(Loadbalancers{
		Loadbalancers: []Loadbalancer{
			{
				Backends: dummyBackends,
			},
		},
		Status: "success",
	})

	mux.HandleFunc("/loadbalancer/"+loadbalancerId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, string(serverResponse))
	})

	got, _ := client.Loadbalancers().ListBackends(loadbalancerId)
	assert.Equal(t, dummyBackends, got)
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
	assert.Equal(t, want, *got)
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
	payload := CreateLoadbalancerRouteParams{}
	_ = faker.FakeData(&payload)

	client, mux, _, teardown := setup(token)
	defer teardown()

	var dummyResponse CreateResponse
	_ = faker.FakeData(&dummyResponse)
	dummyResponse.Status = "success"

	serverResponse, _ := json.Marshal(dummyResponse)

	mux.HandleFunc("/loadbalancer/"+payload.LoadbalancerId+"/route", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, string(serverResponse))
	})

	got, err := client.Loadbalancers().CreateRoute(payload)

	assert.Nil(t, err)
	assert.Equal(t, dummyResponse, *got)
}

func TestLoadbalancerService_CreateRoute_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Loadbalancers().Create(CreateLoadblancerParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestLoadbalancerService_ReadRoute_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	loadbalancerId := faker.UUIDDigit()
	routeId := faker.UUIDDigit()
	var dummyRoute Routes
	_ = faker.FakeData(&dummyRoute)
	dummyRoute.ID = routeId

	serverResponse, _ := json.Marshal(Loadbalancers{
		Loadbalancers: []Loadbalancer{
			{
				Routes: []Routes{dummyRoute},
			},
		},
		Status: "success",
	})

	mux.HandleFunc("/loadbalancer/"+loadbalancerId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, string(serverResponse))
	})

	got, _ := client.Loadbalancers().ReadRoute(loadbalancerId, routeId)
	assert.Equal(t, dummyRoute, *got)
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

	loadbalancerId := faker.UUIDDigit()
	var dummyRoutes []Routes
	for i := 0; i < 3; i++ {
		var route Routes
		_ = faker.FakeData(&route)
		dummyRoutes = append(dummyRoutes, route)
	}

	serverResponse, _ := json.Marshal(Loadbalancers{
		Loadbalancers: []Loadbalancer{
			{
				Routes: dummyRoutes,
			},
		},
		Status: "success",
	})

	mux.HandleFunc("/loadbalancer/"+loadbalancerId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, string(serverResponse))
	})

	got, _ := client.Loadbalancers().ListRoutes(loadbalancerId)
	assert.Equal(t, dummyRoutes, got)
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
	assert.Equal(t, want, *got)
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
