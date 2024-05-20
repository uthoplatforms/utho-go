package utho

import (
	"errors"
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
	for _, r := range vpcs.Vpc {
		if r.ID == vpcId {
			vpc = r
		}
	}
	if len(vpc.ID) == 0 {
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
