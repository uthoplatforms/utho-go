package utho

import (
	"errors"
)

type KubernetesService service

type Kubernetess struct {
	Status  string `json:"status"`
	Message string `json:"message"`
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

func (s *KubernetesService) CreateKubernetes(params CreateKubernetesParams) (*CreateResponse, error) {
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

// func (s *KubernetesService) ReadKubernetes(certId string) (*Certificates, error) {
// 	reqUrl := "certificates"
// 	req, _ := s.client.NewRequest("GET", reqUrl)

// 	var kubernetes Kubernetess
// 	_, err := s.client.Do(req, &kubernetes)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if kubernetes.Status != "success" && kubernetes.Status != "" {
// 		return nil, errors.New(kubernetes.Message)
// 	}

// 	var cert Certificates
// 	for _, r := range kubernetes.Certificates {
// 		if r.ID == certId {
// 			cert = r
// 		}
// 	}
// 	if len(cert.ID) == 0 {
// 		return nil, errors.New("certificate not found")
// 	}

// 	return &cert, nil
// }

// func (s *KubernetesService) ListKubernetess() (*[]Certificates, error) {
// 	reqUrl := "certificates"
// 	req, _ := s.client.NewRequest("GET", reqUrl)

// 	var kubernetes Kubernetess
// 	_, err := s.client.Do(req, &kubernetes)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if kubernetes.Status != "success" && kubernetes.Status != "" {
// 		return nil, errors.New(kubernetes.Message)
// 	}

// 	return &kubernetes.Certificates, nil
// }

type DeleteKubernetesParams struct {
	ClusterId string
	// confirm message"I am aware this action will delete data and cluster permanently"
	Confirm string `json:"confirm"`
}

func (s *KubernetesService) DeleteKubernetes(params DeleteKubernetesParams) (*DeleteResponse, error) {
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

func (s *KubernetesService) CreateKubernetesLoadbalancer(params CreateKubernetesLoadbalancerParams) (*CreateResponse, error) {
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

// func (s *KubernetesService) ReadKubernetesLoadbalancer(kubernetesId, loadbalancerId string) (*Loadbalancers, error) {
// 	reqUrl := "kubernetes/" + kubernetesId
// 	req, _ := s.client.NewRequest("GET", reqUrl)

// 	var kubernetess Kubernetess
// 	_, err := s.client.Do(req, &kubernetess)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if kubernetess.Status != "success" && kubernetess.Status != "" {
// 		return nil, errors.New(kubernetess.Message)
// 	}
// 	var loadbalancers Loadbalancers
// 	for _, r := range kubernetess.Groups[0].LoadBalancers {
// 		if r.ID == loadbalancerId {
// 			loadbalancers = r
// 		}
// 	}
// 	if len(loadbalancers.ID) == 0 {
// 		return nil, errors.New("auto scaling loadbalancer not found")
// 	}

// 	return &loadbalancers, nil
// }

// func (s *KubernetesService) ListKubernetesLoadbalancer(kubernetesId string) (*[]Loadbalancers, error) {
// 	reqUrl := "kubernetes/" + kubernetesId
// 	req, _ := s.client.NewRequest("GET", reqUrl)

// 	var kubernetess Kubernetess
// 	_, err := s.client.Do(req, &kubernetess)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if kubernetess.Status != "success" && kubernetess.Status != "" {
// 		return nil, errors.New(kubernetess.Message)
// 	}

// 	return &kubernetess.Groups[0].Loadbalancers, nil
// }

func (s *KubernetesService) DeleteKubernetesLoadbalancer(kuberneteseId, kubernetesLoadbalancerId string) (*DeleteResponse, error) {
	reqUrl := "kubernetes/" + kuberneteseId + "/loadbalancerpolicy/" + kubernetesLoadbalancerId
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

func (s *KubernetesService) CreateKubernetesSecurityGroup(params CreateKubernetesSecurityGroupParams) (*CreateResponse, error) {
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

// func (s *KubernetesService) ReadKubernetesSecurityGroup(kubernetesId, securitygroupId string) (*SecurityGroups, error) {
// 	reqUrl := "kubernetes/" + kubernetesId
// 	req, _ := s.client.NewRequest("GET", reqUrl)

// 	var kubernetess Kubernetess
// 	_, err := s.client.Do(req, &kubernetess)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if kubernetess.Status != "success" && kubernetess.Status != "" {
// 		return nil, errors.New(kubernetess.Message)
// 	}
// 	var securitygroups SecurityGroups
// 	for _, r := range kubernetess.Groups[0].SecurityGroups {
// 		if r.ID == securitygroupId {
// 			securitygroups = r
// 		}
// 	}
// 	if len(securitygroups.ID) == 0 {
// 		return nil, errors.New("auto scaling securitygroup not found")
// 	}

// 	return &securitygroups, nil
// }

// func (s *KubernetesService) ListKubernetesSecurityGroup(kubernetesId string) (*[]SecurityGroups, error) {
// 	reqUrl := "kubernetes/" + kubernetesId
// 	req, _ := s.client.NewRequest("GET", reqUrl)

// 	var kubernetess Kubernetess
// 	_, err := s.client.Do(req, &kubernetess)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if kubernetess.Status != "success" && kubernetess.Status != "" {
// 		return nil, errors.New(kubernetess.Message)
// 	}

// 	return &kubernetess.Groups[0].SecurityGroups, nil
// }

func (s *KubernetesService) DeleteKubernetesSecurityGroup(kuberneteseId, kubernetesSecurityGroupId string) (*DeleteResponse, error) {
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

func (s *KubernetesService) CreateKubernetesTargetgroup(params CreateKubernetesTargetgroupParams) (*CreateResponse, error) {
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

// func (s *KubernetesService) ReadKubernetesTargetgroup(kubernetesId, targetgroupId string) (*KubernetesTargetGroups, error) {
// 	reqUrl := "kubernetes/" + kubernetesId
// 	req, _ := s.client.NewRequest("GET", reqUrl)

// 	var kubernetess Kubernetess
// 	_, err := s.client.Do(req, &kubernetess)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if kubernetess.Status != "success" && kubernetess.Status != "" {
// 		return nil, errors.New(kubernetess.Message)
// 	}
// 	var targetgroups KubernetesTargetGroups
// 	for _, r := range kubernetess.Groups[0].TargetGroups {
// 		if r.ID == targetgroupId {
// 			targetgroups = r
// 		}
// 	}
// 	if len(targetgroups.ID) == 0 {
// 		return nil, errors.New("auto scaling targetgroup not found")
// 	}

// 	return &targetgroups, nil
// }

// func (s *KubernetesService) ListKubernetesTargetgroup(kubernetesId string) (*[]KubernetesTargetGroups, error) {
// 	reqUrl := "kubernetes/" + kubernetesId
// 	req, _ := s.client.NewRequest("GET", reqUrl)

// 	var kubernetess Kubernetess
// 	_, err := s.client.Do(req, &kubernetess)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if kubernetess.Status != "success" && kubernetess.Status != "" {
// 		return nil, errors.New(kubernetess.Message)
// 	}

// 	return &kubernetess.Groups[0].TargetGroups, nil
// }

func (s *KubernetesService) DeleteKubernetesTargetgroup(kuberneteseId, kubernetesTargetgroupId string) (*DeleteResponse, error) {
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

func (s *KubernetesService) UpdateKubernetesAutoscaleNodepool(params UpdateKubernetesAutoscaleNodepool) (*UpdateResponse, error) {
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

func (s *KubernetesService) UpdateKubernetesStaticNodepool(params UpdateKubernetesStaticNodepool) (*UpdateResponse, error) {
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
