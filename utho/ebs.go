package utho

import (
	"errors"
)

type EBService service

type EBSs struct {
	Ebs     []Ebs  `json:"ebs"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Ebs struct {
	ID         string   `json:"did"`
	Cloudid    string   `json:"cloudid"`
	Primaryd   string   `json:"primaryd"`
	Size       string   `json:"size"`
	Status     string   `json:"status"`
	Extrabill  string   `json:"extrabill"`
	CreatedAt  string   `json:"created_at"`
	DeletedAt  string   `json:"deleted_at"`
	Ebs        string   `json:"ebs"`
	Name       string   `json:"name"`
	Iops       string   `json:"iops"`
	Throughput string   `json:"throughput"`
	Location   Location `json:"location"`
}

type Location struct {
	City    string `json:"city"`
	Country string `json:"country"`
	Dc      string `json:"dc"`
	Dccc    string `json:"dccc"`
}

type CreateEBSParams struct {
	Name       string `json:"name"`
	Dcslug     string `json:"dcslug"`
	Disk       string `json:"disk"`
	Iops       string `json:"iops"`
	Throughput string `json:"throughput"`
	DiskType   string `json:"disk_type"`
}

func (s *EBService) Create(params CreateEBSParams) (*CreateResponse, error) {
	reqUrl := "ebs"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var ebs CreateResponse
	_, err := s.client.Do(req, &ebs)
	if err != nil {
		return nil, err
	}
	if ebs.Status != "success" && ebs.Status != "" {
		return nil, errors.New(ebs.Message)
	}

	return &ebs, nil
}

func (s *EBService) Read(ebsId string) (*Ebs, error) {
	reqUrl := "ebs/" + ebsId
	req, _ := s.client.NewRequest("GET", reqUrl)

	var ebs EBSs
	_, err := s.client.Do(req, &ebs)
	if err != nil {
		return nil, err
	}
	if ebs.Status != "success" && ebs.Status != "" {
		return nil, errors.New(ebs.Message)
	}
	if len(ebs.Ebs) == 0 {
		return nil, errors.New("NotFound")
	}

	return &ebs.Ebs[0], nil
}

func (s *EBService) List() ([]Ebs, error) {
	reqUrl := "ebs"
	req, _ := s.client.NewRequest("GET", reqUrl)

	var ebs EBSs
	_, err := s.client.Do(req, &ebs)
	if err != nil {
		return nil, err
	}
	if ebs.Status != "success" && ebs.Status != "" {
		return nil, errors.New(ebs.Message)
	}

	return ebs.Ebs, nil
}

func (s *EBService) Delete(ebsId string) (*DeleteResponse, error) {
	reqUrl := "ebs/" + ebsId + "/destroy"
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

type AttachEBSParams struct {
	EBSId      string
	ResourceId string `json:"resourceid"`
	Type       string `json:"type"`
}

func (s *EBService) Attach(params AttachEBSParams) (*CreateResponse, error) {
	reqUrl := "ebs/" + params.EBSId + "/attach"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var ebs CreateResponse
	_, err := s.client.Do(req, &ebs)
	if err != nil {
		return nil, err
	}
	if ebs.Status != "success" && ebs.Status != "" {
		return nil, errors.New(ebs.Message)
	}

	return &ebs, nil
}

func (s *EBService) Dettach(params AttachEBSParams) (*CreateResponse, error) {
	reqUrl := "ebs/" + params.EBSId + "/dettach"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var ebs CreateResponse
	_, err := s.client.Do(req, &ebs)
	if err != nil {
		return nil, err
	}
	if ebs.Status != "success" && ebs.Status != "" {
		return nil, errors.New(ebs.Message)
	}

	return &ebs, nil
}
