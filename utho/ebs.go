package utho

import (
	"errors"
)

type EBService service

type EBSs struct {
	Ebs     []Ebs  `json:"ebs"`
	Status  string `json:"status" faker:"oneof: success, failure"`
	Message string `json:"message" faker:"sentence"`
}

type Ebs struct {
	ID         string   `json:"did" faker:"uuid_digit"`
	Cloudid    string   `json:"cloudid" faker:"uuid_digit"`
	Primaryd   string   `json:"primaryd" faker:"oneof: 0, 1"`
	Size       string   `json:"size" faker:"oneof: 10.000, 20.000, 30.000"`
	Status     string   `json:"status" faker:"oneof: Active, Inactive"`
	Extrabill  string   `json:"extrabill" faker:"oneof: 0, 1"`
	CreatedAt  string   `json:"created_at" faker:"timestamp"`
	DeletedAt  string   `json:"deleted_at" faker:"timestamp"`
	Ebs        string   `json:"ebs" faker:"oneof: 0, 1"`
	Name       string   `json:"name" faker:"name"`
	Iops       string   `json:"iops" faker:"oneof: 1000, 2000, 3000"`
	Throughput string   `json:"throughput" faker:"oneof: 125, 250, 500"`
	Location   Location `json:"location"`
}

type Location struct {
	City    string `json:"city" faker:"city"`
	Country string `json:"country" faker:"country"`
	Dc      string `json:"dc" faker:"word"`
	Dccc    string `json:"dccc" faker:"word"`
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
	req, _ := s.client.NewRequest("PUT", reqUrl, &params)

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
	req, _ := s.client.NewRequest("PUT", reqUrl, &params)

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

type ResizeEBSParams struct {
	Disk       string `json:"disk"`
	Iops       string `json:"iops"`
	Throughput string `json:"throughput"`
}

func (s *EBService) Resize(ebsId string, params ResizeEBSParams) (*UpdateResponse, error) {
	reqUrl := "ebs/" + ebsId + "/resize"
	req, _ := s.client.NewRequest("PUT", reqUrl, &params)

	var updateResponse UpdateResponse
	_, err := s.client.Do(req, &updateResponse)
	if err != nil {
		return nil, err
	}
	if updateResponse.Status != "success" && updateResponse.Status != "" {
		return nil, errors.New(updateResponse.Message)
	}

	return &updateResponse, nil
}
