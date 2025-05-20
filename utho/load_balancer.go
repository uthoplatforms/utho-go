package utho

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type AppStatus string

const (
	Pending   AppStatus = "Pending"
	Installed AppStatus = "Installed"
)

type LoadbalancersService service

type Loadbalancers struct {
	Loadbalancers []Loadbalancer `json:"loadbalancers"`
	Status        string         `json:"status" faker:"oneof: success, failure"`
	Message       string         `json:"message" faker:"sentence"`
}

type Loadbalancer struct {
	ID                  string      `json:"id" faker:"uuid_digit"`
	Userid              string      `json:"userid" faker:"uuid_digit"`
	IP                  string      `json:"ip" faker:"ipv4"`
	Name                string      `json:"name" faker:"name"`
	Algorithm           string      `json:"algorithm" faker:"oneof: roundrobin, leastconn"`
	Cookie              string      `json:"cookie" faker:"oneof: 0, 1"`
	Cookiename          string      `json:"cookiename"`
	Redirecthttps       string      `json:"redirecthttps" faker:"oneof: 0, 1"`
	Type                string      `json:"type" faker:"oneof: application, network"`
	Country             string      `json:"country"`
	Cc                  string      `json:"cc"`
	City                string      `json:"city"`
	Backendcount        string      `json:"backendcount"`
	CreatedAt           string      `json:"created_at" faker:"timestamp"`
	Status              string      `json:"status" faker:"oneof: Active, Inactive"`
	AppStatus           string      `json:"app_status" faker:"oneof: Pending, Installed"`
	KubernetesClusterid string      `json:"kubernetes_clusterid" faker:"uuid_digit"`
	Backends            []Backends  `json:"backends"`
	Rules               []Rules     `json:"rules"`
	Acls                []ACLs      `json:"acls"`
	Routes              []Routes    `json:"routes"`
	Frontends           []Frontends `json:"frontends"`
}

type Backends struct {
	ID      string `json:"id" faker:"uuid_digit"`
	Lb      string `json:"lb" faker:"uuid_digit"`
	IP      string `json:"ip" faker:"ipv4"`
	Cloudid string `json:"cloudid" faker:"uuid_digit"`
	Name    string `json:"name" faker:"name"`
	RAM     string `json:"ram"`
	CPU     string `json:"cpu"`
	Disk    string `json:"disk"`
	Country string `json:"country"`
	Cc      string `json:"cc"`
	City    string `json:"city"`
}
type Rules struct {
	ID          string `json:"id" faker:"uuid_digit"`
	Lb          string `json:"lb" faker:"uuid_digit"`
	SrcProto    string `json:"src_proto" faker:"oneof: TCP, UDP"`
	SrcPort     string `json:"src_port"`
	DstProto    string `json:"dst_proto" faker:"oneof: TCP, UDP"`
	DstPort     string `json:"dst_port"`
	Timeadded   string `json:"timeadded" faker:"timestamp"`
	Timeupdated string `json:"timeupdated" faker:"timestamp"`
}
type ACLs struct {
	ID           string `json:"id" faker:"uuid_digit"`
	Name         string `json:"name" faker:"name"`
	ACLCondition string `json:"acl_condition"`
	Value        string `json:"value" faker:"sentence"`
}
type Routes struct {
	ID               string `json:"id" faker:"uuid_digit"`
	ACLID            string `json:"acl_id" faker:"uuid_digit"`
	ACLName          string `json:"acl_name" faker:"name"`
	RoutingCondition string `json:"routing_condition"`
	BackendID        string `json:"backend_id" faker:"uuid_digit"`
}

type Frontends struct {
	ID            string           `json:"id" faker:"uuid_digit"`
	Name          string           `json:"name" faker:"name"`
	Algorithm     string           `json:"algorithm" faker:"oneof: roundrobin, leastconn"`
	Cookie        string           `json:"cookie" faker:"oneof: 0, 1"`
	Cookiename    string           `json:"cookiename"`
	Redirecthttps string           `json:"redirecthttps" faker:"oneof: 0, 1"`
	CertificateID string           `json:"certificate_id" faker:"uuid_digit"`
	Port          string           `json:"port"`
	Proto         string           `json:"proto" faker:"oneof: HTTP, HTTPS"`
	CreatedAt     string           `json:"created_at" faker:"timestamp"`
	UpdatedAt     string           `json:"updated_at" faker:"timestamp"`
	Acls          []ACLs           `json:"acls"`
	Routes        []FrontendRoutes `json:"routes"`
	// Backends      []any    `json:"backends"`
	// Rules         []any    `json:"rules"`
}
type FrontendRoutes struct {
	ID               string `json:"id" faker:"uuid_digit"`
	Lbid             string `json:"lbid" faker:"uuid_digit"`
	ACLID            string `json:"acl_id" faker:"uuid_digit"`
	RoutingCondition string `json:"routing_condition"`
	TargetGroups     string `json:"target_groups" faker:"uuid_digit"`
	FrontendID       string `json:"frontend_id" faker:"uuid_digit"`
	BackendID        string `json:"backend_id" faker:"uuid_digit"`
	ACLName          string `json:"acl_name" faker:"name"`
}

type CreateLoadblancerParams struct {
	Name                string `json:"name"`
	Dcslug              string `json:"dcslug"`
	Type                string `json:"type"`
	Cpumodel            string `json:"cpumodel"`
	VpcID               string `json:"vpc"`
	Firewall            string `json:"firewall"`
	EnablePublicip      string `json:"enable_publicip"`
	KubernetesClusterid string `json:"kubernetes_clusterid"`
}
type CreateLoadbalancerResponse struct {
	Status  string `json:"status"`
	ID      string `json:"loadbalancerid"`
	Message string `json:"message"`
}

func (s *LoadbalancersService) Create(params CreateLoadblancerParams) (*CreateLoadbalancerResponse, error) {
	reqUrl := "loadbalancer/add"
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
	reqUrl := "loadbalancer/" + loadbalancerId + "/destroy"
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

	maxRetries := 5
	sleepDuration := 45 * time.Second
	res, err := s.retryUntilReady(req, maxRetries, sleepDuration)
	if err != nil {
		return nil, err
	}

	if res.Status != "success" && res.Status != "" {
		return nil, errors.New(res.Message)
	}

	return &res, nil
}

type UpdateLoadbalancerACLParams struct {
	LoadbalancerId string
	ACLId          string
	Name           string `json:"name"`
	ConditionType  string `json:"conditionType"`
	FrontendID     string `json:"frontend_id"`
	Value          string `json:"value"`
}

func (s *LoadbalancersService) UpdateACL(params UpdateLoadbalancerACLParams) (*UpdateResponse, error) {
	reqUrl := "loadbalancer/" + params.LoadbalancerId + "/acl/" + params.ACLId
	req, _ := s.client.NewRequest("PUT", reqUrl, &params)

	var updateResponse UpdateResponse
	if _, err := s.client.Do(req, &updateResponse); err != nil {
		return nil, err
	}

	if updateResponse.Status != "success" && updateResponse.Status != "" {
		return nil, errors.New(updateResponse.Message)
	}
	return &updateResponse, nil
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
			break
		}
	}
	if acl.ID == "" {
		return nil, errors.New("Loadbalancer ACL not found")
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
	if len(loadbalancer.Loadbalancers) == 0 {
		return []ACLs{}, nil
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

	maxRetries := 5
	sleepDuration := 45 * time.Second
	res, err := s.retryUntilReady(req, maxRetries, sleepDuration)
	if err != nil {
		return nil, err
	}

	if res.Status != "success" && res.Status != "" {
		return nil, errors.New(res.Message)
	}

	return &res, nil
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
			break
		}
	}
	if frontend.ID == "" {
		return nil, errors.New("Loadbalancer Frontend not found")
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
	if len(loadbalancer.Loadbalancers) == 0 {
		return []Frontends{}, nil
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
	PoolName       string `json:"pool_name"`
}

func (s *LoadbalancersService) CreateBackend(params CreateLoadbalancerBackendParams) (*CreateResponse, error) {
	reqUrl := "loadbalancer/" + params.LoadbalancerId + "/backend"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	maxRetries := 5
	sleepDuration := 45 * time.Second
	res, err := s.retryUntilReady(req, maxRetries, sleepDuration)
	if err != nil {
		return nil, err
	}

	if res.Status != "success" && res.Status != "" {
		return nil, errors.New(res.Message)
	}

	return &res, nil
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
			break
		}
	}
	if backend.ID == "" {
		return nil, errors.New("Loadbalancer Backend not found")
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
	if len(loadbalancer.Loadbalancers) == 0 {
		return []Backends{}, nil
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

	maxRetries := 5
	sleepDuration := 45 * time.Second
	res, err := s.retryUntilReady(req, maxRetries, sleepDuration)
	if err != nil {
		return nil, err
	}

	if res.Status != "success" && res.Status != "" {
		return nil, errors.New(res.Message)
	}

	return &res, nil
}

type UpdateLoadbalancerRouteParams struct {
	LoadbalancerId string
	RouteId        string
	ACLID          string `json:"acl_id"`
	RouteCondition string `json:"route_condition"`
	FrontendID     string `json:"frontend_id"`
	TargetGroups   string `json:"target_groups"`
}

func (s *LoadbalancersService) UpdateRoute(params UpdateLoadbalancerRouteParams) (*UpdateResponse, error) {
	reqUrl := "loadbalancer/" + params.LoadbalancerId + "/route/" + params.RouteId
	req, _ := s.client.NewRequest("PUT", reqUrl, &params)

	var updateResponse UpdateResponse
	if _, err := s.client.Do(req, &updateResponse); err != nil {
		return nil, err
	}

	if updateResponse.Status != "success" && updateResponse.Status != "" {
		return nil, errors.New(updateResponse.Message)
	}
	return &updateResponse, nil
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

	var route Routes
	for _, v := range loadbalancer.Loadbalancers[0].Routes {
		if v.ID == loadbalancerRouteId {
			route = v
			break
		}
	}
	if route.ID == "" {
		return nil, errors.New("Loadbalancer Route not found")
	}

	return &route, nil
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
	if len(loadbalancer.Loadbalancers) == 0 {
		return []Routes{}, nil
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

type UpdateLoadbalancerParams struct {
	LoadbalancerId string
	Dcslug         string `json:"dcslug"`
	Name           string `json:"name"`
	Type           string `json:"type"`
}

func (s *LoadbalancersService) Update(params UpdateLoadbalancerParams) (*UpdateResponse, error) {
	reqUrl := "loadbalancer/" + params.LoadbalancerId + "/update"
	req, _ := s.client.NewRequest("PUT", reqUrl, &params)

	var updateResponse UpdateResponse
	if _, err := s.client.Do(req, &updateResponse); err != nil {
		return nil, err
	}

	if updateResponse.Status != "success" && updateResponse.Status != "" {
		return nil, errors.New(updateResponse.Message)
	}
	return &updateResponse, nil
}

func (s *LoadbalancersService) retryUntilReady(req *http.Request, maxRetries int, sleepDuration time.Duration) (CreateResponse, error) {
	var res CreateResponse

	for i := 0; i < maxRetries; i++ {
		_, err := s.client.Do(req, &res)
		if err != nil {
			return CreateResponse{}, err
		}

		if strings.EqualFold(res.Status, "error") {
			return CreateResponse{}, fmt.Errorf("%s", res.Message)
		}

		if strings.EqualFold(res.AppStatus, string(Installed)) || strings.EqualFold(res.Status, "success") {
			return res, nil
		}

		fmt.Printf("Attempt %d/%d - LB status: %s\n", i+1, maxRetries, res.AppStatus)
		time.Sleep(sleepDuration)
	}

	return CreateResponse{}, fmt.Errorf("AppStatus did not become 'Installed' after %d retries. Last status: %s", maxRetries, res.AppStatus)
}
