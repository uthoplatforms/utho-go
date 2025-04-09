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
	ID             string              `json:"id"`
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
	Cloudid        string              `json:"cloudid"`
	Powerstatus    string              `json:"powerstatus"`
	Dclocation     K8sDclocation       `json:"dclocation"`
	Status         string              `json:"status"`
	WorkerCount    string              `json:"worker_count"`
	LoadBalancers  []K8sLoadbalancers  `json:"load_balancers"`
	TargetGroups   []K8sTargetGroups   `json:"target_groups"`
	SecurityGroups []K8sSecurityGroups `json:"security_groups"`
}
type KubernetesCluster struct {
	Info           KubernetesClusterInfo      `json:"info"`
	Vpc            []VpcDetails               `json:"vpc"`
	Nodepools      map[string]NodepoolDetails `json:"nodepools"`
	LoadBalancers  []K8sLoadbalancers         `json:"load_balancers"`
	TargetGroups   []K8sTargetGroups          `json:"target_groups"`
	SecurityGroups []K8sSecurityGroups        `json:"security_groups"`
	Rcode          string                     `json:"rcode"`
	Status         string                     `json:"status"`
	Message        string                     `json:"message"`
}
type KubernetesClusterInfo struct {
	Cluster KubernetesClusterMetadata `json:"cluster"`
	Master  MasterNodeDetails         `json:"master"`
}
type KubernetesClusterMetadata struct {
	ID              string        `json:"id"`
	Version         string        `json:"version"`
	Label           string        `json:"label"`
	Endpoint        string        `json:"endpoint"`
	Dcslug          string        `json:"dcslug"`
	AutoUpgrade     string        `json:"auto_upgrade"`
	SurgeUpgrade    string        `json:"surge_upgrade"`
	Ipv4            string        `json:"ipv4"`
	ClusterSubnet   string        `json:"cluster_subnet"`
	ServiceSubnet   string        `json:"service_subnet"`
	Tags            string        `json:"tags"`
	CreatedAt       string        `json:"created_at"`
	UpdatedAt       string        `json:"updated_at"`
	DeletedAt       string        `json:"deleted_at"`
	Status          string        `json:"status"`
	Nodepools       string        `json:"nodepools"`
	Vpc             string        `json:"vpc"`
	PublicIpEnabled string        `json:"public_ip_enabled"`
	LoadBalancers   string        `json:"load_balancers"`
	SecurityGroups  string        `json:"security_groups"`
	TargetGroups    string        `json:"target_groups"`
	Userid          string        `json:"userid"`
	Powerstatus     string        `json:"powerstatus"`
	Dclocation      K8sDclocation `json:"dclocation"`
}
type MasterNodeDetails struct {
	Cloudid        string         `json:"cloudid"`
	Hostname       string         `json:"hostname"`
	Ram            string         `json:"ram"`
	Cpu            string         `json:"cpu"`
	Cost           string         `json:"cost"`
	Disksize       string         `json:"disksize"`
	AppStatus      string         `json:"app_status"`
	Dcslug         string         `json:"dcslug"`
	Planid         string         `json:"planid"`
	Ip             string         `json:"ip"`
	PrivateNetwork PrivateNetwork `json:"private_network"`
}
type NodepoolDetails struct {
	Id        string        `json:"id"`
	Size      string        `json:"size"`
	Cost      float64       `json:"cost"`
	Planid    string        `json:"planid"`
	Count     string        `json:"count"`
	AutoScale bool          `json:"auto_scale"`
	MinNodes  int           `json:"min_nodes"`
	MaxNodes  int           `json:"max_nodes"`
	Policies  []interface{} `json:"policies"`
	Workers   []WorkerNode  `json:"workers"`
}
type WorkerNode struct {
	ID             string         `json:"cloudid"`
	Nodepool       string         `json:"nodepool"`
	Hostname       string         `json:"hostname"`
	Ram            string         `json:"ram"`
	Cost           string         `json:"cost"`
	Cpu            string         `json:"cpu"`
	Disksize       string         `json:"disksize"`
	AppStatus      string         `json:"app_status"`
	Ip             string         `json:"ip"`
	Planid         string         `json:"planid"`
	Status         string         `json:"status"`
	PrivateNetwork PrivateNetwork `json:"private_network"`
}
type VpcDetails struct {
	ID         string `json:"id"`
	VpcNetwork string `json:"vpc_network"`
}
type PrivateNetwork struct {
	Ip         string `json:"ip"`
	Vpc        string `json:"vpc"`
	VpcNetwork string `json:"vpc_network"`
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
	NetworkType    string                  `json:"network_type"`
	Firewall       string                  `json:"firewall"`
	Cpumodel       string                  `json:"cpumodel"`
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

func (s *KubernetesService) Read(clusterId string) (*KubernetesCluster, error) {
	reqUrl := "kubernetes/" + clusterId
	req, _ := s.client.NewRequest("GET", reqUrl)

	var kubernetes KubernetesCluster
	_, err := s.client.Do(req, &kubernetes)
	if err != nil {
		return nil, err
	}
	if kubernetes.Info.Cluster.ID != clusterId {
		return nil, errors.New("sorry we unable to find this cluster or you dont have access")
	}

	return &kubernetes, nil
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
	if delResponse.Status != "success" && delResponse.Status != "" {
		return nil, errors.New(delResponse.Message)
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

	var kubernetess KubernetesCluster
	_, err := s.client.Do(req, &kubernetess)
	if err != nil {
		return nil, err
	}
	if kubernetess.Info.Cluster.Status != "Active" && kubernetess.Info.Cluster.Status != "" {
		return nil, errors.New(kubernetess.Message)
	}
	var loadbalancers K8sLoadbalancers
	for _, r := range kubernetess.LoadBalancers {
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

	var kubernetess KubernetesCluster
	_, err := s.client.Do(req, &kubernetess)
	if err != nil {
		return nil, err
	}
	if kubernetess.Status != "success" && kubernetess.Status != "" {
		return nil, errors.New(kubernetess.Message)
	}

	return kubernetess.LoadBalancers, nil
}

func (s *KubernetesService) DeleteLoadbalancer(kubernetesId, kubernetesLoadbalancerId string) (*DeleteResponse, error) {
	reqUrl := "kubernetes/" + kubernetesId + "/loadbalancerpolicy/" + kubernetesLoadbalancerId
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

	var kubernetess KubernetesCluster
	_, err := s.client.Do(req, &kubernetess)
	if err != nil {
		return nil, err
	}
	if kubernetess.Status != "success" && kubernetess.Status != "" {
		return nil, errors.New(kubernetess.Message)
	}
	var securitygroups K8sSecurityGroups
	for _, r := range kubernetess.SecurityGroups {
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

	var kubernetess KubernetesCluster
	_, err := s.client.Do(req, &kubernetess)
	if err != nil {
		return nil, err
	}
	if kubernetess.Status != "success" && kubernetess.Status != "" {
		return nil, errors.New(kubernetess.Message)
	}

	return kubernetess.SecurityGroups, nil
}

func (s *KubernetesService) DeleteSecurityGroup(kuberneteseId, kubernetesSecurityGroupId string) (*DeleteResponse, error) {
	reqUrl := "kubernetes/" + kuberneteseId + "/securitygroup/" + kubernetesSecurityGroupId
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

	var kubernetess KubernetesCluster
	_, err := s.client.Do(req, &kubernetess)
	if err != nil {
		return nil, err
	}
	if kubernetess.Status != "success" && kubernetess.Status != "" {
		return nil, errors.New(kubernetess.Message)
	}

	if len(kubernetess.Info.Cluster.ID) == 0 {
		return nil, errors.New("no Cluster Found")
	}
	var targetgroups K8sTargetGroups
	for _, tg := range kubernetess.TargetGroups {
		if tg.ID == targetgroupId {
			targetgroups = tg
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

	var kubernetess KubernetesCluster
	_, err := s.client.Do(req, &kubernetess)
	if err != nil {
		return nil, err
	}
	if kubernetess.Status != "success" && kubernetess.Status != "" {
		return nil, errors.New(kubernetess.Message)
	}

	return kubernetess.TargetGroups, nil
}

func (s *KubernetesService) DeleteTargetgroup(kuberneteseId, kubernetesTargetgroupId string) (*DeleteResponse, error) {
	reqUrl := "kubernetes/" + kuberneteseId + "/targetgroup/" + kubernetesTargetgroupId
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
	NodePoolId   string
	Count        string `json:"count"`
	Label        string `json:"label"`
	PoolType     string `json:"pool_type"`
	Size         string `json:"size"`
	Policies     string `json:"policies"`
	MinNodes     int    `json:"min_nodes"`
	MaxNodes     int    `json:"max_nodes"`
}

func (s *KubernetesService) UpdateAutoscaleNodepool(params UpdateKubernetesAutoscaleNodepool) (*UpdateResponse, error) {
	reqUrl := "kubernetes/" + params.KubernetesId + "/nodepool/" + params.NodePoolId + "/update"
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
	NodePoolId   string
	Count        string `json:"count"`
	Label        string `json:"label"`
	PoolType     string `json:"pool_type"`
	Size         string `json:"size"`
}

func (s *KubernetesService) UpdateStaticNodepool(params UpdateKubernetesStaticNodepool) (*UpdateResponse, error) {
	reqUrl := "kubernetes/" + params.KubernetesId + "/nodepool/" + params.NodePoolId + "/update"
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
