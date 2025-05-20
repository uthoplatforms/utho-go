package utho

import (
	"errors"
	"fmt"
)

type VpcService service

type Vpcs struct {
	Vpc     []Vpc  `json:"vpc"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Vpc struct {
	ID         string         `json:"id"`
	Total      int            `json:"total"`
	Available  int            `json:"available"`
	Network    string         `json:"network"`
	Name       string         `json:"name"`
	Size       string         `json:"size"`
	Dcslug     string         `json:"dcslug"`
	Dclocation VpcDclocation  `json:"dclocation"`
	IsDefault  string         `json:"is_default"`
	Resources  []VpcResources `json:"resources"`
	Status     string         `json:"status"`
	Message    string         `json:"message"`
}
type VpcDclocation struct {
	Dccc     string `json:"dccc"`
	Location string `json:"location"`
}
type VpcResources struct {
	Type string `json:"type"`
	ID   string `json:"id"`
	Name string `json:"name"`
	IP   string `json:"ip"`
}

type CreateVpcParams struct {
	Dcslug  string `json:"dcslug"`
	Name    string `json:"name"`
	Planid  string `json:"planid"`
	Network string `json:"network"`
	Size    string `json:"size"`
}

func (s *VpcService) Create(params CreateVpcParams) (*CreateResponse, error) {
	reqUrl := "vpc/create"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var vpc CreateResponse
	_, err := s.client.Do(req, &vpc)
	if err != nil {
		return nil, err
	}
	if vpc.Status != "success" && vpc.Status != "" {
		return nil, errors.New(vpc.Message)
	}

	return &vpc, nil
}

func (s *VpcService) Read(vpcId string) (*Vpc, error) {
	reqUrl := "vpc"
	req, _ := s.client.NewRequest("GET", reqUrl)

	var vpcs Vpcs
	_, err := s.client.Do(req, &vpcs)
	if err != nil {
		return nil, err
	}
	if vpcs.Status != "success" && vpcs.Status != "" {
		return nil, errors.New(vpcs.Message)
	}

	var vpc Vpc
	found := false
	for _, r := range vpcs.Vpc {
		if r.ID == vpcId {
			vpc = r
			found = true
			break
		}
	}
	if !found {
		return nil, errors.New("NotFound")
	}

	return &vpc, nil
}

func (s *VpcService) List() ([]Vpc, error) {
	reqUrl := "vpc"
	req, _ := s.client.NewRequest("GET", reqUrl)

	var vpc Vpcs
	_, err := s.client.Do(req, &vpc)
	if err != nil {
		return nil, err
	}
	if vpc.Status != "success" && vpc.Status != "" {
		return nil, errors.New(vpc.Message)
	}

	return vpc.Vpc, nil
}

func (s *VpcService) Delete(vpcId string) (*DeleteResponse, error) {
	reqUrl := "vpc/" + vpcId + "/destroy"
	req, _ := s.client.NewRequest("DELETE", reqUrl)

	var delResponse DeleteResponse
	if _, err := s.client.Do(req, &delResponse); err != nil {
		return nil, err
	}
	if delResponse.Status != "success" && delResponse.Status != "" {
		return nil, errors.New(delResponse.Message)
	}

	return &delResponse, nil
}

type CreateSubnetParams struct {
	Name         string `json:"name"`
	VpcID        string `json:"vpcid"`
	AssignPublic string `json:"assign_public"`
	IsDefault    string `json:"is_default"`
	Type         string `json:"type"`
	Network      string `json:"network"`
	Size         string `json:"size"`
}

func (s *VpcService) CreateSubnet(params CreateSubnetParams) (*CreateResponse, error) {
	reqUrl := "vpc/subnet/create"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var subnet CreateResponse
	_, err := s.client.Do(req, &subnet)
	if err != nil {
		return nil, err
	}
	if subnet.Status != "success" && subnet.Status != "" {
		return nil, errors.New(subnet.Message)
	}

	return &subnet, nil
}

type CreateNatGatewayParams struct {
	Subnet string `json:"subnet"`
	Name   string `json:"name"`
	UserID string `json:"userid"`
}

func (s *VpcService) CreateNatGateway(params CreateNatGatewayParams) (*CreateResponse, error) {
	reqUrl := "vpc/natgateway"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var natGateway CreateResponse
	_, err := s.client.Do(req, &natGateway)
	if err != nil {
		return nil, err
	}
	if natGateway.Status != "success" && natGateway.Status != "" {
		return nil, errors.New(natGateway.Message)
	}

	return &natGateway, nil
}

type AllocateElasticIPParams struct {
	Dcslug string `json:"dcslug"`
}

func (s *VpcService) AllocateElasticIP(params AllocateElasticIPParams) (*CreateResponse, error) {
	reqUrl := "elasticip/allocate"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var elasticIP CreateResponse
	_, err := s.client.Do(req, &elasticIP)
	if err != nil {
		return nil, err
	}
	if elasticIP.Status != "success" && elasticIP.Status != "" {
		return nil, errors.New(elasticIP.Message)
	}

	return &elasticIP, nil
}

type DeallocateElasticIPParams struct {
	IPAddress string
	Subnet    string `json:"subnet"`
	Name      string `json:"name"`
	UserID    string `json:"userid"`
}

func (s *VpcService) DeallocateElasticIP(params DeallocateElasticIPParams) (*BasicResponse, error) {
	reqUrl := "elasticip/" + params.IPAddress + "/deallocate"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var basicResponse BasicResponse
	_, err := s.client.Do(req, &basicResponse)
	if err != nil {
		return nil, err
	}
	if basicResponse.Status != "success" && basicResponse.Status != "" {
		return nil, errors.New(basicResponse.Message)
	}

	return &basicResponse, nil
}

type CreatePeeringConnectionParams struct {
	Name           string `json:"name"`
	Dcslug         string `json:"dcslug"`
	RequesterVpcID string `json:"requester_vpc_id"`
	AccepterVpcID  string `json:"accepter_vpc_id"`
}

func (s *VpcService) CreatePeeringConnection(params CreatePeeringConnectionParams) (*CreateResponse, error) {
	reqUrl := "vpc/peering"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var peering CreateResponse
	_, err := s.client.Do(req, &peering)
	if err != nil {
		return nil, err
	}
	if peering.Status != "success" && peering.Status != "" {
		return nil, errors.New(peering.Message)
	}

	return &peering, nil
}

type UpdateRouteParams struct {
	RouteID              string `json:"route_id"`
	DestinationCidrBlock string `json:"destination_cidr_block"`
	Target               string `json:"target"`
	RouteType            string `json:"route_type"`
}

func (s *VpcService) UpdateRoute(params UpdateRouteParams) (*BasicResponse, error) {
	reqUrl := "vpc/route"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var basicResponse BasicResponse
	_, err := s.client.Do(req, &basicResponse)
	if err != nil {
		return nil, err
	}
	if basicResponse.Status != "success" && basicResponse.Status != "" {
		return nil, errors.New(basicResponse.Message)
	}

	return &basicResponse, nil
}

type CreateInternetGatewayParams struct {
	Subnet string `json:"subnet"`
	Name   string `json:"name"`
	UserID string `json:"userid"`
}

func (s *VpcService) CreateInternetGateway(params CreateInternetGatewayParams) (*CreateResponse, error) {
	reqUrl := "vpc-internetgateway"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var internetGateway CreateResponse
	_, err := s.client.Do(req, &internetGateway)
	if err != nil {
		return nil, err
	}
	if internetGateway.Status != "success" && internetGateway.Status != "" {
		return nil, errors.New(internetGateway.Message)
	}

	return &internetGateway, nil
}

type AttachSubnetToInternetGatewayParams struct {
	GatewayID string
	Subnet    string `json:"subnet"`
}

func (s *VpcService) AttachSubnetToInternetGateway(params AttachSubnetToInternetGatewayParams) (*BasicResponse, error) {
	reqUrl := fmt.Sprintf("vpc-internetgateway/%s/attach", params.GatewayID)
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var basicResponse BasicResponse
	_, err := s.client.Do(req, &basicResponse)
	if err != nil {
		return nil, err
	}
	if basicResponse.Status != "success" && basicResponse.Status != "" {
		return nil, errors.New(basicResponse.Message)
	}

	return &basicResponse, nil
}

type DetachSubnetFromInternetGatewayParams struct {
	GatewayID string
	Subnet    string `json:"subnet"`
}

func (s *VpcService) DetachSubnetFromInternetGateway(params DetachSubnetFromInternetGatewayParams) (*BasicResponse, error) {
	reqUrl := fmt.Sprintf("vpc-internetgateway/%s/dettach", params.GatewayID)
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var basicResponse BasicResponse
	_, err := s.client.Do(req, &basicResponse)
	if err != nil {
		return nil, err
	}
	if basicResponse.Status != "success" && basicResponse.Status != "" {
		return nil, errors.New(basicResponse.Message)
	}

	return &basicResponse, nil
}

type CreateRouteTableParams struct {
	VpcID string `json:"vpc_id"`
}

func (s *VpcService) CreateRouteTable(params CreateRouteTableParams) (*CreateResponse, error) {
	reqUrl := "vpc-route-table"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var routeTable CreateResponse
	_, err := s.client.Do(req, &routeTable)
	if err != nil {
		return nil, err
	}
	if routeTable.Status != "success" && routeTable.Status != "" {
		return nil, errors.New(routeTable.Message)
	}

	return &routeTable, nil
}

type AssociateSubnetWithRouteTableParams struct {
	SubnetID     string `json:"subnet_id"`
	RouteTableID string `json:"route_table_id"`
}

func (s *VpcService) AssociateSubnetWithRouteTable(params AssociateSubnetWithRouteTableParams) (*BasicResponse, error) {
	reqUrl := "vpc/subnet/associate"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var basicResponse BasicResponse
	_, err := s.client.Do(req, &basicResponse)
	if err != nil {
		return nil, err
	}
	if basicResponse.Status != "success" && basicResponse.Status != "" {
		return nil, errors.New(basicResponse.Message)
	}

	return &basicResponse, nil
}

type DissociateSubnetFromRouteTableParams struct {
	SubnetID     string `json:"subnet_id"`
	RouteTableID string `json:"route_table_id"`
}

func (s *VpcService) DissociateSubnetFromRouteTable(params DissociateSubnetFromRouteTableParams) (*BasicResponse, error) {
	reqUrl := "vpc/subnet/dissociate"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var basicResponse BasicResponse
	_, err := s.client.Do(req, &basicResponse)
	if err != nil {
		return nil, err
	}
	if basicResponse.Status != "success" && basicResponse.Status != "" {
		return nil, errors.New(basicResponse.Message)
	}

	return &basicResponse, nil
}
