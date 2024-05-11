package utho

import (
	"errors"
)

type TargetGroupService service

type TargetGroups struct {
	Status       string        `json:"status"`
	Message      string        `json:"message"`
	Targetgroups []TargetGroup `json:"targetgroups"`
}
type TargetGroup struct {
	ID                  string   `json:"id"`
	Name                string   `json:"name"`
	Port                string   `json:"port"`
	Protocol            string   `json:"protocol"`
	HealthCheckPath     string   `json:"health_check_path"`
	HealthCheckInterval string   `json:"health_check_interval"`
	HealthCheckProtocol string   `json:"health_check_protocol"`
	HealthCheckTimeout  string   `json:"health_check_timeout"`
	HealthyThreshold    string   `json:"healthy_threshold"`
	UnhealthyThreshold  string   `json:"unhealthy_threshold"`
	CreatedAt           string   `json:"created_at"`
	UpdatedAt           string   `json:"updated_at"`
	Targets             []Target `json:"targets"`
}
type Target struct {
	Lbid                string `json:"lbid"`
	IP                  string `json:"ip"`
	Cloudid             string `json:"cloudid"`
	Status              string `json:"status"`
	ScalingGroupid      string `json:"scaling_groupid"`
	KubernetesClusterid string `json:"kubernetes_clusterid"`
	BackendPort         string `json:"backend_port"`
	BackendProtocol     string `json:"backend_protocol"`
	TargetgroupID       string `json:"targetgroup_id"`
	FrontendID          string `json:"frontend_id"`
	ID                  string `json:"id"`
}

type CreateTargetGroupParams struct {
	Name                string `json:"name"`
	Protocol            string `json:"protocol"`
	Port                string `json:"port"`
	HealthCheckPath     string `json:"health_check_path"`
	HealthCheckProtocol string `json:"health_check_protocol"`
	HealthCheckInterval string `json:"health_check_interval"`
	HealthCheckTimeout  string `json:"health_check_timeout"`
	HealthyThreshold    string `json:"healthy_threshold"`
	UnhealthyThreshold  string `json:"unhealthy_threshold"`
}

func (s *TargetGroupService) Create(params CreateTargetGroupParams) (*CreateResponse, error) {
	reqUrl := "targetgroup"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var targetgroup CreateResponse
	_, err := s.client.Do(req, &targetgroup)
	if err != nil {
		return nil, err
	}
	if targetgroup.Status != "success" && targetgroup.Status != "" {
		return nil, errors.New(targetgroup.Message)
	}

	return &targetgroup, nil
}

func (s *TargetGroupService) Read(targetGroupId string) (*TargetGroup, error) {
	reqUrl := "targetgroup"
	req, _ := s.client.NewRequest("GET", reqUrl)

	var targetgroup TargetGroups
	_, err := s.client.Do(req, &targetgroup)
	if err != nil {
		return nil, err
	}
	if targetgroup.Status != "success" && targetgroup.Status != "" {
		return nil, errors.New(targetgroup.Message)
	}

	var targetGroup TargetGroup
	for _, r := range targetgroup.Targetgroups {
		if r.ID == targetGroupId {
			targetGroup = r
		}
	}
	if len(targetGroup.ID) == 0 {
		return nil, errors.New("target groupId not found")
	}

	return &targetGroup, nil
}

func (s *TargetGroupService) List() ([]TargetGroup, error) {
	reqUrl := "targetgroup"
	req, _ := s.client.NewRequest("GET", reqUrl)

	var targetgroups TargetGroups
	_, err := s.client.Do(req, &targetgroups)
	if err != nil {
		return nil, err
	}
	if targetgroups.Status != "success" && targetgroups.Status != "" {
		return nil, errors.New(targetgroups.Message)
	}

	return targetgroups.Targetgroups, nil
}

type UpdateTargetGroupParams struct {
	TargetGroupId       string
	Name                string `json:"name"`
	Protocol            string `json:"protocol"`
	Port                string `json:"port"`
	HealthCheckPath     string `json:"health_check_path"`
	HealthCheckProtocol string `json:"health_check_protocol"`
	HealthCheckInterval string `json:"health_check_interval"`
	HealthCheckTimeout  string `json:"health_check_timeout"`
	HealthyThreshold    string `json:"healthy_threshold"`
	UnhealthyThreshold  string `json:"unhealthy_threshold"`
}

func (s *TargetGroupService) Update(params UpdateTargetGroupParams) (*UpdateResponse, error) {
	reqUrl := "targetgroup/" + params.TargetGroupId
	req, _ := s.client.NewRequest("PUT", reqUrl, &params)

	var targetgroup UpdateResponse
	_, err := s.client.Do(req, &targetgroup)
	if err != nil {
		return nil, err
	}
	if targetgroup.Status != "success" && targetgroup.Status != "" {
		return nil, errors.New(targetgroup.Message)
	}

	return &targetgroup, nil
}

func (s *TargetGroupService) Delete(targetGroupId, targetGroupName string) (*DeleteResponse, error) {
	reqUrl := "targetgroup/" + targetGroupId + "?name=" + targetGroupName
	// targetgroup/:id?name=target_group_name
	req, _ := s.client.NewRequest("DELETE", reqUrl)

	var delResponse DeleteResponse
	if _, err := s.client.Do(req, &delResponse); err != nil {
		return nil, err
	}

	return &delResponse, nil
}

type CreateTargetGroupTargetParams struct {
	TargetGroupId   string
	BackendProtocol string `json:"backend_protocol"`
	BackendPort     string `json:"backend_port"`
	IP              string `json:"ip"`
	Cloudid         string `json:"cloudid,omitempty"`
}

func (s *TargetGroupService) CreateTarget(params CreateTargetGroupTargetParams) (*CreateResponse, error) {
	reqUrl := "targetgroup/" + params.TargetGroupId + "/target"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var targetgroup CreateResponse
	_, err := s.client.Do(req, &targetgroup)
	if err != nil {
		return nil, err
	}
	if targetgroup.Status != "success" && targetgroup.Status != "" {
		return nil, errors.New(targetgroup.Message)
	}

	return &targetgroup, nil
}

func (s *TargetGroupService) ReadTarget(targetGroupId, targetId string) (*Target, error) {
	reqUrl := "targetgroup"
	req, _ := s.client.NewRequest("GET", reqUrl)

	var targetgroup TargetGroups
	_, err := s.client.Do(req, &targetgroup)
	if err != nil {
		return nil, err
	}
	if targetgroup.Status != "success" && targetgroup.Status != "" {
		return nil, errors.New(targetgroup.Message)
	}

	var targetGroup TargetGroup
	for _, r := range targetgroup.Targetgroups {
		if r.ID == targetGroupId {
			targetGroup = r
		}
	}
	if len(targetGroup.ID) == 0 {
		return nil, errors.New("target groupId not found")
	}

	var target Target
	for _, r := range targetGroup.Targets {
		if r.ID == targetId {
			target = r
		}
	}
	if len(target.ID) == 0 {
		return nil, errors.New("target groupId not found")
	}

	return &target, nil
}

func (s *TargetGroupService) ListTarget(targetGroupId string) ([]Target, error) {
	reqUrl := "targetgroup"
	req, _ := s.client.NewRequest("GET", reqUrl)

	var targetgroups TargetGroups
	_, err := s.client.Do(req, &targetgroups)
	if err != nil {
		return nil, err
	}
	if targetgroups.Status != "success" && targetgroups.Status != "" {
		return nil, errors.New(targetgroups.Message)
	}

	var targetGroup TargetGroup
	for _, r := range targetgroups.Targetgroups {
		if r.ID == targetGroupId {
			targetGroup = r
		}
	}
	if len(targetGroup.ID) == 0 {
		return nil, errors.New("target groupId not found")
	}

	return targetGroup.Targets, nil
}

func (s *TargetGroupService) DeleteTarget(targetGroupId, targetId string) (*DeleteResponse, error) {
	reqUrl := "targetgroup/" + targetGroupId + "/target/" + targetId
	req, _ := s.client.NewRequest("DELETE", reqUrl)

	var delResponse DeleteResponse
	if _, err := s.client.Do(req, &delResponse); err != nil {
		return nil, err
	}

	return &delResponse, nil
}
