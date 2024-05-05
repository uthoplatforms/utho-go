package utho

import (
	"errors"
)

type DomainService service

type DnsDomains struct {
	Domains []Domain `json:"domains"`
	Status  string   `json:"status,omitempty"`
	Message string   `json:"message,omitempty"`
}
type Domain struct {
	Domain         string   `json:"domain"`
	Status         string   `json:"status"`
	Message        string   `json:"message"`
	Nspoint        string   `json:"nspoint"`
	CreatedAt      string   `json:"created_at"`
	DnsrecordCount string   `json:"dnsrecord_count"`
	Records        []Record `json:"records"`
}
type Record struct {
	ID       string `json:"id"`
	Hostname string `json:"hostname"`
	Type     string `json:"type"`
	Value    string `json:"value"`
	TTL      string `json:"ttl"`
	Priority string `json:"priority"`
}

type CreateDomainParams struct {
	Domain string `json:"domain"`
}

func (s *DomainService) CreateDomain(params CreateDomainParams) (*BasicResponse, error) {
	reqUrl := "dns/adddomain"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var domain BasicResponse
	_, err := s.client.Do(req, &domain)
	if err != nil {
		return nil, err
	}
	if domain.Status != "success" && domain.Status != "" {
		return nil, errors.New(domain.Message)
	}

	return &domain, nil
}

func (s *DomainService) ReadDomain(domainName string) (*Domain, error) {
	reqUrl := "dns/" + domainName
	req, _ := s.client.NewRequest("GET", reqUrl)

	var domain DnsDomains
	_, err := s.client.Do(req, &domain)
	if err != nil {
		return nil, err
	}
	if domain.Status != "success" && domain.Status != "" {
		return nil, errors.New(domain.Message)
	}

	return &domain.Domains[0], nil
}

func (s *DomainService) ListDomains() (*[]Domain, error) {
	reqUrl := "dns"
	req, _ := s.client.NewRequest("GET", reqUrl)

	var domain DnsDomains
	_, err := s.client.Do(req, &domain)
	if err != nil {
		return nil, err
	}
	if domain.Status != "success" && domain.Status != "" {
		return nil, errors.New(domain.Message)
	}

	return &domain.Domains, nil
}

func (s *DomainService) DeleteDomain(domainId string) (*DeleteResponse, error) {
	reqUrl := "dns/" + domainId + "/delete"
	req, _ := s.client.NewRequest("DELETE", reqUrl)

	var delResponse DeleteResponse
	if _, err := s.client.Do(req, &delResponse); err != nil {
		return nil, err
	}

	return &delResponse, nil
}

type CreateDnsRecordParams struct {
	Domain   string
	Type     string `json:"type"`
	Hostname string `json:"hostname"`
	Value    string `json:"value"`
	TTL      string `json:"ttl"`
	Porttype string `json:"porttype"`
	Port     string `json:"port"`
	Priority string `json:"priority"`
	Wight    string `json:"wight"`
}

func (s *DomainService) CreateDnsRecord(params CreateDnsRecordParams) (*CreateResponse, error) {
	reqUrl := "dns/" + params.Domain + "/record/add"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var dnsRecord CreateResponse
	_, err := s.client.Do(req, &dnsRecord)
	if err != nil {
		return nil, err
	}
	if dnsRecord.Status != "success" && dnsRecord.Status != "" {
		return nil, errors.New(dnsRecord.Message)
	}

	return &dnsRecord, nil
}

func (s *DomainService) ReadDnsRecord(domainName, dnsRecordID string) (*Record, error) {
	reqUrl := "dns/" + domainName
	req, _ := s.client.NewRequest("GET", reqUrl)

	var domain DnsDomains
	_, err := s.client.Do(req, &domain)
	if err != nil {
		return nil, err
	}
	if domain.Status != "success" && domain.Status != "" {
		return nil, errors.New(domain.Message)
	}

	var record Record
	for _, dnsRecord := range domain.Domains[0].Records {
		if dnsRecord.ID == dnsRecordID {
			record = dnsRecord
		}

	}

	return &record, nil
}

func (s *DomainService) ListDnsRecords(domainName string) (*[]Record, error) {
	reqUrl := "dns/" + domainName
	req, _ := s.client.NewRequest("GET", reqUrl)

	var domain DnsDomains
	_, err := s.client.Do(req, &domain)
	if err != nil {
		return nil, err
	}
	if domain.Status != "success" && domain.Status != "" {
		return nil, errors.New(domain.Message)
	}

	return &domain.Domains[0].Records, nil
}

func (s *DomainService) DeleteDnsRecord(domainId, recordId string) (*DeleteResponse, error) {
	reqUrl := "dns/" + domainId + "/record/" + recordId + "/delete"
	req, _ := s.client.NewRequest("DELETE", reqUrl)

	var delResponse DeleteResponse
	if _, err := s.client.Do(req, &delResponse); err != nil {
		return nil, err
	}

	return &delResponse, nil
}
