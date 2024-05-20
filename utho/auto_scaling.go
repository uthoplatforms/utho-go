package utho

import (
	"errors"
	"time"
)

type AutoScalingService service

type AutoScalings struct {
	Groups  []Groups `json:"groups"`
	Status  string   `json:"status"`
	Message string   `json:"message"`
}
type Groups struct {
	ID                 string                     `json:"id"`
	Userid             string                     `json:"userid"`
	Name               string                     `json:"name"`
	Dcslug             string                     `json:"dcslug"`
	Minsize            string                     `json:"minsize"`
	Maxsize            string                     `json:"maxsize"`
	Desiredsize        string                     `json:"desiredsize"`
	Planid             string                     `json:"planid"`
	Planname           string                     `json:"planname"`
	InstanceTemplateid string                     `json:"instance_templateid"`
	Image              string                     `json:"image"`
	ImageName          string                     `json:"image_name"`
	Snapshotid         string                     `json:"snapshotid"`
	Status             string                     `json:"status"`
	CreatedAt          string                     `json:"created_at"`
	SuspendedAt        string                     `json:"suspended_at"`
	StoppedAt          string                     `json:"stopped_at"`
	StartedAt          string                     `json:"started_at"`
	DeletedAt          string                     `json:"deleted_at"`
	PublicIPEnabled    string                     `json:"public_ip_enabled"`
	Vpc                []AutoScalingVpc           `json:"vpc"`
	CooldownTill       string                     `json:"cooldown_till"`
	Loadbalancers      []AutoScalingLoadbalancers `json:"load_balancers"`
	TargetGroups       []AutoScalingTargetGroup   `json:"target_groups"`
	SecurityGroups     []SecurityGroup            `json:"security_groups"`
	Backupid           string                     `json:"backupid"`
	Stack              string                     `json:"stack"`
	StackFields        string                     `json:"stack_fields"`
	Instances          []Instances                `json:"instances"`
	Policies           []Policy                   `json:"policies"`
	Schedules          []Schedule                 `json:"schedules"`
	DeletedInstances   []any                      `json:"deleted_instances"`
	Dclocation         Dclocation                 `json:"dclocation"`
	Plan               AutoScalingPlan            `json:"plan"`
}
type AutoScalingVpc struct {
	Total      int                    `json:"total"`
	Available  int                    `json:"available"`
	Network    string                 `json:"network"`
	Name       string                 `json:"name"`
	Size       string                 `json:"size"`
	Dcslug     string                 `json:"dcslug"`
	Dclocation Dclocation             `json:"dclocation"`
	IsDefault  any                    `json:"is_default"`
	Resources  []AutoScalingResources `json:"resources"`
}
type AutoScalingResources struct {
	Type string `json:"type"`
	ID   string `json:"id"`
	Name string `json:"name"`
	IP   string `json:"ip"`
}
type AutoScalingLoadbalancers struct {
	ID   string `json:"lbid"`
	Name string `json:"name"`
	IP   string `json:"ip"`
}
type AutoScalingTargetGroup struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Protocol string `json:"protocol"`
	Port     string `json:"port"`
}
type SecurityGroup struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type Instances struct {
	Cloudid   string `json:"cloudid"`
	Hostname  string `json:"hostname"`
	CreatedAt string `json:"created_at"`
	IP        string `json:"ip"`
	Status    string `json:"status"`
}
type Policy struct {
	ID                 string `json:"id"`
	Userid             string `json:"userid"`
	Product            string `json:"product"`
	Productid          string `json:"productid"`
	Groupid            string `json:"groupid"`
	Name               string `json:"name"`
	Type               string `json:"type"`
	Adjust             string `json:"adjust"`
	Period             string `json:"period"`
	Cooldown           string `json:"cooldown"`
	CooldownTill       string `json:"cooldown_till"`
	Compare            string `json:"compare"`
	Value              string `json:"value"`
	AlertID            string `json:"alert_id"`
	Status             string `json:"status"`
	KubernetesID       string `json:"kubernetes_id"`
	KubernetesNodepool string `json:"kubernetes_nodepool"`
	Cloudid            string `json:"cloudid"`
	Maxsize            string `json:"maxsize"`
	Minsize            string `json:"minsize"`
}
type Schedule struct {
	ID          string `json:"id"`
	Groupid     string `json:"groupid"`
	Name        string `json:"name"`
	Desiredsize string `json:"desiredsize"`
	Recurrence  string `json:"recurrence"`
	StartDate   string `json:"start_date"`
	Status      string `json:"status"`
	Timezone    string `json:"timezone"`
}
type AutoScalingPlan struct {
	Planid         string `json:"planid"`
	RAM            string `json:"ram"`
	CPU            string `json:"cpu"`
	Disk           string `json:"disk"`
	Bandwidth      string `json:"bandwidth"`
	DedicatedVcore string `json:"dedicated_vcore"`
}

type CreateAutoScalingParams struct {
	Name               string                  `json:"name"`
	OsDiskSize         int                     `json:"os_disk_size"`
	Dcslug             string                  `json:"dcslug"`
	Minsize            string                  `json:"minsize"`
	Maxsize            string                  `json:"maxsize"`
	Desiredsize        string                  `json:"desiredsize"`
	Planid             string                  `json:"planid"`
	Planname           string                  `json:"planname"`
	InstanceTemplateid string                  `json:"instance_templateid"`
	PublicIPEnabled    bool                    `json:"public_ip_enabled"`
	Vpc                string                  `json:"vpc"`
	LoadBalancers      string                  `json:"load_balancers"`
	SecurityGroups     string                  `json:"security_groups"`
	Policies           []CreatePoliciesParams  `json:"policies"`
	Schedules          []CreateSchedulesParams `json:"schedules"`
	Stackid            string                  `json:"stack"`
	Stackimage         string                  `json:"stackimage"`
	TargetGroups       string                  `json:"target_groups"`
}
type CreatePoliciesParams struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Compare  string `json:"compare"`
	Value    string `json:"value"`
	Adjust   string `json:"adjust"`
	Period   string `json:"period"`
	Cooldown string `json:"cooldown"`
}
type CreateSchedulesParams struct {
	Name         string    `json:"name"`
	Desiredsize  string    `json:"desiredsize"`
	StartDate    time.Time `json:"start_date"`
	SelectedTime string    `json:"selectedTime"`
	SelectedDate string    `json:"selectedDate"`
}

type CreateAutoScalingResponse struct {
	ID      int    `json:"id"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (s *AutoScalingService) Create(params CreateAutoScalingParams) (*CreateAutoScalingResponse, error) {
	reqUrl := "autoscaling"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var autoscaling CreateAutoScalingResponse
	_, err := s.client.Do(req, &autoscaling)
	if err != nil {
		return nil, err
	}
	if autoscaling.Status != "success" && autoscaling.Status != "" {
		return nil, errors.New(autoscaling.Message)
	}

	return &autoscaling, nil
}

func (s *AutoScalingService) Read(autoscalingId string) (*Groups, error) {
	reqUrl := "autoscaling/" + autoscalingId
	req, _ := s.client.NewRequest("GET", reqUrl)

	var autoscalings AutoScalings
	_, err := s.client.Do(req, &autoscalings)
	if err != nil {
		return nil, err
	}
	if autoscalings.Status != "success" && autoscalings.Status != "" {
		return nil, errors.New(autoscalings.Message)
	}
	if len(autoscalings.Groups) == 0 {
		return nil, errors.New("NotFound")
	}

	return &autoscalings.Groups[0], nil
}

func (s *AutoScalingService) List() ([]Groups, error) {
	reqUrl := "autoscaling"
	req, _ := s.client.NewRequest("GET", reqUrl)

	var autoscalings AutoScalings
	_, err := s.client.Do(req, &autoscalings)
	if err != nil {
		return nil, err
	}
	if autoscalings.Status != "success" && autoscalings.Status != "" {
		return nil, errors.New(autoscalings.Message)
	}

	return autoscalings.Groups, nil
}

type UpdateAutoScalingParams struct {
	AutoScalingId string
	Name          string `json:"name"`
	Minsize       string `json:"minsize"`
	Maxsize       string `json:"maxsize"`
	Desiredsize   string `json:"desiredsize"`
}

func (s *AutoScalingService) Update(params UpdateAutoScalingParams) (*UpdateResponse, error) {
	reqUrl := "autoscaling/" + params.AutoScalingId
	req, _ := s.client.NewRequest("PUT", reqUrl, &params)

	var autoscaling UpdateResponse
	_, err := s.client.Do(req, &autoscaling)
	if err != nil {
		return nil, err
	}
	if autoscaling.Status != "success" && autoscaling.Status != "" {
		return nil, errors.New(autoscaling.Message)
	}

	return &autoscaling, nil
}

func (s *AutoScalingService) Delete(autoscalingId, autoscalingName string) (*DeleteResponse, error) {
	reqUrl := "autoscaling/" + autoscalingId + "?name=" + autoscalingName
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

// Auto Scaling Policy
type CreateAutoScalingPolicyParams struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Compare   string `json:"compare"`
	Value     string `json:"value"`
	Adjust    string `json:"adjust"`
	Period    string `json:"period"`
	Cooldown  string `json:"cooldown"`
	Product   string `json:"product"`
	Productid string `json:"productid"`
}

func (s *AutoScalingService) CreatePolicy(params CreateAutoScalingPolicyParams) (*CreateResponse, error) {
	reqUrl := "autoscaling/policy"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var autoscaling CreateResponse
	_, err := s.client.Do(req, &autoscaling)
	if err != nil {
		return nil, err
	}
	if autoscaling.Status != "success" && autoscaling.Status != "" {
		return nil, errors.New(autoscaling.Message)
	}

	return &autoscaling, nil
}

func (s *AutoScalingService) ReadPolicy(autoscalingId, policyId string) (*Policy, error) {
	reqUrl := "autoscaling/" + autoscalingId
	req, _ := s.client.NewRequest("GET", reqUrl)

	var autoscalings AutoScalings
	_, err := s.client.Do(req, &autoscalings)
	if err != nil {
		return nil, err
	}
	if autoscalings.Status != "success" && autoscalings.Status != "" {
		return nil, errors.New(autoscalings.Message)
	}
	var policies Policy
	for _, r := range autoscalings.Groups[0].Policies {
		if r.ID == policyId {
			policies = r
		}
	}
	if len(policies.ID) == 0 {
		return nil, errors.New("auto scaling policy not found")
	}

	return &policies, nil
}

func (s *AutoScalingService) ListPolicies(autoscalingId string) ([]Policy, error) {
	reqUrl := "autoscaling/" + autoscalingId
	req, _ := s.client.NewRequest("GET", reqUrl)

	var autoscalings AutoScalings
	_, err := s.client.Do(req, &autoscalings)
	if err != nil {
		return nil, err
	}
	if autoscalings.Status != "success" && autoscalings.Status != "" {
		return nil, errors.New(autoscalings.Message)
	}

	return autoscalings.Groups[0].Policies, nil
}

type UpdateAutoScalingPolicyParams struct {
	AutoScalingPolicyId string
	Name                string `json:"name"`
	Type                string `json:"type"`
	Compare             string `json:"compare"`
	Value               string `json:"value"`
	Adjust              string `json:"adjust"`
	Period              string `json:"period"`
	Cooldown            string `json:"cooldown"`
}

func (s *AutoScalingService) UpdatePolicy(params UpdateAutoScalingPolicyParams) (*UpdateResponse, error) {
	reqUrl := "autoscaling/policy/" + params.AutoScalingPolicyId
	req, _ := s.client.NewRequest("PUT", reqUrl, &params)

	var autoscaling UpdateResponse
	_, err := s.client.Do(req, &autoscaling)
	if err != nil {
		return nil, err
	}
	if autoscaling.Status != "success" && autoscaling.Status != "" {
		return nil, errors.New(autoscaling.Message)
	}

	return &autoscaling, nil
}

func (s *AutoScalingService) DeletePolicy(autoScalingPolicyId string) (*DeleteResponse, error) {
	reqUrl := "autoscaling/policy/" + autoScalingPolicyId
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

// Auto Scaling Schedule
type CreateAutoScalingScheduleParams struct {
	AutoScalingId string
	Name          string `json:"name"`
	Desiredsize   string `json:"desiredsize"`
	Recurrence    string `json:"recurrence"`
	StartDate     string `json:"start_date"`
}

func (s *AutoScalingService) CreateSchedule(params CreateAutoScalingScheduleParams) (*CreateResponse, error) {
	reqUrl := "autoscaling/" + params.AutoScalingId + "/schedulepolicy"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var autoscaling CreateResponse
	_, err := s.client.Do(req, &autoscaling)
	if err != nil {
		return nil, err
	}
	if autoscaling.Status != "success" && autoscaling.Status != "" {
		return nil, errors.New(autoscaling.Message)
	}

	return &autoscaling, nil
}

func (s *AutoScalingService) ReadSchedule(autoscalingId, scheduleId string) (*Schedule, error) {
	reqUrl := "autoscaling/" + autoscalingId
	req, _ := s.client.NewRequest("GET", reqUrl)

	var autoscalings AutoScalings
	_, err := s.client.Do(req, &autoscalings)
	if err != nil {
		return nil, err
	}
	if autoscalings.Status != "success" && autoscalings.Status != "" {
		return nil, errors.New(autoscalings.Message)
	}
	var schedules Schedule
	for _, r := range autoscalings.Groups[0].Schedules {
		if r.ID == scheduleId {
			schedules = r
		}
	}
	if len(schedules.ID) == 0 {
		return nil, errors.New("auto scaling schedule not found")
	}

	return &schedules, nil
}

func (s *AutoScalingService) ListSchedules(autoscalingId string) ([]Schedule, error) {
	reqUrl := "autoscaling/" + autoscalingId
	req, _ := s.client.NewRequest("GET", reqUrl)

	var autoscalings AutoScalings
	_, err := s.client.Do(req, &autoscalings)
	if err != nil {
		return nil, err
	}
	if autoscalings.Status != "success" && autoscalings.Status != "" {
		return nil, errors.New(autoscalings.Message)
	}

	return autoscalings.Groups[0].Schedules, nil
}

type UpdateAutoScalingScheduleParams struct {
	AutoScalingeId        string
	AutoScalingScheduleId string
	Name                  string `json:"name"`
	Desiredsize           string `json:"desiredsize"`
	Recurrence            string `json:"recurrence"`
	StartDate             string `json:"start_date"`
}

func (s *AutoScalingService) UpdateSchedule(params UpdateAutoScalingScheduleParams) (*UpdateResponse, error) {
	reqUrl := "autoscaling/" + params.AutoScalingeId + "/schedulepolicy/" + params.AutoScalingScheduleId
	req, _ := s.client.NewRequest("PUT", reqUrl, &params)

	var autoscaling UpdateResponse
	_, err := s.client.Do(req, &autoscaling)
	if err != nil {
		return nil, err
	}
	if autoscaling.Status != "success" && autoscaling.Status != "" {
		return nil, errors.New(autoscaling.Message)
	}

	return &autoscaling, nil
}

func (s *AutoScalingService) DeleteSchedule(autoScalingeId, autoScalingScheduleId string) (*DeleteResponse, error) {
	reqUrl := "autoscaling/" + autoScalingeId + "/schedulepolicy/" + autoScalingScheduleId
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

// Auto Scaling Loadbalancer
type CreateAutoScalingLoadbalancerParams struct {
	AutoScalingId  string
	LoadbalancerId string
}

func (s *AutoScalingService) CreateLoadbalancer(params CreateAutoScalingLoadbalancerParams) (*CreateResponse, error) {
	reqUrl := "autoscaling/" + params.AutoScalingId + "/loadbalancer/" + params.LoadbalancerId
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var autoscaling CreateResponse
	_, err := s.client.Do(req, &autoscaling)
	if err != nil {
		return nil, err
	}
	if autoscaling.Status != "success" && autoscaling.Status != "" {
		return nil, errors.New(autoscaling.Message)
	}

	return &autoscaling, nil
}

func (s *AutoScalingService) ReadLoadbalancer(autoscalingId, loadbalancerId string) (*AutoScalingLoadbalancers, error) {
	reqUrl := "autoscaling/" + autoscalingId
	req, _ := s.client.NewRequest("GET", reqUrl)

	var autoscalings AutoScalings
	_, err := s.client.Do(req, &autoscalings)
	if err != nil {
		return nil, err
	}
	if autoscalings.Status != "success" && autoscalings.Status != "" {
		return nil, errors.New(autoscalings.Message)
	}
	var loadbalancers AutoScalingLoadbalancers
	for _, r := range autoscalings.Groups[0].Loadbalancers {
		if r.ID == loadbalancerId {
			loadbalancers = r
		}
	}
	if len(loadbalancers.ID) == 0 {
		return nil, errors.New("auto scaling loadbalancer not found")
	}

	return &loadbalancers, nil
}

func (s *AutoScalingService) ListLoadbalancers(autoscalingId string) ([]AutoScalingLoadbalancers, error) {
	reqUrl := "autoscaling/" + autoscalingId
	req, _ := s.client.NewRequest("GET", reqUrl)

	var autoscalings AutoScalings
	_, err := s.client.Do(req, &autoscalings)
	if err != nil {
		return nil, err
	}
	if autoscalings.Status != "success" && autoscalings.Status != "" {
		return nil, errors.New(autoscalings.Message)
	}

	return autoscalings.Groups[0].Loadbalancers, nil
}

func (s *AutoScalingService) DeleteLoadbalancer(autoScalingeId, autoScalingLoadbalancerId string) (*DeleteResponse, error) {
	reqUrl := "autoscaling/" + autoScalingeId + "/loadbalancerpolicy/" + autoScalingLoadbalancerId
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

// Auto Scaling Security Group
type CreateAutoScalingSecurityGroupParams struct {
	AutoScalingId              string
	AutoScalingSecurityGroupId string
}

func (s *AutoScalingService) CreateSecurityGroup(params CreateAutoScalingSecurityGroupParams) (*CreateResponse, error) {
	reqUrl := "autoscaling/" + params.AutoScalingId + "/securitygroup/" + params.AutoScalingSecurityGroupId
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var autoscaling CreateResponse
	_, err := s.client.Do(req, &autoscaling)
	if err != nil {
		return nil, err
	}
	if autoscaling.Status != "success" && autoscaling.Status != "" {
		return nil, errors.New(autoscaling.Message)
	}

	return &autoscaling, nil
}

func (s *AutoScalingService) ReadSecurityGroup(autoscalingId, securitygroupId string) (*SecurityGroup, error) {
	reqUrl := "autoscaling/" + autoscalingId
	req, _ := s.client.NewRequest("GET", reqUrl)

	var autoscalings AutoScalings
	_, err := s.client.Do(req, &autoscalings)
	if err != nil {
		return nil, err
	}
	if autoscalings.Status != "success" && autoscalings.Status != "" {
		return nil, errors.New(autoscalings.Message)
	}
	var securitygroups SecurityGroup
	for _, r := range autoscalings.Groups[0].SecurityGroups {
		if r.ID == securitygroupId {
			securitygroups = r
		}
	}
	if len(securitygroups.ID) == 0 {
		return nil, errors.New("auto scaling securitygroup not found")
	}

	return &securitygroups, nil
}

func (s *AutoScalingService) ListSecurityGroups(autoscalingId string) ([]SecurityGroup, error) {
	reqUrl := "autoscaling/" + autoscalingId
	req, _ := s.client.NewRequest("GET", reqUrl)

	var autoscalings AutoScalings
	_, err := s.client.Do(req, &autoscalings)
	if err != nil {
		return nil, err
	}
	if autoscalings.Status != "success" && autoscalings.Status != "" {
		return nil, errors.New(autoscalings.Message)
	}

	return autoscalings.Groups[0].SecurityGroups, nil
}

func (s *AutoScalingService) DeleteSecurityGroup(autoScalingeId, autoScalingSecurityGroupId string) (*DeleteResponse, error) {
	reqUrl := "autoscaling/" + autoScalingeId + "/securitygroup/" + autoScalingSecurityGroupId
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

// Auto Scaling Target group
type CreateAutoScalingTargetgroupParams struct {
	AutoScalingId            string
	AutoScalingTargetgroupId string
}

func (s *AutoScalingService) CreateTargetgroup(params CreateAutoScalingTargetgroupParams) (*CreateResponse, error) {
	reqUrl := "autoscaling/" + params.AutoScalingId + "/targetgroup/" + params.AutoScalingTargetgroupId
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var autoscaling CreateResponse
	_, err := s.client.Do(req, &autoscaling)
	if err != nil {
		return nil, err
	}
	if autoscaling.Status != "success" && autoscaling.Status != "" {
		return nil, errors.New(autoscaling.Message)
	}

	return &autoscaling, nil
}

func (s *AutoScalingService) ReadTargetgroup(autoscalingId, targetgroupId string) (*AutoScalingTargetGroup, error) {
	reqUrl := "autoscaling/" + autoscalingId
	req, _ := s.client.NewRequest("GET", reqUrl)

	var autoscalings AutoScalings
	_, err := s.client.Do(req, &autoscalings)
	if err != nil {
		return nil, err
	}
	if autoscalings.Status != "success" && autoscalings.Status != "" {
		return nil, errors.New(autoscalings.Message)
	}
	var targetgroups AutoScalingTargetGroup
	for _, r := range autoscalings.Groups[0].TargetGroups {
		if r.ID == targetgroupId {
			targetgroups = r
		}
	}
	if len(targetgroups.ID) == 0 {
		return nil, errors.New("auto scaling targetgroup not found")
	}

	return &targetgroups, nil
}

func (s *AutoScalingService) ListTargetgroups(autoscalingId string) ([]AutoScalingTargetGroup, error) {
	reqUrl := "autoscaling/" + autoscalingId
	req, _ := s.client.NewRequest("GET", reqUrl)

	var autoscalings AutoScalings
	_, err := s.client.Do(req, &autoscalings)
	if err != nil {
		return nil, err
	}
	if autoscalings.Status != "success" && autoscalings.Status != "" {
		return nil, errors.New(autoscalings.Message)
	}

	return autoscalings.Groups[0].TargetGroups, nil
}

func (s *AutoScalingService) DeleteTargetgroup(autoScalingeId, autoScalingTargetgroupId string) (*DeleteResponse, error) {
	reqUrl := "autoscaling/" + autoScalingeId + "/targetgroup/" + autoScalingTargetgroupId
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
