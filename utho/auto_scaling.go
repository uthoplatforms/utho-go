package utho

import (
	"errors"
	"time"
)

type AutoScalingService service

type AutoScalings struct {
	Groups  []Groups `json:"groups" faker:"slice_len=1"`
	Status  string   `json:"status" faker:"oneof:success,error"`
	Message string   `json:"message" faker:"sentence"`
}
type Groups struct {
	ID                 string                     `json:"id" faker:"oneof: 00000,11111,22222,33333"`
	Userid             string                     `json:"userid" faker:"oneof: 00000,11111,22222,33333"`
	Name               string                     `json:"name"`
	Dcslug             string                     `json:"dcslug"`
	Minsize            string                     `json:"minsize" faker:"oneof:1,2,3,4,5"`
	Maxsize            string                     `json:"maxsize" faker:"oneof:1,2,3,4,5"`
	Desiredsize        string                     `json:"desiredsize" faker:"oneof:1,2,3,4,5"`
	Planid             string                     `json:"planid" faker:"oneof: 00000,11111,22222,33333"`
	Planname           string                     `json:"planname"`
	InstanceTemplateid string                     `json:"instance_templateid"`
	Image              string                     `json:"image"`
	ImageName          string                     `json:"image_name"`
	Snapshotid         string                     `json:"snapshotid" faker:"oneof: 00000,11111,22222,33333"`
	Status             string                     `json:"status" faker:"oneof:Active,Stopped"`
	CreatedAt          string                     `json:"created_at" faker:"date"`
	SuspendedAt        string                     `json:"suspended_at" faker:"date"`
	StoppedAt          string                     `json:"stopped_at" faker:"date"`
	StartedAt          string                     `json:"started_at" faker:"date"`
	DeletedAt          string                     `json:"deleted_at" faker:"date"`
	PublicIPEnabled    string                     `json:"public_ip_enabled" faker:"oneof:0,1"`
	Vpc                []AutoScalingVpc           `json:"vpc" faker:"slice_len=1"`
	CooldownTill       string                     `json:"cooldown_till" faker:"date"`
	Loadbalancers      []AutoScalingLoadbalancers `json:"load_balancers" faker:"slice_len=1"`
	TargetGroups       []AutoScalingTargetGroup   `json:"target_groups" faker:"slice_len=1"`
	SecurityGroups     []SecurityGroup            `json:"security_groups" faker:"slice_len=1"`
	Backupid           string                     `json:"backupid" faker:"oneof: 00000,11111,22222,33333"`
	Stack              string                     `json:"stack" faker:"oneof: 00000,11111,22222,33333"`
	StackFields        string                     `json:"stack_fields"`
	Instances          []Instances                `json:"instances" faker:"slice_len=1"`
	Policies           []Policy                   `json:"policies" faker:"slice_len=1"`
	Schedules          []Schedule                 `json:"schedules" faker:"slice_len=1"`
	DeletedInstances   []any                      `json:"deleted_instances"`
	Dclocation         AutoScalingDclocation      `json:"dclocation"`
	Plan               AutoScalingPlan            `json:"plan"`
}
type AutoScalingVpc struct {
	ID               string                 `json:"id" faker:"oneof: 00000,11111,22222,33333"`
	Total            int                    `json:"total" faker:"boundary_start=1,boundary_end=255"`
	Available        int                    `json:"available" faker:"boundary_start=1,boundary_end=255"`
	Network          string                 `json:"network" faker:"ipv4"`
	Name             string                 `json:"name"`
	Size             string                 `json:"size" faker:"oneof:24,16,8"`
	Dcslug           string                 `json:"dcslug"`
	Dclocation       AutoScalingDclocation  `json:"dclocation"`
	IsDefault        string                 `json:"is_default" faker:"oneof:0,1"`
	IsVpc            string                 `json:"is_vpc" faker:"oneof:0,1"`
	Routetable       []any                  `json:"routetable"`
	Subnets          []any                  `json:"subnets"`
	InternetGateways []any                  `json:"internet_gateways"`
	Resources        []AutoScalingResources `json:"resources" faker:"slice_len=1"`
}
type AutoScalingResources struct {
	Type         string                `json:"type"`
	ResourceID   string                `json:"resource_id" faker:"oneof: 00000,11111,22222,33333"`
	ID           string                `json:"id" faker:"oneof: 00000,11111,22222,33333"`
	Name         string                `json:"name"`
	RAM          string                `json:"ram" faker:"oneof:1024,2048,4096,8192,16384"`
	CPU          string                `json:"cpu" faker:"oneof:1,2,4,8,16"`
	Disksize     string                `json:"disksize" faker:"oneof:20,50,80,100"`
	DedicatedCPU any                   `json:"dedicated_cpu"`
	Dclocation   AutoScalingDclocation `json:"dclocation"`
	IP           string                `json:"ip" faker:"ipv4"`
	PrivateIP    string                `json:"private_ip" faker:"ipv4"`
	Vpc          string                `json:"vpc"`
	Source       string                `json:"source"`
}
type AutoScalingLoadbalancers struct {
	ID   string `json:"lbid" faker:"oneof: 00000,11111,22222,33333"`
	Name string `json:"name"`
	IP   string `json:"ip" faker:"ipv4"`
}
type AutoScalingTargetGroup struct {
	ID       string `json:"id" faker:"oneof: 00000,11111,22222,33333"`
	Name     string `json:"name"`
	Protocol string `json:"protocol"`
	Port     string `json:"port" faker:"oneof:80,443,8080"`
}
type SecurityGroup struct {
	ID   string `json:"id" faker:"oneof: 00000,11111,22222,33333"`
	Name string `json:"name"`
}
type Instances struct {
	ID        string `json:"cloudid" faker:"oneof: 00000,11111,22222,33333"`
	Hostname  string `json:"hostname"`
	CreatedAt string `json:"created_at" faker:"date"`
	IP        string `json:"ip" faker:"ipv4"`
	Status    string `json:"status" faker:"oneof:Active,Stopped"`
}
type Policy struct {
	ID                 string `json:"id" faker:"oneof: 00000,11111,22222,33333"`
	Userid             string `json:"userid" faker:"oneof: 00000,11111,22222,33333"`
	Product            string `json:"product"`
	Productid          string `json:"productid" faker:"oneof: 00000,11111,22222,33333"`
	Groupid            string `json:"groupid" faker:"oneof: 00000,11111,22222,33333"`
	Name               string `json:"name"`
	Type               string `json:"type"`
	Adjust             string `json:"adjust" faker:"oneof:1,2,3"`
	Period             string `json:"period" faker:"oneof:5m,10m,15m"`
	Cooldown           string `json:"cooldown" faker:"oneof:60,120,300"`
	CooldownTill       string `json:"cooldown_till" faker:"date"`
	Compare            string `json:"compare" faker:"oneof:above,below"`
	Value              string `json:"value" faker:"oneof:1,2,3"`
	AlertID            string `json:"alert_id" faker:"oneof: 00000,11111,22222,33333"`
	Status             string `json:"status" faker:"oneof:1,0"`
	KubernetesID       string `json:"kubernetes_id" faker:"oneof: 00000,11111,22222,33333"`
	KubernetesNodepool string `json:"kubernetes_nodepool"`
	Cloudid            string `json:"cloudid" faker:"oneof: 00000,11111,22222,33333"`
	Maxsize            string `json:"maxsize" faker:"oneof:1,2,3,4,5"`
	Minsize            string `json:"minsize" faker:"oneof:1,2,3,4,5"`
}
type Schedule struct {
	ID          string `json:"id" faker:"oneof: 00000,11111,22222,33333"`
	Groupid     string `json:"groupid" faker:"oneof: 00000,11111,22222,33333"`
	Name        string `json:"name"`
	Desiredsize string `json:"desiredsize" faker:"oneof:1,2,3,4,5"`
	Recurrence  string `json:"recurrence"`
	StartDate   string `json:"start_date" faker:"date"`
	Status      string `json:"status" faker:"oneof:1,0"`
	Timezone    string `json:"timezone"`
}
type AutoScalingPlan struct {
	Planid         string `json:"planid" faker:"oneof: 00000,11111,22222,33333"`
	RAM            string `json:"ram" faker:"oneof:1024,2048,4096,8192,16384"`
	CPU            string `json:"cpu" faker:"oneof:1,2,4,8,16"`
	Disk           string `json:"disk" faker:"oneof:20,50,80,100"`
	Bandwidth      string `json:"bandwidth" faker:"oneof:100,500,1000,2000"`
	DedicatedVcore string `json:"dedicated_vcore" faker:"oneof:0,1"`
}
type AutoScalingDclocation struct {
	Location string `json:"location"`
	Country  string `json:"country,omitempty"`
	DC       string `json:"dc,omitempty"`
	Dccc     string `json:"dccc"`
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
	PublicIPEnabled    string                  `json:"public_ip_enabled"`
	Vpc                string                  `json:"vpc"`
	LoadBalancers      string                  `json:"load_balancers"`
	SecurityGroups     string                  `json:"security_groups"`
	Policies           []CreatePoliciesParams  `json:"policies"`
	Schedules          []CreateSchedulesParams `json:"schedules"`
	Stackid            string                  `json:"stackid"`
	Stackimage         string                  `json:"stackimage"`
	TargetGroups       string                  `json:"target_groups"`
}
type CreatePoliciesParams struct {
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

type CreateSchedulesParams struct {
	Name         string    `json:"name"`
	Desiredsize  string    `json:"desiredsize"`
	Recurrence   string    `json:"recurrence"`
	StartDate    time.Time `json:"start_date"`
	SelectedTime string    `json:"selectedTime"`
	SelectedDate string    `json:"selectedDate"`
	Adjust       string    `json:"adjust"`
	Period       string    `json:"period"`
	Cooldown     string    `json:"cooldown"`
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
	Stackid       string `json:"stackid"`
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
	Adjust                string `json:"adjust"`
	Period                string `json:"period"`
	Cooldown              string `json:"cooldown"`
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
	if len(autoscalings.Groups) == 0 || len(autoscalings.Groups[0].SecurityGroups) == 0 {
		return nil, errors.New("no security groups found")
	}

	for _, r := range autoscalings.Groups[0].SecurityGroups {
		if r.ID == securitygroupId {
			return &r, nil
		}
	}

	return nil, errors.New("security group not found")
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
