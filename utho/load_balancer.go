package utho

import (
	"errors"
)

type LoadbalancersService service

type Loadbalancers struct {
	Loadbalancers []Loadbalancer `json:"loadbalancers"`
	Status        string         `json:"status"`
	Message       string         `json:"message"`
}

type Loadbalancer struct {
	ID            string      `json:"id"`
	Userid        string      `json:"userid"`
	IP            string      `json:"ip"`
	Name          string      `json:"name"`
	Algorithm     string      `json:"algorithm"`
	Cookie        string      `json:"cookie"`
	Cookiename    string      `json:"cookiename"`
	Redirecthttps string      `json:"redirecthttps"`
	Type          string      `json:"type"`
	Country       string      `json:"country"`
	Cc            string      `json:"cc"`
	City          string      `json:"city"`
	Backendcount  string      `json:"backendcount"`
	CreatedAt     string      `json:"created_at"`
	Status        string      `json:"status"`
	Backends      []Backends  `json:"backends"`
	Rules         []Rules     `json:"rules"`
	Acls          []ACLs      `json:"acls"`
	Routes        []Routes    `json:"routes"`
	Frontends     []Frontends `json:"frontends"`
	// ScalingGroups []any       `json:"scaling_groups"`
}
type Backends struct {
	ID      string `json:"id"`
	Lb      string `json:"lb"`
	IP      string `json:"ip"`
	Cloudid string `json:"cloudid"`
	Name    string `json:"name"`
	RAM     string `json:"ram"`
	CPU     string `json:"cpu"`
	Disk    string `json:"disk"`
	Country string `json:"country"`
	Cc      string `json:"cc"`
	City    string `json:"city"`
}
type Rules struct {
	ID          string `json:"id"`
	Lb          string `json:"lb"`
	SrcProto    string `json:"src_proto"`
	SrcPort     string `json:"src_port"`
	DstProto    string `json:"dst_proto"`
	DstPort     string `json:"dst_port"`
	Timeadded   string `json:"timeadded"`
	Timeupdated string `json:"timeupdated"`
}
type ACLs struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	ACLCondition string `json:"acl_condition"`
	Value        string `json:"value"`
}
type Routes struct {
	ID               string `json:"id"`
	ACLID            string `json:"acl_id"`
	ACLName          string `json:"acl_name"`
	RoutingCondition string `json:"routing_condition"`
	BackendID        string `json:"backend_id"`
}

type Frontends struct {
	ID            string           `json:"id"`
	Name          string           `json:"name"`
	Algorithm     string           `json:"algorithm"`
	Cookie        string           `json:"cookie"`
	Cookiename    string           `json:"cookiename"`
	Redirecthttps string           `json:"redirecthttps"`
	CertificateID string           `json:"certificate_id"`
	Port          string           `json:"port"`
	Proto         string           `json:"proto"`
	CreatedAt     string           `json:"created_at"`
	UpdatedAt     string           `json:"updated_at"`
	Acls          []ACLs           `json:"acls"`
	Routes        []FrontendRoutes `json:"routes"`
	// Backends      []any    `json:"backends"`
	// Rules         []any    `json:"rules"`
}
type FrontendRoutes struct {
	ID               string `json:"id"`
	Lbid             string `json:"lbid"`
	ACLID            string `json:"acl_id"`
	RoutingCondition string `json:"routing_condition"`
	TargetGroups     string `json:"target_groups"`
	FrontendID       string `json:"frontend_id"`
	BackendID        string `json:"backend_id"`
	ACLName          string `json:"acl_name"`
}

type CreateLoadbalancerParams struct {
	Dcslug string `json:"dcslug"`
	Type   string `json:"type"`
	Name   string `json:"name"`
}
type CreateLoadbalancerResponse struct {
	Status  string `json:"status"`
	ID      string `json:"loadbalancerid"`
	Message string `json:"message"`
}

func (s *LoadbalancersService) Create(params CreateLoadbalancerParams) (*CreateLoadbalancerResponse, error) {
	reqUrl := "loadbalancer"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var loadbalancer CreateLoadbalancerResponse
	_, err := s.client.Do(req, &loadbalancer)
	if err != nil {
		return nil, err
	}
	if loadbalancer.Status != "success" && loadbalancer.Status != "" {
		return nil, errors.New(loadbalancer.Message)
	}

	return &loadbalancer, nil
}

func (s *LoadbalancersService) Read(loadbalancerId string) (*Loadbalancer, error) {
	reqUrl := "loadbalancer/" + loadbalancerId
	req, _ := s.client.NewRequest("GET", reqUrl)

	var loadbalancer Loadbalancers
	_, err := s.client.Do(req, &loadbalancer)
	if err != nil {
		return nil, err
	}
	if loadbalancer.Status != "success" && loadbalancer.Status != "" {
		return nil, errors.New(loadbalancer.Message)
	}
	if len(loadbalancer.Loadbalancers) == 0 {
		return nil, errors.New("NotFound")
	}

	return &loadbalancer.Loadbalancers[0], nil
}

func (s *LoadbalancersService) List() ([]Loadbalancer, error) {
	reqUrl := "loadbalancer"
	req, _ := s.client.NewRequest("GET", reqUrl)

	var loadbalancer Loadbalancers
	_, err := s.client.Do(req, &loadbalancer)
	if err != nil {
		return nil, err
	}
	if loadbalancer.Status != "success" && loadbalancer.Status != "" {
		return nil, errors.New(loadbalancer.Message)
	}

	return loadbalancer.Loadbalancers, nil
}

func (s *LoadbalancersService) Delete(loadbalancerId string) (*DeleteResponse, error) {
	reqUrl := "loadbalancer/" + loadbalancerId
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

type CreateLoadbalancerACLParams struct {
	LoadbalancerId string
	Name           string `json:"name"`
	ConditionType  string `json:"conditionType"`
	FrontendID     string `json:"frontend_id"`
	Value          string `json:"value"`
}

func (s *LoadbalancersService) CreateACL(params CreateLoadbalancerACLParams) (*CreateResponse, error) {
	reqUrl := "loadbalancer/" + params.LoadbalancerId + "/acl"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var loadbalancerACL CreateResponse
	_, err := s.client.Do(req, &loadbalancerACL)
	if err != nil {
		return nil, err
	}
	if loadbalancerACL.Status != "success" && loadbalancerACL.Status != "" {
		return nil, errors.New(loadbalancerACL.Message)
	}

	return &loadbalancerACL, nil
}

func (s *LoadbalancersService) ReadACL(loadbalancerId, loadbalancerACLId string) (*ACLs, error) {
	reqUrl := "loadbalancer/" + loadbalancerId
	req, _ := s.client.NewRequest("GET", reqUrl)

	var loadbalancer Loadbalancers
	_, err := s.client.Do(req, &loadbalancer)
	if err != nil {
		return nil, err
	}
	if loadbalancer.Status != "success" && loadbalancer.Status != "" {
		return nil, errors.New(loadbalancer.Message)
	}

	var acl ACLs
	for _, v := range loadbalancer.Loadbalancers[0].Acls {
		if v.ID == loadbalancerACLId {
			acl = v
		}
	}

	return &acl, nil
}

func (s *LoadbalancersService) ListACLs(loadbalancerId string) ([]ACLs, error) {
	reqUrl := "loadbalancer/" + loadbalancerId
	req, _ := s.client.NewRequest("GET", reqUrl)

	var loadbalancer Loadbalancers
	_, err := s.client.Do(req, &loadbalancer)
	if err != nil {
		return nil, err
	}
	if loadbalancer.Status != "success" && loadbalancer.Status != "" {
		return nil, errors.New(loadbalancer.Message)
	}

	return loadbalancer.Loadbalancers[0].Acls, nil
}

func (s *LoadbalancersService) DeleteACL(loadbalancerId, loadbalancerACLId string) (*DeleteResponse, error) {
	reqUrl := "loadbalancer/" + loadbalancerId + "/acl/" + loadbalancerACLId
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

type CreateLoadbalancerFrontendParams struct {
	LoadbalancerId string
	Name           string `json:"name"`
	Proto          string `json:"proto"`
	Port           string `json:"port"`
	CertificateID  string `json:"certificate_id,omitempty"`
	Algorithm      string `json:"algorithm"`
	Redirecthttps  string `json:"redirecthttps,omitempty"`
	Cookie         string `json:"cookie"`
}

func (s *LoadbalancersService) CreateFrontend(params CreateLoadbalancerFrontendParams) (*CreateResponse, error) {
	reqUrl := "loadbalancer/" + params.LoadbalancerId + "/frontend"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var loadbalancerFrontend CreateResponse
	_, err := s.client.Do(req, &loadbalancerFrontend)
	if err != nil {
		return nil, err
	}
	if loadbalancerFrontend.Status != "success" && loadbalancerFrontend.Status != "" {
		return nil, errors.New(loadbalancerFrontend.Message)
	}

	return &loadbalancerFrontend, nil
}

type UpdateLoadbalancerFrontendParams struct {
	LoadbalancerId string
	Name           string `json:"name"`
	Proto          string `json:"proto"`
	Port           string `json:"port"`
	CertificateID  string `json:"certificate_id,omitempty"`
	Algorithm      string `json:"algorithm"`
	Redirecthttps  string `json:"redirecthttps,omitempty"`
	Cookie         string `json:"cookie"`
}

func (s *LoadbalancersService) UpdateFrontend(params UpdateLoadbalancerFrontendParams, loadbalancerId, loadbalancerFrontendId string) (*UpdateResponse, error) {
	reqUrl := "loadbalancer/" + loadbalancerId + "/frontend/" + loadbalancerFrontendId
	req, _ := s.client.NewRequest("PUT", reqUrl, &params)

	var frontend UpdateResponse
	if _, err := s.client.Do(req, &frontend); err != nil {
		return nil, err
	}

	if frontend.Status != "success" && frontend.Status != "" {
		return nil, errors.New(frontend.Message)
	}
	return &frontend, nil
}

func (s *LoadbalancersService) ReadFrontend(loadbalancerId, loadbalancerFrontendId string) (*Frontends, error) {
	reqUrl := "loadbalancer/" + loadbalancerId
	req, _ := s.client.NewRequest("GET", reqUrl)

	var loadbalancer Loadbalancers
	_, err := s.client.Do(req, &loadbalancer)
	if err != nil {
		return nil, err
	}
	if loadbalancer.Status != "success" && loadbalancer.Status != "" {
		return nil, errors.New(loadbalancer.Message)
	}

	var frontend Frontends
	for _, v := range loadbalancer.Loadbalancers[0].Frontends {
		if v.ID == loadbalancerFrontendId {
			frontend = v
		}
	}

	return &frontend, nil
}

func (s *LoadbalancersService) ListFrontends(loadbalancerId string) ([]Frontends, error) {
	reqUrl := "loadbalancer/" + loadbalancerId
	req, _ := s.client.NewRequest("GET", reqUrl)

	var loadbalancer Loadbalancers
	_, err := s.client.Do(req, &loadbalancer)
	if err != nil {
		return nil, err
	}
	if loadbalancer.Status != "success" && loadbalancer.Status != "" {
		return nil, errors.New(loadbalancer.Message)
	}

	return loadbalancer.Loadbalancers[0].Frontends, nil
}

func (s *LoadbalancersService) DeleteFrontend(loadbalancerId, loadbalancerFrontendId string) (*DeleteResponse, error) {
	reqUrl := "loadbalancer/" + loadbalancerId + "/frontend/" + loadbalancerFrontendId
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

type CreateLoadbalancerBackendParams struct {
	LoadbalancerId string
	Type           string `json:"type"`
	FrontendID     string `json:"frontend_id"`
	BackendPort    string `json:"backend_port"`
	Cloudid        string `json:"cloudid,omitempty"`
	IP             string `json:"ip,omitempty"`
}

func (s *LoadbalancersService) CreateBackend(params CreateLoadbalancerBackendParams) (*CreateResponse, error) {
	reqUrl := "loadbalancer/" + params.LoadbalancerId + "/backend"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var loadbalancerBackend CreateResponse
	_, err := s.client.Do(req, &loadbalancerBackend)
	if err != nil {
		return nil, err
	}
	if loadbalancerBackend.Status != "success" && loadbalancerBackend.Status != "" {
		return nil, errors.New(loadbalancerBackend.Message)
	}

	return &loadbalancerBackend, nil
}

func (s *LoadbalancersService) ReadBackend(loadbalancerId, loadbalancerBackendId string) (*Backends, error) {
	reqUrl := "loadbalancer/" + loadbalancerId
	req, _ := s.client.NewRequest("GET", reqUrl)

	var loadbalancer Loadbalancers
	_, err := s.client.Do(req, &loadbalancer)
	if err != nil {
		return nil, err
	}
	if loadbalancer.Status != "success" && loadbalancer.Status != "" {
		return nil, errors.New(loadbalancer.Message)
	}

	var backend Backends
	for _, v := range loadbalancer.Loadbalancers[0].Backends {
		if v.ID == loadbalancerBackendId {
			backend = v
		}
	}

	return &backend, nil
}

func (s *LoadbalancersService) ListBackends(loadbalancerId string) ([]Backends, error) {
	reqUrl := "loadbalancer/" + loadbalancerId
	req, _ := s.client.NewRequest("GET", reqUrl)

	var loadbalancer Loadbalancers
	_, err := s.client.Do(req, &loadbalancer)
	if err != nil {
		return nil, err
	}
	if loadbalancer.Status != "success" && loadbalancer.Status != "" {
		return nil, errors.New(loadbalancer.Message)
	}

	return loadbalancer.Loadbalancers[0].Backends, nil
}

func (s *LoadbalancersService) DeleteBackend(loadbalancerId, loadbalancerBackendId string) (*DeleteResponse, error) {
	reqUrl := "loadbalancer/" + loadbalancerId + "/backend/" + loadbalancerBackendId
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

type CreateLoadbalancerRouteParams struct {
	LoadbalancerId string
	FrontendID     string `json:"frontend_id"`
	ACLID          string `json:"acl_id"`
	RouteCondition string `json:"route_condition"`
	TargetGroups   string `json:"target_groups"`
}

func (s *LoadbalancersService) CreateRoute(params CreateLoadbalancerRouteParams) (*CreateResponse, error) {
	reqUrl := "loadbalancer/" + params.LoadbalancerId + "/route"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var loadbalancerRoute CreateResponse
	_, err := s.client.Do(req, &loadbalancerRoute)
	if err != nil {
		return nil, err
	}
	if loadbalancerRoute.Status != "success" && loadbalancerRoute.Status != "" {
		return nil, errors.New(loadbalancerRoute.Message)
	}

	return &loadbalancerRoute, nil
}

func (s *LoadbalancersService) ReadRoute(loadbalancerId, loadbalancerRouteId string) (*Routes, error) {
	reqUrl := "loadbalancer/" + loadbalancerId
	req, _ := s.client.NewRequest("GET", reqUrl)

	var loadbalancer Loadbalancers
	_, err := s.client.Do(req, &loadbalancer)
	if err != nil {
		return nil, err
	}
	if loadbalancer.Status != "success" && loadbalancer.Status != "" {
		return nil, errors.New(loadbalancer.Message)
	}

	var backend Routes
	for _, v := range loadbalancer.Loadbalancers[0].Routes {
		if v.ID == loadbalancerRouteId {
			backend = v
		}
	}

	return &backend, nil
}

func (s *LoadbalancersService) ListRoutes(loadbalancerId string) ([]Routes, error) {
	reqUrl := "loadbalancer/" + loadbalancerId
	req, _ := s.client.NewRequest("GET", reqUrl)

	var loadbalancer Loadbalancers
	_, err := s.client.Do(req, &loadbalancer)
	if err != nil {
		return nil, err
	}
	if loadbalancer.Status != "success" && loadbalancer.Status != "" {
		return nil, errors.New(loadbalancer.Message)
	}

	return loadbalancer.Loadbalancers[0].Routes, nil
}

func (s *LoadbalancersService) DeleteRoute(loadbalancerId, loadbalancerRouteId string) (*DeleteResponse, error) {
	reqUrl := "loadbalancer/" + loadbalancerId + "/route/" + loadbalancerRouteId
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
