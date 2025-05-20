package utho

import (
	"errors"
)

type SqsService service

type Sqss struct {
	Sqs     []Sqs  `json:"sqs"`
	Status  string `json:"status"`
	Message string `json:"message"`
}
type Sqs struct {
	ID        string `json:"id"`
	Userid    string `json:"userid"`
	Cloudid   string `json:"cloudid"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	IP        string `json:"ip"`
	Count     string `json:"count"`
}

type CreateSqsParams struct {
	Dcslug string `json:"dcslug"`
	Planid string `json:"planid"`
	Name   string `json:"name"`
}

func (s *SqsService) Create(params CreateSqsParams) (*CreateResponse, error) {
	reqUrl := "sqs"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var sqs CreateResponse
	_, err := s.client.Do(req, &sqs)
	if err != nil {
		return nil, err
	}
	if sqs.Status != "success" && sqs.Status != "" {
		return nil, errors.New(sqs.Message)
	}

	return &sqs, nil
}

func (s *SqsService) Read(sqsId string) (*Sqs, error) {
	reqUrl := "sqs/" + sqsId
	req, _ := s.client.NewRequest("GET", reqUrl)

	var sqs Sqss
	_, err := s.client.Do(req, &sqs)
	if err != nil {
		return nil, err
	}
	if sqs.Status != "success" && sqs.Status != "" {
		return nil, errors.New(sqs.Message)
	}
	if len(sqs.Sqs) == 0 {
		return nil, errors.New("NotFound")
	}

	return &sqs.Sqs[0], nil
}

func (s *SqsService) List() ([]Sqs, error) {
	reqUrl := "sqs"
	req, _ := s.client.NewRequest("GET", reqUrl)

	var sqs Sqss
	_, err := s.client.Do(req, &sqs)
	if err != nil {
		return nil, err
	}
	if sqs.Status != "success" && sqs.Status != "" {
		return nil, errors.New(sqs.Message)
	}

	return sqs.Sqs, nil
}

func (s *SqsService) Delete(sqsId, sqsName string) (*DeleteResponse, error) {
	reqUrl := "sqs/" + sqsId + "/destroy?confirm=" + sqsName
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

type CreateQueueParams struct {
	SqsID                  string
	Name                   string `json:"name"`
	FifoQueue              string `json:"FifoQueue"`
	VisibilityTimeout      string `json:"VisibilityTimeout"`
	MessageRetentionPeriod string `json:"MessageRetentionPeriod"`
	MaximumMessageSize     string `json:"maximumMessageSize"`
}

func (s *SqsService) CreateQueue(params CreateQueueParams) (*CreateResponse, error) {
	reqUrl := "sqs/" + params.SqsID + "/queue"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var queue CreateResponse
	_, err := s.client.Do(req, &queue)
	if err != nil {
		return nil, err
	}
	if queue.Status != "success" && queue.Status != "" {
		return nil, errors.New(queue.Message)
	}

	return &queue, nil
}
