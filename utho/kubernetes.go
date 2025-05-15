package utho

import (
	"context"
	"errors"
	"fmt"
)

type KubernetesService service

type Kubernetes struct {
	K8s     []K8s  `json:"k8s"`
	Status  string `json:"status"`
	Message string `json:"message"`
}
type K8s struct {
	ID             int                 `json:"id,string"`
	CreatedAt      string              `json:"created_at"`
	Dcslug         string              `json:"dcslug"`
	RefID          string              `json:"ref_id"`
	Nodepool       string              `json:"nodepool"`
	Hostname       string              `json:"hostname"`
	RAM            int                 `json:"ram,string"`
	CPU            int                 `json:"cpu,string"`
	Disksize       int                 `json:"disksize,string"`
	AppStatus      string              `json:"app_status"`
	IP             string              `json:"ip"`
	Cloudid        int                 `json:"cloudid,string"`
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
	ID              int           `json:"id,string"`
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
	Userid          int           `json:"userid,string"`
	Powerstatus     string        `json:"powerstatus"`
	Dclocation      K8sDclocation `json:"dclocation"`
}
type MasterNodeDetails struct {
	Cloudid        int            `json:"cloudid,string"`
	Hostname       string         `json:"hostname"`
	Ram            int            `json:"ram,string"`
	Cpu            int            `json:"cpu,string"`
	Cost           string         `json:"cost"`
	Disksize       int            `json:"disksize,string"`
	AppStatus      string         `json:"app_status"`
	Dcslug         string         `json:"dcslug"`
	Planid         int            `json:"planid,string"`
	Ip             string         `json:"ip"`
	PrivateNetwork PrivateNetwork `json:"private_network"`
}
type NodepoolDetails struct {
	ID        string        `json:"id"`
	Size      string        `json:"size"`
	Cost      float64       `json:"cost"`
	Planid    int           `json:"planid,string"`
	Count     int           `json:"count,string"`
	AutoScale bool          `json:"auto_scale"`
	MinNodes  int           `json:"min_nodes,string"`
	MaxNodes  int           `json:"max_nodes,string"`
	Policies  []interface{} `json:"policies"`
	Workers   []WorkerNode  `json:"workers"`
}

type WorkerNode struct {
	ID             int            `json:"cloudid,string"`
	Nodepool       string         `json:"nodepool"`
	Hostname       string         `json:"hostname"`
	Ram            int            `json:"ram,string"`
	Cost           string         `json:"cost"`
	Cpu            int            `json:"cpu,string"`
	Disksize       int            `json:"disksize,string"`
	AppStatus      string         `json:"app_status"`
	Ip             string         `json:"ip"`
	Planid         int            `json:"planid,string"`
	Status         string         `json:"status"`
	PrivateNetwork PrivateNetwork `json:"private_network"`
}
type VpcDetails struct {
	ID         int    `json:"id,string"`
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
	ID   int    `json:"lbid,string"`
	Name string `json:"name"`
	IP   string `json:"ip"`
}
type K8sTargetGroups struct {
	ID       int    `json:"id,string"`
	Name     string `json:"name"`
	Protocol any    `json:"protocol"`
	Port     string `json:"port"`
}
type K8sSecurityGroups struct {
	ID   int    `json:"id,string"`
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

type CreateKubernetesNodePoolParams struct {
	ClusterId int
	Nodepools []CreateNodepoolsDetails `json:"nodepools"`
}
type CreateNodepoolsDetails struct {
	Label string              `json:"label"`
	Size  string              `json:"size"`
	Count string              `json:"count"`
	Ebs   []CreateNodePoolEbs `json:"ebs"`
}
type CreateNodePoolEbs struct {
	Disk string `json:"disk"`
	Type string `json:"type"`
}

func (k *KubernetesService) Create(ctx context.Context, params CreateKubernetesParams) (*CreateResponse, error) {
	reqUrl := "kubernetes/deploy"
	req, _ := k.client.NewRequest("POST", reqUrl, &params)

	var kubernetes CreateResponse
	_, err := k.client.Do(req, &kubernetes)
	if err != nil {
		return nil, err
	}
	if kubernetes.Status != "success" && kubernetes.Status != "" {
		return nil, errors.New(kubernetes.Message)
	}

	return &kubernetes, nil
}

func (k *KubernetesService) Read(ctx context.Context, clusterId int) (*KubernetesCluster, error) {
	reqUrl := fmt.Sprintf("kubernetes/%d", clusterId)
	req, _ := k.client.NewRequest("GET", reqUrl)

	var kubernetes KubernetesCluster
	_, err := k.client.Do(req, &kubernetes)
	if err != nil {
		return nil, err
	}
	if kubernetes.Info.Cluster.ID != clusterId {
		return nil, errors.New("sorry we unable to find this cluster or you dont have access")
	}

	return &kubernetes, nil
}

func (k *KubernetesService) List(ctx context.Context) ([]K8s, error) {
	reqUrl := "kubernetes"
	req, _ := k.client.NewRequest("GET", reqUrl)

	var kubernetes Kubernetes
	_, err := k.client.Do(req, &kubernetes)
	if err != nil {
		return nil, err
	}
	if kubernetes.Status != "success" && kubernetes.Status != "" {
		return nil, errors.New(kubernetes.Message)
	}

	return kubernetes.K8s, nil
}

type DeleteKubernetesParams struct {
	ClusterId int
	// confirm message"I am aware this action will delete data and cluster permanently"
	Confirm string `json:"confirm"`
}

func (k *KubernetesService) Delete(ctx context.Context, params DeleteKubernetesParams) (*DeleteResponse, error) {
	reqUrl := fmt.Sprintf("kubernetes/%d/destroy", params.ClusterId)
	req, _ := k.client.NewRequest("DELETE", reqUrl)

	var delResponse DeleteResponse
	if _, err := k.client.Do(req, &delResponse); err != nil {
		return nil, err
	}
	if delResponse.Status != "success" && delResponse.Status != "" {
		return nil, errors.New(delResponse.Message)
	}

	return &delResponse, nil
}

type CreateKubernetesLoadbalancerParams struct {
	ClusterId      int
	LoadbalancerId int
}

func (k *KubernetesService) CreateLoadbalancer(ctx context.Context, params CreateKubernetesLoadbalancerParams) (*CreateResponse, error) {
	reqUrl := fmt.Sprintf("kubernetes/%d/loadbalancer/%d", params.ClusterId, params.LoadbalancerId)
	req, _ := k.client.NewRequest("POST", reqUrl, &params)

	var kubernetes CreateResponse
	_, err := k.client.Do(req, &kubernetes)
	if err != nil {
		return nil, err
	}
	if kubernetes.Status != "success" && kubernetes.Status != "" {
		return nil, errors.New(kubernetes.Message)
	}

	return &kubernetes, nil
}

func (k *KubernetesService) ReadLoadbalancer(ctx context.Context, clusterId, loadbalancerId int) (*K8sLoadbalancers, error) {
	reqUrl := fmt.Sprintf("kubernetes/%d", clusterId)
	req, _ := k.client.NewRequest("GET", reqUrl)

	var kubernetess KubernetesCluster
	_, err := k.client.Do(req, &kubernetess)
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
	if loadbalancers.ID == 0 {
		return nil, errors.New("kubernetess loadbalancer not found")
	}

	return &loadbalancers, nil
}

func (k *KubernetesService) ListLoadbalancers(ctx context.Context, clusterId int) ([]K8sLoadbalancers, error) {
	reqUrl := fmt.Sprintf("kubernetes/%d", clusterId)
	req, _ := k.client.NewRequest("GET", reqUrl)

	var kubernetess KubernetesCluster
	_, err := k.client.Do(req, &kubernetess)
	if err != nil {
		return nil, err
	}
	if kubernetess.Status != "success" && kubernetess.Status != "" {
		return nil, errors.New(kubernetess.Message)
	}

	return kubernetess.LoadBalancers, nil
}

func (k *KubernetesService) DeleteLoadbalancer(ctx context.Context, clusterId, kubernetesLoadbalancerId int) (*DeleteResponse, error) {
	reqUrl := fmt.Sprintf("kubernetes/%d/loadbalancer/%d", clusterId, kubernetesLoadbalancerId)
	req, _ := k.client.NewRequest("DELETE", reqUrl)

	var delResponse DeleteResponse
	if _, err := k.client.Do(req, &delResponse); err != nil {
		return nil, err
	}
	if delResponse.Status != "success" && delResponse.Status != "" {
		return nil, errors.New(delResponse.Message)
	}

	return &delResponse, nil
}

type CreateKubernetesSecurityGroupParams struct {
	ClusterId                 int
	KubernetesSecurityGroupId int
}

func (k *KubernetesService) CreateSecurityGroup(ctx context.Context, params CreateKubernetesSecurityGroupParams) (*CreateResponse, error) {
	reqUrl := fmt.Sprintf("kubernetes/%d", params.ClusterId)
	req, _ := k.client.NewRequest("POST", reqUrl, &params)

	var kubernetes CreateResponse
	_, err := k.client.Do(req, &kubernetes)
	if err != nil {
		return nil, err
	}
	if kubernetes.Status != "success" && kubernetes.Status != "" {
		return nil, errors.New(kubernetes.Message)
	}

	return &kubernetes, nil
}

func (k *KubernetesService) ReadSecurityGroup(ctx context.Context, clusterId, securitygroupId int) (*K8sSecurityGroups, error) {
	reqUrl := fmt.Sprintf("kubernetes/%d", clusterId)
	req, _ := k.client.NewRequest("GET", reqUrl)

	var kubernetess KubernetesCluster
	_, err := k.client.Do(req, &kubernetess)
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
	if securitygroups.ID == 0 {
		return nil, errors.New("kubernetess securitygroup not found")
	}

	return &securitygroups, nil
}

func (k *KubernetesService) ListSecurityGroups(ctx context.Context, clusterId int) ([]K8sSecurityGroups, error) {
	reqUrl := fmt.Sprintf("kubernetes/%d", clusterId)
	req, _ := k.client.NewRequest("GET", reqUrl)

	var kubernetess KubernetesCluster
	_, err := k.client.Do(req, &kubernetess)
	if err != nil {
		return nil, err
	}
	if kubernetess.Status != "success" && kubernetess.Status != "" {
		return nil, errors.New(kubernetess.Message)
	}

	return kubernetess.SecurityGroups, nil
}

func (k *KubernetesService) DeleteSecurityGroup(ctx context.Context, clusterId, kubernetesSecurityGroupId int) (*DeleteResponse, error) {
	reqUrl := fmt.Sprintf("kubernetes/%d/securitygroup/%d", clusterId, kubernetesSecurityGroupId)
	req, _ := k.client.NewRequest("DELETE", reqUrl)

	var delResponse DeleteResponse
	if _, err := k.client.Do(req, &delResponse); err != nil {
		return nil, err
	}
	if delResponse.Status != "success" && delResponse.Status != "" {
		return nil, errors.New(delResponse.Message)
	}

	return &delResponse, nil
}

type CreateKubernetesTargetgroupParams struct {
	ClusterId               int
	KubernetesTargetgroupId int
}

func (k *KubernetesService) CreateTargetgroup(ctx context.Context, params CreateKubernetesTargetgroupParams) (*CreateResponse, error) {
	reqUrl := fmt.Sprintf("kubernetes/%d/targetgroup/%d", params.ClusterId, params.KubernetesTargetgroupId)
	req, _ := k.client.NewRequest("POST", reqUrl, &params)

	var kubernetes CreateResponse
	_, err := k.client.Do(req, &kubernetes)
	if err != nil {
		return nil, err
	}
	if kubernetes.Status != "success" && kubernetes.Status != "" {
		return nil, errors.New(kubernetes.Message)
	}

	return &kubernetes, nil
}

func (k *KubernetesService) ReadTargetgroup(ctx context.Context, clusterId, targetgroupId int) (*K8sTargetGroups, error) {
	reqUrl := fmt.Sprintf("kubernetes/%d", clusterId)
	req, _ := k.client.NewRequest("GET", reqUrl)

	var kubernetess KubernetesCluster
	_, err := k.client.Do(req, &kubernetess)
	if err != nil {
		return nil, err
	}
	if kubernetess.Status != "success" && kubernetess.Status != "" {
		return nil, errors.New(kubernetess.Message)
	}

	if kubernetess.Info.Cluster.ID == 0 {
		return nil, errors.New("no Cluster Found")
	}
	var targetgroups K8sTargetGroups
	for _, tg := range kubernetess.TargetGroups {
		if tg.ID == targetgroupId {
			targetgroups = tg
		}
	}
	if targetgroups.ID == 0 {
		return nil, errors.New("kubernetess targetgroup not found")
	}

	return &targetgroups, nil
}

func (k *KubernetesService) ListTargetgroups(ctx context.Context, clusterId int) ([]K8sTargetGroups, error) {
	reqUrl := fmt.Sprintf("kubernetes/%d", clusterId)
	req, _ := k.client.NewRequest("GET", reqUrl)

	var kubernetess KubernetesCluster
	_, err := k.client.Do(req, &kubernetess)
	if err != nil {
		return nil, err
	}
	if kubernetess.Status != "success" && kubernetess.Status != "" {
		return nil, errors.New(kubernetess.Message)
	}

	return kubernetess.TargetGroups, nil
}

func (k *KubernetesService) DeleteTargetgroup(ctx context.Context, clusterId, kubernetesTargetgroupId int) (*DeleteResponse, error) {
	reqUrl := fmt.Sprintf("kubernetes/%d/targetgroup/%d", clusterId, kubernetesTargetgroupId)
	req, _ := k.client.NewRequest("DELETE", reqUrl)

	var delResponse DeleteResponse
	if _, err := k.client.Do(req, &delResponse); err != nil {
		return nil, err
	}
	if delResponse.Status != "success" && delResponse.Status != "" {
		return nil, errors.New(delResponse.Message)
	}

	return &delResponse, nil
}

func (k *KubernetesService) PowerOff(ctx context.Context, clusterId int) (*BasicResponse, error) {
	reqUrl := fmt.Sprintf("kubernetes/%d/stop", clusterId)
	req, _ := k.client.NewRequest("POST", reqUrl)

	var basicResponse BasicResponse
	_, err := k.client.Do(req, &basicResponse)
	if err != nil {
		return nil, err
	}
	if basicResponse.Status != "success" && basicResponse.Status != "" {
		return nil, errors.New(basicResponse.Message)
	}

	return &basicResponse, nil
}

func (k *KubernetesService) PowerOn(ctx context.Context, clusterId int) (*BasicResponse, error) {
	reqUrl := fmt.Sprintf("kubernetes/%d/start", clusterId)
	req, _ := k.client.NewRequest("POST", reqUrl)

	var basicResponse BasicResponse
	_, err := k.client.Do(req, &basicResponse)
	if err != nil {
		return nil, err
	}
	if basicResponse.Status != "success" && basicResponse.Status != "" {
		return nil, errors.New(basicResponse.Message)
	}

	return &basicResponse, nil
}

// NodePool
func (k *KubernetesService) CreateNodePool(ctx context.Context, params CreateKubernetesNodePoolParams) (*CreateResponse, error) {
	reqUrl := fmt.Sprintf("kubernetes/%d/nodepool/add", params.ClusterId)
	req, _ := k.client.NewRequest("POST", reqUrl, &params)

	var kubernetes CreateResponse
	_, err := k.client.Do(req, &kubernetes)
	if err != nil {
		return nil, err
	}
	if kubernetes.Status != "success" && kubernetes.Status != "" {
		return nil, errors.New(kubernetes.Message)
	}

	return &kubernetes, nil
}

func (k *KubernetesService) ReadNodePool(ctx context.Context, clusterId int, nodePoolId string) (*NodepoolDetails, error) {
	reqUrl := fmt.Sprintf("kubernetes/%d", clusterId)
	req, _ := k.client.NewRequest("GET", reqUrl)

	var kubernetess KubernetesCluster
	_, err := k.client.Do(req, &kubernetess)
	if err != nil {
		return nil, err
	}
	if kubernetess.Status != "success" && kubernetess.Status != "" {
		return nil, errors.New(kubernetess.Message)
	}

	if kubernetess.Info.Cluster.ID == 0 {
		return nil, errors.New("no Cluster Found")
	}
	var nodepools NodepoolDetails
	for id, np := range kubernetess.Nodepools {
		if id == nodePoolId {
			np.ID = id
			nodepools = np
		}
	}
	if len(nodepools.ID) == 0 {
		return nil, errors.New("kubernetess NodePool not found")
	}

	return &nodepools, nil
}

func (k *KubernetesService) ListNodePools(ctx context.Context, clusterId int) ([]NodepoolDetails, error) {
	reqUrl := fmt.Sprintf("kubernetes/%d", clusterId)
	req, _ := k.client.NewRequest("GET", reqUrl)

	var kubernetess KubernetesCluster
	_, err := k.client.Do(req, &kubernetess)
	if err != nil {
		return nil, err
	}
	if kubernetess.Status != "success" && kubernetess.Status != "" {
		return nil, errors.New(kubernetess.Message)
	}

	nodepools := make([]NodepoolDetails, 0, len(kubernetess.Nodepools))
	for id, np := range kubernetess.Nodepools {
		np.ID = id
		nodepools = append(nodepools, np)
	}
	return nodepools, nil
}

type UpdateKubernetesAutoscaleNodepool struct {
	ClusterId  int
	NodePoolId string
	Count      int    `json:"count"`
	Label      string `json:"label,omitempty"`
	PoolType   string `json:"pool_type,omitempty"`
	Size       int    `json:"size,omitempty"`
	Policies   string `json:"policies,omitempty"`
	MinNodes   int    `json:"min_nodes,omitempty"`
	MaxNodes   int    `json:"max_nodes,omitempty"`
}

type UpdateKubernetesAutoscaleNodepoolResponse struct {
	ID        string        `json:"id"`
	Size      string        `json:"size"`
	Cost      float64       `json:"cost"`
	Planid    string        `json:"planid"`
	Count     int           `json:"count"`
	AutoScale bool          `json:"auto_scale"`
	MinNodes  int           `json:"min_nodes"`
	MaxNodes  int           `json:"max_nodes"`
	Policies  []interface{} `json:"policies"`
	Workers   []WorkerNode  `json:"workers"`

	Status  string `json:"status"`
	Message string `json:"message"`
}

// UpdateAutoscaleNodepool
func (k *KubernetesService) UpdateNodePool(ctx context.Context, params UpdateKubernetesAutoscaleNodepool) (*UpdateKubernetesAutoscaleNodepoolResponse, error) {
	reqUrl := fmt.Sprintf("kubernetes/%d/nodepool/%s/update", params.ClusterId, params.NodePoolId)
	req, _ := k.client.NewRequest("POST", reqUrl)

	var kubernetes UpdateKubernetesAutoscaleNodepoolResponse
	_, err := k.client.Do(req, &kubernetes)
	if err != nil {
		return nil, err
	}
	if kubernetes.Status != "success" && kubernetes.Status != "" {
		return nil, errors.New(kubernetes.Message)
	}

	return &kubernetes, nil
}

type UpdateKubernetesStaticNodepool struct {
	ClusterId  int
	NodePoolId int
	Count      string `json:"count"`
	Label      string `json:"label"`
	PoolType   string `json:"pool_type"`
	Size       string `json:"size"`
}

func (k *KubernetesService) UpdateStaticNodepool(ctx context.Context, params UpdateKubernetesStaticNodepool) (*UpdateResponse, error) {
	reqUrl := fmt.Sprintf("kubernetes/%d/nodepool/%d/update", params.ClusterId, params.NodePoolId)
	req, _ := k.client.NewRequest("POST", reqUrl)

	var kubernetes UpdateResponse
	_, err := k.client.Do(req, &kubernetes)
	if err != nil {
		return nil, err
	}
	if kubernetes.Status != "success" && kubernetes.Status != "" {
		return nil, errors.New(kubernetes.Message)
	}

	return &kubernetes, nil
}

type DeleteNodeParams struct {
	ClusterId int
	PoolId    string
	NodeId    string
}

func (k *KubernetesService) DeleteNode(ctx context.Context, params DeleteNodeParams) (*DeleteResponse, error) {
	reqUrl := fmt.Sprintf("kubernetes/%d/nodepool/%s/node/%s", params.ClusterId, params.PoolId, params.NodeId)
	req, _ := k.client.NewRequest("DELETE", reqUrl)

	var delResponse DeleteResponse
	if _, err := k.client.Do(req, &delResponse); err != nil {
		return nil, err
	}
	if delResponse.Status != "success" && delResponse.Status != "" {
		return nil, errors.New(delResponse.Message)
	}

	return &delResponse, nil
}
