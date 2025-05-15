package utho

import (
	"errors"
)

type DomainService service

type DnsDomains struct {
	Domains []Domain `json:"domains" faker:"v4slice"`
	Status  string   `json:"status,omitempty" faker:"status"`
	Message string   `json:"message,omitempty" faker:"message"`
}
type Domain struct {
	Domain         string      `json:"domain" faker:"domain_name"`
	Status         string      `json:"status" faker:"status"`
	Message        string      `json:"message" faker:"message"`
	Nspoint        string      `json:"nspoint" faker:"word"`
	CreatedAt      string      `json:"created_at" faker:"date_time"`
	DnsrecordCount string      `json:"dnsrecord_count" faker:"digit"`
	Records        []DnsRecord `json:"records" faker:"v4slice"`
}
type DnsRecord struct {
	ID       string `json:"id" faker:"uuid_digit"`
	Hostname string `json:"hostname" faker:"domain_name"`
	Type     string `json:"type" faker:"dns_type"`
	Value    string `json:"value" faker:"ipv4"`
	TTL      string `json:"ttl" faker:"digit"`
	Priority string `json:"priority" faker:"digit"`
	Porttype string `json:"porttype"`
	Port     string `json:"port"`
	Weight   string `json:"weight"`
}

type CreateDomainParams struct {
	Domain string `json:"domain" faker:"domain_name"`
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
	if len(domain.Domains) == 0 {
		return nil, errors.New("domain not found")
	}

	return &domain.Domains[0], nil
}

func (s *DomainService) ListDomains() ([]Domain, error) {
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

	return domain.Domains, nil
}

func (s *DomainService) DeleteDomain(domainName string) (*DeleteResponse, error) {
	reqUrl := "dns/" + domainName + "/delete"
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

type CreateDnsRecordParams struct {
	Domain   string
	Type     string `json:"type"`
	Hostname string `json:"hostname"`
	Value    string `json:"value"`
	TTL      string `json:"ttl"`
	Porttype string `json:"porttype"`
	Port     string `json:"port"`
	Priority string `json:"priority"`
	Weight   string `json:"weight"`
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

func (s *DomainService) ReadDnsRecord(domainName, dnsRecordID string) (*DnsRecord, error) {
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
	if len(domain.Domains) == 0 {
		return nil, errors.New("domain not found")
	}

	var record DnsRecord
	found := false
	for _, dnsRecord := range domain.Domains[0].Records {
		if dnsRecord.ID == dnsRecordID {
			record = dnsRecord
			found = true
			break
		}
	}

	if !found {
		return nil, errors.New("dns record not found")
	}

	return &record, nil
}

func (s *DomainService) ListDnsRecords(domainName string) ([]DnsRecord, error) {
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
	if len(domain.Domains) == 0 {
		return []DnsRecord{}, nil
	}

	return domain.Domains[0].Records, nil
}

func (s *DomainService) DeleteDnsRecord(domainName, recordId string) (*DeleteResponse, error) {
	reqUrl := "dns/" + domainName + "/record/" + recordId + "/delete"
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
