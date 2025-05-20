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

func TestVpcService_Create_happyPath(t *testing.T) {
	token := "token"
	var payload CreateVpcParams
	_ = faker.FakeData(&payload)

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp CreateResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/vpc/create", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	got, err := client.Vpc().Create(payload)

	var want CreateResponse
	_ = json.Unmarshal(respBytes, &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestVpcService_Create_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Vpc().Create(CreateVpcParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestVpcService_Read_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	var fakeVpc Vpc
	_ = faker.FakeData(&fakeVpc)
	fakeVpc.Status = "success"
	serverResp := struct {
		Vpc []Vpc `json:"vpc"`
	}{
		Vpc: []Vpc{fakeVpc},
	}
	serverResponse, _ := json.Marshal(serverResp)
	expectedResponse, _ := json.Marshal(fakeVpc)

	mux.HandleFunc("/vpc", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		w.Write(serverResponse)
	})

	var want Vpc
	_ = json.Unmarshal(expectedResponse, &want)

	got, _ := client.Vpc().Read(fakeVpc.ID)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestVpcService_Read_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	vpc, err := client.Vpc().Read("someId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if vpc != nil {
		t.Errorf("Was not expecting any vpc to be returned, instead got %v", vpc)
	}
}

func TestVpcService_List_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	var fakeVpcs []Vpc
	for i := 0; i < 2; i++ {
		var v Vpc
		_ = faker.FakeData(&v)
		fakeVpcs = append(fakeVpcs, v)
	}
	serverResp := struct {
		Vpc []Vpc `json:"vpc"`
	}{
		Vpc: fakeVpcs,
	}
	serverResponse, _ := json.Marshal(serverResp)
	expectedResponse, _ := json.Marshal(fakeVpcs)

	mux.HandleFunc("/vpc", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		w.Write(serverResponse)
	})

	var want []Vpc
	_ = json.Unmarshal(expectedResponse, &want)

	got, _ := client.Vpc().List()
	if len(got) != len(want) {
		t.Errorf("Was expecting %d vpcs to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestVpcService_List_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	vpcs, err := client.Vpc().List()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if vpcs != nil {
		t.Errorf("Was not expecting any vpcs to be returned, instead got %v", vpcs)
	}
}

func TestVpcService_Delete_happyPath(t *testing.T) {
	token := "token"
	var fakeVpc Vpc
	_ = faker.FakeData(&fakeVpc)
	vpcId := fakeVpc.ID

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp DeleteResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	fakeResp.Message = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/vpc/"+vpcId+"/destroy", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	want := fakeResp

	got, _ := client.Vpc().Delete(vpcId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestVpcService_Delete_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	delResponse, err := client.Vpc().Delete("someVpcId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
	}
}

func TestVpcService_CreateSubnet_happyPath(t *testing.T) {
	token := "token"
	var payload CreateSubnetParams
	_ = faker.FakeData(&payload)

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp CreateResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/vpc/subnet/create", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	got, err := client.Vpc().CreateSubnet(payload)

	var want CreateResponse
	_ = json.Unmarshal(respBytes, &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestVpcService_CreateSubnet_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Vpc().CreateSubnet(CreateSubnetParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestVpcService_CreateNatGateway_happyPath(t *testing.T) {
	token := "token"
	var payload CreateNatGatewayParams
	_ = faker.FakeData(&payload)

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp CreateResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/vpc/natgateway", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	got, err := client.Vpc().CreateNatGateway(payload)

	var want CreateResponse
	_ = json.Unmarshal(respBytes, &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestVpcService_CreateNatGateway_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Vpc().CreateNatGateway(CreateNatGatewayParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestVpcService_AllocateElasticIP_happyPath(t *testing.T) {
	token := "token"
	var payload AllocateElasticIPParams
	_ = faker.FakeData(&payload)

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp CreateResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/elasticip/allocate", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	got, err := client.Vpc().AllocateElasticIP(payload)

	var want CreateResponse
	_ = json.Unmarshal(respBytes, &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestVpcService_AllocateElasticIP_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Vpc().AllocateElasticIP(AllocateElasticIPParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestVpcService_DeallocateElasticIP_happyPath(t *testing.T) {
	token := "token"
	var payload DeallocateElasticIPParams
	_ = faker.FakeData(&payload)
	payload.IPAddress = faker.IPv4()

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp BasicResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc(fmt.Sprintf("/elasticip/%s/deallocate", payload.IPAddress), func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	got, err := client.Vpc().DeallocateElasticIP(payload)

	var want BasicResponse
	_ = json.Unmarshal(respBytes, &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestVpcService_DeallocateElasticIP_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Vpc().DeallocateElasticIP(DeallocateElasticIPParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestVpcService_CreatePeeringConnection_happyPath(t *testing.T) {
	token := "token"
	var payload CreatePeeringConnectionParams
	_ = faker.FakeData(&payload)

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp CreateResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/vpc/peering", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	got, err := client.Vpc().CreatePeeringConnection(payload)

	var want CreateResponse
	_ = json.Unmarshal(respBytes, &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestVpcService_CreatePeeringConnection_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Vpc().CreatePeeringConnection(CreatePeeringConnectionParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestVpcService_UpdateRoute_happyPath(t *testing.T) {
	token := "token"
	var payload UpdateRouteParams
	_ = faker.FakeData(&payload)

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp BasicResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/vpc/route", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	got, err := client.Vpc().UpdateRoute(payload)

	var want BasicResponse
	_ = json.Unmarshal(respBytes, &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestVpcService_UpdateRoute_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Vpc().UpdateRoute(UpdateRouteParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestVpcService_CreateInternetGateway_happyPath(t *testing.T) {
	token := "token"
	var payload CreateInternetGatewayParams
	_ = faker.FakeData(&payload)

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp CreateResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/vpc-internetgateway", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	got, err := client.Vpc().CreateInternetGateway(payload)

	var want CreateResponse
	_ = json.Unmarshal(respBytes, &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestVpcService_CreateInternetGateway_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Vpc().CreateInternetGateway(CreateInternetGatewayParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestVpcService_AttachSubnetToInternetGateway_happyPath(t *testing.T) {
	token := "token"
	var payload AttachSubnetToInternetGatewayParams
	_ = faker.FakeData(&payload)
	payload.GatewayID = faker.UUIDDigit()

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp BasicResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc(fmt.Sprintf("/vpc-internetgateway/%s/attach", payload.GatewayID), func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	got, err := client.Vpc().AttachSubnetToInternetGateway(payload)

	var want BasicResponse
	_ = json.Unmarshal(respBytes, &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestVpcService_AttachSubnetToInternetGateway_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Vpc().AttachSubnetToInternetGateway(AttachSubnetToInternetGatewayParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestVpcService_DetachSubnetFromInternetGateway_happyPath(t *testing.T) {
	token := "token"
	var payload DetachSubnetFromInternetGatewayParams
	_ = faker.FakeData(&payload)
	payload.GatewayID = faker.UUIDDigit()

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp BasicResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc(fmt.Sprintf("/vpc-internetgateway/%s/dettach", payload.GatewayID), func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	got, err := client.Vpc().DetachSubnetFromInternetGateway(payload)

	var want BasicResponse
	_ = json.Unmarshal(respBytes, &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestVpcService_DetachSubnetFromInternetGateway_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Vpc().DetachSubnetFromInternetGateway(DetachSubnetFromInternetGatewayParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestVpcService_CreateRouteTable_happyPath(t *testing.T) {
	token := "token"
	var payload CreateRouteTableParams
	_ = faker.FakeData(&payload)

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp CreateResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/vpc-route-table", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	got, err := client.Vpc().CreateRouteTable(payload)

	var want CreateResponse
	_ = json.Unmarshal(respBytes, &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestVpcService_CreateRouteTable_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Vpc().CreateRouteTable(CreateRouteTableParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestVpcService_AssociateSubnetWithRouteTable_happyPath(t *testing.T) {
	token := "token"
	var payload AssociateSubnetWithRouteTableParams
	_ = faker.FakeData(&payload)

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp BasicResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/vpc/subnet/associate", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	got, err := client.Vpc().AssociateSubnetWithRouteTable(payload)

	var want BasicResponse
	_ = json.Unmarshal(respBytes, &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestVpcService_AssociateSubnetWithRouteTable_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Vpc().AssociateSubnetWithRouteTable(AssociateSubnetWithRouteTableParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestVpcService_DissociateSubnetFromRouteTable_happyPath(t *testing.T) {
	token := "token"
	var payload DissociateSubnetFromRouteTableParams
	_ = faker.FakeData(&payload)

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp BasicResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/vpc/subnet/dissociate", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	got, err := client.Vpc().DissociateSubnetFromRouteTable(payload)

	var want BasicResponse
	_ = json.Unmarshal(respBytes, &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestVpcService_DissociateSubnetFromRouteTable_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Vpc().DissociateSubnetFromRouteTable(DissociateSubnetFromRouteTableParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}
