package utho

import (
	"errors"
)

type FirewallService service

type Firewalls struct {
	Firewalls []Firewall `json:"firewalls"`
	Status    string     `json:"status"`
	Message   string     `json:"message"`
}

type Firewall struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	CreatedAt    string `json:"created_at"`
	Rulecount    string `json:"rulecount"`
	Serverscount string `json:"serverscount"`
	Rules        []Rule `json:"rules"`
}
type Rule struct {
	ID         string `json:"id"`
	Firewallid string `json:"firewallid"`
	Type       string `json:"type"`
	Service    string `json:"service"`
	Protocol   string `json:"protocol"`
	Port       string `json:"port"`
	Addresses  string `json:"addresses"`
}

type CreateFirewallParams struct {
	Name string `json:"name"`
}

func (s *FirewallService) CreateFirewall(params CreateFirewallParams) (*CreateResponse, error) {
	reqUrl := "firewall/create"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var firewall CreateResponse
	_, err := s.client.Do(req, &firewall)
	if err != nil {
		return nil, err
	}
	if firewall.Status != "success" && firewall.Status != "" {
		return nil, errors.New(firewall.Message)
	}

	return &firewall, nil
}

func (s *FirewallService) ReadFirewall(firewallId string) (*Firewall, error) {
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

	return &firewall.Firewalls[0], nil
}

func (s *FirewallService) ListFirewalls() (*[]Firewall, error) {
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

	return &firewall.Firewalls, nil
}

func (s *FirewallService) DeleteFirewall(firewallId string) (*DeleteResponse, error) {
	reqUrl := "firewall/" + firewallId + "/destroy"
	req, _ := s.client.NewRequest("DELETE", reqUrl)

	var delResponse DeleteResponse
	if _, err := s.client.Do(req, &delResponse); err != nil {
		return nil, err
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

func (s *FirewallService) ReadFirewallRule(firewallId, firewallRuleId string) (*Rule, error) {
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

	var rule Rule
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

func (s *FirewallService) ListFirewallRules() (*[]Rule, error) {
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

	return &firewall.Firewalls[0].Rules, nil
}

func (s *FirewallService) DeleteFirewallRule(firewallId, firewallRuleId string) (*DeleteResponse, error) {
	reqUrl := "firewall/" + firewallId + "/rule/" + firewallRuleId + "/delete"
	req, _ := s.client.NewRequest("DELETE", reqUrl)

	var delResponse DeleteResponse
	if _, err := s.client.Do(req, &delResponse); err != nil {
		return nil, err
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

	return &delResponse, nil
}
