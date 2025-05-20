package utho

import (
	"errors"
)

type ObjectStorageService service

type Buckets struct {
	Buckets []Bucket `json:"buckets"`
	Status  string   `json:"status,omitempty"`
	Message string   `json:"message,omitempty"`
}
type Bucket struct {
	Name           string           `json:"name"`
	Access         string           `json:"access"`
	Dcslug         string           `json:"dcslug"`
	Size           string           `json:"size"`
	Status         string           `json:"status"`
	CreatedAt      string           `json:"created_at"`
	ObjectCount    string           `json:"object_count"`
	VersionEnabled bool             `json:"version_enabled"`
	CurrentSize    string           `json:"current_size"`
	Permissions    []Permissions    `json:"permissions"`
	Dclocation     BucketDclocation `json:"dclocation"`
}
type BucketDclocation struct {
	Location string `json:"location"`
	Country  string `json:"country"`
	Dc       string `json:"dc"`
	Dccc     string `json:"dccc"`
}
type Permissions struct {
	Bucket     string `json:"bucket"`
	Name       string `json:"name"`
	Accesskey  string `json:"accesskey"`
	Permission string `json:"permission"`
	Status     any    `json:"status" faker:"-"`
	CreatedAt  string `json:"created_at"`
}

type AccessKeys struct {
	AccessKeys []AccessKey `json:"accesskeys"`
	Status     string      `json:"status,omitempty"`
	Message    string      `json:"message,omitempty"`
}
type AccessKey struct {
	Name      string `json:"name"`
	Accesskey string `json:"accesskey"`
	Dcslug    string `json:"dcslug"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

type Objects struct {
	Objects []Object `json:"objects"`
	Path    string   `json:"path"`
	Status  string   `json:"status,omitempty"`
	Message string   `json:"message,omitempty"`
}
type Object struct {
	Size string `json:"size"`
	Type string `json:"type"`
	Name string `json:"name"`
}

type PlanList struct {
	Pricing []Pricing `json:"pricing"`
	Rcode   string    `json:"rcode"`
	Status  string    `json:"status,omitempty"`
	Message string    `json:"message,omitempty"`
}
type Pricing struct {
	ID             string `json:"id"`
	UUID           string `json:"uuid"`
	Type           string `json:"type"`
	Slug           string `json:"slug"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Disk           string `json:"disk"`
	RAM            string `json:"ram"`
	CPU            string `json:"cpu"`
	Bandwidth      string `json:"bandwidth"`
	IsFeatured     string `json:"is_featured"`
	DedicatedVcore string `json:"dedicated_vcore"`
	Price          int    `json:"price"`
	Monthly        string `json:"monthly"`
}

type CreateBucketParams struct {
	Dcslug  string `json:"dcslug"`
	Billing string `json:"billing"`
	Size    string `json:"size"`
	Price   string `json:"price"`
	Name    string `json:"name"`
}

func (s *ObjectStorageService) CreateBucket(params CreateBucketParams) (*CreateResponse, error) {
	reqUrl := "objectstorage/bucket/create"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var bucket CreateResponse
	_, err := s.client.Do(req, &bucket)
	if err != nil {
		return nil, err
	}
	if bucket.Status != "success" && bucket.Status != "" {
		return nil, errors.New(bucket.Message)
	}

	return &bucket, nil
}

func (s *ObjectStorageService) ReadBucket(dcslug, bucketName string) (*Bucket, error) {
	reqUrl := "objectstorage/" + dcslug + "/bucket"
	req, _ := s.client.NewRequest("GET", reqUrl)

	var buckets Buckets
	_, err := s.client.Do(req, &buckets)
	if err != nil {
		return nil, err
	}
	if buckets.Status != "success" && buckets.Status != "" {
		return nil, errors.New(buckets.Message)
	}

	var bucket Bucket
	for _, v := range buckets.Buckets {
		if v.Name == bucketName {
			bucket = v
			break
		}
	}
	if len(bucket.Name) == 0 {
		return nil, errors.New("bucket not found")
	}

	return &bucket, nil
}

func (s *ObjectStorageService) ListBuckets(dcslug string) ([]Bucket, error) {
	reqUrl := "objectstorage/" + dcslug + "/bucket"
	req, _ := s.client.NewRequest("GET", reqUrl)

	var buckets Buckets
	_, err := s.client.Do(req, &buckets)
	if err != nil {
		return nil, err
	}
	if buckets.Status != "success" && buckets.Status != "" {
		return nil, errors.New(buckets.Message)
	}

	return buckets.Buckets, nil
}

func (s *ObjectStorageService) DeleteBucket(dcslug, bucketName string) (*DeleteResponse, error) {
	reqUrl := "objectstorage/" + dcslug + "/bucket/" + bucketName + "/delete"
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

type CreateAccessKeyParams struct {
	Dcslug        string
	AccesskeyName string `json:"accesskey"`
}
type CreateAccessKeyResponse struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	Accesskey string `json:"accesskey"`
	Secretkey string `json:"secretkey"`
}

func (s *ObjectStorageService) CreateAccessKey(params CreateAccessKeyParams) (*CreateAccessKeyResponse, error) {
	reqUrl := "objectstorage/" + params.Dcslug + "/accesskey/create"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var accesskey CreateAccessKeyResponse
	_, err := s.client.Do(req, &accesskey)
	if err != nil {
		return nil, err
	}
	if accesskey.Status != "success" && accesskey.Status != "" {
		return nil, errors.New(accesskey.Message)
	}

	return &accesskey, nil
}

func (s *ObjectStorageService) ReadAccessKey(dcslug, accesskeyName string) (*AccessKey, error) {
	reqUrl := "objectstorage/" + dcslug + "/accesskeys"
	req, _ := s.client.NewRequest("GET", reqUrl)

	var accesskeys AccessKeys
	_, err := s.client.Do(req, &accesskeys)
	if err != nil {
		return nil, err
	}
	if accesskeys.Status != "success" && accesskeys.Status != "" {
		return nil, errors.New(accesskeys.Message)
	}

	var accesskey AccessKey
	for _, v := range accesskeys.AccessKeys {
		if v.Accesskey == accesskeyName {
			accesskey = v
			break
		}
	}
	if len(accesskey.Name) == 0 {
		return nil, errors.New("access key not found")
	}

	return &accesskey, nil
}

func (s *ObjectStorageService) ListAccessKeys(dcslug string) ([]AccessKey, error) {
	reqUrl := "objectstorage/" + dcslug + "/accesskeys"
	req, _ := s.client.NewRequest("GET", reqUrl)

	var accesskeys AccessKeys
	_, err := s.client.Do(req, &accesskeys)
	if err != nil {
		return nil, err
	}
	if accesskeys.Status != "success" && accesskeys.Status != "" {
		return nil, errors.New(accesskeys.Message)
	}
	if len(accesskeys.AccessKeys) == 0 {
		return []AccessKey{}, nil
	}
	return accesskeys.AccessKeys, nil
}

type UpdateBucketAccessControlParams struct {
	Dcslug     string
	BucketName string
	Policy     string `json:"policy"`
}

func (s *ObjectStorageService) UpdateBucketAccessControl(params UpdateBucketAccessControlParams) (*CreateResponse, error) {
	reqUrl := "objectstorage/" + params.Dcslug + "/bucket/" + params.BucketName + "/policy/" + params.Policy
	req, _ := s.client.NewRequest("POST", reqUrl, &params)
	var bucket CreateResponse
	_, err := s.client.Do(req, &bucket)
	if err != nil {
		return nil, err
	}
	if bucket.Status != "success" && bucket.Status != "" {
		return nil, errors.New(bucket.Message)
	}

	return &bucket, nil
}

type UpdateBucketAccessKeyPermissionParams struct {
	Dcslug         string
	BucketName     string
	PermissionName string
	AccessKeyId    string
}

func (s *ObjectStorageService) UpdateBucketAccessKeyPermission(params UpdateBucketAccessKeyPermissionParams) (*CreateResponse, error) {
	reqUrl := "objectstorage/" + params.Dcslug + "/bucket/" + params.BucketName + "/permission/" + params.PermissionName + "/accesskey/" + params.AccessKeyId
	req, _ := s.client.NewRequest("POST", reqUrl, nil)

	var permission CreateResponse
	_, err := s.client.Do(req, &permission)
	if err != nil {
		return nil, err
	}
	if permission.Status != "success" && permission.Status != "" {
		return nil, errors.New(permission.Message)
	}

	return &permission, nil
}

type CreateDirectroyParams struct {
	Dcslug     string
	BucketName string
	Path       string `json:"path"`
}

func (s *ObjectStorageService) CreateDirectroy(params CreateDirectroyParams) (*CreateResponse, error) {
	reqUrl := "objectstorage/" + params.Dcslug + "/bucket/" + params.BucketName + "/createdirectory"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var accesskey CreateResponse
	_, err := s.client.Do(req, &accesskey)
	if err != nil {
		return nil, err
	}
	if accesskey.Status != "success" && accesskey.Status != "" {
		return nil, errors.New(accesskey.Message)
	}

	return &accesskey, nil
}

func (s *ObjectStorageService) ListBucketObjectsAndDirectories(dcslug, bucketName, path string) ([]Object, error) {
	reqUrl := "objectstorage/" + dcslug + "/bucket/" + bucketName + "/objects?path=" + path
	req, _ := s.client.NewRequest("GET", reqUrl)

	var objects Objects
	_, err := s.client.Do(req, &objects)
	if err != nil {
		return nil, err
	}
	if objects.Status != "success" && objects.Status != "" {
		return nil, errors.New(objects.Message)
	}

	return objects.Objects, nil
}

func (s *ObjectStorageService) DeleteDirectroy(dcslug, bucketName, directoryName string) (*DeleteResponse, error) {
	reqUrl := "objectstorage/" + dcslug + "/bucket/" + bucketName + "/delete/object?path=" + directoryName
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

type GetSharableUrlOfObject struct {
	URL     string `json:"url"`
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

func (s *ObjectStorageService) GetSharableUrlOfObject(dcslug, bucketName, path string) (*GetSharableUrlOfObject, error) {
	reqUrl := "objectstorage/" + dcslug + "/bucket/" + bucketName + "/download?path=" + path
	req, _ := s.client.NewRequest("GET", reqUrl)

	var object GetSharableUrlOfObject
	_, err := s.client.Do(req, &object)
	if err != nil {
		return nil, err
	}
	if object.Status != "success" && object.Status != "" {
		return nil, errors.New(object.Message)
	}

	return &object, nil
}

func (s *ObjectStorageService) ListSubscriptionPlanPricing() ([]Pricing, error) {
	reqUrl := "pricing/objectstorage"
	req, _ := s.client.NewRequest("GET", reqUrl)

	var planList PlanList
	_, err := s.client.Do(req, &planList)
	if err != nil {
		return nil, err
	}
	if planList.Status != "success" && planList.Status != "" {
		return nil, errors.New(planList.Message)
	}
	if len(planList.Pricing) == 0 {
		return []Pricing{}, nil
	}
	return planList.Pricing, nil
}

type UploadFileParams struct {
	Dcslug     string
	BucketName string
	File       interface{}
}

func (s *ObjectStorageService) UploadFile(params UploadFileParams) (*CreateResponse, error) {
	reqUrl := "objectstorage/" + params.Dcslug + "/bucket/" + params.BucketName + "/upload/internal"
	req, _ := s.client.NewRequest("POST", reqUrl, nil)

	var response CreateResponse
	_, err := s.client.Do(req, &response)
	if err != nil {
		return nil, err
	}
	if response.Status != "success" && response.Status != "" {
		return nil, errors.New(response.Message)
	}
	return &response, nil
}

type ModifyAccessKeyParams struct {
	Dcslug    string
	Name      string
	Accesskey string `json:"accesskey"`
	Status    string `json:"status"`
}

func (s *ObjectStorageService) ModifyAccessKey(params ModifyAccessKeyParams) (*UpdateResponse, error) {
	reqUrl := "objectstorage/" + params.Dcslug + "/accesskey/" + params.Name + "/status"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var updateResponse UpdateResponse
	_, err := s.client.Do(req, &updateResponse)
	if err != nil {
		return nil, err
	}
	if updateResponse.Status != "success" && updateResponse.Status != "" {
		return nil, errors.New(updateResponse.Message)
	}
	return &updateResponse, nil
}
