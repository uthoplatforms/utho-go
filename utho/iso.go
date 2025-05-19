package utho

import (
	"errors"
)

type ISOService service

type ISOs struct {
	ISOs    []ISO  `json:"isos"`
	Status  string `json:"status" faker:"oneof: success, failure"`
	Message string `json:"message" faker:"sentence"`
}
type ISO struct {
	Name       string        `json:"name" faker:"word"`
	File       string        `json:"file" faker:"word"`
	Size       float64       `json:"size" faker:"boundary_start=1000, boundary_end=10000"`
	AddedAt    string        `json:"added_at" faker:"timestamp"`
	Download   string        `json:"download" faker:"oneof: 50, 100"`
	Dc         string        `json:"dc" faker:"word"`
	Dclocation ISODclocation `json:"dclocation"`
}
type ISODclocation struct {
	Dccc     string `json:"dccc" faker:"word"`
	Location string `json:"location" faker:"city"`
}

type CreateISOParams struct {
	Dcslug string `json:"dcslug" faker:"word"`
	URL    string `json:"url" faker:"url"`
	Name   string `json:"name" faker:"word"`
}

func (s *ISOService) Create(params CreateISOParams) (*CreateResponse, error) {
	reqUrl := "iso/add"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var iso CreateResponse
	_, err := s.client.Do(req, &iso)
	if err != nil {
		return nil, err
	}
	if iso.Status != "success" && iso.Status != "" {
		return nil, errors.New(iso.Message)
	}

	return &iso, nil
}

func (s *ISOService) List() ([]ISO, error) {
	reqUrl := "iso"
	req, _ := s.client.NewRequest("GET", reqUrl)

	var iso ISOs
	_, err := s.client.Do(req, &iso)
	if err != nil {
		return nil, err
	}
	if iso.Status != "success" && iso.Status != "" {
		return nil, errors.New(iso.Message)
	}

	return iso.ISOs, nil
}

func (s *ISOService) Delete(isoId string) (*DeleteResponse, error) {
	reqUrl := "iso/" + isoId + "/delete"
	req, _ := s.client.NewRequest("DELETE", reqUrl)

	var delResponse DeleteResponse
	_, err := s.client.Do(req, &delResponse)
	if err != nil {
		return nil, err
	}
	if delResponse.Status != "success" && delResponse.Status != "" {
		return nil, errors.New(delResponse.Message)
	}

	return &delResponse, nil
}
