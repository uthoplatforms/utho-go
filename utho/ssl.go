package utho

import (
	"errors"
)

type SslService service

type Ssls struct {
	Certificates []Certificates `json:"certificates"`
	Status       string         `json:"status"`
	Message      string         `json:"message"`
}
type Certificates struct {
	ID               string `json:"id"`
	Userid           string `json:"userid"`
	Name             string `json:"name"`
	PrivateKey       string `json:"private_key"`
	CertificateKey   string `json:"certificate_key"`
	CertificateChain string `json:"certificate_chain"`
	Type             string `json:"type"`
	State            string `json:"state"`
	DNSNames         string `json:"dns_names"`
	CreatedAt        string `json:"created_at"`
	ExpireAt         string `json:"expire_at"`
	Sha1Fingerprint  string `json:"sha1_fingerprint"`
	Issuer           string `json:"issuer"`
	IsDeleted        string `json:"is_deleted"`
	DeletedAt        string `json:"deleted_at"`
}

type CreateSslParams struct {
	Name             string `json:"name"`
	Type             string `json:"type"`
	CertificateKey   string `json:"certificate_key"`
	PrivateKey       string `json:"private_key"`
	CertificateChain string `json:"certificateChain"`
}

func (s *SslService) Create(params CreateSslParams) (*CreateResponse, error) {
	reqUrl := "certificates"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var ssl CreateResponse
	_, err := s.client.Do(req, &ssl)
	if err != nil {
		return nil, err
	}
	if ssl.Status != "success" && ssl.Status != "" {
		return nil, errors.New(ssl.Message)
	}

	return &ssl, nil
}

func (s *SslService) Read(certId string) (*Certificates, error) {
	reqUrl := "certificates"
	req, _ := s.client.NewRequest("GET", reqUrl)

	var ssl Ssls
	_, err := s.client.Do(req, &ssl)
	if err != nil {
		return nil, err
	}
	if ssl.Status != "success" && ssl.Status != "" {
		return nil, errors.New(ssl.Message)
	}

	var cert Certificates
	for _, r := range ssl.Certificates {
		if r.ID == certId {
			cert = r
		}
	}
	if len(cert.ID) == 0 {
		return nil, errors.New("certificate not found")
	}

	return &cert, nil
}

func (s *SslService) List() ([]Certificates, error) {
	reqUrl := "certificates"
	req, _ := s.client.NewRequest("GET", reqUrl)

	var ssl Ssls
	_, err := s.client.Do(req, &ssl)
	if err != nil {
		return nil, err
	}
	if ssl.Status != "success" && ssl.Status != "" {
		return nil, errors.New(ssl.Message)
	}

	return ssl.Certificates, nil
}

func (s *SslService) Delete(certId string) (*DeleteResponse, error) {
	reqUrl := "certificates/" + certId
	req, _ := s.client.NewRequest("DELETE", reqUrl)

	var delResponse DeleteResponse
	if _, err := s.client.Do(req, &delResponse); err != nil {
		return nil, err
	}

	return &delResponse, nil
}
