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
	ID                 string                    `json:"id"`
	Userid             string                    `json:"userid"`
	Name               string                    `json:"name"`
	Dcslug             string                    `json:"dcslug"`
	Minsize            string                    `json:"minsize"`
	Maxsize            string                    `json:"maxsize"`
	Desiredsize        string                    `json:"desiredsize"`
	Planid             string                    `json:"planid"`
	Planname           string                    `json:"planname"`
	InstanceTemplateid string                    `json:"instance_templateid"`
	Image              string                    `json:"image"`
	ImageName          string                    `json:"image_name"`
	Snapshotid         string                    `json:"snapshotid"`
	Status             string                    `json:"status"`
	CreatedAt          string                    `json:"created_at"`
	SuspendedAt        string                    `json:"suspended_at"`
	StoppedAt          string                    `json:"stopped_at"`
	StartedAt          string                    `json:"started_at"`
	DeletedAt          string                    `json:"deleted_at"`
	PublicIPEnabled    string                    `json:"public_ip_enabled"`
	Vpc                string                    `json:"vpc"`
	CooldownTill       string                    `json:"cooldown_till"`
	LoadBalancers      []any                     `json:"load_balancers"`
	TargetGroups       []AutoScalingTargetGroups `json:"target_groups"`
	SecurityGroups     []SecurityGroups          `json:"security_groups"`
	Backupid           string                    `json:"backupid"`
	Stack              string                    `json:"stack"`
	StackFields        string                    `json:"stack_fields"`
	Instances          []Instances               `json:"instances"`
	Policies           []Policies                `json:"policies"`
	Schedules          []Schedules               `json:"schedules"`
	DeletedInstances   []any                     `json:"deleted_instances"`
	Dclocation         Dclocation                `json:"dclocation"`
	Plan               AutoScalingPlan           `json:"plan"`
}
type AutoScalingTargetGroups struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Protocol string `json:"protocol"`
	Port     string `json:"port"`
}
type SecurityGroups struct {
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
type Policies struct {
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
type Schedules struct {
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
	Stack              string                  `json:"stack"`
	Stackid            string                  `json:"stackid"`
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

func (s *AutoScalingService) CreateAutoScaling(params CreateAutoScalingParams) (*CreateResponse, error) {
	reqUrl := "autoscaling"
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

func (s *AutoScalingService) ReadAutoScaling(autoscalingId string) (*Groups, error) {
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

	return &autoscalings.Groups[0], nil
}

func (s *AutoScalingService) ListAutoScaling() (*[]Groups, error) {
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

	return &autoscalings.Groups, nil
}

type UpdateAutoScalingParams struct {
	AutoScalingId string
	Name          string `json:"name"`
	Minsize       string `json:"minsize"`
	Maxsize       string `json:"maxsize"`
	Desiredsize   string `json:"desiredsize"`
}

func (s *AutoScalingService) UpdateAutoScaling(params UpdateAutoScalingParams) (*UpdateResponse, error) {
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

func (s *AutoScalingService) DeleteAutoScaling(autoscalingId, autoscalingName string) (*DeleteResponse, error) {
	reqUrl := "autoscaling/" + autoscalingId + "?name=" + autoscalingName
	req, _ := s.client.NewRequest("DELETE", reqUrl)

	var delResponse DeleteResponse
	if _, err := s.client.Do(req, &delResponse); err != nil {
		return nil, err
	}

	return &delResponse, nil
}

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

func (s *AutoScalingService) CreateAutoScalingPolicy(params CreateAutoScalingPolicyParams) (*CreateResponse, error) {
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

func (s *AutoScalingService) ReadAutoScalingPolicy(autoscalingId, policyId string) (*Policies, error) {
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
	var policies Policies
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

func (s *AutoScalingService) ListAutoScalingPolicy(autoscalingId string) (*[]Policies, error) {
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

	return &autoscalings.Groups[0].Policies, nil
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

func (s *AutoScalingService) UpdateAutoScalingPolicy(params UpdateAutoScalingPolicyParams) (*UpdateResponse, error) {
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

func (s *AutoScalingService) DeleteAutoScalingPolicy(autoScalingPolicyId string) (*DeleteResponse, error) {
	reqUrl := "autoscaling/policy/" + autoScalingPolicyId
	req, _ := s.client.NewRequest("DELETE", reqUrl)

	var delResponse DeleteResponse
	if _, err := s.client.Do(req, &delResponse); err != nil {
		return nil, err
	}

	return &delResponse, nil
}

type CreateAutoScalingScheduleParams struct {
	AutoScalingId string
	Name          string `json:"name"`
	Desiredsize   string `json:"desiredsize"`
	Recurrence    string `json:"recurrence"`
	StartDate     string `json:"start_date"`
}

func (s *AutoScalingService) CreateAutoScalingSchedule(params CreateAutoScalingScheduleParams) (*CreateResponse, error) {
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

func (s *AutoScalingService) ReadAutoScalingSchedule(autoscalingId, scheduleId string) (*Schedules, error) {
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
	var schedules Schedules
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

func (s *AutoScalingService) ListAutoScalingSchedule(autoscalingId string) (*[]Schedules, error) {
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

	return &autoscalings.Groups[0].Schedules, nil
}

type UpdateAutoScalingScheduleParams struct {
	AutoScalingeId        string
	AutoScalingScheduleId string
	Name                  string `json:"name"`
	Desiredsize           string `json:"desiredsize"`
	Recurrence            string `json:"recurrence"`
	StartDate             string `json:"start_date"`
}

func (s *AutoScalingService) UpdateAutoScalingSchedule(params UpdateAutoScalingScheduleParams) (*UpdateResponse, error) {
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

func (s *AutoScalingService) DeleteAutoScalingSchedule(autoScalingeId, autoScalingScheduleId string) (*DeleteResponse, error) {
	reqUrl := "autoscaling/" + autoScalingeId + "/schedulepolicy/" + autoScalingScheduleId
	req, _ := s.client.NewRequest("DELETE", reqUrl)

	var delResponse DeleteResponse
	if _, err := s.client.Do(req, &delResponse); err != nil {
		return nil, err
	}

	return &delResponse, nil
}

// type CreateAutoScalingLoadbalancerParams struct {
// 	AutoScalingId  string
// 	LoadbalancerId string
// }

// func (s *AutoScalingService) CreateAutoScalingLoadbalancer(params CreateAutoScalingLoadbalancerParams) (*CreateResponse, error) {
// 	reqUrl := "autoscaling/" + params.AutoScalingId + "/loadbalancer/" + params.LoadbalancerId
// 	req, _ := s.client.NewRequest("POST", reqUrl, &params)

// 	var autoscaling CreateResponse
// 	_, err := s.client.Do(req, &autoscaling)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if autoscaling.Status != "success" && autoscaling.Status != "" {
// 		return nil, errors.New(autoscaling.Message)
// 	}

// 	return &autoscaling, nil
// }

// func (s *AutoScalingService) ReadAutoScalingLoadbalancer(autoscalingId, loadbalancerId string) (*Loadbalancers, error) {
// 	reqUrl := "autoscaling/" + autoscalingId
// 	req, _ := s.client.NewRequest("GET", reqUrl)

// 	var autoscalings AutoScalings
// 	_, err := s.client.Do(req, &autoscalings)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if autoscalings.Status != "success" && autoscalings.Status != "" {
// 		return nil, errors.New(autoscalings.Message)
// 	}
// 	var loadbalancers Loadbalancers
// 	for _, r := range autoscalings.Groups[0].LoadBalancers {
// 		if r.ID == loadbalancerId {
// 			loadbalancers = r
// 		}
// 	}
// 	if len(loadbalancers.ID) == 0 {
// 		return nil, errors.New("auto scaling loadbalancer not found")
// 	}

// 	return &loadbalancers, nil
// }

// func (s *AutoScalingService) ListAutoScalingLoadbalancer(autoscalingId string) (*[]Loadbalancers, error) {
// 	reqUrl := "autoscaling/" + autoscalingId
// 	req, _ := s.client.NewRequest("GET", reqUrl)

// 	var autoscalings AutoScalings
// 	_, err := s.client.Do(req, &autoscalings)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if autoscalings.Status != "success" && autoscalings.Status != "" {
// 		return nil, errors.New(autoscalings.Message)
// 	}

// 	return &autoscalings.Groups[0].Loadbalancers, nil
// }

// type UpdateAutoScalingLoadbalancerParams struct {
// 	AutoScalingeId            string
// 	AutoScalingLoadbalancerId string
// 	Name                      string `json:"name"`
// 	Desiredsize               string `json:"desiredsize"`
// 	Recurrence                string `json:"recurrence"`
// 	StartDate                 string `json:"start_date"`
// }

// func (s *AutoScalingService) UpdateAutoScalingLoadbalancer(params UpdateAutoScalingLoadbalancerParams) (*UpdateResponse, error) {
// 	reqUrl := "autoscaling/" + params.AutoScalingeId + "/loadbalancerpolicy/" + params.AutoScalingLoadbalancerId
// 	req, _ := s.client.NewRequest("PUT", reqUrl, &params)

// 	var autoscaling UpdateResponse
// 	_, err := s.client.Do(req, &autoscaling)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if autoscaling.Status != "success" && autoscaling.Status != "" {
// 		return nil, errors.New(autoscaling.Message)
// 	}

// 	return &autoscaling, nil
// }

// func (s *AutoScalingService) DeleteAutoScalingLoadbalancer(autoScalingeId, autoScalingLoadbalancerId string) (*DeleteResponse, error) {
// 	reqUrl := "autoscaling/" + autoScalingeId + "/loadbalancerpolicy/" + autoScalingLoadbalancerId
// 	req, _ := s.client.NewRequest("DELETE", reqUrl)

// 	var delResponse DeleteResponse
// 	if _, err := s.client.Do(req, &delResponse); err != nil {
// 		return nil, err
// 	}

// 	return &delResponse, nil
// }

type CreateAutoScalingSecurityGroupParams struct {
	AutoScalingId              string
	AutoScalingSecurityGroupId string
}

func (s *AutoScalingService) CreateAutoScalingSecurityGroup(params CreateAutoScalingSecurityGroupParams) (*CreateResponse, error) {
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

func (s *AutoScalingService) ReadAutoScalingSecurityGroup(autoscalingId, securitygroupId string) (*SecurityGroups, error) {
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
	var securitygroups SecurityGroups
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

func (s *AutoScalingService) ListAutoScalingSecurityGroup(autoscalingId string) (*[]SecurityGroups, error) {
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

	return &autoscalings.Groups[0].SecurityGroups, nil
}

func (s *AutoScalingService) DeleteAutoScalingSecurityGroup(autoScalingeId, autoScalingSecurityGroupId string) (*DeleteResponse, error) {
	reqUrl := "autoscaling/" + autoScalingeId + "/securitygroup/" + autoScalingSecurityGroupId
	req, _ := s.client.NewRequest("DELETE", reqUrl)

	var delResponse DeleteResponse
	if _, err := s.client.Do(req, &delResponse); err != nil {
		return nil, err
	}

	return &delResponse, nil
}

type CreateAutoScalingTargetgroupParams struct {
	AutoScalingId            string
	AutoScalingTargetgroupId string
}

func (s *AutoScalingService) CreateAutoScalingTargetgroup(params CreateAutoScalingTargetgroupParams) (*CreateResponse, error) {
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

func (s *AutoScalingService) ReadAutoScalingTargetgroup(autoscalingId, targetgroupId string) (*AutoScalingTargetGroups, error) {
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
	var targetgroups AutoScalingTargetGroups
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

func (s *AutoScalingService) ListAutoScalingTargetgroup(autoscalingId string) (*[]AutoScalingTargetGroups, error) {
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

	return &autoscalings.Groups[0].TargetGroups, nil
}

func (s *AutoScalingService) DeleteAutoScalingTargetgroup(autoScalingeId, autoScalingTargetgroupId string) (*DeleteResponse, error) {
	reqUrl := "autoscaling/" + autoScalingeId + "/targetgroup/" + autoScalingTargetgroupId
	req, _ := s.client.NewRequest("DELETE", reqUrl)

	var delResponse DeleteResponse
	if _, err := s.client.Do(req, &delResponse); err != nil {
		return nil, err
	}

	return &delResponse, nil
}
