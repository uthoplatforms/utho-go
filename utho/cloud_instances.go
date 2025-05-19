package utho

import (
	"encoding/json"
	"errors"
)

type CloudInstancesService service

type CloudInstances struct {
	CloudInstance []CloudInstance `json:"cloud" faker:"slice_len=2"`
	Meta          Meta            `json:"meta"`
	Status        string          `json:"status,omitempty" faker:"oneof:success,error"`
	Message       string          `json:"message,omitempty" faker:"sentence"`
}
type CloudInstance struct {
	ID                string                   `json:"cloudid" faker:"oneof: 00000,11111,22222,33333"`
	Hostname          string                   `json:"hostname"`
	CPU               string                   `json:"cpu" faker:"oneof:1,2,4,8,16"`
	RAM               string                   `json:"ram" faker:"oneof:1024,2048,4096,8192,16384"`
	DiscountType      string                   `json:"discount_type" faker:"oneof:Percentage,Flat"`
	DiscountValue     string                   `json:"discount_value" faker:"oneof:,10,20,30"`
	ManagedOs         string                   `json:"managed_os,omitempty"`
	ManagedFull       string                   `json:"managed_full,omitempty"`
	ManagedOnetime    string                   `json:"managed_onetime,omitempty"`
	PlanDisksize      int                      `json:"plan_disksize" faker:"boundary_start=20,boundary_end=500"`
	Disksize          int                      `json:"disksize" faker:"boundary_start=20,boundary_end=500"`
	Ha                string                   `json:"ha" faker:"oneof:0,1"`
	Status            string                   `json:"status" faker:"oneof:Active,Stopped,Pending"`
	Iso               string                   `json:"iso,omitempty"`
	IP                string                   `json:"ip" faker:"ipv4"`
	Billingcycle      string                   `json:"billingcycle" faker:"oneof:hourly,monthly"`
	Cost              float64                  `json:"cost" faker:"amount"`
	Vmcost            float64                  `json:"vmcost" faker:"amount"`
	Imagecost         int                      `json:"imagecost" faker:"boundary_start=0,boundary_end=10"`
	Backupcost        float64                  `json:"backupcost" faker:"amount"`
	Hourlycost        float64                  `json:"hourlycost" faker:"amount"`
	Cloudhourlycost   float64                  `json:"cloudhourlycost" faker:"amount"`
	Imagehourlycost   int                      `json:"imagehourlycost" faker:"boundary_start=0,boundary_end=10"`
	Backuphourlycost  float64                  `json:"backuphourlycost" faker:"amount"`
	Creditrequired    float64                  `json:"creditrequired" faker:"amount"`
	Creditreserved    int                      `json:"creditreserved" faker:"boundary_start=0,boundary_end=10"`
	Nextinvoiceamount float64                  `json:"nextinvoiceamount" faker:"amount"`
	Nextinvoicehours  string                   `json:"nextinvoicehours" faker:"oneof:1,2,3,4,5,6,7,8,9,10"`
	Consolepassword   string                   `json:"consolepassword" faker:"password"`
	Powerstatus       string                   `json:"powerstatus" faker:"oneof:Running,Stopped"`
	CreatedAt         string                   `json:"created_at" faker:"date"`
	UpdatedAt         string                   `json:"updated_at" faker:"date"`
	Nextduedate       string                   `json:"nextduedate" faker:"date"`
	Bandwidth         string                   `json:"bandwidth" faker:"oneof:100,500,1000,2000"`
	BandwidthUsed     int                      `json:"bandwidth_used" faker:"boundary_start=0,boundary_end=1000"`
	BandwidthFree     int                      `json:"bandwidth_free" faker:"boundary_start=0,boundary_end=1000"`
	Features          Features                 `json:"features"`
	Image             Image                    `json:"image"`
	Dclocation        Dclocation               `json:"dclocation"`
	V4                V4Public                 `json:"v4"`
	Networks          Networks                 `json:"networks"`
	V4Private         V4Private                `json:"v4private"`
	Storages          []Storages               `json:"storages,omitempty" faker:"slice_len=1"`
	Storage           Storage                  `json:"storage"`
	DiskUsed          int                      `json:"disk_used" faker:"boundary_start=0,boundary_end=500"`
	DiskFree          int                      `json:"disk_free" faker:"boundary_start=0,boundary_end=500"`
	DiskUsedp         int                      `json:"disk_usedp" faker:"boundary_start=0,boundary_end=100"`
	Backups           []any                    `json:"backups,omitempty"`
	Snapshots         []Snapshots              `json:"snapshots,omitempty" faker:"slice_len=1"`
	Firewalls         []CloudInstanceFirewalls `json:"firewalls,omitempty" faker:"slice_len=1"`
	GpuAvailable      string                   `json:"gpu_available,omitempty" faker:"oneof:0,1"`
	Gpus              []any                    `json:"gpus,omitempty"`
	Rescue            int                      `json:"rescue" faker:"boundary_start=0,boundary_end=1"`
}
type Features struct {
	Backups string `json:"backups" faker:"oneof:0,1"`
}
type Image struct {
	Name         string `json:"name" faker:"word"`
	Distribution string `json:"distribution" faker:"word"`
	Version      string `json:"version" faker:"semver"`
	Image        string `json:"image" faker:"word"`
	Cost         string `json:"cost" faker:"amount"`
}
type Networks struct {
	Public  Public  `json:"public"`
	Private Private `json:"private"`
}
type Public struct {
	V4 V4PublicArray `json:"v4" faker:"slice_len=1"`
}
type V4Public struct {
	IPAddress string `json:"ip_address,omitempty" faker:"ipv4"`
	Netmask   string `json:"netmask,omitempty" faker:"ipv4_netmask"`
	Gateway   string `json:"gateway,omitempty" faker:"ipv4"`
	Type      string `json:"type,omitempty" faker:"oneof:public"`
	Nat       bool   `json:"nat,omitempty" faker:"bool"`
	Primary   string `json:"primary,omitempty" faker:"oneof:1,0"`
	Rdns      string `json:"rdns,omitempty" faker:"domain_name"`
	Mac       string `json:"mac,omitempty"`
}

type Private struct {
	V4 []V4Private `json:"v4" faker:"slice_len=1"`
}
type V4Private struct {
	Noip        int    `json:"noip" faker:"boundary_start=0,boundary_end=10"`
	IPAddress   string `json:"ip_address,omitempty" faker:"ipv4"`
	NatPublicIP string `json:"nat_publicip,omitempty" faker:"ipv4"`
	VpcName     string `json:"vpc_name,omitempty" faker:"word"`
	Network     string `json:"network,omitempty" faker:"ipv4"`
	VpcID       string `json:"vpc_id,omitempty" faker:"uuid_digit"`
	Netmask     string `json:"netmask,omitempty" faker:"ipv4_netmask"`
	Gateway     string `json:"gateway,omitempty" faker:"ipv4"`
	Type        string `json:"type,omitempty" faker:"oneof:private"`
	Mac         string `json:"mac,omitempty"`
	Primary     string `json:"primary,omitempty" faker:"oneof:1,0"`
}
type Storages struct {
	ID        string `json:"id" faker:"oneof: 00000,11111,22222,33333"`
	Size      int    `json:"size" faker:"boundary_start=20,boundary_end=500"`
	DiskUsed  string `json:"disk_used" faker:"oneof:1GB,50GB,100GB"`
	DiskFree  string `json:"disk_free" faker:"oneof:1GB,50GB,100GB"`
	DiskUsedp int    `json:"disk_usedp" faker:"boundary_start=0,boundary_end=100"`
	CreatedAt string `json:"created_at" faker:"date"`
	Bus       string `json:"bus" faker:"oneof:virtio,sata"`
	Type      string `json:"type" faker:"oneof:ssd,hdd"`
}
type Storage struct {
	ID        string `json:"id" faker:"oneof: 00000,11111,22222,33333"`
	Size      int    `json:"size" faker:"boundary_start=20,boundary_end=500"`
	DiskUsed  string `json:"disk_used" faker:"oneof:1GB,50GB,100GB"`
	DiskFree  string `json:"disk_free" faker:"oneof:1GB,50GB,100GB"`
	DiskUsedp int    `json:"disk_usedp" faker:"boundary_start=0,boundary_end=100"`
	CreatedAt string `json:"created_at" faker:"date"`
	Bus       string `json:"bus" faker:"oneof:virtio,sata"`
	Type      string `json:"type" faker:"oneof:ssd,hdd"`
}
type Snapshot struct {
	ID        string `json:"id" faker:"oneof: 00000,11111,22222,33333"`
	Size      string `json:"size" faker:"oneof:1GB,50GB,100GB"`
	CreatedAt string `json:"created_at" faker:"date"`
	Note      string `json:"note" faker:"sentence"`
	Name      string `json:"name" faker:"word"`
}
type CloudInstanceFirewall struct {
	ID        string `json:"id" faker:"oneof: 00000,11111,22222,33333"`
	Name      string `json:"name" faker:"word"`
	CreatedAt string `json:"created_at" faker:"date"`
}
type Meta struct {
	Total       int `json:"total" faker:"boundary_start=1,boundary_end=10"`
	Totalpages  int `json:"totalpages" faker:"boundary_start=1,boundary_end=5"`
	Currentpage int `json:"currentpage" faker:"boundary_start=1,boundary_end=5"`
}
type Snapshots struct {
	ID        string `json:"id" faker:"oneof: 00000,11111,22222,33333"`
	Size      string `json:"size" faker:"oneof:1GB,50GB,100GB"`
	CreatedAt string `json:"created_at" faker:"date"`
	Note      string `json:"note" faker:"sentence"`
	Name      string `json:"name" faker:"word"`
}
type CloudInstanceFirewalls struct {
	ID        string `json:"id" faker:"oneof: 00000,11111,22222,33333"`
	Name      string `json:"name" faker:"word"`
	CreatedAt string `json:"created_at" faker:"date"`
}

type OsImages struct {
	OsImages []OsImage `json:"images" faker:"slice_len=5"`
	Status   string    `json:"status,omitempty" faker:"oneof:success,error"`
	Message  string    `json:"message,omitempty" faker:"sentence"`
}
type OsImage struct {
	Distro       string `json:"distro" faker:"word"`
	Distribution string `json:"distribution" faker:"word"`
	Version      string `json:"version" faker:"semver"`
	Image        string `json:"image" faker:"word"`
	Cost         int    `json:"cost" faker:"boundary_start=0,boundary_end=10"`
}

type Plans struct {
	Plans   []Plan `json:"plans" faker:"slice_len=3"`
	Status  string `json:"status,omitempty" faker:"oneof:success,error"`
	Message string `json:"message,omitempty" faker:"sentence"`
}
type Plan struct {
	ID        string  `json:"id" faker:"oneof: 00000,11111,22222,33333"`
	Type      string  `json:"type" faker:"oneof:ramcpu,disk"`
	Disk      string  `json:"disk" faker:"oneof:20GB,50GB,100GB"`
	RAM       string  `json:"ram" faker:"oneof:1024,2048,4096,8192,16384"`
	CPU       string  `json:"cpu" faker:"oneof:1,2,4,8,16"`
	Bandwidth string  `json:"bandwidth" faker:"oneof:100,500,1000,2000"`
	Slug      string  `json:"slug" faker:"slug"`
	Price     float64 `json:"price" faker:"amount"`
	Monthly   float64 `json:"monthly" faker:"amount"`
	Plantype  string  `json:"plantype" faker:"oneof:cloud,dedicated"`
}

type CreateCloudInstanceParams struct {
	Dcslug         string          `json:"dcslug"`
	Image          string          `json:"image"`
	Planid         string          `json:"planid"`
	VpcId          string          `json:"vpc"`
	EnablePublicip string          `json:"enable_publicip"`
	SubnetRequired string          `json:"subnetRequired"`
	Cpumodel       string          `json:"cpumodel"`
	Auth           string          `json:"auth,omitempty"`
	RootPassword   string          `json:"root_password,omitempty"`
	Firewall       string          `json:"firewall"`
	Enablebackup   string          `json:"enablebackup,omitempty"`
	Support        string          `json:"support,omitempty"`
	Management     string          `json:"management,omitempty"`
	Billingcycle   string          `json:"billingcycle,omitempty"`
	Backupid       string          `json:"backupid,omitempty"`
	Snapshotid     string          `json:"snapshotid,omitempty"`
	Sshkeys        string          `json:"sshkeys,omitempty"`
	Cloud          []CloudHostname `json:"cloud"`
}

type CloudHostname struct {
	Hostname string `json:"hostname"`
}

type CreateCloudInstanceResponse struct {
	ID       string `json:"cloudid" faker:"oneof: 00000,11111,22222,33333"`
	Password string `json:"password" faker:"password"`
	Ipv4     string `json:"ipv4" faker:"ipv4"`
	Status   string `json:"status,omitempty" faker:"oneof:success,error"`
	Message  string `json:"message,omitempty" faker:"sentence"`
}

func (s *CloudInstancesService) Create(params CreateCloudInstanceParams) (*CreateCloudInstanceResponse, error) {
	reqUrl := "cloud/deploy"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var cloudInstances CreateCloudInstanceResponse
	_, err := s.client.Do(req, &cloudInstances)
	if err != nil {
		return nil, err
	}
	if cloudInstances.Status != "success" && cloudInstances.Status != "" {
		return nil, errors.New(cloudInstances.Message)
	}

	return &cloudInstances, nil
}

func (s *CloudInstancesService) Read(instanceId string) (*CloudInstance, error) {
	reqUrl := "cloud/" + instanceId
	req, _ := s.client.NewRequest("GET", reqUrl)

	var cloudInstances CloudInstances
	_, err := s.client.Do(req, &cloudInstances)
	if err != nil {
		return nil, err
	}
	if cloudInstances.Status != "success" && cloudInstances.Status != "" {
		return nil, errors.New(cloudInstances.Message)
	}
	if len(cloudInstances.CloudInstance) == 0 {
		return nil, errors.New("NotFound")
	}

	return &cloudInstances.CloudInstance[0], nil
}

func (s *CloudInstancesService) List() ([]CloudInstance, error) {
	reqUrl := "cloud"
	req, _ := s.client.NewRequest("GET", reqUrl)

	var cloudInstances CloudInstances
	_, err := s.client.Do(req, &cloudInstances)
	if err != nil {
		return nil, err
	}
	if cloudInstances.Status != "success" && cloudInstances.Status != "" {
		return nil, errors.New(cloudInstances.Message)
	}

	return cloudInstances.CloudInstance, nil
}

type DeleteCloudInstanceParams struct {
	// Please provide confirm string as follow: "I am aware this action will delete data and server permanently"
	Confirm string `json:"confirm"`
}

func (s *CloudInstancesService) Delete(cloudInstancesId string, deleteCloudInstanceParams DeleteCloudInstanceParams) (*DeleteResponse, error) {
	reqUrl := "cloud/" + cloudInstancesId + "/destroy"

	req, _ := s.client.NewRequest("DELETE", reqUrl, deleteCloudInstanceParams)

	var delResponse DeleteResponse
	if _, err := s.client.Do(req, &delResponse); err != nil {
		return nil, err
	}
	if delResponse.Status != "success" && delResponse.Status != "" {
		return nil, errors.New(delResponse.Message)
	}

	return &delResponse, nil
}

func (s *CloudInstancesService) ListOsImages() ([]OsImage, error) {
	reqUrl := "cloud/images"
	req, _ := s.client.NewRequest("GET", reqUrl)

	var osImages OsImages
	_, err := s.client.Do(req, &osImages)
	if err != nil {
		return nil, err
	}
	if osImages.Status != "success" && osImages.Status != "" {
		return nil, errors.New(osImages.Message)
	}

	return osImages.OsImages, nil
}

func (s *CloudInstancesService) ListResizePlans(instanceId string) ([]Plan, error) {
	reqUrl := "cloud/" + instanceId + "/resizeplans"
	req, _ := s.client.NewRequest("GET", reqUrl)

	var plans Plans
	_, err := s.client.Do(req, &plans)
	if err != nil {
		return nil, err
	}
	if plans.Status != "success" && plans.Status != "" {
		return nil, errors.New(plans.Message)
	}

	return plans.Plans, nil
}

type CreateSnapshotParams struct {
	Name string `json:"name"`
}

func (s *CloudInstancesService) CreateSnapshot(instanceId string, params CreateSnapshotParams) (*CreateBasicResponse, error) {
	reqUrl := "cloud/" + instanceId + "/snapshot/create"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var snapshot CreateBasicResponse
	_, err := s.client.Do(req, &snapshot)
	if err != nil {
		return nil, err
	}
	if snapshot.Status != "success" && snapshot.Status != "" {
		return nil, errors.New(snapshot.Message)
	}

	return &snapshot, nil
}

func (s *CloudInstancesService) DeleteSnapshot(cloudInstanceId, snapshotId string) (*DeleteResponse, error) {
	reqUrl := "cloud/" + cloudInstanceId + "/snapshot/" + snapshotId + "/delete"
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

func (s *CloudInstancesService) EnableBackup(instanceId string) (*BasicResponse, error) {
	reqUrl := "cloud/" + instanceId + "/backups/enable"
	req, _ := s.client.NewRequest("POST", reqUrl, nil)

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

func (s *CloudInstancesService) DisableBackup(instanceId string) (*BasicResponse, error) {
	reqUrl := "cloud/" + instanceId + "/backups/disable"
	req, _ := s.client.NewRequest("POST", reqUrl, nil)

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

type UpdateBillingCycleParams struct {
	Billingcycle string `json:"billingcycle"`
}

func (s *CloudInstancesService) UpdateBillingCycle(cloudid string, params UpdateBillingCycleParams) (*BasicResponse, error) {
	reqUrl := "cloud/" + cloudid + "/billingcycle"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

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

func (s *CloudInstancesService) HardReboot(instanceId string) (*BasicResponse, error) {
	reqUrl := "cloud/" + instanceId + "/hardreboot"
	req, _ := s.client.NewRequest("POST", reqUrl, nil)

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

func (s *CloudInstancesService) PowerCycle(instanceId string) (*BasicResponse, error) {
	reqUrl := "cloud/" + instanceId + "/powercycle"
	req, _ := s.client.NewRequest("POST", reqUrl, nil)

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

func (s *CloudInstancesService) PowerOff(instanceId string) (*BasicResponse, error) {
	reqUrl := "cloud/" + instanceId + "/poweroff"
	req, _ := s.client.NewRequest("POST", reqUrl, nil)

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

func (s *CloudInstancesService) PowerOn(instanceId string) (*BasicResponse, error) {
	reqUrl := "cloud/" + instanceId + "/poweron"
	req, _ := s.client.NewRequest("POST", reqUrl, nil)

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

type RebuildCloudInstanceParams struct {
	Image string `json:"image"`
	// Please provide confirm string as follow: "I am aware this action will delete data permanently and build a fresh server"
	Confirm string `json:"confirm"`
}

func (s *CloudInstancesService) Rebuild(instanceId string, rebuildCloudInstanceParams RebuildCloudInstanceParams) (*BasicResponse, error) {
	reqUrl := "cloud/" + instanceId + "/rebuild"
	req, _ := s.client.NewRequest("POST", reqUrl, rebuildCloudInstanceParams)

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

type ResetPasswordResponse struct {
	Password string `json:"password" faker:"password"`
	Status   string `json:"status,omitempty" faker:"oneof:success,error"`
	Message  string `json:"message,omitempty" faker:"sentence"`
}

func (s *CloudInstancesService) ResetPassword(instanceId string) (*ResetPasswordResponse, error) {
	reqUrl := "cloud/" + instanceId + "/resetpassword"
	req, _ := s.client.NewRequest("POST", reqUrl, nil)

	var resetPasswordResponse ResetPasswordResponse
	_, err := s.client.Do(req, &resetPasswordResponse)
	if err != nil {
		return nil, err
	}
	if resetPasswordResponse.Status != "success" && resetPasswordResponse.Status != "" {
		return nil, errors.New(resetPasswordResponse.Message)
	}

	return &resetPasswordResponse, nil
}

type ResizeCloudInstanceParams struct {
	Type string `json:"type"`
	Plan string `json:"plan"`
}

func (s *CloudInstancesService) Resize(instanceId string, resizeCloudInstanceParams ResizeCloudInstanceParams) (*BasicResponse, error) {
	reqUrl := "cloud/" + instanceId + "/resize"
	req, _ := s.client.NewRequest("POST", reqUrl, resizeCloudInstanceParams)

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

func (s *CloudInstancesService) RestoreSnapshot(instanceId, snapshotId string) (*BasicResponse, error) {
	reqUrl := "cloud/" + instanceId + "/snapshot/" + snapshotId + "/restore"
	req, _ := s.client.NewRequest("POST", reqUrl, nil)

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

type UpdateStorageParams struct {
	Bus  string `json:"bus"`
	Type string `json:"type"`
}

func (s *CloudInstancesService) UpdateStorage(cloudid, storageid string, params UpdateStorageParams) (*BasicResponse, error) {
	reqUrl := "cloud/" + cloudid + "/storage/" + storageid + "/update"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

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

func (s *CloudInstancesService) AssignPublicIP(cloudid string) (*BasicResponse, error) {
	reqUrl := "cloud/" + cloudid + "/assignpublicip"
	req, _ := s.client.NewRequest("POST", reqUrl, nil)

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

type UpdateRDNSParams struct {
	Rdns string `json:"rdns"`
}

func (s *CloudInstancesService) UpdateRDNS(cloudId, ipAddress string, params UpdateRDNSParams) (*BasicResponse, error) {
	reqUrl := "cloud/" + cloudId + "/updaterdns/" + ipAddress
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

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

func (s *CloudInstancesService) EnableRescue(cloudid string) (*BasicResponse, error) {
	reqUrl := "cloud/" + cloudid + "/enablerescue"
	req, _ := s.client.NewRequest("POST", reqUrl, nil)

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

func (s *CloudInstancesService) DisableRescue(cloudid string) (*BasicResponse, error) {
	reqUrl := "cloud/" + cloudid + "/disablerescue"
	req, _ := s.client.NewRequest("POST", reqUrl, nil)

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

type MountISOParams struct {
	Iso string `json:"iso"`
}

func (s *CloudInstancesService) MountISO(cloudid string, params MountISOParams) (*BasicResponse, error) {
	reqUrl := "cloud/" + cloudid + "/mountiso"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

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

func (s *CloudInstancesService) UnmountISO(cloudid string) (*BasicResponse, error) {
	reqUrl := "cloud/" + cloudid + "/umountiso"
	req, _ := s.client.NewRequest("POST", reqUrl, nil)

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

// Custom type to handle unmarshaling of V4Public
type V4PublicArray []V4Public

func (v *V4PublicArray) UnmarshalJSON(data []byte) error {
	var single []V4Public
	var nested [][]V4Public

	// Try unmarshaling as a single array
	if err := json.Unmarshal(data, &single); err == nil {
		*v = single
		return nil
	}

	// Try unmarshaling as a nested array
	if err := json.Unmarshal(data, &nested); err == nil {
		for _, inner := range nested {
			*v = append(*v, inner...)
		}
		return nil
	}

	return errors.New("invalid format for V4PublicArray")
}
