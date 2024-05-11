package utho

import (
	"errors"
)

type KubernetesService service

type Kubernetes struct {
	K8s     []K8s  `json:"k8s"`
	Status  string `json:"status"`
	Message string `json:"message"`
}
type K8s struct {
	Cloudid        string              `json:"cloudid"`
	CreatedAt      string              `json:"created_at"`
	Dcslug         string              `json:"dcslug"`
	RefID          string              `json:"ref_id"`
	Nodepool       string              `json:"nodepool"`
	Hostname       string              `json:"hostname"`
	RAM            string              `json:"ram"`
	CPU            string              `json:"cpu"`
	Disksize       string              `json:"disksize"`
	AppStatus      string              `json:"app_status"`
	IP             string              `json:"ip"`
	ID             string              `json:"id"`
	Powerstatus    string              `json:"powerstatus"`
	Dclocation     K8sDclocation       `json:"dclocation"`
	Status         string              `json:"status"`
	WorkerCount    string              `json:"worker_count"`
	LoadBalancers  []K8sLoadbalancers  `json:"load_balancers"`
	TargetGroups   []K8sTargetGroups   `json:"target_groups"`
	SecurityGroups []K8sSecurityGroups `json:"security_groups"`
}
type K8sDclocation struct {
	Location string `json:"location"`
	Country  string `json:"country"`
	Dc       string `json:"dc"`
	Dccc     string `json:"dccc"`
}
type K8sLoadbalancers struct {
	ID   string `json:"lbid"`
	Name string `json:"name"`
	IP   string `json:"ip"`
}
type K8sTargetGroups struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Protocol any    `json:"protocol"`
	Port     string `json:"port"`
}
type K8sSecurityGroups struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateKubernetesParams struct {
	Dcslug         string                  `json:"dcslug"`
	ClusterLabel   string                  `json:"cluster_label"`
	ClusterVersion string                  `json:"cluster_version"`
	Nodepools      []CreateNodepoolsParams `json:"nodepools"`
	Auth           string                  `json:"auth"`
	Vpc            string                  `json:"vpc"`
	SecurityGroups string                  `json:"security_groups"`
}
type CreateNodepoolsParams struct {
	Label    string                           `json:"label"`
	Size     string                           `json:"size"`
	PoolType string                           `json:"pool_type"`
	MaxCount string                           `json:"maxCount,omitempty"`
	Count    string                           `json:"count"`
	Policies []CreateKubernetesPoliciesParams `json:"policies,omitempty"`
}
type CreateKubernetesPoliciesParams struct {
	Adjust   int    `json:"adjust"`
	Compare  string `json:"compare"`
	Cooldown int    `json:"cooldown"`
	Period   string `json:"period"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Value    string `json:"value"`
	Product  string `json:"product"`
	Maxsize  string `json:"maxsize"`
	Minsize  string `json:"minsize"`
}

func (s *KubernetesService) Create(params CreateKubernetesParams) (*CreateResponse, error) {
	reqUrl := "kubernetes/deploy"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var kubernetes CreateResponse
	_, err := s.client.Do(req, &kubernetes)
	if err != nil {
		return nil, err
	}
	if kubernetes.Status != "success" && kubernetes.Status != "" {
		return nil, errors.New(kubernetes.Message)
	}

	return &kubernetes, nil
}

func (s *KubernetesService) Read(clusterId string) (*K8s, error) {
	reqUrl := "kubernetes"
	req, _ := s.client.NewRequest("GET", reqUrl)

	var kubernetes Kubernetes
	_, err := s.client.Do(req, &kubernetes)
	if err != nil {
		return nil, err
	}
	if kubernetes.Status != "success" && kubernetes.Status != "" {
		return nil, errors.New(kubernetes.Message)
	}

	var k8s K8s
	for _, r := range kubernetes.K8s {
		if r.ID == clusterId {
			k8s = r
		}
	}
	if len(k8s.ID) == 0 {
		return nil, errors.New("kubernetess loadbalancer not found")
	}
	return &k8s, nil
}

func (s *KubernetesService) List() ([]K8s, error) {
	reqUrl := "kubernetes"
	req, _ := s.client.NewRequest("GET", reqUrl)

	var kubernetes Kubernetes
	_, err := s.client.Do(req, &kubernetes)
	if err != nil {
		return nil, err
	}
	if kubernetes.Status != "success" && kubernetes.Status != "" {
		return nil, errors.New(kubernetes.Message)
	}

	return kubernetes.K8s, nil
}

type DeleteKubernetesParams struct {
	ClusterId string
	// confirm message"I am aware this action will delete data and cluster permanently"
	Confirm string `json:"confirm"`
}

func (s *KubernetesService) Delete(params DeleteKubernetesParams) (*DeleteResponse, error) {
	reqUrl := "kubernetes/" + params.ClusterId + "/destroy"
	req, _ := s.client.NewRequest("DELETE", reqUrl)

	var delResponse DeleteResponse
	if _, err := s.client.Do(req, &delResponse); err != nil {
		return nil, err
	}

	return &delResponse, nil
}

type CreateKubernetesLoadbalancerParams struct {
	KubernetesId   string
	LoadbalancerId string
}

func (s *KubernetesService) CreateLoadbalancer(params CreateKubernetesLoadbalancerParams) (*CreateResponse, error) {
	reqUrl := "kubernetes/" + params.KubernetesId + "/loadbalancer/" + params.LoadbalancerId
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var kubernetes CreateResponse
	_, err := s.client.Do(req, &kubernetes)
	if err != nil {
		return nil, err
	}
	if kubernetes.Status != "success" && kubernetes.Status != "" {
		return nil, errors.New(kubernetes.Message)
	}

	return &kubernetes, nil
}

func (s *KubernetesService) ReadLoadbalancer(kubernetesId, loadbalancerId string) (*K8sLoadbalancers, error) {
	reqUrl := "kubernetes/" + kubernetesId
	req, _ := s.client.NewRequest("GET", reqUrl)

	var kubernetess Kubernetes
	_, err := s.client.Do(req, &kubernetess)
	if err != nil {
		return nil, err
	}
	if kubernetess.Status != "success" && kubernetess.Status != "" {
		return nil, errors.New(kubernetess.Message)
	}
	var loadbalancers K8sLoadbalancers
	for _, r := range kubernetess.K8s[0].LoadBalancers {
		if r.ID == loadbalancerId {
			loadbalancers = r
		}
	}
	if len(loadbalancers.ID) == 0 {
		return nil, errors.New("kubernetess loadbalancer not found")
	}

	return &loadbalancers, nil
}

func (s *KubernetesService) ListLoadbalancers(kubernetesId string) ([]K8sLoadbalancers, error) {
	reqUrl := "kubernetes/" + kubernetesId
	req, _ := s.client.NewRequest("GET", reqUrl)

	var kubernetess Kubernetes
	_, err := s.client.Do(req, &kubernetess)
	if err != nil {
		return nil, err
	}
	if kubernetess.Status != "success" && kubernetess.Status != "" {
		return nil, errors.New(kubernetess.Message)
	}

	return kubernetess.K8s[0].LoadBalancers, nil
}

func (s *KubernetesService) DeleteLoadbalancer(kubernetesId, kubernetesLoadbalancerId string) (*DeleteResponse, error) {
	reqUrl := "kubernetes/" + kubernetesId + "/loadbalancerpolicy/" + kubernetesLoadbalancerId
	req, _ := s.client.NewRequest("DELETE", reqUrl)

	var delResponse DeleteResponse
	if _, err := s.client.Do(req, &delResponse); err != nil {
		return nil, err
	}

	return &delResponse, nil
}

type CreateKubernetesSecurityGroupParams struct {
	KubernetesId              string
	KubernetesSecurityGroupId string
}

func (s *KubernetesService) CreateSecurityGroup(params CreateKubernetesSecurityGroupParams) (*CreateResponse, error) {
	reqUrl := "kubernetes/" + params.KubernetesId + "/securitygroup/" + params.KubernetesSecurityGroupId
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var kubernetes CreateResponse
	_, err := s.client.Do(req, &kubernetes)
	if err != nil {
		return nil, err
	}
	if kubernetes.Status != "success" && kubernetes.Status != "" {
		return nil, errors.New(kubernetes.Message)
	}

	return &kubernetes, nil
}

func (s *KubernetesService) ReadSecurityGroup(kubernetesId, securitygroupId string) (*K8sSecurityGroups, error) {
	reqUrl := "kubernetes/" + kubernetesId
	req, _ := s.client.NewRequest("GET", reqUrl)

	var kubernetess Kubernetes
	_, err := s.client.Do(req, &kubernetess)
	if err != nil {
		return nil, err
	}
	if kubernetess.Status != "success" && kubernetess.Status != "" {
		return nil, errors.New(kubernetess.Message)
	}
	var securitygroups K8sSecurityGroups
	for _, r := range kubernetess.K8s[0].SecurityGroups {
		if r.ID == securitygroupId {
			securitygroups = r
		}
	}
	if len(securitygroups.ID) == 0 {
		return nil, errors.New("kubernetess securitygroup not found")
	}

	return &securitygroups, nil
}

func (s *KubernetesService) ListSecurityGroups(kubernetesId string) ([]K8sSecurityGroups, error) {
	reqUrl := "kubernetes/" + kubernetesId
	req, _ := s.client.NewRequest("GET", reqUrl)

	var kubernetess Kubernetes
	_, err := s.client.Do(req, &kubernetess)
	if err != nil {
		return nil, err
	}
	if kubernetess.Status != "success" && kubernetess.Status != "" {
		return nil, errors.New(kubernetess.Message)
	}

	return kubernetess.K8s[0].SecurityGroups, nil
}

func (s *KubernetesService) DeleteSecurityGroup(kuberneteseId, kubernetesSecurityGroupId string) (*DeleteResponse, error) {
	reqUrl := "kubernetes/" + kuberneteseId + "/securitygroup/" + kubernetesSecurityGroupId
	req, _ := s.client.NewRequest("DELETE", reqUrl)

	var delResponse DeleteResponse
	if _, err := s.client.Do(req, &delResponse); err != nil {
		return nil, err
	}

	return &delResponse, nil
}

type CreateKubernetesTargetgroupParams struct {
	KubernetesId            string
	KubernetesTargetgroupId string
}

func (s *KubernetesService) CreateTargetgroup(params CreateKubernetesTargetgroupParams) (*CreateResponse, error) {
	reqUrl := "kubernetes/" + params.KubernetesId + "/targetgroup/" + params.KubernetesTargetgroupId
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var kubernetes CreateResponse
	_, err := s.client.Do(req, &kubernetes)
	if err != nil {
		return nil, err
	}
	if kubernetes.Status != "success" && kubernetes.Status != "" {
		return nil, errors.New(kubernetes.Message)
	}

	return &kubernetes, nil
}

func (s *KubernetesService) ReadTargetgroup(kubernetesId, targetgroupId string) (*K8sTargetGroups, error) {
	reqUrl := "kubernetes/" + kubernetesId
	req, _ := s.client.NewRequest("GET", reqUrl)

	var kubernetess Kubernetes
	_, err := s.client.Do(req, &kubernetess)
	if err != nil {
		return nil, err
	}
	if kubernetess.Status != "success" && kubernetess.Status != "" {
		return nil, errors.New(kubernetess.Message)
	}
	var targetgroups K8sTargetGroups
	for _, r := range kubernetess.K8s[0].TargetGroups {
		if r.ID == targetgroupId {
			targetgroups = r
		}
	}
	if len(targetgroups.ID) == 0 {
		return nil, errors.New("kubernetess targetgroup not found")
	}

	return &targetgroups, nil
}

func (s *KubernetesService) ListTargetgroups(kubernetesId string) ([]K8sTargetGroups, error) {
	reqUrl := "kubernetes/" + kubernetesId
	req, _ := s.client.NewRequest("GET", reqUrl)

	var kubernetess Kubernetes
	_, err := s.client.Do(req, &kubernetess)
	if err != nil {
		return nil, err
	}
	if kubernetess.Status != "success" && kubernetess.Status != "" {
		return nil, errors.New(kubernetess.Message)
	}

	return kubernetess.K8s[0].TargetGroups, nil
}

func (s *KubernetesService) DeleteTargetgroup(kuberneteseId, kubernetesTargetgroupId string) (*DeleteResponse, error) {
	reqUrl := "kubernetes/" + kuberneteseId + "/targetgroup/" + kubernetesTargetgroupId
	req, _ := s.client.NewRequest("DELETE", reqUrl)

	var delResponse DeleteResponse
	if _, err := s.client.Do(req, &delResponse); err != nil {
		return nil, err
	}

	return &delResponse, nil
}

func (s *KubernetesService) PowerOff(kubernetesId string) (*BasicResponse, error) {
	reqUrl := "kubernetes/" + kubernetesId + "/stop"
	req, _ := s.client.NewRequest("POST", reqUrl)

	var basicResponse BasicResponse
	_, err := s.client.Do(req, &basicResponse)
	if err != nil {
		return nil, err
	}
	if basicResponse.Status != "success" && basicResponse.Status != "" {
		return nil, errors.New(basicResponse.Message)
	}

	return &basicResponse, nil
}

func (s *KubernetesService) PowerOn(kubernetesId string) (*BasicResponse, error) {
	reqUrl := "kubernetes/" + kubernetesId + "/start"
	req, _ := s.client.NewRequest("POST", reqUrl)

	var basicResponse BasicResponse
	_, err := s.client.Do(req, &basicResponse)
	if err != nil {
		return nil, err
	}
	if basicResponse.Status != "success" && basicResponse.Status != "" {
		return nil, errors.New(basicResponse.Message)
	}

	return &basicResponse, nil
}

type UpdateKubernetesAutoscaleNodepool struct {
	KubernetesId string
	NodeId       string
	Count        string `json:"count"`
	Label        string `json:"label"`
	PoolType     string `json:"pool_type"`
	Size         string `json:"size"`
	Policies     string `json:"policies"`
}

func (s *KubernetesService) UpdateAutoscaleNodepool(params UpdateKubernetesAutoscaleNodepool) (*UpdateResponse, error) {
	reqUrl := "kubernetes/" + params.KubernetesId + "/nodepool/" + params.NodeId + "/update"
	req, _ := s.client.NewRequest("POST", reqUrl)

	var kubernetes UpdateResponse
	_, err := s.client.Do(req, &kubernetes)
	if err != nil {
		return nil, err
	}
	if kubernetes.Status != "success" && kubernetes.Status != "" {
		return nil, errors.New(kubernetes.Message)
	}

	return &kubernetes, nil
}

type UpdateKubernetesStaticNodepool struct {
	KubernetesId string
	NodeId       string
	Count        string `json:"count"`
	Label        string `json:"label"`
	PoolType     string `json:"pool_type"`
	Size         string `json:"size"`
}

func (s *KubernetesService) UpdateStaticNodepool(params UpdateKubernetesStaticNodepool) (*UpdateResponse, error) {
	reqUrl := "kubernetes/" + params.KubernetesId + "/nodepool/" + params.NodeId + "/update"
	req, _ := s.client.NewRequest("POST", reqUrl)

	var kubernetes UpdateResponse
	_, err := s.client.Do(req, &kubernetes)
	if err != nil {
		return nil, err
	}
	if kubernetes.Status != "success" && kubernetes.Status != "" {
		return nil, errors.New(kubernetes.Message)
	}

	return &kubernetes, nil
}
