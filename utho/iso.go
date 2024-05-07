package utho

import (
	"errors"
)

type ISOService service

type ISOs struct {
	ISOs    []ISO  `json:"isos"`
	Status  string `json:"status"`
	Message string `json:"message"`
}
type ISO struct {
	Name       string        `json:"name"`
	File       string        `json:"file"`
	Size       float64       `json:"size"`
	AddedAt    string        `json:"added_at"`
	Download   string        `json:"download"`
	Dc         string        `json:"dc"`
	Dclocation ISODclocation `json:"dclocation"`
}
type ISODclocation struct {
	Dccc     string `json:"dccc"`
	Location string `json:"location"`
}

type CreateISOParams struct {
	Dcslug string `json:"dcslug"`
	URL    string `json:"url"`
	Name   string `json:"name"`
}

func (s *ISOService) CreateISO(params CreateISOParams) (*CreateResponse, error) {
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

func (s *ISOService) ListISOs() ([]ISO, error) {
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

func (s *ISOService) DeleteISO(isoId string) (*DeleteResponse, error) {
	reqUrl := "iso/" + isoId + "/delete"
	req, _ := s.client.NewRequest("DELETE", reqUrl)

	var delResponse DeleteResponse
	_, err := s.client.Do(req, &delResponse)
	if err != nil {
		return nil, err
	}

	return &delResponse, nil
}
