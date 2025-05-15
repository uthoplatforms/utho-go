package utho

import (
	"errors"
)

type FirewallService service

type Firewalls struct {
	Firewalls []Firewall `json:"firewalls" faker:"-"`
	Status    string     `json:"status" faker:"oneof: success, failure"`
	Message   string     `json:"message" faker:"sentence"`
}

type Firewall struct {
	ID           string         `json:"id" faker:"uuid_digit"`
	Name         string         `json:"name" faker:"name"`
	CreatedAt    string         `json:"created_at" faker:"timestamp"`
	Rulecount    string         `json:"rulecount" faker:"oneof: 0, 5, 10"`
	Serverscount string         `json:"serverscount" faker:"oneof: 0, 1, 50"`
	Rules        []FirewallRule `json:"rules" faker:"-"`
}
type FirewallRule struct {
	ID         string `json:"id" faker:"uuid_digit"`
	Firewallid string `json:"firewallid" faker:"uuid_digit"`
	Type       string `json:"type" faker:"oneof: incoming, outgoing"`
	Service    string `json:"service" faker:"word"`
	Protocol   string `json:"protocol" faker:"oneof: TCP, UDP, ICMP"`
	Port       string `json:"port" faker:"oneof: 22, 80, 443"`
	Addresses  string `json:"addresses" faker:"ipv4"`
}

type CreateFirewallParams struct {
	Name string `json:"name" faker:"name"`
}
type CreateFirewallResponse struct {
	ID      string `json:"firewallid" faker:"uuid_digit"`
	Status  string `json:"status" faker:"oneof: success, failure"`
	Message string `json:"message" faker:"sentence"`
}

func (s *FirewallService) Create(params CreateFirewallParams) (*CreateFirewallResponse, error) {
	reqUrl := "firewall/create"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var firewall CreateFirewallResponse
	_, err := s.client.Do(req, &firewall)
	if err != nil {
		return nil, err
	}
	if firewall.Status != "success" && firewall.Status != "" {
		return nil, errors.New(firewall.Message)
	}

	return &firewall, nil
}

func (s *FirewallService) Read(firewallId string) (*Firewall, error) {
	reqUrl := "firewall/" + firewallId
	req, _ := s.client.NewRequest("GET", reqUrl)

	var firewall Firewalls
	_, err := s.client.Do(req, &firewall)
	if err != nil {
		return nil, err
	}
	if firewall.Status != "success" && firewall.Status != "" {
		return nil, errors.New(firewall.Message)
	}
	if len(firewall.Firewalls) == 0 {
		return nil, errors.New("NotFound")
	}

	return &firewall.Firewalls[0], nil
}

func (s *FirewallService) List() ([]Firewall, error) {
	reqUrl := "firewall"
	req, _ := s.client.NewRequest("GET", reqUrl)

	var firewall Firewalls
	_, err := s.client.Do(req, &firewall)
	if err != nil {
		return nil, err
	}
	if firewall.Status != "success" && firewall.Status != "" {
		return nil, errors.New(firewall.Message)
	}

	return firewall.Firewalls, nil
}

func (s *FirewallService) Delete(firewallId string) (*DeleteResponse, error) {
	reqUrl := "firewall/" + firewallId + "/destroy"
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

type CreateFirewallRuleParams struct {
	FirewallId string
	Type       string `json:"type"`
	Service    string `json:"service"`
	Protocol   string `json:"protocol"`
	Port       string `json:"port"`
	Addresses  string `json:"addresses"`
}

func (s *FirewallService) CreateFirewallRule(params CreateFirewallRuleParams) (*CreateResponse, error) {
	reqUrl := "firewall/" + params.FirewallId + "/rule/add"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var firewallRule CreateResponse
	_, err := s.client.Do(req, &firewallRule)
	if err != nil {
		return nil, err
	}
	if firewallRule.Status != "success" && firewallRule.Status != "" {
		return nil, errors.New(firewallRule.Message)
	}

	return &firewallRule, nil
}

func (s *FirewallService) ReadFirewallRule(firewallId, firewallRuleId string) (*FirewallRule, error) {
	reqUrl := "firewall/" + firewallId
	req, _ := s.client.NewRequest("GET", reqUrl)

	var firewall Firewalls
	_, err := s.client.Do(req, &firewall)
	if err != nil {
		return nil, err
	}
	if firewall.Status != "success" && firewall.Status != "" {
		return nil, errors.New(firewall.Message)
	}
	if len(firewall.Firewalls) == 0 {
		return nil, errors.New("NotFound")
	}

	var rule FirewallRule
	for _, r := range firewall.Firewalls[0].Rules {
		if r.ID == firewallRuleId {
			rule = r
		}
	}
	if len(rule.ID) == 0 {
		return nil, errors.New("firewall rule not found")
	}

	return &rule, nil
}

func (s *FirewallService) ListFirewallRules(firewallId string) ([]FirewallRule, error) {
	reqUrl := "firewall/" + firewallId
	req, _ := s.client.NewRequest("GET", reqUrl)

	var firewall Firewalls
	_, err := s.client.Do(req, &firewall)
	if err != nil {
		return nil, err
	}
	if firewall.Status != "success" && firewall.Status != "" {
		return nil, errors.New(firewall.Message)
	}
	if len(firewall.Firewalls) == 0 {
		return []FirewallRule{}, nil
	}

	return firewall.Firewalls[0].Rules, nil
}

func (s *FirewallService) DeleteFirewallRule(firewallId, firewallRuleId string) (*DeleteResponse, error) {
	reqUrl := "firewall/" + firewallId + "/rule/" + firewallRuleId + "/delete"
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

type AddCloudInsanceToFirewallParams struct {
	FirewallId string
	Cloudid    string `json:"cloudid"`
}

func (s *FirewallService) AddCloudInsanceToFirewall(params AddCloudInsanceToFirewallParams) (*CreateResponse, error) {
	reqUrl := "firewall/" + params.FirewallId + "/server/add"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var firewallRule CreateResponse
	_, err := s.client.Do(req, &firewallRule)
	if err != nil {
		return nil, err
	}
	if firewallRule.Status != "success" && firewallRule.Status != "" {
		return nil, errors.New(firewallRule.Message)
	}

	return &firewallRule, nil
}

func (s *FirewallService) DeleteCloudInsanceFromFirewall(firewallId, firewallRuleId string) (*DeleteResponse, error) {
	reqUrl := "firewall/" + firewallId + "/server/" + firewallRuleId + "/delete"
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
