package utho

import (
	"errors"
)

type VpcService service

type Vpcs struct {
	Vpcs    []Vpc  `json:"vpcs"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Vpc struct {
	Total      int            `json:"total"`
	Available  int            `json:"available"`
	Network    string         `json:"network"`
	Name       string         `json:"name"`
	Size       string         `json:"size"`
	Dcslug     string         `json:"dcslug"`
	Dclocation VpcDclocation  `json:"dclocation"`
	IsDefault  any            `json:"is_default"`
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

func (s *VpcService) CreateVpc(params CreateVpcParams) (*CreateResponse, error) {
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

func (s *VpcService) ReadVpc(vpcId string) (*Vpc, error) {
	reqUrl := "vpc"
	req, _ := s.client.NewRequest("GET", reqUrl)

	var vpc Vpc
	_, err := s.client.Do(req, &vpc)
	if err != nil {
		return nil, err
	}
	if vpc.Status != "success" && vpc.Status != "" {
		return nil, errors.New(vpc.Message)
	}

	return &vpc, nil
}

// func (s *VpcService) ListVpcs() (*[]Vpc, error) {
// 	reqUrl := "vpc"
// 	req, _ := s.client.NewRequest("GET", reqUrl)

// 	var vpc Vpcs
// 	_, err := s.client.Do(req, &vpc)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if vpc.Status != "success" && vpc.Status != "" {
// 		return nil, errors.New(vpc.Message)
// 	}

// 	return &vpc.Vpcs, nil
// }

func (s *VpcService) DeleteVpc(vpcId string) (*DeleteResponse, error) {
	reqUrl := "vpc/" + vpcId + "/destroy"
	req, _ := s.client.NewRequest("DELETE", reqUrl)

	var delResponse DeleteResponse
	if _, err := s.client.Do(req, &delResponse); err != nil {
		return nil, err
	}

	return &delResponse, nil
}
